package store

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type ExecutionLogEntry struct {
	Command []string `json:"command"`
	Out     string   `json:"out"`
}

func (e *ExecutionLogEntry) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("value is not []byte: %T", value)
	}

	return json.Unmarshal(b, &e)
}

func (e ExecutionLogEntry) Value() (driver.Value, error) {
	return json.Marshal(e)
}
