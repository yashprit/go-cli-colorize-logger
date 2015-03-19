package main

import (
	"errors"
	"github.com/yashprit/go-cli-colorize-logger"
)

func main() {
	log.Info("Hello %s %d", "hi", 1)
	log.Debug("This is debug %f", 1.0)
	log.Warn("This will work, but please check %s", "not a register user")
	log.Error("%s", errors.New("Something wents wrong"))
}
