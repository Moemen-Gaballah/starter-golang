package Application

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"starter-golang/Models"
)

type Request struct {
	Context    *gin.Context
	DB         *gorm.DB
	Connection *sql.DB
	User       *Models.User
	IsAuth     bool
	IsAdmin    bool
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
	CloseConnection(&req)
	req.Context.JSON(code, body)
}

// init new request closure
func NewRequest(c *gin.Context) *Request {
	request := req()
	req := request(c)
	return &req
}

func (req *Request) Auth() *Request {
	req.IsAuth = false
	req.IsAdmin = false
	authHeader := req.Context.GetHeader("Authorization")
	if authHeader != "" {
		req.DB.Where("token = ? ", authHeader).First(&req.User)

		if req.User.ID != 0 {
			req.IsAuth = true

			if req.User.Group == "admin" {
				req.IsAdmin = true
			}
		}

	}

	return req
}
