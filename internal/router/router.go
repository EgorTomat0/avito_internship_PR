package router

import (
	"avito_internship_PR/internal/config"
	prTransformer "avito_internship_PR/internal/pr/application/transformer"
	prHandlers "avito_internship_PR/internal/pr/infrastructure/controllers/http"
	prRepo "avito_internship_PR/internal/pr/infrastructure/repository"
	prUseCase "avito_internship_PR/internal/pr/usecase"
	teamTransformer "avito_internship_PR/internal/team/application/transformer"
	teamHandlers "avito_internship_PR/internal/team/infrastructure/controllers/http"
	teamRepo "avito_internship_PR/internal/team/infrastructure/repository"
	teamUseCase "avito_internship_PR/internal/team/usecase"
	userTransformer "avito_internship_PR/internal/user/application/transformer"
	userHandlers "avito_internship_PR/internal/user/infrastructure/controllers/http"
	userRepo "avito_internship_PR/internal/user/infrastructure/repository"
	userUseCase "avito_internship_PR/internal/user/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB, cfg *config.Config) *gin.Engine {
	r := gin.Default()

	// Repositories
	userRepository := userRepo.NewUserRepository(db)
	teamRepository := teamRepo.NewTeamRepository(db, userRepository)
	prRepository := prRepo.NewPullRequestRepository(db)

	// Transformers
	userTransformerInstance := userTransformer.NewUserTransformer()
	prTransformerForUser := userTransformer.NewPullRequestTransformer()
	teamTransformerInstance := teamTransformer.NewTeamTransformer()
	prTransformerInstance := prTransformer.NewPullRequestTransformer()

	// UseCases
	setIsActiveUseCase := userUseCase.NewSetIsActiveUseCase(userRepository)
	getReviewsUseCase := userUseCase.NewGetReviewsUseCase(prRepository, *prTransformerForUser)
	createTeamUseCase := teamUseCase.NewCreateTeamUseCase(teamRepository)
	getTeamUseCase := teamUseCase.NewGetTeamUseCase(teamRepository)
	createPRUseCase := prUseCase.NewCreatePRUseCase(prRepository)
	mergePRUseCase := prUseCase.NewMergePRUseCase(prRepository)
	reassignReviewerUseCase := prUseCase.NewReassignReviewerUseCase(prRepository)

	// Handlers
	userHandler := userHandlers.NewUserHandler(setIsActiveUseCase, getReviewsUseCase, userTransformerInstance)
	teamHandler := teamHandlers.NewTeamHandler(createTeamUseCase, getTeamUseCase, teamTransformerInstance)
	prHandler := prHandlers.NewPullRequestHandler(createPRUseCase, mergePRUseCase, reassignReviewerUseCase, prTransformerInstance)

	// Routes
	// Teams
	r.POST("/team/add", teamHandler.CreateTeam)
	r.GET("/team/get", teamHandler.GetTeam)

	// Users
	r.POST("/users/setIsActive", userHandler.SetIsActive)
	r.GET("/users/getReview", userHandler.GetReviews)

	// Pull Requests
	r.POST("/pullRequest/create", prHandler.Create)
	r.POST("/pullRequest/merge", prHandler.Merge)
	r.POST("/pullRequest/reassign", prHandler.Reassign)

	return r
}
