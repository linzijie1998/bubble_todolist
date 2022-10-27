package main

import (
	"bubble_todolist/dao"
	"bubble_todolist/models"
	"bubble_todolist/routers"
)

func main() {
	// 创建数据库: CREATE DATABASE bubble;
	// 连接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	// 程序退出关闭数据库
	defer dao.Close()

	// 模型绑定
	dao.DB.AutoMigrate(&models.Todo{})

	// Router
	router := routers.SetupRouter()
	if err = router.Run(":8080"); err != nil {
		return
	}
}
