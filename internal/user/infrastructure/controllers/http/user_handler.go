package http

import (
	"errors"
	"net/http"
	"strings"

	errPkg "avito_internship_PR/internal/error"
	"avito_internship_PR/internal/user/application/dto"
	"avito_internship_PR/internal/user/application/dto/input"
	userTransformer "avito_internship_PR/internal/user/application/transformer"
	"avito_internship_PR/internal/user/usecase"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	setIsActiveUseCase *usecase.SetIsActiveUseCase
	getReviewsUseCase  *usecase.GetReviewsUseCase
	userTransformer    *userTransformer.UserTransformer
}

func NewUserHandler(
	setIsActiveUseCase *usecase.SetIsActiveUseCase,
	getReviewsUseCase *usecase.GetReviewsUseCase,
	userTransformer *userTransformer.UserTransformer,
) *UserHandler {
	return &UserHandler{
		setIsActiveUseCase: setIsActiveUseCase,
		getReviewsUseCase:  getReviewsUseCase,
		userTransformer:    userTransformer,
	}
}

func (h *UserHandler) SetIsActive(c *gin.Context) {
	var req dto.SetIsActiveRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		handledErr := errPkg.HandleValidationError(err)
		c.JSON(http.StatusBadRequest, handledErr)
		return
	}
	validate := errPkg.GetValidator()
	if err := validate.Struct(&req); err != nil {
		handledErr := errPkg.HandleValidationError(err)
		c.JSON(http.StatusBadRequest, handledErr)
		return
	}
	user, err := h.setIsActiveUseCase.Execute(input.SetIsActiveInput{
		UserId:   req.UserId,
		IsActive: req.IsActive,
	})
	if err != nil {
		var customErr errPkg.CustomError
		if errors.As(err, &customErr) {
			statusCode := http.StatusBadRequest
			switch customErr.Code {
			case errPkg.ErrNotFound:
				statusCode = http.StatusNotFound
			}
			c.JSON(statusCode, errPkg.ErrResponse{Error: customErr})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "unknown error"})
		}
		return
	}
	userResponse := h.userTransformer.ToResponse(&user)
	c.JSON(http.StatusOK, userResponse)
}

func (h *UserHandler) GetReviews(c *gin.Context) {
	userID := strings.TrimSpace(c.Query("user_id"))
	validate := errPkg.GetValidator()
	if err := validate.Var(userID, "required,uuid"); err != nil {
		handledErr := errPkg.HandleValidationError(err)
		c.JSON(http.StatusBadRequest, handledErr)
		return
	}
	output, _ := h.getReviewsUseCase.Execute(input.GetReviewsInput{UserId: userID})
	response := dto.GetReviewsResponse{
		UserID:       output.UserID,
		PullRequests: output.PullRequests,
	}
	c.JSON(http.StatusOK, response)
}
