/*
Copyright 2018 Cai Gwatkin

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
