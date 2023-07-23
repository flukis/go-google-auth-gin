package main

import (
	"expenset/internals/presenter"
	"expenset/pkg/config"

	"github.com/gin-gonic/gin"
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

	presenter.NewBaseHandler().Route(&r.RouterGroup)
	presenter.NewOAuthHandler(config.GoogleOauthCfg, config.GoogleOauthUrlApi).Route(&r.RouterGroup)

	r.Run(":3000")
}
