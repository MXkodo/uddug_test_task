package models

import (
	"encoding/json"
	"time"
)

type Transaction struct {
	Value     int       `json:"Value"`
	Timestamp time.Time `json:"Timestamp"`
}

func (t *Transaction) UnmarshalJSON(data []byte) error {
	type Alias struct {
		Value     int
		Timestamp int64 `json:"Timestamp"`
	}

	var alias Alias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}

	t.Timestamp = time.Unix(alias.Timestamp, 0)
	t.Value = alias.Value

	return nil
}
func (t Transaction) MarshalJSON() ([]byte, error) {
	type Alias Transaction
	return json.Marshal(&struct {
		Value     int   `json:"Value"`
		Timestamp int64 `json:"Timestamp"`
		*Alias
	}{
		Value:     t.Value,
		Timestamp: t.Timestamp.Unix(),
		Alias:     (*Alias)(&t),
	})
}
