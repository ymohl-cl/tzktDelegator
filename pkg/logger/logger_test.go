package logger

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	evaluableMessage = "avaluable message to the tests"
)

func TestNewWithConfig(t *testing.T) {
	t.Run("Should be ok with an info level", func(t *testing.T) {
		l, err := NewWithConfig(Config{Level: InfoLevel, Encoding: ConsoleEncoding})
		if assert.NoError(t, err) {
			assert.NotNil(t, l)
		}
	})
	t.Run("Should be ok with a warning level", func(t *testing.T) {
		l, err := NewWithConfig(Config{Level: WarnLevel, Encoding: ConsoleEncoding})
		if assert.NoError(t, err) {
			assert.NotNil(t, l)
		}
	})
	t.Run("Should be ok with a debug level", func(t *testing.T) {
		l, err := NewWithConfig(Config{Level: DebugLevel, Encoding: ConsoleEncoding})
		if assert.NoError(t, err) {
			assert.NotNil(t, l)
		}
	})
	t.Run("Should be ok with an error level", func(t *testing.T) {
		l, err := NewWithConfig(Config{Level: ErrorLevel, Encoding: ConsoleEncoding})
		if assert.NoError(t, err) {
			assert.NotNil(t, l)
		}
	})
	t.Run("Should be ok with an json encoding", func(t *testing.T) {
		l, err := NewWithConfig(Config{Level: "error", Encoding: JSONEncoding})
		if assert.NoError(t, err) {
			assert.NotNil(t, l)
		}
	})
	t.Run("Should be ok with an console encoding", func(t *testing.T) {
		l, err := NewWithConfig(Config{Level: "error", Encoding: ConsoleEncoding})
		if assert.NoError(t, err) {
			assert.NotNil(t, l)
		}
	})
}

// Capture the log to return it like a string with the Do method.
// Just use it for the unit tests
// Example:
// capture := NewCapture()
// log := logger{driver: c.driver}
// str := capture.Do(func(log.Info("ceci est un message de test")))
// str == "ceci est un message de test".
type LogCapture struct {
	writer *bufio.Writer
	buffer *bytes.Buffer
	driver *zap.Logger
}

func NewLogCapture() LogCapture {
	var capture LogCapture

	capture.buffer = &bytes.Buffer{}
	capture.writer = bufio.NewWriter(capture.buffer)
	cfg := zap.NewDevelopmentEncoderConfig()
	cfg.TimeKey = ""
	cfg.EncodeTime = nil
	encoder := zapcore.NewConsoleEncoder(cfg)
	capture.driver = zap.New(
		zapcore.NewCore(encoder, zapcore.AddSync(capture.writer), zapcore.DebugLevel))

	return capture
}

func (c LogCapture) Do(t *testing.T, f func()) string {
	c.buffer.Reset()
	f()
	if err := c.writer.Flush(); err != nil {
		t.Fatal(err)
	}

	return c.buffer.String()
}

func TestLogger_Info(t *testing.T) {
	capture := NewLogCapture()

	t.Run("should print log message", func(t *testing.T) {
		log := &logger{
			driver: capture.driver,
		}
		expectedLog := fmt.Sprintf("INFO\t%s\n", evaluableMessage)
		str := capture.Do(t, func() { log.Info(evaluableMessage) })
		assert.EqualValues(t, expectedLog, str)
	})
}

func TestLogger_Warn(t *testing.T) {
	capture := NewLogCapture()

	t.Run("should print log message", func(t *testing.T) {
		log := &logger{
			driver: capture.driver,
		}
		expectedLog := fmt.Sprintf("WARN\t%s\n", evaluableMessage)
		str := capture.Do(t, func() { log.Warn(evaluableMessage) })
		assert.EqualValues(t, expectedLog, str)
	})
}

func TestLogger_Debug(t *testing.T) {
	capture := NewLogCapture()

	t.Run("should print log message", func(t *testing.T) {
		log := &logger{
			driver: capture.driver,
		}
		expectedLog := fmt.Sprintf("DEBUG\t%s\n", evaluableMessage)
		str := capture.Do(t, func() { log.Debug(evaluableMessage) })
		assert.EqualValues(t, expectedLog, str)
	})
}

func TestLogger_Error(t *testing.T) {
	capture := NewLogCapture()

	t.Run("should print log message", func(t *testing.T) {
		log := &logger{
			driver: capture.driver,
		}
		expectedLog := fmt.Sprintf("ERROR\t%s\n", evaluableMessage)
		str := capture.Do(t, func() { log.Error(errors.New(evaluableMessage)) })
		assert.EqualValues(t, expectedLog, str)
	})
}

func TestLogger_WithStackTrace(t *testing.T) {
	capture := NewLogCapture()

	t.Run("should print log message", func(t *testing.T) {
		log := &logger{
			driver: capture.driver,
		}
		expectedLog := fmt.Sprintf(
			"ERROR\t%s\ngithub.com/ymohl-cl/tzktDelegator/pkg/logger.TestLogger_WithStackTrace.func",
			evaluableMessage)
		str := capture.Do(t, func() { log.WithStackTrace().Error(errors.New(evaluableMessage)) })
		assert.Contains(t, str, expectedLog)
	})
}

func TestLogger_WithField(t *testing.T) {
	capture := NewLogCapture()

	t.Run("should print log message with customs field", func(t *testing.T) {
		log := &logger{
			driver: capture.driver,
		}
		expectedLog := fmt.Sprintf("WARN\t%s\t{\"custom_key\": \"value\"}\n", evaluableMessage)
		str := capture.Do(t, func() { log.WithField("custom_key", "value").Warn(evaluableMessage) })
		assert.EqualValues(t, expectedLog, str)
	})
}

// SyncError struct which implement zapcore.WriteSyncer to return an error in close situation.
type SyncError struct {
	io.Writer
}

// Sync implement zapcore.WriteSyncer.
func (s SyncError) Sync() error { return errors.New("error to sync writer") }

func TestLogger_Close(t *testing.T) {
	t.Run("Should return an error because sync writer was failed", func(t *testing.T) {
		// specific initialization to specific error statement.
		var capture LogCapture
		var mockSync SyncError
		var log logger

		capture.buffer = &bytes.Buffer{}
		capture.writer = bufio.NewWriter(capture.buffer)
		mockSync.Writer = capture.writer

		cfg := zap.NewDevelopmentEncoderConfig()
		cfg.TimeKey = ""
		cfg.EncodeTime = nil

		encoder := zapcore.NewConsoleEncoder(cfg)
		capture.driver = zap.New(zapcore.NewCore(encoder, &mockSync, zapcore.DebugLevel))
		log.driver = capture.driver

		message := "error to sync writer"
		expectedLog := fmt.Sprintf("ERROR\t%s\n", message)
		str := capture.Do(t, func() { log.Close() })
		assert.EqualValues(t, expectedLog, str)
	})
	t.Run("Should be ok", func(t *testing.T) {
		capture := NewLogCapture()
		log := &logger{
			driver: capture.driver,
		}
		str := capture.Do(t, func() { log.Close() })
		assert.EqualValues(t, "", str)
	})
}
