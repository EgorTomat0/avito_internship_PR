package transformer

import (
	"avito_internship_PR/internal/user/application/dto"
	"avito_internship_PR/internal/user/application/dto/output"
)

type UserTransformer struct{}

func NewUserTransformer() *UserTransformer {
	return &UserTransformer{}
}

func (t *UserTransformer) ToResponse(user *output.SetIsActiveOutput) dto.UserResponse {
	return dto.UserResponse{User: user.User}
}
