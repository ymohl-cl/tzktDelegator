package logger_test

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/ymohl-cl/tzktDelegator/pkg/logger"
)

func ExampleNew() {
	var err error
	var l logger.Logger

	viper.Set("LOGGER_LEVEL", "info")
	viper.Set("LOGGER_ENCODING", "console")
	defer func() {
		viper.Reset()
	}()

	if l, err = logger.New(); err != nil {
		fmt.Println(err.Error())

		return
	}

	defer l.Close()

	// l.Info("your message")
	fmt.Println("done")

	// Output:
	// done
}

func ExampleNewWithConfig() {
	var err error
	var l logger.Logger

	if l, err = logger.NewWithConfig(
		logger.Config{Level: "info", Encoding: logger.JSONEncoding}); err != nil {
		fmt.Println(err.Error())

		return
	}

	defer l.Close()

	// l.Info("your message")
	fmt.Println("done")

	// Output:
	// done
}

func ExampleNewWithConfig_enum() {
	var err error
	var l logger.Logger

	viper.Set("LOGGER_LEVEL", "info")
	viper.Set("LOGGER_ENCODING", "console")
	defer func() {
		viper.Reset()
	}()

	c, err := logger.NewConfig()
	if err != nil {
		fmt.Println(err.Error())

		return
	}

	if l, err = logger.NewWithConfig(c); err != nil {
		fmt.Println(err.Error())

		return
	}

	defer l.Close()

	// l.Info("your message")
	fmt.Println("done")

	// Output:
	// done
}
