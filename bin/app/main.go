package main

import (
	"mgo-skeleton/bin/configs"
	"mgo-skeleton/bin/modules/auth/routes"
	"mgo-skeleton/bin/pkg/database/postgres"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	configs.InitEnvironments()

	db := postgres.ConnectPostgresql()

	configCors := cors.DefaultConfig()
	configCors.AllowAllOrigins = true
	configCors.AllowHeaders = append(configCors.AllowHeaders, "Content-Type", "Content-Length", "Accept-Encoding", "X-XSRF-TOKEN", "X-CSRF-Token", "Authorization", "X-M2M-Origin", "Access-Control-Allow-Origin", "Access-Control-Allow-Methods", "Access-Control-Allow-Headers", "Access-Control-Allow-Credentials", "Origin", "Accept", "X-Requested-With", "access-control-allow-origin", "access-control-allow-methods", "access-control-allow-headers")
	configCors.AllowMethods = append(configCors.AllowMethods, "Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
	configCors.AllowCredentials = true

	r := gin.Default()

	r.Use(cors.New(configCors))

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	routerGroup := r.Group("/api")

	routes.AuthRoute(routerGroup, db)

	r.Run(":8080")
}
