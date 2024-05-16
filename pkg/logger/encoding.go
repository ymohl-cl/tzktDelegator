package logger

import "github.com/pkg/errors"

// Supported encoding.
const (
	JSONEncoding    Encoding = "json"
	ConsoleEncoding Encoding = "console"
)

const (
	formatWrongEncodingMessage = "unsupported log %s encoding, choice between %s or %s"
)

// Encoding logger.
type Encoding string

// Valid validates the encoding value
func (e Encoding) Valid() error {
	switch e {
	case JSONEncoding:
		return nil
	case ConsoleEncoding:
		return nil
	default:
		return errors.Errorf(
			formatWrongEncodingMessage,
			e,
			string(JSONEncoding),
			string(ConsoleEncoding),
		)
	}
}

func ParseEncoding(encode string) (Encoding, error) {
	switch encode {
	case string(JSONEncoding):
		return JSONEncoding, nil
	case string(ConsoleEncoding):
		return ConsoleEncoding, nil
	default:
		return "", errors.Errorf(
			formatWrongEncodingMessage,
			encode,
			string(JSONEncoding),
			string(ConsoleEncoding),
		)
	}
}
