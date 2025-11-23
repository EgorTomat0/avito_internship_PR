package usecase

import (
	"avito_internship_PR/internal/pr/application/dto"
	"avito_internship_PR/internal/pr/application/dto/input"
	"avito_internship_PR/internal/pr/application/dto/output"
	prRepo "avito_internship_PR/internal/pr/infrastructure/repository"
)

type CreatePRUseCase struct {
	prRepo prRepo.IPullRequestRepository
}

func NewCreatePRUseCase(prRepo prRepo.IPullRequestRepository) *CreatePRUseCase {
	return &CreatePRUseCase{
		prRepo: prRepo,
	}
}

func (uc *CreatePRUseCase) Execute(createPRInput input.CreatePRInput) (output.CreatePROutput, error) {
	prId, err := uc.prRepo.Create(createPRInput.PullRequest)
	if err != nil {
		return output.CreatePROutput{}, err
	}
	pr, err := uc.prRepo.GetById(prId)
	if err != nil {
		return output.CreatePROutput{}, err
	}
	assignedReviewers := make([]string, len(pr.Reviewers))
	for i, reviewer := range pr.Reviewers {
		assignedReviewers[i] = reviewer.UserID
	}
	return output.CreatePROutput{PR: &dto.PullRequestDto{
		PullRequestID:     pr.ID,
		PullRequestName:   pr.PullRequestName,
		AuthorID:          pr.AuthorID,
		AssignedReviewers: assignedReviewers,
		Status:            pr.Status,
	}}, nil
}
