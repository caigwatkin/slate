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

package log

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
	pkg_context "slate/internal/pkg/context"
	"strings"
	"time"
)

type Client interface {
	Debug(ctx context.Context, message string, fields ...Field)
	Info(ctx context.Context, message string, fields ...Field)
	Warn(ctx context.Context, message string, fields ...Field)
	Error(ctx context.Context, message string, fields ...Field)
	Fatal(ctx context.Context, message string, fields ...Field)
}

type client struct {
	debug       bool
	loggerDebug *log.Logger
	loggerInfo  *log.Logger
	loggerWarn  *log.Logger
	loggerError *log.Logger
	loggerFatal *log.Logger
}

// NewClient for logging
func NewClient(enableDebug bool) Client {
	return client{
		debug:       enableDebug,
		loggerDebug: log.New(os.Stdout, fmt.Sprintf("\x1b[%dmDEBUG ", green), log.Ldate|log.Ltime|log.Lmicroseconds),
		loggerInfo:  log.New(os.Stdout, fmt.Sprintf("\x1b[%dmINFO  ", cyan), log.Ldate|log.Ltime|log.Lmicroseconds),
		loggerWarn:  log.New(os.Stderr, fmt.Sprintf("\x1b[%dmWARN  ", yellow), log.Ldate|log.Ltime|log.Lmicroseconds),
		loggerError: log.New(os.Stderr, fmt.Sprintf("\x1b[%dmERROR ", red), log.Ldate|log.Ltime|log.Lmicroseconds),
		loggerFatal: log.New(os.Stderr, fmt.Sprintf("\x1b[%dmFATAL ", red), log.Ldate|log.Ltime|log.Lmicroseconds),
	}
}

// Debug log at debug level
func (c client) Debug(ctx context.Context, message string, fields ...Field) {
	if c.debug {
		c.output(ctx, severityDebug, message, fields)
	}
}

// Info log at info level
func (c client) Info(ctx context.Context, message string, fields ...Field) {
	c.output(ctx, severityInfo, message, fields)
}

// Warn log at warn level
func (c client) Warn(ctx context.Context, message string, fields ...Field) {
	c.output(ctx, severityWarn, message, fields)
}

// Error log at error level
func (c client) Error(ctx context.Context, message string, fields ...Field) {
	c.output(ctx, severityError, message, fields)
}

// Fatal log at fatal level
func (c client) Fatal(ctx context.Context, message string, fields ...Field) {
	c.output(ctx, severityFatal, message, fields)
}

// Field to log
type Field struct {
	s string
}

// FmtAny using JSON marshaller with indenting
//
// Use this to format for logging any object that can be JSON unmarshalled
// If JSON unmarshalling fails, the object will be formatted as a verb
func FmtAny(value interface{}, name string) Field {
	blob, err := json.MarshalIndent(value, "\t", "\t")
	if err != nil {
		return Field{fmt.Sprintf("%q: %v", name, value)}
	}
	return Field{fmt.Sprintf("%q: %s", name, blob)}
}

// FmtBool as name/value pair for logging
func FmtBool(value bool, name string) Field {
	return Field{fmt.Sprintf("%q: %t", name, value)}
}

// FmtBools as name/value pair for logging
func FmtBools(values []bool, name string) Field {
	if len(values) == 0 {
		return Field{fmt.Sprintf("%q: []", name)}
	}
	f := make([]string, len(values))
	for i, v := range values {
		f[i] = fmt.Sprintf("%t", v)
	}
	return Field{fmt.Sprintf("%q: [\n\t\t%s\n\t]", name, strings.Join(f, ",\n\t\t"))}
}

// FmtByte as name/value pair for logging
func FmtByte(value byte, name string) Field {
	return Field{fmt.Sprintf("%q: %q", name, value)}
}

// FmtBytes as name/value pair for logging
func FmtBytes(value []byte, name string) Field {
	return Field{fmt.Sprintf("%q: %q", name, value)}
}

// FmtDuration as name/value pair for logging
func FmtDuration(value time.Duration, name string) Field {
	return Field{fmt.Sprintf("%q: %q", name, value)}
}

// FmtError as name/value pair for logging
func FmtError(value error) Field {
	return Field{fmt.Sprintf("%q: {\n\t\t%q: %q,\n\t\t%q: \"%+v\"\n\t}", "error", "friendly", value, "trace", value)}
}

// FmtFloat64 as name/value pair for logging
func FmtFloat64(value float64, name string) Field {
	return Field{fmt.Sprintf("%q: %f", name, value)}
}

// FmtInt as name/value pair for logging
func FmtInt(value int, name string) Field {
	return Field{fmt.Sprintf("%q: %d", name, value)}
}

// FmtInt64 as name/value pair for logging
func FmtInt64(value int64, name string) Field {
	return Field{fmt.Sprintf("%q: %d", name, value)}
}

// FmtString as name/value pair for logging
func FmtString(value string, name string) Field {
	return Field{fmt.Sprintf("%q: %q", name, value)}
}

// FmtStrings as name/value pair for logging
func FmtStrings(values []string, name string) Field {
	if len(values) == 0 {
		return Field{fmt.Sprintf("%q: []", name)}
	}
	f := make([]string, len(values))
	for i, v := range values {
		f[i] = fmt.Sprintf("%q", v)
	}
	return Field{fmt.Sprintf("%q: [\n\t\t%s\n\t]", name, strings.Join(f, ",\n\t\t"))}
}

// FmtTime as name/value pair for logging
func FmtTime(value time.Time, name string) Field {
	return Field{fmt.Sprintf("%q: %q", name, value.Format(time.RFC3339Nano))}
}

const (
	red    = 31
	green  = 32
	yellow = 33
	cyan   = 36

	severityDebug = iota
	severityInfo  = iota
	severityWarn  = iota
	severityError = iota
	severityFatal = iota
)

func (c client) output(ctx context.Context, severity int, message string, fields []Field) {
	pc, _, line, _ := runtime.Caller(2)
	funcName := runtime.FuncForPC(pc).Name()
	message = fmtLog(pkg_context.CorrelationID(ctx), message, funcName, line, fields)
	switch severity {
	case severityDebug:
		c.loggerDebug.Println(message)
	case severityInfo:
		c.loggerInfo.Println(message)
	case severityWarn:
		c.loggerWarn.Println(message)
	case severityError:
		c.loggerError.Println(message)
	case severityFatal:
		c.loggerFatal.Panicln(message)
	}
}

func fmtLog(correlationID, message, funcName string, line int, fields []Field) string {
	if len(fields) > 0 {
		return fmt.Sprintf("%s %s %s:%d %s\x1b[0m", message, correlationID, funcName, line, fmtFields(fields))
	}
	return fmt.Sprintf("%s %s %s:%d\x1b[0m", message, correlationID, funcName, line)
}

func fmtFields(fields []Field) string {
	var fs []string
	for _, field := range fields {
		fs = append(fs, field.s)
	}
	return fmt.Sprintf("{\n\t%s\n}", strings.Join(fs, ",\n\t"))
}
