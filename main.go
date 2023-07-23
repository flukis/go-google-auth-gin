package main

import (
	"context"
	"expenset/internals/presenter"
	register "expenset/internals/service/auth"
	"expenset/internals/storer/auth"
	"expenset/pkg/config"
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
	pool, err := pgxpool.New(context.Background(), os.Getenv("POSTGRESQL_URL"))
	if err != nil {
		panic("failed to connect db")
	}

	authWriter := auth.NewWriter(pool)
	authReader := auth.NewReader(pool)

	authRegister := register.NewRegister(authWriter, authReader)

	presenter.NewBaseHandler().Route(&r.RouterGroup)
	presenter.NewOAuthHandler(config.GoogleOauthCfg, config.GoogleOauthUrlApi, authRegister).Route(&r.RouterGroup)

	r.Run(":3000")
}
