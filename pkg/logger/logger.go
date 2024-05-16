// Package logger package to wrap a log driver
package logger

import (
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger driver to log any things
// Call the close method before close you app
//
//go:generate mockery --name=Logger --output=mocks --filename=logger.go --outpkg=mocks
type Logger interface {
	Close()
	Info(msg string)
	Warn(msg string)
	Debug(msg string)
	Error(err error)
	WithStackTrace() Logger
	WithField(key, value string) Logger
}

type logger struct {
	driver *zap.Logger
}

// New instance of Logger
func New() (Logger, error) {
	var err error
	var c Config
	var l Logger

	if c, err = NewConfig(); err != nil {
		return nil, err
	}
	if l, err = NewWithConfig(c); err != nil {
		return nil, err
	}

	return l, err
}

// NewWithConfig instance of logger with manual configuration.
func NewWithConfig(c Config) (Logger, error) {
	var err error
	var log logger

	level, err := c.Level.convertToZapLevel()
	if err != nil {
		return nil, err
	}

	if err = c.Encoding.Valid(); err != nil {
		return nil, err
	}

	if log.driver, err = (zap.Config{
		Level:             zap.NewAtomicLevelAt(level),
		Encoding:          string(c.Encoding),
		OutputPaths:       []string{"stderr"},
		ErrorOutputPaths:  []string{"stderr"},
		DisableStacktrace: true,
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:    "message",
			LevelKey:      "level",
			EncodeLevel:   zapcore.CapitalLevelEncoder,
			TimeKey:       "time",
			EncodeTime:    zapcore.ISO8601TimeEncoder,
			CallerKey:     "caller",
			StacktraceKey: "stacktrace",
			EncodeCaller:  zapcore.ShortCallerEncoder,
		},
	}.Build()); err != nil {
		return nil, err
	}

	return &log, nil
}

// Close method call Sync (kind of flush buffer).
func (l *logger) Close() {
	if err := l.driver.Sync(); err != nil {
		if strings.Contains(err.Error(), "bad file descriptor") ||
			strings.Contains(err.Error(), "invalid argument") {
			// ignore because the stderr should not sync.
			return
		}
		l.driver.Error(err.Error())
	}
}

// Info level message.
func (l *logger) Info(msg string) {
	l.driver.WithOptions(zap.AddCallerSkip(1)).Info(msg)
}

// Warn level message.
func (l *logger) Warn(msg string) {
	l.driver.WithOptions(zap.AddCallerSkip(1)).Warn(msg)
}

// Debug level message.
func (l *logger) Debug(msg string) {
	l.driver.WithOptions(zap.AddCallerSkip(1)).Debug(msg)
}

// Error level message.
func (l *logger) Error(err error) {
	l.driver.WithOptions(zap.AddCallerSkip(1)).Error(err.Error())
}

// WithStackTrace add log's stack trace
func (l *logger) WithStackTrace() Logger {
	return &logger{driver: l.driver.WithOptions(zap.AddStacktrace(zapcore.DebugLevel))}
}

// WithField to add specific keys - values
// Example usage:
// logger.WithField("key-1", "value-1").With.Field("key-2", "value-2").Info(
// "message with context")
func (l *logger) WithField(key, value string) Logger {
	return &logger{driver: l.driver.With(zap.Any(key, value))}
}
