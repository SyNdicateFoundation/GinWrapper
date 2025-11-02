package https_core

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Fn        gin.HandlerFunc
	Method    string
	Addresses []string
	Protected bool
}

var (
	Responses = map[string]Response{
		"not-found-screen": {
			Fn: func(c *gin.Context) {
				c.HTML(http.StatusNotFound, "not-found.html", nil)
			},
			Method: "GET",
		},
	}
)
