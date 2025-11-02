package https_core

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	fn        gin.HandlerFunc
	method    string
	addresses []string
	protected bool
}

var (
	Responses = map[string]Response{
		"not-found-screen": {
			fn: func(c *gin.Context) {
				c.HTML(http.StatusNotFound, "not-found.html", nil)
			},
			method: "GET",
		},
	}
)
