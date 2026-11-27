package httpserver

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	authApp "github.com/toby-minigames/server/internal/modules/auth/application"
	authHandler "github.com/toby-minigames/server/internal/modules/auth/handler"
	gamesApp "github.com/toby-minigames/server/internal/modules/games/application"
	gamesHandler "github.com/toby-minigames/server/internal/modules/games/handler"
	gamesInfra "github.com/toby-minigames/server/internal/modules/games/infrastructure"
	usersApp "github.com/toby-minigames/server/internal/modules/users/application"
	usersHandler "github.com/toby-minigames/server/internal/modules/users/handler"
	usersInfra "github.com/toby-minigames/server/internal/modules/users/infrastructure"
	"github.com/toby-minigames/server/internal/shared/database"
	appErrors "github.com/toby-minigames/server/internal/shared/errors"
	"github.com/toby-minigames/server/internal/shared/middleware"
	"github.com/toby-minigames/server/internal/shared/response"
	"github.com/toby-minigames/server/internal/shared/security"
	"gorm.io/gorm"
)

// NewRouter construye y configura el motor HTTP Gin con todas las rutas /api/v1.
func NewRouter(db *gorm.DB, jwtMgr *security.JWTManager, environment string) *gin.Engine {
	if environment != "development" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery(), middleware.CORS())

	router.GET("/health", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
		defer cancel()
		if err := database.Ping(ctx, db); err != nil {
			response.Error(c, appErrors.Internal("database healthcheck failed"))
			return
		}
		response.Success(c, http.StatusOK, gin.H{"status": "ok"}, "Service is healthy")
	})

	playerRepo := usersInfra.NewPostgresPlayerRepository(db)
	profileRepo := usersInfra.NewPostgresGameProfileRepository(db)
	gameRepo := gamesInfra.NewPostgresGameRepository(db)
	categoryRepo := gamesInfra.NewPostgresCategoryRepository(db)

	authSvc := authApp.NewAuthService(playerRepo, jwtMgr)
	userSvc := usersApp.NewUserService(playerRepo, profileRepo, gameRepo)
	gameSvc := gamesApp.NewGameService(gameRepo, categoryRepo)

	v1 := router.Group("/api/v1")
	authHandler.NewAuthHandler(authSvc).RegisterRoutes(v1)
	usersHandler.NewUserHandler(userSvc).RegisterRoutes(v1, middleware.AuthRequired(jwtMgr), middleware.PlayerOnly())
	gamesHandler.NewGameHandler(gameSvc).RegisterRoutes(v1)

	return router
}
