package handlers

import (
	"mgo-skeleton/bin/modules/team/models"
	"mgo-skeleton/bin/modules/team/services"
	"mgo-skeleton/bin/pkg/helpers"
	"net/http"
	"strconv"

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

func (h *teamHandler) Detail(c *gin.Context) {
	idStr := c.Param("id")
	idInt, _ := strconv.Atoi(idStr)
	response, err := h.services.Detail(idInt)

	if err != nil {
		helpers.ErrorHandler(c, err)
		return
	}

	res := helpers.Response(helpers.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "success find detail user",
		Data:       response,
	})

	c.JSON(http.StatusOK, res)
}

func (h *teamHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	idInt, _ := strconv.Atoi(idStr)

	userId, _ := c.Get("userId")
	userIdInt := userId.(int)

	if err := h.services.Delete(idInt, userIdInt); err != nil {
		helpers.ErrorHandler(c, err)
		return
	}

	res := helpers.Response(helpers.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "success delete user team",
	})

	c.JSON(http.StatusOK, res)
}
