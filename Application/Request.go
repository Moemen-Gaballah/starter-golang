package App

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Request struct {
	Context    *gin.Context
	DB         *gorm.DB
	Connection *sql.DB
}

type ShareResources interface {
	Share()
}

func (req *Request) Share() {

}

// handel request closure data
func req() func(c *gin.Context) Request {
	return func(c *gin.Context) Request {
		var request Request
		request.Context = c
		connectToDatabase(&request)
		return request
	}
}

// Response response
func (req Request) Response(code int, body interface{}) {
	closeConnection(&req)
	req.Context.JSON(code, body)
}

// init new request closure
func NewRequest(c *gin.Context) Request {
	request := req()
	req := request(c)
	return req
}
