package repository

import (
	"avito_internship_PR/internal/db/models"
	"avito_internship_PR/internal/team/application/dto"
)

type ITeamRepository interface {
	Create(team *dto.TeamDto) (*models.Team, error)
	GetByName(teamName string) (*models.Team, error)
}
