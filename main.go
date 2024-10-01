package main

import (
	"sample/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.LoadEnv()
	utils.ConnectDB()
	app := gin.Default()
	route(app)
	app.Run()
}
