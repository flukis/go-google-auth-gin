package presenter

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseHandler struct {
}

func NewBaseHandler() *BaseHandler {
	return &BaseHandler{}
}

func (h *BaseHandler) Route(r *gin.RouterGroup) {
	r.GET("/", h.HomePage)
}

// type Books struct {
// 	Title string `json:"title"`
// 	Link  string `json:"link"`
// }

func (h *BaseHandler) HomePage(ctx *gin.Context) {
	_, err := ctx.Cookie("oauthstate")
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			ctx.HTML(http.StatusOK, "login.html", gin.H{
				"title": "Login With Google",
			})
		default:
			ctx.Status(http.StatusInternalServerError)
			ctx.Abort()
			return
		}
		return
	}
	// var b = []Books{
	// 	{Title: "Catatan 1", Link: "/1"},
	// 	{Title: "Catatan 2", Link: "/2"},
	// }
	// ctx.HTML(http.StatusOK, "home.html", gin.H{
	// 	"title": "Expenset",
	// 	"books": b,
	// })
}
