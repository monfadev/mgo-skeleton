package routes

import (
	"mgo-skeleton/bin/modules/auth/handlers"
	"mgo-skeleton/bin/modules/auth/repositories"
	"mgo-skeleton/bin/modules/auth/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthRoute(r *gin.RouterGroup, db *gorm.DB) {
	authRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthServices(authRepository)
	authHandler := handlers.NewAuthHandler(authService)

	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)

}
