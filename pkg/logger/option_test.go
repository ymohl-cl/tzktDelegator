package logger

import (
	"testing"
)

func TestWithLevel(t *testing.T) {
	/*	t.Run("Should return the default level", func(t *testing.T) {
		c, err := NewConfig()
		if assert.NoError(t, err) {
			assert.Equal(t, ErrorLevel, c.Level)
		}
	})*/
	/*
			t.Run("Should return the info level", func(t *testing.T) {
				c := NewConfig(WithLevel(InfoLevel))
				if assert.NoError(t, c.Validate()) {
				assert.Equal(t, InfoLevel, c.Level)
			})
			t.Run("Should return the error level", func(t *testing.T) {
				c := NewConfig(WithLevel(DebugLevel))
				assert.Equal(t, DebugLevel, c.Level)
			})
		}

		func TestWithEncoding(t *testing.T) {
			t.Run("Should return the default encoding", func(t *testing.T) {
				c := NewConfig()
				assert.Equal(t, JSONEncoding, c.Encoding)
			})

			t.Run("Should return the console level", func(t *testing.T) {
				c := NewConfig(WithEncoding(ConsoleEncoding))
				assert.Equal(t, ConsoleEncoding, c.Encoding)
			})
	*/
}
