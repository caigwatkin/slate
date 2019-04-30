package log

import (
	"context"
	"fmt"
	"log"
	"os"
)

type Client interface {
	Debug(ctx context.Context, message string, fields ...Field)
	Info(ctx context.Context, message string, fields ...Field)
	Notice(ctx context.Context, message string, fields ...Field)
	Warn(ctx context.Context, message string, fields ...Field)
	Error(ctx context.Context, message string, fields ...Field)
	Fatal(ctx context.Context, message string, fields ...Field)
}

// NewClient for logging
func NewClient(enableDebug bool) Client {
	return client{
		debug:       enableDebug,
		loggerDebug: log.New(os.Stdout, fmt.Sprintf("\x1b[%dmDEBUG ", green), log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile),
		loggerInfo:  log.New(os.Stdout, fmt.Sprintf("\x1b[%dmINFO  ", cyan), log.Ldate|log.Ltime|log.Lmicroseconds),
		loggerWarn:  log.New(os.Stderr, fmt.Sprintf("\x1b[%dmWARN  ", yellow), log.Ldate|log.Ltime|log.Lmicroseconds),
		loggerError: log.New(os.Stderr, fmt.Sprintf("\x1b[%dmERROR ", red), log.Ldate|log.Ltime|log.Lmicroseconds),
		loggerFatal: log.New(os.Stderr, fmt.Sprintf("\x1b[%dmFATAL ", red), log.Ldate|log.Ltime|log.Lmicroseconds),
	}
}

type client struct {
	debug       bool
	loggerDebug *log.Logger
	loggerInfo  *log.Logger
	loggerWarn  *log.Logger
	loggerError *log.Logger
	loggerFatal *log.Logger
}
