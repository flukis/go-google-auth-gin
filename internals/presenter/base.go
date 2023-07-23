package presenter

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseHandler struct {
}

func NewBaseHandler() *BaseHandler {
	return &BaseHandler{}
}

func (h *BaseHandler) Route(r *gin.RouterGroup) {
	r.GET("/", h.LoginPage)
}

func (h *BaseHandler) LoginPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", gin.H{
		"title": "Login With Google",
	})
}
