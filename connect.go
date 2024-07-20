package main

import (
	"fmt"
	"time"
  	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
  	// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
  	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
	"net/http"
)

//構造体Userの宣言
type User struct{
	ID	uint	`json:"ID"`
	title	string	`json:"title"`
	description string	`json:"description"`
	filepath	string	`json:"filepath"`
	tags	string	`json:"tags"`
	created_at	time.Time	`json:"created_at"`
}

var db *gorm.DB

func main(){
	var err error
	// github.com/mattn/go-sqlite3
	//SQLiteを開く　※PATHはあとで設定
	db, err := gorm.Open(sqlite.Open("PATH"), &gorm.Config{})
	//つながらないとエラー返す
	if err != nil {
		panic("failed to connect to database")
	}

	//POSTを受け取る
	router:=gin.Default()
	router.POST("/somePost",posting)
	router.Run(":8080")
}

func posting(c*gin.Context){
	var user User
	//構造に合わなければエラー返す
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	//挿入
	result := db.Create(&user) // pass pointer of data to Create
	//できなければエラー返す
	if result.Error != nil{
		c.JSON(500,gin.H{"error":result.Error.Error()})
		return
	}
	//成功した際に送信
	c.JSON(200,gin.H{"message":"success"})
}
