package tasks

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// Global variable for converting field to desired data type
type InfoResult int

// Method for convert field
func (r *InfoResult) UnmarshalJSON(b []byte) error {
	if b[0] == '"' {
		b := b[1 : len(b)-1]
		if len(b) == 0 {
			*r = 0
			return nil
		}
		n, err := strconv.Atoi(string(b))
		if err != nil {
			return err
		}
		*r = InfoResult(n)
	} else if b[0] == '[' {
		res := []interface{}{}
		if err := json.Unmarshal(b, &res); err != nil {
			return err
		}
		if n, ok := res[0].(float64); ok {
			*r = InfoResult(n)
		} else {
			return fmt.Errorf("could not unmarshal %v into int", res[0])
		}
	}

	return nil
}

// Detailed information about task
type RecordTask struct {
	// Audit ID
	AuditID string `json:"auditId"`

	// Completed
	Completed bool `json:"completed"`

	// Error
	Error string `json:"error"`

	// List of logs
	Log []string `json:"log"`

	// Final result
	Result InfoResult `json:"result"`

	// Stage
	Stage string `json:"stage"`

	// Status
	Status string `json:"status"`

	// Update time
	UpdateTime uint64 `json:"updateTime"`

	// Updated time
	UpdatedTime uint64 `json:"updatedTime"`
}

// List of tasks
type ListTasks []RecordTask
