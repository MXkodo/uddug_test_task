package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/MXkodo/uddug_test_task/handlers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGroupTransactions(t *testing.T) {
	router := gin.Default()
	router.POST("/group-transactions", handlers.GroupTransactions)

	requestBody := `{"transactions": [{"timestamp": 1672444800}, {"timestamp": 1672444800}, {"timestamp": 1672444800}], "interval": "hour"}`
	req, err := http.NewRequest("POST", "/group-transactions", strings.NewReader(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	expectedBody := `{"transactions":[{"Timestamp":1672444800, "Value":0}]}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}
