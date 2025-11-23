package input

import "avito_internship_PR/internal/db/models"

type UserInput struct {
	UserId      string
	UserName    string
	IsActive    bool
	TeamMembers []models.TeamMember
}
