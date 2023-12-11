package main

import (
	"URLshorter/store"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	store.InitStore()

	engine.Run(":8081")
}
