package validator

import (
	"encoding/json"
	"fmt"
)

// Errors validation errors
type Errors struct {
	Errors VErrors `json:"fields"`
}

// Error error interface implemetntation
func (e Errors) Error() string {
	var errStr string
	for _, v := range e.Errors {
		errStr = v
		break
	}
	if len(e.Errors) > 1 {
		errStr = fmt.Sprintf("%s and %d more errors", errStr, len(e.Errors)-1)
	}
	return errStr
}

// MarshalJSON custom json marshaling
func (e *Errors) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.Errors)
}

// VErrors validation error
type VErrors map[string]string
