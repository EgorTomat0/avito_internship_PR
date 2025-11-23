package http

import (
	"errors"
	"net/http"
	"strings"

	errPkg "avito_internship_PR/internal/error"
	"avito_internship_PR/internal/team/application/dto"
	"avito_internship_PR/internal/team/application/dto/input"
	"avito_internship_PR/internal/team/application/transformer"
	"avito_internship_PR/internal/team/usecase"

	"github.com/gin-gonic/gin"
)

type TeamHandler struct {
	createTeamUseCase *usecase.CreateTeamUseCase
	getTeamUseCase    *usecase.GetTeamUseCase
	teamTransformer   *transformer.TeamTransformer
}

func NewTeamHandler(
	createTeamUseCase *usecase.CreateTeamUseCase,
	getTeamUseCase *usecase.GetTeamUseCase,
	teamTransformer *transformer.TeamTransformer,
) *TeamHandler {
	return &TeamHandler{
		createTeamUseCase: createTeamUseCase,
		getTeamUseCase:    getTeamUseCase,
		teamTransformer:   teamTransformer,
	}
}

func (h *TeamHandler) CreateTeam(c *gin.Context) {
	var req dto.CreateTeamRequest
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
	team, err := h.createTeamUseCase.Execute(h.teamTransformer.ToCreateTeamDto(&req))
	if err != nil {
		var customErr errPkg.CustomError
		if errors.As(err, &customErr) {
			statusCode := http.StatusBadRequest
			switch customErr.Code {
			case errPkg.ErrTeamExists:
				statusCode = http.StatusBadRequest
			}
			c.JSON(statusCode, errPkg.ErrResponse{Error: customErr})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "unknown error"})
		}
		return
	}
	teamResponse := h.teamTransformer.ToResponse(&team)
	c.JSON(http.StatusCreated, dto.CreateTeamResponse{
		Team: teamResponse,
	})
}

func (h *TeamHandler) GetTeam(c *gin.Context) {
	teamName := strings.TrimSpace(c.Query("team_name"))
	validate := errPkg.GetValidator()
	if err := validate.Var(teamName, "required"); err != nil {
		handledErr := errPkg.HandleValidationError(err)
		c.JSON(http.StatusBadRequest, handledErr)
		return
	}
	team, err := h.getTeamUseCase.Execute(input.GetTeamInput{TeamName: teamName})
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
	teamResponse := h.teamTransformer.ToResponse(&dto.TeamDto{
		Name:    team.Team.Name,
		Members: team.Team.Members,
	})
	c.JSON(http.StatusOK, dto.CreateTeamResponse{
		Team: teamResponse,
	})
}
