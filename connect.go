package main

import (
	"fmt"
	"time"
  	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
  	// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
  	"gorm.io/gorm"
)

func main(){
	// github.com/mattn/go-sqlite3
	db, err := gorm.Open(sqlite.Open("PATH"), &gorm.Config{})

	user := User{ID:x, filepath: "http:ほにゃらら", tags: "Jinzhu",  Created_at: time.Now()}

	result := db.Create(&user) // pass pointer of data to Create

	user.ID             // returns inserted data's primary key
	result.Error        // returns error
	result.RowsAffected // returns inserted records count
}