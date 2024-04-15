package helpers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context, err error) {
	var statusCode int

	switch err.(type) {
	case *NotFoundError:
		statusCode = http.StatusNotFound

		if notFoundErr, ok := err.(*NotFoundError); ok {
			fmt.Println("\n")
			log.Println("--------Error Handler-------------")
			log.Printf("NotFoundError, Message is %s and Developer %s", err.Error(), notFoundErr.MessageDev)
			log.Println("---------------------\n")
		}

	case *BadRequestError:
		statusCode = http.StatusBadRequest

		if badRequestErr, ok := err.(*BadRequestError); ok {
			fmt.Println("\n")
			fmt.Println("--------Error Handler-------------")
			log.Printf("BadRequestError, Message is %s and Developer %s", err.Error(), badRequestErr.MessageDev)
			fmt.Println("---------------------\n")
		}

	case *InternalServerError:
		statusCode = http.StatusInternalServerError

		if internalServerErr, ok := err.(*BadRequestError); ok {
			fmt.Println("\n")
			log.Println("--------Error Handler-------------")
			log.Printf("InternalServerError, Message is %s and Developer", err.Error(), internalServerErr.MessageDev)
			log.Println("---------------------\n")
		}

	case *UnauthorizedError:
		statusCode = http.StatusUnauthorized

		if UnauthorizedErr, ok := err.(*BadRequestError); ok {

			fmt.Println("\n")
			log.Println("--------Error Handler-------------")
			log.Printf("UnauthorizedError, Message is %s and Developer", err.Error(), UnauthorizedErr.MessageDev)
			log.Println("---------------------\n")
		}

	}

	response := Response(ResponseParams{StatusCode: statusCode, Message: err.Error()})

	c.JSON(statusCode, response)
}
