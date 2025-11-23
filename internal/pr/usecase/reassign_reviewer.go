package usecase

import (
	"avito_internship_PR/internal/pr/application/dto"
	"avito_internship_PR/internal/pr/application/dto/input"
	"avito_internship_PR/internal/pr/application/dto/output"
	prRepo "avito_internship_PR/internal/pr/infrastructure/repository"
)

type ReassignReviewerUseCase struct {
	prRepo prRepo.IPullRequestRepository
}

func NewReassignReviewerUseCase(prRepo prRepo.IPullRequestRepository) *ReassignReviewerUseCase {
	return &ReassignReviewerUseCase{
		prRepo: prRepo,
	}
}

func (uc *ReassignReviewerUseCase) Execute(reassignInput input.ReassignReviewerInput) (output.ReassignReviewerOutput, error) {
	pr, err := uc.prRepo.ReassignReviewer(reassignInput.PullRequestID, reassignInput.OldUserID)
	if err != nil {
		return output.ReassignReviewerOutput{}, err
	}
	assignedReviewers := make([]string, len(pr.Reviewers))
	for i, reviewer := range pr.Reviewers {
		assignedReviewers[i] = reviewer.UserID
	}
	return output.ReassignReviewerOutput{PR: &dto.PullRequestDto{
		PullRequestID:     pr.ID,
		PullRequestName:   pr.PullRequestName,
		AuthorID:          pr.AuthorID,
		AssignedReviewers: assignedReviewers,
	}, ReplacedBy: pr.Reviewers[len(pr.Reviewers)-1].UserID}, nil
}
