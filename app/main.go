package main

import (
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Todo struct {
	gorm.Model
	Content string `form:"content" binding:"required"`
}

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "test_user"
	PASS := "test_user"
	// コンテナ名:ポート番号を指定する
	PROTOCOL := "tcp(deploypractice-mysql:3306)"
	DBNAME := "deploypractice"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&Todo{})
	return db
}

func dbInsert(content string) {
	db := gormConnect()
	defer db.Close()

	db.Create(&Todo{Content: content})
}

func dbGetAllTodo() []Todo {
	db := gormConnect()
	defer db.Close()

	var todos []Todo

	db.Order("created_at desc").Find(&todos)
	return todos
}

func DeleteTodo(todoId int) {
	var todo []Todo
	db := gormConnect()
	db.Delete(&todo, todoId)
	defer db.Close()
}

func dbFindTodo(todoID int) []Todo {
	var todo []Todo
	db := gormConnect()
	db.First(&todo, todoID)
	defer db.Close()
	return todo
}

func dbUpdate(id int, content string) {
	db := gormConnect()
	var todo Todo
	db.First(&todo, id)
	todo.Content = content
	db.Save(&todo)
	db.Close()
}

func main() {
	r := gin.Default()

	// cors対策
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"},
		AllowMethods:     []string{"POST", "GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// 全件取得
	r.GET("/getAllTodos", func(c *gin.Context) {
		resultTodos := dbGetAllTodo()
		c.JSON(200, resultTodos)
	})

	// 一覧
	r.GET("/", func(c *gin.Context) {
		todos := dbGetAllTodo()
		c.JSON(200, todos)
	})

	// 登録
	r.POST("/new", func(c *gin.Context) {
		content := c.PostForm("content")
		dbInsert(content)
		todos := dbGetAllTodo()
		c.JSON(200, todos)
	})
	// 削除
	r.POST("/deleteTodo", func(c *gin.Context) {
		todoIdStr := c.PostForm("todoId")
		todoID, _ := strconv.Atoi(todoIdStr)
		DeleteTodo(todoID)
	})

	// 編集
	r.POST("/getOneTodo", func(c *gin.Context) {
		todoIdStr := c.PostForm("id")
		todoID, _ := strconv.Atoi(todoIdStr)

		resultTodo := dbFindTodo(todoID)
		c.JSON(200, resultTodo)
	})

	// 更新
	r.POST("/updateTodo", func(c *gin.Context) {
		todoIdStr := c.PostForm("id")
		todoID, _ := strconv.Atoi(todoIdStr)
		content := c.PostForm("content")
		dbUpdate(todoID, content)
		c.Redirect(302, "/")
	})

	r.Run(":3000")
}
