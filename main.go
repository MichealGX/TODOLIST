package main

import (
	"TODOLIST/initx"
	"fmt"
)

func main() {
	fmt.Println("project init")
	// 创建一个默认的Gin引擎
	initx.Init()
}
