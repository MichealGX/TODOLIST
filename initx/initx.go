package initx

import (
	"TODOLIST/controllers"
	"TODOLIST/router"
	"log"
)

func Init() {
	//链接数据库
	controllers.DatabaseLink()
	defer controllers.Db.Close()

	// 设置路由
	r := router.SetupRouter()

	// 启动服务器
	if err := r.Run(":3000"); err != nil {
		log.Fatal("Unable to start server:", err)
	}
}
