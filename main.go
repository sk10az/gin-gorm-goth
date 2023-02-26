package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sk10az/gin-gorm-goth/routes"
)

var router = gin.Default()

func getRoutes() {
	routes.AuthRoutes(router)
}

func main() {
	getRoutes()
	router.Run(":3000")
}
