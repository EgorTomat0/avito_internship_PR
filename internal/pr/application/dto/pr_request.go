package dto

import "time"

type CreatePRRequest struct {
	PullRequestID   string `json:"pull_request_id" binding:"required" validate:"required,uuid"`
	PullRequestName string `json:"pull_request_name" binding:"required" validate:"required,min=1,max=255"`
	AuthorID        string `json:"author_id" binding:"required" validate:"required,uuid"`
}

type CreatePRResponse struct {
	PR PullRequestResponse `json:"pr"`
}

type MergePRRequest struct {
	PullRequestID string `json:"pull_request_id" binding:"required" validate:"required,uuid"`
}

type MergePRResponse struct {
	PR       PullRequestResponse `json:"pr"`
	MergedAt time.Time           `json:"merged_at"`
}

type ReassignReviewerRequest struct {
	PullRequestID string `json:"pull_request_id" binding:"required" validate:"required,uuid"`
	OldReviewerID string `json:"old_reviewer_id" binding:"required" validate:"required,uuid"`
}

type ReassignReviewerResponse struct {
	PR         PullRequestResponse `json:"pr"`
	ReplacedBy string              `json:"replaced_by"`
}
