package router

import (
	"TODOLIST/controllers"
	"github.com/gin-gonic/gin"
)

// SetupRouter 设置路由
func SetupRouter() *gin.Engine {
	// 创建一个默认的 Gin 引擎
	r := gin.Default()

	// 定义路由
	r.POST("/todolist/manager/addItem", controllers.AddItem)
	r.DELETE("/todolist/{id}/manager/deleteItem", controllers.DeleteItem)
	r.PUT("/todolist/manager/modifyItem", controllers.ModifyItem)
	r.POST("/todolist/manager/searchItem", controllers.SearchItem)
	r.PUT("/todolist/manager/outputItems", controllers.OutputItem)

	// 返回路由引擎
	return r
}
