package Auth

import (
	"github.com/bykovme/gotrans"
	"github.com/gin-gonic/gin"
	"starter-golang/Application"
)

func CreateUser(c *gin.Context) {
	r, auth := Application.AuthRequest(c)
	if !auth {
		return
	}
	r.Ok(gin.H{
		"message": gotrans.T("hello_world"),
		"Lang":    r.Lang,
	})
}
