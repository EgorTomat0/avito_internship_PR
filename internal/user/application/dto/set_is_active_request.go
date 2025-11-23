package dto

type SetIsActiveRequest struct {
	UserId   string `json:"user_id" binding:"required" validate:"required,uuid"`
	IsActive bool   `json:"is_active"`
}
