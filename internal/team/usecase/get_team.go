package usecase

import (
	"avito_internship_PR/internal/team/application/dto"
	"avito_internship_PR/internal/team/application/dto/input"
	"avito_internship_PR/internal/team/application/dto/output"
	"avito_internship_PR/internal/team/infrastructure/repository"
)

type GetTeamUseCase struct {
	teamRepo repository.ITeamRepository
}

func NewGetTeamUseCase(teamRepo repository.ITeamRepository) *GetTeamUseCase {
	return &GetTeamUseCase{
		teamRepo: teamRepo,
	}
}

func (uc *GetTeamUseCase) Execute(input input.GetTeamInput) (output.GetTeamOutput, error) {
	team, err := uc.teamRepo.GetByName(input.TeamName)
	if err != nil {
		return output.GetTeamOutput{}, err
	}
	teamMembers := make([]dto.TeamMemberDto, len(team.TeamMembers))
	for i, member := range team.TeamMembers {
		teamMembers[i] = dto.TeamMemberDto{
			UserID:   member.UserID,
			Username: member.User.Username,
			IsActive: member.User.IsActive,
		}
	}
	return output.GetTeamOutput{Team: &dto.TeamDto{
		Name:    team.TeamName,
		Members: teamMembers,
	}}, nil
}
