# Logger

Logger is a wrapper tool to log driver. Actually use the [zap.logger driver](https://github.com/uber-go/zap) [see the architecture decision]

## Install

## Configuration

To configure, please provide the following environment variables

``` bash
export LOGGER_LEVEL=<debug/info/warning/error>
```

## Usage

``` go
func ExampleNew() {
  var err error
  var l logger.Logger

  /* setup your env values */
  os.Setenv("LOGGER_LEVEL", "info")
  defer func() {
    os.Unsetenv("LOGGER_LEVEL")
  }()

  if l, err = logger.New("appname"); err != nil {
    fmt.Println(err.Error())
    return
  }

  defer l.Close()

  // l.Info("your message")
  fmt.Println("done")

  // Output:
  // done
}
```

``` go
func ExampleNewWithConfig() {
  var err error
  var l logger.Logger

  if l, err = logger.NewWithConfig("appname", logger.Config{Level: "info"}); err != nil {
    fmt.Println(err.Error())
    return
  }

  defer l.Close()

  // l.Info("your message")
  fmt.Println("done")

  // Output:
  // done
}
```

A mock package is provided to your unit tests and implement the [testify/mock](https://github.com/stretchr/testify) package.

``` go
import (
  "testing"

  "github.com/stretchr/testify/mock"

  "github/winning-number/lotto/pkg/logger/mocks"
)

func TestAny(t *testing.T) {
  t.Run("Example", func(t *testing.T) {
    logger := &mocks.Logger{}

    logger.On("Info", "test input").Return()
    // funcThatDoesSomething(logger)
    logger.AssertExpectations(t)
  })
}
```
