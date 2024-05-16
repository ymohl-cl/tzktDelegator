// Package config provides a simple way to load the configuration.
// The configuration is loaded from the environment variables and the .env file
// It uses the viper library to load the configuration
package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

const (
	ParamEnvFile = "envfile"
	ParamHelp    = "help"
	ParamH       = "h"
)

var (
	ErrNoError = fmt.Errorf("no error")
)

// Load loads the configuration from the environment variables and the .env file
// The .env file is specified by the -envfile flag and is optional
// The Load method call flag.Parse(), so you could record any other flag in your application
// If the the -help or -h flag is specified, the Load method will print the help message
// and method will return an error type of ErrNoError to properly stop the application
func Load() (err error) {
	appname := os.Args[0]

	envfile := flag.String(ParamEnvFile, "", fmt.Sprintf("./%s -envfile=.env", appname))
	help := flag.Bool(ParamHelp, false, fmt.Sprintf("./%s -help", appname))
	helpMin := flag.Bool(ParamH, false, fmt.Sprintf("./%s -h", appname))

	flag.Parse()
	if *help || *helpMin {
		flag.PrintDefaults()

		return ErrNoError
	}
	if *envfile != "" {
		viper.SetConfigFile(*envfile)
		if err = viper.ReadInConfig(); err != nil {
			return err
		}
	}
	viper.AutomaticEnv()

	return nil
}
