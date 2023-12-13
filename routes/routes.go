package routes

import (
	"URLshorter/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter(engine *gin.Engine) {
	engine.GET("/", handler.Welcome)
	engine.POST("/create-short-url", handler.CreatShortURL)
	engine.GET("/:short_url", handler.HandleShortUrlRedirect)

}
