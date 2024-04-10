package middlewares

import (
	"fmt"
	"mgo-skeleton/bin/pkg/helpers"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenStr := ctx.GetHeader("Authorization")

		if tokenStr == "" {
			helpers.ErrorHandler(ctx, &helpers.UnauthorizedError{Message: "Unauthorized", MessageDev: "token is empty"})
			ctx.Abort()
			return
		}

		tokenArr := strings.Split(tokenStr, " ")
		if len(tokenArr) != 2 {
			helpers.ErrorHandler(ctx, &helpers.UnauthorizedError{Message: "Unauthorized", MessageDev: "token array split length is not 2"})
			ctx.Abort()
			return
		}

		userId, err := helpers.ValidateToken(tokenArr[1])
		if err != nil {
			helpers.ErrorHandler(ctx, &helpers.UnauthorizedError{Message: err.Error(), MessageDev: "validate token is empty"})
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
