package services

import (
	"time"

	"github.com/MXkodo/uddug_test_task/models"
)

func GroupTransactionsByInterval(transactions []*models.Transaction, interval time.Duration) []*models.Transaction {
	grouped := make(map[time.Time]*models.Transaction)

	for _, tx := range transactions {
		rounded := tx.Timestamp.Truncate(interval)
		if existing, ok := grouped[rounded]; !ok || tx.Timestamp.After(existing.Timestamp) {
			grouped[rounded] = tx
		}
	}

	var result []*models.Transaction
	for _, tx := range grouped {
		result = append(result, tx)
	}

	return result
}
