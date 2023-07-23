package presenter

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

type OAuthHandler struct {
	cfg *oauth2.Config
	url string
}

func NewOAuthHandler(cfg *oauth2.Config, url string) *OAuthHandler {
	return &OAuthHandler{cfg, url}
}

func (h *OAuthHandler) Route(r *gin.RouterGroup) {
	authGroup := r.Group("/auth/google")

	authGroup.GET("/login", h.Login)
	authGroup.GET("/callback", h.Callback)
}

func (h *OAuthHandler) Login(ctx *gin.Context) {
	var expirationTime = time.Now().Add(20 * time.Minute)
	bin := make([]byte, 16)
	rand.Read(bin)
	state := base64.URLEncoding.EncodeToString(bin)
	cookie := http.Cookie{
		Name:    "oauthstate",
		Value:   state,
		Expires: expirationTime,
	}
	http.SetCookie(ctx.Writer, &cookie)

	u := h.cfg.AuthCodeURL(state)
	ctx.Redirect(http.StatusTemporaryRedirect, u)
}

func (h *OAuthHandler) Callback(ctx *gin.Context) {
	oauthState, err := ctx.Cookie("oauthstate")
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Abort()
		return
	}

	state := ctx.Query("state")
	code := ctx.Query("code")

	if oauthState != state {
		ctx.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	token, err := h.cfg.Exchange(ctx, code)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Abort()
		return
	}

	res, err := http.Get(h.url + token.AccessToken)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Abort()
		return
	}
	defer res.Body.Close()

	contents, err := io.ReadAll(res.Body)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		ctx.Abort()
		return
	}

	fmt.Fprintf(ctx.Writer, "UserInfo: %s\n", contents)
	ctx.Status(http.StatusOK)
}
