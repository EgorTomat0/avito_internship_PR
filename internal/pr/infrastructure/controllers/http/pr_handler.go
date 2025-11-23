package http

import (
	"errors"
	"net/http"

	errPkg "avito_internship_PR/internal/error"
	"avito_internship_PR/internal/pr/application/dto"
	"avito_internship_PR/internal/pr/application/dto/input"
	"avito_internship_PR/internal/pr/application/transformer"
	"avito_internship_PR/internal/pr/usecase"

	"github.com/gin-gonic/gin"
)

type PullRequestHandler struct {
	createPRUseCase         *usecase.CreatePRUseCase
	mergePRUseCase          *usecase.MergePRUseCase
	reassignReviewerUseCase *usecase.ReassignReviewerUseCase
	prTransformer           *transformer.PullRequestTransformer
}

func NewPullRequestHandler(
	createPRUseCase *usecase.CreatePRUseCase,
	mergePRUseCase *usecase.MergePRUseCase,
	reassignReviewerUseCase *usecase.ReassignReviewerUseCase,
	prTransformer *transformer.PullRequestTransformer,
) *PullRequestHandler {
	return &PullRequestHandler{
		createPRUseCase:         createPRUseCase,
		mergePRUseCase:          mergePRUseCase,
		reassignReviewerUseCase: reassignReviewerUseCase,
		prTransformer:           prTransformer,
	}
}

func (h *PullRequestHandler) Create(c *gin.Context) {
	var req dto.CreatePRRequest
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
	pr, err := h.createPRUseCase.Execute(input.CreatePRInput{PullRequest: dto.PullRequestDto{
		PullRequestID:   req.PullRequestID,
		PullRequestName: req.PullRequestName,
		AuthorID:        req.AuthorID,
	}})
	if err != nil {
		var customErr errPkg.CustomError
		if errors.As(err, &customErr) {
			statusCode := http.StatusBadRequest
			switch customErr.Code {
			case errPkg.ErrPRExists:
				statusCode = http.StatusConflict
			case errPkg.ErrNotFound:
				statusCode = http.StatusNotFound
			}
			c.JSON(statusCode, errPkg.ErrResponse{Error: customErr})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "unknown error"})
		}
		return
	}
	prResponse := h.prTransformer.ToResponse(pr.PR)

	c.JSON(http.StatusCreated, dto.CreatePRResponse{
		PR: prResponse,
	})
}

func (h *PullRequestHandler) Merge(c *gin.Context) {
	var req dto.MergePRRequest
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
	pr, err := h.mergePRUseCase.Execute(input.MergePRInput{PullRequestID: req.PullRequestID})
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
	prResponse := h.prTransformer.ToResponse(pr.PR)
	c.JSON(http.StatusOK, dto.MergePRResponse{
		PR:       prResponse,
		MergedAt: pr.MergedAt,
	})
}

func (h *PullRequestHandler) Reassign(c *gin.Context) {
	var req dto.ReassignReviewerRequest
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
	pr, err := h.reassignReviewerUseCase.Execute(input.ReassignReviewerInput{
		PullRequestID: req.PullRequestID,
		OldUserID:     req.OldReviewerID,
	})
	if err != nil {
		var customErr errPkg.CustomError
		if errors.As(err, &customErr) {
			statusCode := http.StatusBadRequest
			switch customErr.Code {
			case errPkg.ErrNotFound:
				statusCode = http.StatusNotFound
			case errPkg.ErrPRMerged:
				statusCode = http.StatusConflict
			case errPkg.ErrNoCandidate:
				statusCode = http.StatusConflict
			case errPkg.ErrNotAssigned:
				statusCode = http.StatusConflict
			}
			c.JSON(statusCode, errPkg.ErrResponse{Error: customErr})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unknown error"})
	}
	prResponse := h.prTransformer.ToResponse(pr.PR)
	c.JSON(http.StatusOK, dto.ReassignReviewerResponse{
		PR:         prResponse,
		ReplacedBy: pr.ReplacedBy,
	})
}
