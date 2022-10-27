package controller

import (
	"bubble_todolist/logic"
	"bubble_todolist/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
 * url -> controller -> logic -> model
 */

func IndexHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}

func CreateTask(ctx *gin.Context) {
	// 前段页面填写待办事项 点击提交到这里
	// 1. 从请求中拿数据
	var todo models.Todo
	if err := ctx.BindJSON(&todo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	// 2. 存入数据库&返回响应
	err := logic.CreateTask(&todo)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 2000,
			"msg":  "success",
			"data": todo,
		})
	}
}

func GetTodoList(ctx *gin.Context) {
	// 查看todo表的所有数据
	todoList, err := logic.GetTodoList()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, todoList)
	}
}

func UpdateTask(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{"error": "无效id"})
		return
	}
	todo, err := logic.GetTask(id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}
	if err = ctx.BindJSON(&todo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}
	if err = logic.UpdateTask(todo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 2000,
			"msg":  "success",
			"data": todo,
		})
	}
}

func DeleteTask(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{"error": "无效id"})
		return
	}
	if err := logic.DeleteTask(id); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 2000,
			"msg":  "success",
		})
	}
}
