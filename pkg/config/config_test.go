package config

import (
	"flag"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func resetFlag() {
	// https://go.dev/src/flag/flag_test.go
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	flag.CommandLine.SetOutput(io.Discard)
	flag.CommandLine.Usage = func() { flag.Usage() }
	flag.Usage = flag.PrintDefaults
}

func TestLoad(t *testing.T) {
	var cpy []string
	_ = copy(cpy, os.Args)
	defer copy(os.Args, cpy)

	t.Run("should load the configuration", func(t *testing.T) {
		os.Args = append([]string{}, os.Args[0])
		resetFlag()

		err := Load()
		assert.NoError(t, err)
	})
	t.Run("should return an error if the .env file is not found", func(t *testing.T) {
		os.Args = append([]string{}, os.Args[0], "-envfile=tata.env")
		resetFlag()

		err := Load()
		if assert.Error(t, err) {
			assert.ErrorContains(t, err, "no such file or directory")
		}
	})
	t.Run("should print the helper message -help", func(t *testing.T) {
		os.Args = append([]string{}, os.Args[0], "-help")
		resetFlag()

		err := Load()
		if assert.Error(t, err) {
			assert.EqualError(t, err, ErrNoError.Error())
		}
	})
	t.Run("should print the helper message with -h", func(t *testing.T) {
		os.Args = append([]string{}, os.Args[0], "-h")
		resetFlag()

		err := Load()
		if assert.Error(t, err) {
			assert.EqualError(t, err, ErrNoError.Error())
		}
	})
}
