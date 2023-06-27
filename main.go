package main

import (
	"github.com/gin-gonic/gin"
	"github.com/harshitha-git-hub/gin-gorm-rest/config"
	"github.com/harshitha-git-hub/gin-gorm-rest/routes"
)

func main() {
	router := gin.New()

	config.Connect()

	routes.UserRoute(router)

	router.Run(":8040")
}
