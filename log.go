package baiduyun

import (
	"fmt"
)

type Logger interface {
	Info(format string, args ...interface{})
	Error(format string, args ...interface{})
}

func (r *Client) SetLogger(logger Logger) {
	r.logger = logger
}

func (r *Client) info(format string, args ...interface{}) {
	if r.logger != nil {
		r.logger.Info(format, args...)
	}
}

func (r *Client) error(format string, args ...interface{}) {
	if r.logger != nil {
		r.logger.Error(format, args...)
	}
}

type defaultLogger struct{}

func NewDefaultLogger() Logger {
	return &defaultLogger{}
}

func (r *defaultLogger) Info(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}

func (r *defaultLogger) Error(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
