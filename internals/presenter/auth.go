package presenter

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"expenset/internals/service/auth"
	"expenset/pkg/utils"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

type OAuthHandler struct {
	cfg         *oauth2.Config
	url         string
	authService auth.Register
}

func NewOAuthHandler(cfg *oauth2.Config, url string, authService auth.Register) *OAuthHandler {
	return &OAuthHandler{cfg, url, authService}
}

func (h *OAuthHandler) Route(r *gin.RouterGroup) {
	authGroup := r.Group("/auth/google")

	authGroup.GET("/login", h.Login)
	authGroup.GET("/callback", h.Callback)
}

func (h *OAuthHandler) Login(ctx *gin.Context) {
	var expirationTime = time.Now().Add(60 * 24 * time.Hour)
	bin := make([]byte, 16)
	rand.Read(bin)
	state := base64.URLEncoding.EncodeToString(bin)
	cookie := http.Cookie{
		Path:     "/",
		Name:     "oauthstate",
		Value:    state,
		Expires:  expirationTime,
		HttpOnly: true,
		Secure:   true,
	}
	http.SetCookie(ctx.Writer, &cookie)

	u := h.cfg.AuthCodeURL(state)
	ctx.Redirect(http.StatusTemporaryRedirect, u)
}

type OauthData struct {
	ID      string `json:"id"`
	Email   string `json:"email"`
	Picture string `json:"picture"`
}

func (h *OAuthHandler) Callback(ctx *gin.Context) {
	oauthState, err := ctx.Cookie("oauthstate")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &utils.ApiResponse{
			Message: err.Error(),
			Data:    nil,
			Meta:    nil,
		})
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
		ctx.JSON(http.StatusBadRequest, &utils.ApiResponse{
			Message: err.Error(),
			Data:    nil,
			Meta:    nil,
		})
		ctx.Abort()
		return
	}

	res, err := http.Get(h.url + token.AccessToken)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &utils.ApiResponse{
			Message: err.Error(),
			Data:    nil,
			Meta:    nil,
		})
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

	var data OauthData
	err = json.Unmarshal(contents, &data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &utils.ApiResponse{
			Message: err.Error(),
			Data:    nil,
			Meta:    nil,
		})
		ctx.Abort()
		return
	}

	registerData := auth.RegistrationRequest{
		ID:    data.ID,
		Email: data.Email,
	}

	resp := h.authService.Registration(ctx, registerData)
	if resp.Error != nil {
		ctx.JSON(resp.Error.Code, &utils.ApiResponse{
			Message: resp.Error.Error.Error(),
			Data:    nil,
			Meta:    nil,
		})
		ctx.Abort()
		return
	}

	ctx.Redirect(http.StatusTemporaryRedirect, "/")
}
