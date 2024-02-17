package main

import (
	"github.com/gin-gonic/gin"
	"starter-golang/Application"
)

// Ok Response
func (req Application.Request) Ok(body interface{}) {
	req.Response(200, body)
}

// Created Response
func (req Application.Request) Created(body interface{}) {
	req.Response(201, body)
}

// NotAuth Response
func (req Application.Request) NotAuth() {
	req.Response(401, gin.H{
		"message": "You are not auth",
	})
}
