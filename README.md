# go-cli-logger

> Simple and minimal command line logger in go

![ScreenShot](https://raw.github.com/yashprit/go-cli-colorise-logger/master/output.png)

## Install

```sh
$ go get github.com/yashprit/go-cli-colorise-logger
```

## Usage

```go
  clog.Info("Hello %s %d", "hi", 1)
  clog.IsDebug = true
  clog.Debug("This is debug %f", 1.0)
  clog.Warn("This will work, but please check %s", "not a register user")
  clog.Error("%s", errors.New("Something wents wrong"))
```

## Documentation
Follow godoc.org for details [docs](http://godoc.org/github.com/yashprit/go-cli-colorize-logger)

## Methods

### Debug([message], [arguments])
debug method of clog, argument is optional if you have nou used any placeholder


### Info([message], [arguments])
info method of clog, argument is optional if you have nou used any placeholder

### Warn([error], [arguments])
warn method of clog, argument is optional if you have nou used any placeholder

### Error([error], [arguments])
error method of clog, argument is optional if you have nou used any placeholder

## Contribute or Re port Issue

For bugs and feature requests, [please create an issue](github.com/yashprit/go-cli-colorise-logger/issue).


## License

[unlicense](http://unlicense.org/) © [Yashprit](yashprit.github.io)
