package usecase

import (
	prRepo "avito_internship_PR/internal/pr/infrastructure/repository"
	"avito_internship_PR/internal/user/application/dto/input"
	"avito_internship_PR/internal/user/application/dto/output"
	"avito_internship_PR/internal/user/application/transformer"
)

type GetReviewsUseCase struct {
	prRepo      prRepo.IPullRequestRepository
	prTransform transformer.PullRequestTransformer
}

func NewGetReviewsUseCase(prRepo prRepo.IPullRequestRepository, prTransform transformer.PullRequestTransformer) *GetReviewsUseCase {
	return &GetReviewsUseCase{
		prRepo:      prRepo,
		prTransform: prTransform,
	}
}

func (uc *GetReviewsUseCase) Execute(getReviewsInput input.GetReviewsInput) (output.GetReviewsOutput, error) {
	var response output.GetReviewsOutput
	prs, err := uc.prRepo.GetPullRequestsByReviewerUserID(getReviewsInput.UserId)
	if err != nil {
		return output.GetReviewsOutput{}, err
	}
	for _, pr := range prs {
		response.PullRequests = append(response.PullRequests, uc.prTransform.ToShortResponse(&pr))
	}
	response.UserID = getReviewsInput.UserId
	return response, nil
}
