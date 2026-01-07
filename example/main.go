package main

import (
	"example/handlers"

	"github.com/dmtruong4697/errx/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(middleware.ErrorHandler())

	r.GET("/users/:id", handlers.GetUser)

	r.Run(":8080")
}
