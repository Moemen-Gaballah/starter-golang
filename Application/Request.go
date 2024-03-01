package Application

import (
	"database/sql"
	"github.com/bykovme/gotrans"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"gorm.io/gorm"
	"starter-golang/Models"
)

type Request struct {
	Context         *gin.Context
	DB              *gorm.DB
	Connection      *sql.DB
	User            *Models.User
	IsAuth          bool
	IsAdmin         bool
	Lang            string
	ValidationError error
}

type ShareResources interface {
	Share()
}

func (req *Request) Share() {

}

// handel request closure data
func req() func(c *gin.Context) *Request {
	return func(c *gin.Context) *Request {
		var request Request
		request.Context = c
		connectToDatabase(&request)
		setLang(&request)
		return &request
	}
}

func (req *Request) ValidateRequest(errors validation.Errors) *Request {
	req.ValidationError = errors.Filter()
	return req
}

func (req *Request) Fails() bool {
	if req.ValidationError != nil {
		req.BadRequest(req.ValidationError)
		return true
	}
	return false
}

func setLang(req *Request) {
	lang := gotrans.DetectLanguage(req.Context.GetHeader("Accept-Language"))
	gotrans.SetDefaultLocale(lang) // Setting default locale
	req.Lang = lang
}

// Response response
func (req Request) Response(code int, body map[string]interface{}) {
	CloseConnection(&req)
	req.Context.JSON(code, body)
}

// init new request closure
func NewRequest(c *gin.Context) *Request {
	request := req()
	return request(c)
}

func NewRequestWithAuth(c *gin.Context) *Request {
	return NewRequest(c).Auth()
}

func AuthRequest(c *gin.Context) (*Request, bool) {
	r := NewRequestWithAuth(c)
	if !r.IsAuth {
		r.NotAuth()
		return r, false
	}
	return r, true
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
