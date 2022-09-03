package controller

import (
	"my/app/config"
	"my/app/models"

	"github.com/gin-gonic/gin"
)

func HomeController(c *gin.Context) {
	todo := []models.Todo{}
	config.DB.Find(&todo)

	c.JSON(200, &todo)
}

func CreateTodo(c *gin.Context) {
	var todo models.Todo
	c.BindJSON(&todo)
	config.DB.Create(&todo)
	c.JSON(200, &todo)

}
func DeleteTodo(c *gin.Context) {
	var todo models.Todo

	config.DB.Where("id=?", c.Param("id")).Delete(&todo)

	c.JSON(200, &todo)

}

func UpdateTodo(c *gin.Context) {
	var todo models.Todo

	config.DB.Where("id=?", c.Param("id")).First(&todo)
	c.BindJSON(&todo)
	config.DB.Save(&todo)

	c.JSON(200, &todo)

}
func ReadTodo(c *gin.Context) {
	var todo models.Todo

	config.DB.Where("id=?", c.Param("id")).First(&todo)

	c.JSON(200, &todo)

}
