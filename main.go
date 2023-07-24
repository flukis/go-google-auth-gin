package main

import (
	"context"
	"expenset/internals/presenter"
	"expenset/internals/service/auth"
	"expenset/internals/storer/account"
	"expenset/pkg/config"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	config.InitGoogleConfig()
	if err != nil {
		panic("error loading .env file")
	}
	r := gin.Default()
	r.LoadHTMLGlob("views/**/*")
	r.Static("/static/", "static")
	r.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "404.html", gin.H{
			"title": "Page Not Found",
		})
	})

	pool, err := pgxpool.New(context.Background(), os.Getenv("POSTGRESQL_URL"))
	if err != nil {
		panic("failed to connect db")
	}

	authWriter := account.NewWriter(pool)
	authReader := account.NewReader(pool)

	authRegister := auth.NewRegister(authWriter, authReader)

	presenter.NewBaseHandler().Route(&r.RouterGroup)
	presenter.NewOAuthHandler(config.GoogleOauthCfg, config.GoogleOauthUrlApi, authRegister).Route(&r.RouterGroup)

	r.Run(":3000")
}
