package repository

import (
	models2 "avito_internship_PR/internal/db/models"
	"avito_internship_PR/internal/pr/application/dto"
)

type IPullRequestRepository interface {
	GetById(prId string) (*models2.PullRequest, error)
	Create(pr dto.PullRequestDto) (string, error)
	Merge(prID string) (*models2.PullRequest, error)
	ReassignReviewer(prID string, oldReviewerID string) (*models2.PullRequest, error)
	GetPullRequestsByReviewerUserID(userID string) ([]models2.PullRequest, error)
}
