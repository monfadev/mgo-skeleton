package helpers

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ResponseFilterParams(c *gin.Context) *FilterParams {
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "20")
	search := c.Query("search")

	pageNumber, _ := strconv.Atoi(page)
	limitNumber, _ := strconv.Atoi(limit)
	offset := (pageNumber - 1) * limitNumber

	fmt.Println("ResponseFilterParams")
	fmt.Println("page", page)
	fmt.Println("limit", limit)
	fmt.Println("search", search)
	fmt.Println("offset embedded", offset)

	return &FilterParams{
		Page:   pageNumber,
		Limit:  limitNumber,
		Offset: offset,
		Search: search,
	}
}
