package main

import (
	"sample/api/controller"

	"github.com/gin-gonic/gin"
)

func route(app *gin.Engine) {
	app.GET("/ping", controller.Test)
	app.POST("/addContact", controller.Create)
	app.GET("/getAll", controller.Get)
}
