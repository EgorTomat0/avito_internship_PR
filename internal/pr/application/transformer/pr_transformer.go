package transformer

import (
	"avito_internship_PR/internal/pr/application/dto"
)

type PullRequestTransformer struct{}

func NewPullRequestTransformer() *PullRequestTransformer {
	return &PullRequestTransformer{}
}

func (t *PullRequestTransformer) ToResponse(pr *dto.PullRequestDto) dto.PullRequestResponse {
	return dto.PullRequestResponse{
		PullRequestID:     pr.PullRequestID,
		PullRequestName:   pr.PullRequestName,
		AuthorID:          pr.AuthorID,
		Status:            string(pr.Status),
		AssignedReviewers: pr.AssignedReviewers,
	}
}
