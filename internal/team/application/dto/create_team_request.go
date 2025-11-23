package dto

type CreateTeamRequest struct {
	TeamName string              `json:"team_name" binding:"required" validate:"required,min=1,max=255"`
	Members  []TeamMemberRequest `json:"members" binding:"required,dive" validate:"required,min=1,dive"`
}

type TeamMemberRequest struct {
	UserID   string `json:"user_id" binding:"required" validate:"required,uuid"`
	Username string `json:"username" binding:"required" validate:"required,min=1,max=255"`
	IsActive bool   `json:"is_active"`
}
