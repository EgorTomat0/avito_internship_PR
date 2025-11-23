package dto

import "avito_internship_PR/internal/pr/application/dto"

type GetReviewsResponse struct {
	UserID       string                         `json:"user_id"`
	PullRequests []dto.PullRequestShortResponse `json:"pull_requests"`
}
