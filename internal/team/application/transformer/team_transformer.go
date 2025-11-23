package transformer

import (
	teamDto "avito_internship_PR/internal/team/application/dto"
)

type TeamTransformer struct{}

func NewTeamTransformer() *TeamTransformer {
	return &TeamTransformer{}
}

func (t *TeamTransformer) ToResponse(team *teamDto.TeamDto) teamDto.TeamResponse {
	members := make([]teamDto.TeamMemberResponse, len(team.Members))
	for i, member := range team.Members {
		members[i] = teamDto.TeamMemberResponse(member)
	}

	return teamDto.TeamResponse{
		TeamName: team.Name,
		Members:  members,
	}
}

func (t *TeamTransformer) ToCreateTeamDto(team *teamDto.CreateTeamRequest) teamDto.TeamDto {
	members := make([]teamDto.TeamMemberDto, len(team.Members))
	for i, member := range team.Members {
		members[i] = teamDto.TeamMemberDto(member)
	}

	return teamDto.TeamDto{
		Name:    team.TeamName,
		Members: members,
	}
}
