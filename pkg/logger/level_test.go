package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"
)

func TestConvertToZapLevel(t *testing.T) {

	t.Run("Should be ok with error level", func(t *testing.T) {
		l := ErrorLevel
		expectedLevel := zapcore.ErrorLevel
		level, err := l.convertToZapLevel()
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedLevel, level)
		}
	})
	t.Run("Should be ok with debug level", func(t *testing.T) {
		l := DebugLevel
		expectedLevel := zapcore.DebugLevel
		level, err := l.convertToZapLevel()
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedLevel, level)
		}
	})
	t.Run("Should be ok with warn level", func(t *testing.T) {
		l := WarnLevel
		expectedLevel := zapcore.WarnLevel
		level, err := l.convertToZapLevel()
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedLevel, level)
		}
	})
	t.Run("Should be ok with info level", func(t *testing.T) {
		l := InfoLevel
		expectedLevel := zapcore.InfoLevel
		level, err := l.convertToZapLevel()
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedLevel, level)
		}
	})
	t.Run("Should be ko with too level", func(t *testing.T) {
		l := Level("toto")
		expectedErr := "unsupported log toto level, choice between info, warning, debug or error"
		level, err := l.convertToZapLevel()
		if assert.Error(t, err) {
			assert.EqualError(t, err, expectedErr)
			assert.Empty(t, level)
		}
	})
}
