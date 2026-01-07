package handlers

import (
	"strconv"

	"example/errors"

	"github.com/dmtruong4697/errx"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	if id != 1 {
		c.Error(
			errx.Wrap(errors.ErrUserNotFound).
				WithMeta("user_id", id),
		)
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"data": gin.H{
			"id":   1,
			"name": "John Cena",
		},
	})
}
