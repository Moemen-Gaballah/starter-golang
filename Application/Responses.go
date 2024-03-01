package Application

import (
	"github.com/gin-gonic/gin"
)

// Ok Response
func (req Request) Ok(body interface{}) {
	req.Response(200, body)
}

// Created Response
func (req Request) Created(body interface{}) {
	req.Response(201, body)
}

// NotAuth Response
func (req Request) NotAuth() {
	req.Response(401, gin.H{
		"message": "You are not auth",
	})
}

func (req Request) BadRequest(err interface{}) {
	req.Response(422, err)
}

func (req Request) UserNotFound() {
	req.Response(404, gin.H{
		"message": "We not found this user in our system",
	})
}
