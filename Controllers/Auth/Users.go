package Auth

import (
	"github.com/gin-gonic/gin"
	"starter-golang/Application"
)

func Me(c *gin.Context) {
	r, auth := Application.AuthRequest(c)
	if !auth {
		return
	}
	r.Ok(r.User.Transform())
}
