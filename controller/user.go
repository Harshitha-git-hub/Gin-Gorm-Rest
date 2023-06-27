package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/harshitha-git-hub/gin-gorm-rest/config"
	"github.com/harshitha-git-hub/gin-gorm-rest/models"
)

func GetUsers(c *gin.Context) {
	users := []models.User{}
	config.DB.Find(&users)
	c.JSON(200, &users)
}

func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	config.DB.Create(&user)
	c.JSON(200, &user)
}

func DeleteUser(c *gin.Context) {
	var user models.User
	config.DB.Where("ID = ?", c.Param("ID")).Delete(&user)
	c.JSON(200, &user)
}

func UpdateUser(c *gin.Context) {
	var user models.User
	config.DB.Where("ID = ?", c.Param("ID")).First(&user)
	c.BindJSON(&user)
	config.DB.Save(&user)
	c.JSON(200, &user)
}
