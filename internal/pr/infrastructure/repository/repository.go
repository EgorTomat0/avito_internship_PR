package repository

import (
	"log"
	"time"

	"gorm.io/gorm"

	models2 "avito_internship_PR/internal/db/models"
	errPkg "avito_internship_PR/internal/error"
	"avito_internship_PR/internal/pr/application/dto"
)

type PullRequestRepository struct {
	db *gorm.DB
}

func NewPullRequestRepository(db *gorm.DB) *PullRequestRepository {
	return &PullRequestRepository{db: db}
}

func (r *PullRequestRepository) GetById(prId string) (*models2.PullRequest, error) {
	var pr models2.PullRequest
	err := r.db.
		Preload("Author").
		Preload("Reviewers").
		Preload("Reviewers.User").
		Preload("Reviewers.Team").
		Preload("Team").
		Where("id = ?", prId).
		First(&pr).Error
	if err != nil {
		return nil, errPkg.NotFound
	}

	return &pr, nil
}

func (r *PullRequestRepository) Create(prDto dto.PullRequestDto) (string, error) {
	var authorTeamMember models2.TeamMember
	var team models2.Team
	var teamMembers []models2.TeamMember
	err := r.db.Where("user_id = ?", prDto.AuthorID).First(&authorTeamMember).Error
	if err != nil {
		log.Printf("ERROR IN PR REPO: %s", err)
		return "", errPkg.NotFound
	}
	err = r.db.Where("id = ?", authorTeamMember.TeamID).First(&team).Error
	if err != nil {
		log.Printf("ERROR IN PR REPO: %s", err)
		return "", errPkg.NotFound
	}
	err = r.db.Where("team_id = ? AND user_id <> ?", team.ID, prDto.AuthorID).Limit(2).Find(&teamMembers).Error
	if err != nil {
		log.Printf("ERROR IN PR REPO: %s", err)
		return "", errPkg.NotFound
	}

	pr := models2.PullRequest{
		ID:              prDto.PullRequestID,
		PullRequestName: prDto.PullRequestName,
		AuthorID:        prDto.AuthorID,
		Status:          models2.PRStatusOpen,
		TeamID:          team.ID,
		Reviewers:       teamMembers,
	}

	if err := r.db.Create(&pr).Error; err != nil {
		log.Printf("ERROR IN PR REPO: %s", err)
		return "", errPkg.PrExists
	}

	return pr.ID, nil
}

func (r *PullRequestRepository) Merge(prID string) (*models2.PullRequest, error) {
	var pr models2.PullRequest
	now := time.Now()
	err := r.db.Model(models2.PullRequest{}).
		Where("id = ? AND merged_at IS NULL", prID).
		Updates(models2.PullRequest{
			Status:   models2.PRStatusMerged,
			MergedAt: now,
		}).Error
	if err != nil {
		log.Printf("ERROR IN PR REPO: %s", err)
		return nil, err
	}
	err = r.db.Preload("Reviewers").Where("id = ?", prID).First(&pr).Error
	if err != nil {
		return nil, errPkg.NotFound
	}
	pr.Status = models2.PRStatusMerged
	return &pr, nil
}

func (r *PullRequestRepository) ReassignReviewer(prID string, oldReviewerID string) (*models2.PullRequest, error) {
	var pr models2.PullRequest
	var newReviewer models2.TeamMember
	var status string
	r.db.Raw("SELECT status FROM pull_requests WHERE id = ?", prID).Scan(&status)
	if status == string(models2.PRStatusMerged) {
		return nil, errPkg.PrMerged
	}
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.
			Preload("Reviewers").
			Preload("Reviewers.User").
			Preload("Author").
			Where("id = ?", prID).
			First(&pr).Error; err != nil {
			return errPkg.NotFound
		}
		var oldTeamMember models2.TeamMember
		if err := tx.
			Where("user_id = ? AND team_id = ?", oldReviewerID, pr.TeamID).
			First(&oldTeamMember).Error; err != nil {
			return errPkg.NotFound
		}
		found := false
		for _, reviewer := range pr.Reviewers {
			if reviewer.ID == oldTeamMember.ID {
				found = true
				break
			}
		}
		if !found {
			return errPkg.NotAssigned
		}
		excludedUserIDs := []string{pr.AuthorID}
		for _, reviewer := range pr.Reviewers {
			excludedUserIDs = append(excludedUserIDs, reviewer.UserID)
		}
		var candidateTeamMembers []models2.TeamMember
		if err := tx.
			Preload("User").
			Where("team_id = ? AND user_id NOT IN (?)", pr.TeamID, excludedUserIDs).
			Find(&candidateTeamMembers).Error; err != nil {
			return errPkg.NoCandidate
		}
		var activeCandidates []models2.TeamMember
		for _, candidate := range candidateTeamMembers {
			if candidate.User.IsActive {
				activeCandidates = append(activeCandidates, candidate)
			}
		}
		if len(activeCandidates) == 0 {
			return errPkg.NoCandidate
		}
		newReviewer = activeCandidates[0]
		if err := tx.Model(&pr).Association("Reviewers").Delete(&oldTeamMember); err != nil {
			return err
		}
		if err := tx.Model(&pr).Association("Reviewers").Append(&newReviewer); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	err = r.db.
		Preload("Reviewers").
		Preload("Reviewers.User").
		Preload("Reviewers.Team").
		Preload("Author").
		Preload("Team").
		Where("id = ?", prID).
		First(&pr).Error
	if err != nil {
		log.Printf("ERROR IN PR REPO: %s", err)
		return nil, errPkg.NotFound
	}
	return &pr, nil
}

func (r *PullRequestRepository) GetPullRequestsByReviewerUserID(userID string) ([]models2.PullRequest, error) {
	var pullRequests []models2.PullRequest
	err := r.db.
		Joins("JOIN pull_request_reviewers ON pull_requests.id = pull_request_reviewers.pull_request_id").
		Joins("JOIN team_members ON pull_request_reviewers.team_member_id = team_members.id").
		Where("team_members.user_id = ?", userID).
		Preload("Reviewers").
		Preload("Reviewers.User").
		Preload("Reviewers.Team").
		Preload("Author").
		Preload("Team").
		Find(&pullRequests).Error

	if err != nil {
		log.Printf("ERROR IN PR REPO: %s", err)
		return nil, err
	}
	return pullRequests, nil
}
