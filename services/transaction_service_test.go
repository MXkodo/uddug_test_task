package services_test

import (
	"testing"
	"time"

	"github.com/MXkodo/uddug_test_task/models"
	"github.com/MXkodo/uddug_test_task/services"
)

func TestGroupTransactionsByInterval(t *testing.T) {
	transactions := []*models.Transaction{
	{Timestamp: time.Date(2024, 3, 23, 10, 0, 0, 0, time.UTC)},
	{Timestamp: time.Date(2024, 3, 23, 10, 30, 0, 0, time.UTC)},
	{Timestamp: time.Date(2024, 3, 23, 10, 45, 0, 0, time.UTC)},
}

	groupedTransactions := services.GroupTransactionsByInterval(transactions, time.Hour)

	if len(groupedTransactions) != 1 {
		t.Errorf("Expected 1 transaction, got %d", len(groupedTransactions))
	}
}
