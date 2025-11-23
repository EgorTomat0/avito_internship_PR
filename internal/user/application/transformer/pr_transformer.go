package transformer

import (
	"avito_internship_PR/internal/db/models"
	prDto "avito_internship_PR/internal/pr/application/dto"
)

type PullRequestTransformer struct{}

func NewPullRequestTransformer() *PullRequestTransformer {
	return &PullRequestTransformer{}
}

func (t *PullRequestTransformer) ToShortResponse(pr *models.PullRequest) prDto.PullRequestShortResponse {
	return prDto.PullRequestShortResponse{
		PullRequestID:   pr.ID,
		PullRequestName: pr.PullRequestName,
		AuthorID:        pr.AuthorID,
		Status:          string(pr.Status),
	}
}
