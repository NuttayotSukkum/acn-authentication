package rest

import (
	"github.com/NuttayotSukkum/acn/acn-authentication/internal/handlers"
	"github.com/NuttayotSukkum/acn/acn-authentication/internal/pkg/middleware"
	"github.com/NuttayotSukkum/acn/acn-authentication/internal/repositories"
	"github.com/NuttayotSukkum/acn/acn-authentication/internal/services"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"log"
)

func UserInitRouter(e *echo.Echo, db *gorm.DB) {
	repo := repositories.NewMessageRepository(db)
	svc := services.NewUserService(repo)
	userHandler := handlers.NewUserHandler(svc)
	api := e.Group("/api")
	userRegister := api.Group("/authentication")
	userVerify := api.Group("/verification", middleware.ValidateTokenMiddleware)
	v1 := userRegister.Group("/v1")
	v1UserVerify := userVerify.Group("/v1")
	log.Println("v1:", v1)
	v1.POST("/register", userHandler.RegisterUser)
	v1.POST("/login", userHandler.LoginHandler)
	v1UserVerify.GET("/:id", userHandler.VerifyUserHandler)
}
