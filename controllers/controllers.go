package controllers

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" // 导入 MySQL 驱动程序包，但不直接使用它，因此使用下划线标识符将其命名为匿名导入
	"github.com/google/uuid"
	"log"
	"net/http"
)

var Db *sql.DB

func DatabaseLink() {
	// 连接到 MySQL 数据库
	var err error
	Db, err = sql.Open("mysql", "root:123456@tcp(localhost:3306)/TodolistMysql")
	if err != nil {
		log.Fatal("Could not connect to database:", err)
	}
}

type Todo struct {
	ID      string
	Content string `json:"content"`
}

// todoRequest  dto
type todoRequest struct {
	Content string `json:"content"`
}

// AddItem 添加一条记录
func AddItem(c *gin.Context) {
	// 实现添加记录的逻辑
	var todo Todo
	var err error
	db := Db
	if err = c.BindJSON(&todo); err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  "Invalid request payload",
			"id":   nil,
		})
		return
	}

	// 生成唯一的 UUID
	todo.ID, err = generateUniqueUUID(db)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("INSERT INTO TodolistTable (id, content) VALUES (?, ?)", todo.ID, todo.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 1,
			"msg":  "Failed to add todo",
			"id":   nil,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"code": 0,
		"msg":  "success",
		"id":   todo.ID,
	})
}

// DeleteItem 删除记录
func DeleteItem(c *gin.Context) {
	// 实现删除记录的逻辑
}

// ModifyItem 修改记录
func ModifyItem(c *gin.Context) {
	// 实现修改记录的逻辑
}

// SearchItem 查找记录
func SearchItem(c *gin.Context) {
	// 实现查找记录的逻辑
}

// OutputItem 输出全部记录
func OutputItem(c *gin.Context) {
	// 实现输出全部记录的逻辑
}

// generateUniqueUUID 生成唯一的 UUID
func generateUniqueUUID(db *sql.DB) (string, error) {
	for {
		id := uuid.New().String()

		// 检查数据库中是否已存在相同的 UUID
		err := db.QueryRow("SELECT id FROM TodolistTable WHERE id = ?", id).Scan(&id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				// UUID 不存在，可以使用
				return id, nil
			}
			// 其他数据库查询错误，返回错误
			return "", err
		}
		// 如果已存在相同的 UUID，重新生成
	}
}
