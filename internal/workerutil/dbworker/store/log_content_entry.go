package store

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type LogContentEntry struct {
	Command []string `json:"command"`
	Out     string   `json:"out"`
}

func (e *LogContentEntry) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("value is not []byte: %T", value)
	}

	return json.Unmarshal(b, &e)
}

func (e LogContentEntry) Value() (driver.Value, error) {
	return json.Marshal(e)
}
