package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Age  int
}

func main() {
	// step1. connect to the mysql
	dsn := "root:theless14more@tcp(127.0.0.1:3306)/gro-up?charset=utf8mb4&parseTime=True&loc=Local"
	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		panic("failed to connect database!")
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err == nil {
		fmt.Println(err)
	}
	// step2. define a struct to reflect the table structure
	// step3. generate a table
	_ = db.AutoMigrate(&User{})
}
