package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

var (
	DB *gorm.DB
)

// Todo Model
type Todo struct {
	gorm.Model
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func initMySQL() (err error) {
	dsn := "root:123456@tcp(localhost:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}

func main() {
	// 创建数据库: CREATE DATABASE bubble;
	// 连接数据库
	err := initMySQL()
	if err != nil {
		panic(err)
	}
	// 程序退出关闭数据库
	defer func(DB *gorm.DB) {
		err := DB.Close()
		if err != nil {
			panic(err)
		}
	}(DB)

	// 模型绑定
	DB.AutoMigrate(&Todo{})

	router := gin.Default()
	router.Static("/static", "static")
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	v1Group := router.Group("v1")
	{
		// 添加
		v1Group.POST("/todo", func(ctx *gin.Context) {
			// 前段页面填写待办事项 点击提交到这里
			// 1. 从请求中拿数据
			var todo Todo
			if err = ctx.BindJSON(&todo); err != nil {
				ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
				return
			}
			// 2. 存入数据库&返回响应
			if err = DB.Create(&todo).Error; err != nil {
				ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"code": 2000,
					"msg":  "success",
					"data": todo,
				})
			}
		})
		// 查看所有待办事项
		v1Group.GET("/todo", func(ctx *gin.Context) {
			// 查看todo表的所有数据
			var todoList []Todo
			if err = DB.Find(&todoList).Error; err != nil {
				ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, todoList)
			}
		})
		// 查看某个待办事项
		v1Group.GET("/todo/:id", func(ctx *gin.Context) {

		})
		// 修改
		v1Group.PUT("/todo/:id", func(ctx *gin.Context) {
			id, ok := ctx.Params.Get("id")
			if !ok {
				ctx.JSON(http.StatusOK, gin.H{"error": "无效id"})
				return
			}
			var todo Todo
			if err = DB.Where("id=?", id).First(&todo).Error; err != nil {
				ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
			}
			if err = ctx.BindJSON(&todo); err != nil {
				ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
			}
			if err = DB.Save(&todo).Error; err != nil {
				ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"code": 2000,
					"msg":  "success",
					"data": todo,
				})
			}
		})
		// 删除
		v1Group.DELETE("/todo/:id", func(ctx *gin.Context) {
			id, ok := ctx.Params.Get("id")
			if !ok {
				ctx.JSON(http.StatusOK, gin.H{"error": "无效id"})
				return
			}
			if err = DB.Where("id=?", id).Delete(&Todo{}).Error; err != nil {
				ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"code": 2000,
					"msg":  "success",
				})
			}
		})
	}
	if err = router.Run(":8080"); err != nil {
		return
	}
}
