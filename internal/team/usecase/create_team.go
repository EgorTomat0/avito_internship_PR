package usecase

import (
	"avito_internship_PR/internal/team/application/dto"
	teamRepo "avito_internship_PR/internal/team/infrastructure/repository"
)

type CreateTeamUseCase struct {
	teamRepo teamRepo.ITeamRepository
}

func NewCreateTeamUseCase(teamRepo teamRepo.ITeamRepository) *CreateTeamUseCase {
	return &CreateTeamUseCase{
		teamRepo: teamRepo,
	}
}

func (uc *CreateTeamUseCase) Execute(input dto.TeamDto) (dto.TeamDto, error) {
	var response dto.TeamDto
	team, err := uc.teamRepo.Create(&input)
	if err != nil {
		return dto.TeamDto{}, err
	}
	for _, member := range team.TeamMembers {
		response.Members = append(response.Members, dto.TeamMemberDto{
			UserID:   member.UserID,
			Username: member.User.Username,
			IsActive: member.User.IsActive,
		})
	}
	response.Name = team.TeamName
	return response, nil
}
