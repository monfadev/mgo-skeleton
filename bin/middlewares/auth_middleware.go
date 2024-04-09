package middlewares

import (
	"fmt"
	"mgo-skeleton/bin/pkg/helpers"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader("Authorization")

		if !strings.Contains(authorizationHeader, "Bearer") {
			helpers.ErrorHandler(ctx, &helpers.UnauthorizedError{Message: "Unauthorized"})
			return
		}

		tokenStr := strings.Replace(authorizationHeader, "Bearer ", "", -1)

		if tokenStr == "" {
			helpers.ErrorHandler(ctx, &helpers.UnauthorizedError{Message: "Unauthorized"})
			ctx.Abort()
			return
		}

		userId, err := helpers.ValidateToken(tokenStr)
		if err != nil {
			helpers.ErrorHandler(ctx, &helpers.UnauthorizedError{Message: err.Error()})
			ctx.Abort()
			return
		}

		fmt.Println("\n\n\n")
		fmt.Printf("userId without pointer is %v", userId)
		fmt.Println("\n")
		fmt.Printf("userId with pointer is %v", *userId)
		fmt.Println("\n\n\n")

		ctx.Set("userId", *userId)
		ctx.Next()
	}

}
