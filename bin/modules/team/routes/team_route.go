package routes

import (
	"mgo-skeleton/bin/middlewares"
	"mgo-skeleton/bin/modules/team/handlers"
	"mgo-skeleton/bin/modules/team/repositories"
	"mgo-skeleton/bin/modules/team/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TeamRoute(r *gin.RouterGroup, db *gorm.DB) {
	teamRepository := repositories.NewTeamRepository(db)
	teamService := services.NewTeamService(teamRepository)
	teamHandler := handlers.NewTeamHandler(teamService)

	r.Use(middlewares.JWTMiddleware())

	r.POST("/team", teamHandler.Create)
	r.GET("/:id/team", teamHandler.Detail)
	r.DELETE("/:id/team", teamHandler.Delete)
}
