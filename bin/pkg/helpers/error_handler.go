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
		fmt.Println("\n")
		log.Println("--------Error Handler-------------")
		log.Printf("NotFoundError, Message Developer is %s", err.Error())
		log.Println("---------------------\n")

	case *BadRequestError:
		statusCode = http.StatusBadRequest
		fmt.Println("\n")
		fmt.Println("--------Error Handler-------------")
		log.Printf("BadRequestError, Message Developer is %s", err.Error())
		fmt.Println("---------------------\n")

	case *InternalServerError:
		statusCode = http.StatusInternalServerError
		fmt.Println("\n")
		log.Println("--------Error Handler-------------")
		log.Printf("InternalServerError, Message Developer is %s", err.Error())
		log.Println("---------------------\n")

	case *UnauthorizedError:
		statusCode = http.StatusUnauthorized
		fmt.Println("\n")
		log.Println("--------Error Handler-------------")
		log.Printf("UnauthorizedError, Message Developer is %s", err.Error())
		log.Println("---------------------\n")

	}

	response := Response(ResponseParams{StatusCode: statusCode, Message: err.Error()})

	c.JSON(statusCode, response)
}
