package messaging

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/weflux/fluxworks/logging"
	"github.com/weflux/fluxworks/types"
)

type Logger struct {
	logger *logging.Logger
}

func NewLogger(logger *logging.Logger) *Logger {
	return &Logger{logger: logger}
}

func (l *Logger) Error(msg string, err error, fields watermill.LogFields) {
	f := types.M{}
	for k, v := range fields {
		f[k] = v
	}
	l.logger.Error(msg, err, f)
}

func (l *Logger) Info(msg string, fields watermill.LogFields) {
	f := types.M{}
	for k, v := range fields {
		f[k] = v
	}
	l.logger.Info(msg, f)
}

func (l *Logger) Debug(msg string, fields watermill.LogFields) {
	f := types.M{}
	for k, v := range fields {
		f[k] = v
	}
	l.logger.Debug(msg, f)
}

func (l *Logger) Trace(msg string, fields watermill.LogFields) {
	f := types.M{}
	for k, v := range fields {
		f[k] = v
	}
	l.logger.Debug(msg, f)
}

func (l *Logger) With(fields watermill.LogFields) watermill.LoggerAdapter {
	f := types.M{}
	for k, v := range fields {
		f[k] = v
	}
	logger := l.logger.With(f)
	return &Logger{logger: logger}
}
