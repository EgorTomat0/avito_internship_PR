package output

import (
	"avito_internship_PR/internal/pr/application/dto"
)

type GetReviewsOutput struct {
	UserID       string
	PullRequests []dto.PullRequestShortResponse
}
