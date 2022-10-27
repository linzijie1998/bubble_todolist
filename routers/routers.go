package routers

import (
	"bubble_todolist/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() (router *gin.Engine) {
	router = gin.Default()
	router.Static("/static", "static")
	router.LoadHTMLGlob("templates/*")

	router.GET("/", controller.IndexHandler)

	v1Group := router.Group("v1")
	{
		// 添加
		v1Group.POST("/todo", controller.CreateTask)
		// 查看所有待办事项
		v1Group.GET("/todo", controller.GetTodoList)
		// 修改
		v1Group.PUT("/todo/:id", controller.UpdateTask)
		// 删除
		v1Group.DELETE("/todo/:id", controller.DeleteTask)
	}
	return
}
