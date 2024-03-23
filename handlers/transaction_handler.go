package handlers

import (
	"net/http"
	"time"

	"github.com/MXkodo/uddug_test_task/models"
	"github.com/MXkodo/uddug_test_task/services"
	"github.com/gin-gonic/gin"
)

const (
	MonthInterval = 30 * 24 * time.Hour
	WeekInterval  = 7 * 24 * time.Hour
	DayInterval   = 24 * time.Hour
	HourInterval  = time.Hour
)

type TransactionRequest struct {
	Transactions []*models.Transaction `json:"transactions"`
	Interval     string                `json:"interval"`
}

func GroupTransactions(c *gin.Context) {
	var request TransactionRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var interval time.Duration
	switch request.Interval {
	case "month":
		interval = MonthInterval
	case "week":
		interval = WeekInterval
	case "day":
		interval = DayInterval
	case "hour":
		interval = HourInterval
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interval"})
		return
	}

	groupedTransactions := services.GroupTransactionsByInterval(request.Transactions, interval)

	c.JSON(http.StatusOK, gin.H{
		"transactions": groupedTransactions,
	})
}
