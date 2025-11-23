package repository

import (
	"log"

	"gorm.io/gorm"

	"avito_internship_PR/internal/db/models"
	errorPkg "avito_internship_PR/internal/error"
	"avito_internship_PR/internal/team/application/dto"
	"avito_internship_PR/internal/user/application/dto/input"
	"avito_internship_PR/internal/user/infrastructure/repository"
)

type TeamRepository struct {
	db       *gorm.DB
	userRepo repository.IUserRepository
}

func NewTeamRepository(db *gorm.DB, userRepo repository.IUserRepository) *TeamRepository {
	return &TeamRepository{db: db, userRepo: userRepo}
}

func (r *TeamRepository) GetByName(teamName string) (*models.Team, error) {
	var team models.Team
	err := r.db.
		Preload("TeamMembers.User").
		Where("team_name = ?", teamName).
		First(&team).Error
	if err != nil {
		return nil, errorPkg.NotFound
	}
	return &team, nil
}

func (r *TeamRepository) Create(teamDto *dto.TeamDto) (*models.Team, error) {
	team := models.Team{
		TeamName: teamDto.Name,
	}
	if err := r.db.Create(&team).Error; err != nil {
		return nil, errorPkg.TeamExists
	}
	for _, member := range teamDto.Members {
		_, err := r.userRepo.CreateOrUpdate(input.UserInput{
			UserId:   member.UserID,
			UserName: member.Username,
			IsActive: member.IsActive,
			TeamMembers: []models.TeamMember{
				{
					UserID: member.UserID,
					TeamID: team.ID,
				},
			},
		})
		if err != nil {
			log.Printf("ERROR IN USER REPO: %s", err)
			return nil, err
		}
	}
	err := r.db.Preload("TeamMembers").Preload("TeamMembers.User").Where("id = ?", team.ID).First(&team).Error
	if err != nil {
		log.Printf("ERROR IN USER REPO: %s", err)
		return nil, err
	}
	return &team, nil
}
