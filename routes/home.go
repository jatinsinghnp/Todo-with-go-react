package routes

import (
	"my/app/controller"

	"github.com/gin-gonic/gin"
)

func TodoRoute(router *gin.Engine) {
	router.GET("", controller.HomeController)
	router.POST("create", controller.CreateTodo)
	router.GET("getTodo/:id", controller.ReadTodo)
	router.PATCH("updateTodo/:id", controller.UpdateTodo)
	router.DELETE("deleteTodo/:id", controller.DeleteTodo)
}
