package status

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
)

type Status struct {
	At       string                 `json:"at,omitempty"`
	Code     int                    `json:"code,omitempty"`
	Message  string                 `json:"message,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

func NewStatus(code int, message string) *Status {
	return newStatus(1, code, message, nil)
}

func NewStatusWithMetadata(code int, message string, metadata map[string]interface{}) *Status {
	return newStatus(1, code, message, metadata)
}

func NewStatusf(code int, format string, args ...interface{}) *Status {
	return newStatus(1, code, fmt.Sprintf(format, args...), nil)
}

func newStatus(atSkip, code int, message string, metadata map[string]interface{}) *Status {
	pc, _, lineNumber, _ := runtime.Caller(atSkip + 1)
	return &Status{
		At:       fmt.Sprintf("%s:%d", runtime.FuncForPC(pc).Name(), lineNumber),
		Code:     code,
		Message:  fmt.Sprintf("%s: %s", http.StatusText(code), message),
		Metadata: metadata,
	}
}

func (s Status) Error() string {
	return fmt.Sprintf("%q: %q, %q: %q, %q: %q", "code", s.Code, "message", s.Message, "at", s.At)
}

func (s Status) Render() []byte {
	v := map[string]interface{}{
		"code": s.Code,
	}
	switch s.Code {
	default:
		v["message"] = http.StatusText(s.Code)
	case http.StatusBadRequest,
		http.StatusUnprocessableEntity:
		v["message"] = s.Message
		if s.Metadata != nil {
			v["metadata"] = s.Metadata
		}
	}
	b, err := json.Marshal(v)
	if err != nil {
		return nil
	}
	return b
}
