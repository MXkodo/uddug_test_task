package routes

import (
	"github.com/MXkodo/uddug_test_task/handlers"
	"github.com/gin-gonic/gin"
)

func SetupTransactionRoutes(router *gin.Engine) {
	router.POST("/group-transactions", handlers.GroupTransactions)
}
