package main

import (
	"URLshorter/routes"
	"URLshorter/store"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	store.InitStore()

	routes.SetupRouter(engine)

	errRun := engine.Run(":8081")
	if errRun != nil {
		panic(fmt.Sprintf("Error while running server: {%s}", errRun.Error()))
	}
}
