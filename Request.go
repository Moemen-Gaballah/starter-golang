package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Request struct {
	Context    *gin.Context
	DB         *gorm.DB
	Connection *sql.DB
}

// handel request closure data
func req() func(c *gin.Context) Request {
	return func(c *gin.Context) Request {
		var request Request
		request.Context = c
		request.DB, request.Connection = connectToDatabase()
		return request
	}
}

// Response response
func (req Request) Response(code int, body interface{}) {
	req.closeConnection()
	req.Context.JSON(code, body)
}

// init new request closure
func newRequest(c *gin.Context) Request {
	request := req()
	req := request(c)
	return req
}

// connect to database
func connectToDatabase() (*gorm.DB, *sql.DB) {
	dsn := "root@tcp(127.0.0.1:3306)/golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connect to database", err.Error())
	}

	connection, err := db.DB()
	if err != nil {
		fmt.Println("Error on get database connection", err.Error())
	}

	return db, connection
}

// close database connection
func (req Request) closeConnection() {
	req.Connection.Close()
}
