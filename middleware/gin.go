package middleware

import (
	"net/http"

	"github.com/dmtruong4697/errx"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		lastErr := c.Errors.Last()
		if lastErr == nil {
			return
		}

		e, ok := lastErr.Err.(*errx.Error)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
			})
			return
		}

		c.JSON(e.HTTPStatus, errx.ToResponse(e))
	}
}
