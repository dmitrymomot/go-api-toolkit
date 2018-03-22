package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	uuid "github.com/satori/go.uuid"
)

var config Configer

// Setup the package
func Setup(c Configer) {
	config = c
}

// Resp function builds successful JSON response
func Resp(c *gin.Context, r Responder) {
	c.JSON(http.StatusOK, r)
}

// Err sends error as JSON response
func Err(c *gin.Context, code int, msg interface{}, prevErr error) {
	r := &Response{}
	if code != 0 {
		r.HTTPStatus = code
	} else {
		r.HTTPStatus = http.StatusInternalServerError
	}
	err := Error{
		ID:        uuid.Must(uuid.NewV1()).String(),
		Code:      code,
		Title:     http.StatusText(code),
		Detail:    msg,
		PrevError: prevErr,
	}
	r.AddError(err)
	if prevErr != nil {
		log.Println(err)
	}
	c.AbortWithStatusJSON(code, r)
}
