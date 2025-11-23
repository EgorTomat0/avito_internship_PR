package dto

import "avito_internship_PR/internal/db/models"

type PullRequestDto struct {
	PullRequestID     string
	PullRequestName   string
	AuthorID          string
	Status            models.PRStatus
	AssignedReviewers []string
}
