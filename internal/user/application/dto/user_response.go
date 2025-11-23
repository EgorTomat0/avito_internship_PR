package dto

import "avito_internship_PR/internal/user/application/dto/output"

type UserResponse struct {
	User output.UserOutput `json:"user"`
}
