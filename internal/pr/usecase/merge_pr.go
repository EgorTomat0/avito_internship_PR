package usecase

import (
	"avito_internship_PR/internal/pr/application/dto"
	"avito_internship_PR/internal/pr/application/dto/input"
	"avito_internship_PR/internal/pr/application/dto/output"
	prRepo "avito_internship_PR/internal/pr/infrastructure/repository"
)

type MergePRUseCase struct {
	prRepo prRepo.IPullRequestRepository
}

func NewMergePRUseCase(prRepo prRepo.IPullRequestRepository) *MergePRUseCase {
	return &MergePRUseCase{
		prRepo: prRepo,
	}
}

func (uc *MergePRUseCase) Execute(mergePRInput input.MergePRInput) (output.MergePROutput, error) {
	pr, err := uc.prRepo.Merge(mergePRInput.PullRequestID)
	if err != nil {
		return output.MergePROutput{}, err
	}
	assignedReviewers := make([]string, len(pr.Reviewers))
	for i, reviewer := range pr.Reviewers {
		assignedReviewers[i] = reviewer.UserID
	}
	return output.MergePROutput{PR: &dto.PullRequestDto{
		PullRequestID:     pr.ID,
		PullRequestName:   pr.PullRequestName,
		AuthorID:          pr.AuthorID,
		Status:            pr.Status,
		AssignedReviewers: assignedReviewers,
	}, MergedAt: pr.MergedAt}, nil
}
