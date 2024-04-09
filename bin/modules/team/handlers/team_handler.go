package handlers

import (
	"mgo-skeleton/bin/modules/team/models"
	"mgo-skeleton/bin/modules/team/services"
	"mgo-skeleton/bin/pkg/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type teamHandler struct {
	services services.TeamService
}

func NewTeamHandler(s services.TeamService) *teamHandler {
	return &teamHandler{
		services: s,
	}
}

func (h *teamHandler) Create(c *gin.Context) {
	var user models.TeamRequest

	if err := c.ShouldBindJSON(&user); err != nil {
		helpers.ErrorHandler(c, &helpers.BadRequestError{Message: err.Error()})
	}

	userID, _ := c.Get("userId")
	user.UserId = userID.(int)

	if err := h.services.Create(&user); err != nil {
		helpers.ErrorHandler(c, err)
		return
	}

	res := helpers.Response(helpers.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "create user successfully",
	})

	c.JSON(http.StatusCreated, res)

}
