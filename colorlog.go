package clog

//learning:: if you need to export method, start it with capital letter

import "fmt"
import "github.com/mgutz/ansi"

/**
 * Enable debugging
 * @type {Boolean}
 */
var IsDebug = false

/**
 * Internal method to make code dry
 * @param {string} color string       color of string
 * @param {string} key   string       message with placeholder
 * @param {interface} args interface  veriadic params
 */
func colorIt(color string, key string, args ...interface{}) {
	p := fmt.Sprintf(key, args...)
	msg := ansi.Color(p, color)
	fmt.Println(msg)
}

/**
 * Info method
 *
 * @param {string} key   string       message with placeholder
 * @param {interface} args interface  veriadic params
 */
func Info(key string, args ...interface{}) {
	colorIt("green", key, args...)
}

/**
 * Debug method
 *
 * @param {string} key   string       message with placeholder
 * @param {interface} args interface  veriadic params
 */
func Debug(key string, args ...interface{}) {
	if IsDebug {
		colorIt("blue", key, args...)
	}
}

/**
 * Warn method
 *
 * @param {string} key   string       message with placeholder
 * @param {interface} args interface  veriadic params
 */
func Warn(key string, args ...interface{}) {
	colorIt("yellow", key, args...)
}

/**
 * Error method
 *
 * @param {string} key   string       message with placeholder
 * @param {interface} args interface  veriadic params
 */
func Error(key string, args ...interface{}) {
	colorIt("red", key, args...)
}
