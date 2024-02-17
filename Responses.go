package main

import "github.com/gin-gonic/gin"

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
