package handlers

import (
	"mgo-skeleton/bin/modules/auth/models"
	"mgo-skeleton/bin/modules/auth/services"
	"mgo-skeleton/bin/pkg/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authHandler struct {
	services services.AuthServices
}

func NewAuthHandler(s services.AuthServices) *authHandler {
	return &authHandler{
		services: s,
	}
}

func (h *authHandler) Register(c *gin.Context) {
	var register models.RegisterRequest

	if err := c.ShouldBindJSON(&register); err != nil {
		helpers.ErrorHandler(c, &helpers.BadRequestError{Message: err.Error()})
		return
	}

	if err := h.services.Register(&register); err != nil {
		helpers.ErrorHandler(c, err)
		return
	}

	res := helpers.Response(helpers.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "Register successfully",
	})

	c.JSON(http.StatusCreated, res)

}

func (h *authHandler) Login(c *gin.Context) {
	var login models.LoginRequest

	if err := c.ShouldBindJSON(&login); err != nil {
		helpers.ErrorHandler(c, &helpers.BadRequestError{Message: err.Error()})
		return
	}

	result, err := h.services.Login(&login)

	if err != nil {
		helpers.ErrorHandler(c, err)
		return
	}

	res := helpers.Response(helpers.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "login successfully",
		Data:       result,
	})

	c.JSON(http.StatusOK, res)

}
