package main

import (
	"errors"
	"github.com/yashprit/go-cli-colorize-logger"
)

func main() {
	clog.Info("Hello %s %d", "hi", 1)
	clog.IsDebug = true
	clog.Debug("This is debug %f", 1.0)
	clog.Warn("This will work, but please check %s", "not a register user")
	clog.Error("%s", errors.New("Something wents wrong"))
}
