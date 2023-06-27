package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/harshitha-git-hub/gin-gorm-rest/controller"
)

func UserRoute(router *gin.Engine) {
	router.GET("/", controller.GetUsers)
	router.POST("/", controller.CreateUser)
	router.DELETE("/:ID", controller.DeleteUser)
	router.PUT("/:ID", controller.UpdateUser)

}
