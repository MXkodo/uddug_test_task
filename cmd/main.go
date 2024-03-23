package main

import (
	"github.com/MXkodo/uddug_test_task/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.SetupTransactionRoutes(router)
	router.Run(":8080")
}
