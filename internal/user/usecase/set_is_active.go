package usecase

import (
	"avito_internship_PR/internal/user/application/dto/input"
	"avito_internship_PR/internal/user/application/dto/output"
	"avito_internship_PR/internal/user/infrastructure/repository"
)

type SetIsActiveUseCase struct {
	userRepo repository.IUserRepository
}

func NewSetIsActiveUseCase(userRepo repository.IUserRepository) *SetIsActiveUseCase {
	return &SetIsActiveUseCase{
		userRepo: userRepo,
	}
}

func (uc *SetIsActiveUseCase) Execute(setIsActiveInput input.SetIsActiveInput) (output.SetIsActiveOutput, error) {
	user, err := uc.userRepo.SetUserActiveStatus(setIsActiveInput.UserId, setIsActiveInput.IsActive)
	if err != nil {
		return output.SetIsActiveOutput{}, err
	}
	return output.SetIsActiveOutput{
		User: output.UserOutput{
			UserId:   user.ID,
			Username: user.Username,
			TeamName: user.TeamMembers[0].Team.TeamName,
			IsActive: user.IsActive,
		},
	}, nil
}
