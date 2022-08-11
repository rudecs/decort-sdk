package tasks

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type TaskResult int

func (r *TaskResult) UnmarshalJSON(b []byte) error {
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
		*r = TaskResult(n)
	} else if b[0] == '[' {
		res := []interface{}{}
		if err := json.Unmarshal(b, &res); err != nil {
			return err
		}
		if n, ok := res[0].(float64); ok {
			*r = TaskResult(n)
		} else {
			return fmt.Errorf("could not unmarshal %v into int", res[0])
		}
	}

	return nil
}

//AsyncTask represents a long task completion status
type AsyncTask struct {
	AuditID     string     `json:"auditId"`
	Completed   bool       `json:"completed"`
	Error       string     `json:"error"`
	Log         []string   `json:"log"`
	Result      TaskResult `json:"result"`
	Stage       string     `json:"stage"`
	Status      string     `json:"status"`
	UpdateTime  uint64     `json:"updateTime"`
	UpdatedTime uint64     `json:"updatedTime"`
}

type TasksList []AsyncTask
