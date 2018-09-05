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

// Init logger
func Init(debug bool) {
	l.debug = debug
}

type Field struct {
	s string
}

// Debug severity log
func Debug(ctx context.Context, message string, fields ...Field) {
	if l.debug {
		output(ctx, severityDebug, message, fields)
	}
}

// Info severity log
func Info(ctx context.Context, message string, fields ...Field) {
	output(ctx, severityInfo, message, fields)
}

// Warn severity log
func Warn(ctx context.Context, message string, fields ...Field) {
	output(ctx, severityWarn, message, fields)
}

// Error severity log
func Error(ctx context.Context, message string, fields ...Field) {
	output(ctx, severityError, message, fields)
}

// Fatal severity log
func Fatal(ctx context.Context, message string, fields ...Field) {
	output(ctx, severityFatal, message, fields)
}

// FmtAny into field
func FmtAny(value interface{}, name string) Field {
	blob, err := json.MarshalIndent(value, "\t", "\t")
	if err != nil {
		message := "Failed encoding value as JSON"
		pc, _, line, _ := runtime.Caller(0)
		funcName := runtime.FuncForPC(pc).Name()
		l.loggerError.Println(fmtLog(pkg_context.CorrelationIDBackground, message, funcName, line, nil))
		return Field{fmt.Sprintf("%q: %q", name, message)}
	}
	return Field{fmt.Sprintf("%q: %s", name, blob)}
}

// FmtBool into field
func FmtBool(value bool, name string) Field {
	return Field{fmt.Sprintf("%q: %t", name, value)}
}

// FmtByte into field
func FmtByte(value byte, name string) Field {
	return Field{fmt.Sprintf("%q: %q", name, value)}
}

// FmtBytes into field
func FmtBytes(value []byte, name string) Field {
	return Field{fmt.Sprintf("%q: %q", name, value)}
}

// FmtDuration into field
func FmtDuration(value time.Duration, name string) Field {
	return Field{fmt.Sprintf("%q: %q", name, value)}
}

// FmtError into field
func FmtError(value error) Field {
	return Field{fmt.Sprintf("%q: {\n\t\t%q: %q,\n\t\t%q: \"%+v\"\n\t}", "error", "friendly", value, "trace", value)}
}

// FmtFloat64 into field
func FmtFloat64(value float64, name string) Field {
	return Field{fmt.Sprintf("%q: %f", name, value)}
}

// FmtInt into field
func FmtInt(value int, name string) Field {
	return Field{fmt.Sprintf("%q: %d", name, value)}
}

// FmtInt64 into field
func FmtInt64(value int64, name string) Field {
	return Field{fmt.Sprintf("%q: %d", name, value)}
}

// FmtString into field
func FmtString(value string, name string) Field {
	return Field{fmt.Sprintf("%q: %q", name, value)}
}

// FmtStrings into field
func FmtStrings(values []string, name string) Field {
	f := make([]string, len(values))
	for i, v := range values {
		f[i] = fmt.Sprintf("%q", v)
	}
	return Field{fmt.Sprintf("%q: [\n\t\t%s\n\t]", name, strings.Join(f, ",\n\t\t"))}
}

// FmtTime into field
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

var (
	l = &logger{
		loggerDebug: log.New(os.Stdout, fmt.Sprintf("\x1b[%dmDEBUG ", green), log.Ldate|log.Ltime|log.Lmicroseconds),
		loggerInfo:  log.New(os.Stdout, fmt.Sprintf("\x1b[%dmINFO  ", cyan), log.Ldate|log.Ltime|log.Lmicroseconds),
		loggerWarn:  log.New(os.Stderr, fmt.Sprintf("\x1b[%dmWARN  ", yellow), log.Ldate|log.Ltime|log.Lmicroseconds),
		loggerError: log.New(os.Stderr, fmt.Sprintf("\x1b[%dmERROR ", red), log.Ldate|log.Ltime|log.Lmicroseconds),
		loggerFatal: log.New(os.Stderr, fmt.Sprintf("\x1b[%dmFATAL ", red), log.Ldate|log.Ltime|log.Lmicroseconds),
	}
)

type logger struct {
	debug       bool
	loggerDebug *log.Logger
	loggerInfo  *log.Logger
	loggerWarn  *log.Logger
	loggerError *log.Logger
	loggerFatal *log.Logger
}

func output(ctx context.Context, severity int, message string, fields []Field) {
	pc, _, line, _ := runtime.Caller(2)
	funcName := runtime.FuncForPC(pc).Name()
	message = fmtLog(pkg_context.CorrelationID(ctx), message, funcName, line, fields)
	switch severity {
	case severityDebug:
		l.loggerDebug.Println(message)
	case severityInfo:
		l.loggerInfo.Println(message)
	case severityWarn:
		l.loggerWarn.Println(message)
	case severityError:
		l.loggerError.Println(message)
	case severityFatal:
		l.loggerFatal.Panicln(message)
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
