package logger

import (
	"github.com/pkg/errors"
	"go.uber.org/zap/zapcore"
)

// Supported levels.
const (
	InfoLevel  Level = "info"
	WarnLevel  Level = "warning"
	DebugLevel Level = "debug"
	ErrorLevel Level = "error"
)

const (
	formatWrongLevelMessage = "unsupported log %s level, choice between %s, %s, %s or %s"
)

// Level print.
type Level string

func (l Level) convertToZapLevel() (zapcore.Level, error) {
	switch l {
	case InfoLevel:
		return zapcore.InfoLevel, nil
	case WarnLevel:
		return zapcore.WarnLevel, nil
	case DebugLevel:
		return zapcore.DebugLevel, nil
	case ErrorLevel:
		return zapcore.ErrorLevel, nil
	default:
		return 0, errors.Errorf(
			formatWrongLevelMessage,
			l,
			string(InfoLevel),
			string(WarnLevel),
			string(DebugLevel),
			string(ErrorLevel),
		)
	}
}

func ParseLevel(l string) (Level, error) {
	switch l {
	case string(InfoLevel):
		return InfoLevel, nil
	case string(WarnLevel):
		return WarnLevel, nil
	case string(DebugLevel):
		return DebugLevel, nil
	case string(ErrorLevel):
		return ErrorLevel, nil
	default:
		return "", errors.Errorf(
			formatWrongLevelMessage,
			l,
			string(InfoLevel),
			string(WarnLevel),
			string(DebugLevel),
			string(ErrorLevel),
		)
	}
}
