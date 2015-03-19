package logger

import "fmt"

/**
 * Enable debugging
 * @type {Boolean}
 */
var isDebug = false

func info(key string, args ...interface{}) {
	p := fmt.Sprintf(key, args)
	fmt.Println(p)
}
