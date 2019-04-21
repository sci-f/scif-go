// Copyright (C) 2019 Vanessa Sochat.

// This program is free software: you can redistribute it and/or modify it
// under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or (at your
// option) any later version.

// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.  See the GNU Affero General Public
// License for more details.

// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package logger

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"runtime"
	"strconv"
	"strings"
)

type messageLevel int

const (
	fatal messageLevel = iota - 4 // fatal    : -4
	error                         // error    : -3
	warn                          // warn     : -2
	log                           // log      : -1
	_                             // SKIP     : 0
	info                          // info     : 1
	debug                         // debug    : 5
)

// messageLabels describe the levels above
var messageLabels = map[messageLevel]string{
	fatal: "FATAL",
	error: "ERROR",
	warn:  "WARNING",
	log:   "LOG",
	info:  "INFO",
	debug: "DEBUG",
}

// messageColors describe the levels above (and make them pretty!)
var messageColors = map[messageLevel]string{
	fatal: "\x1b[31m",
	error: "\x1b[31m",
	warn:  "\x1b[33m",
	info:  "\x1b[34m",
	log:   "\x1b[30m",
	debug: "\x1b[32m",
}

// String ensures that we can print a messageLevel
func (l messageLevel) String() string {
	str, ok := messageLabels[l]
	if !ok {
		str = "????"
	}
	return str
}

var colorReset = "\x1b[0m"
var loggerLevel messageLevel

func init() {
	_level, ok := os.LookupEnv("SCIF_MESSAGELEVEL")
	if !ok {
		loggerLevel = debug
	} else {
		_levelint, err := strconv.Atoi(_level)
		if err != nil {
			loggerLevel = debug
		} else {
			loggerLevel = messageLevel(_levelint)
		}
	}
}

// prefix will add the prefix (color, name) to message
func prefix(level messageLevel) string {

	// Default to no color
	messageColor, ok := messageColors[level]
	if !ok {
		messageColor = "\x1b[0m"
	}

	// This section builds and returns the prefix for levels < debug
	if loggerLevel < debug {
		return fmt.Sprintf("%s%-8s%s ", messageColor, level.String()+":", colorReset)
	}

	pc, _, _, ok := runtime.Caller(3)
	details := runtime.FuncForPC(pc)

	var funcName string
	if ok && details == nil {
		fmt.Printf("Unable to get details of calling function\n")
		funcName = "UNKNOWN CALLING FUNC"
	} else {
		funcNameSplit := strings.Split(details.Name(), ".")
		funcName = funcNameSplit[len(funcNameSplit)-1] + "()"
	}

	return fmt.Sprintf("%s%-8s%-19s%-30s", messageColor, level, colorReset, funcName)
}

// writef is the shared function to handle printing the error
func writef(level messageLevel, format string, a ...interface{}) {
	if loggerLevel < level {
		return
	}

	message := fmt.Sprintf(format, a...)
	message = strings.TrimSuffix(message, "\n")

	fmt.Fprintf(os.Stderr, "%s%s\n", prefix(level), message)
}

// Debugf means that DEBUG level message is sent to the log.
func Debugf(format string, a ...interface{}) {
	writef(debug, format, a...)
}

// Errorf means that ERROR level message is sent to the log and the error returned
func Errorf(format string, a ...interface{}) {
	writef(error, format, a...)
}

// Exitf throws a Fatal error and exits!
func Exitf(format string, a ...interface{}) {
	writef(fatal, format, a...)
	os.Exit(255)
}

// Infof sends an INFO level message to the log.
func Infof(format string, a ...interface{}) {
	writef(info, format, a...)
}

// Warningf sends a WARNING level message to the log.
func Warningf(format string, a ...interface{}) {
	writef(warn, format, a...)
}

// SetLevel explicitly sets the loggerLevel
func SetLevel(l int) {
	loggerLevel = messageLevel(l)
}

// DisableColor for the logger (is used by the command line client)
func DisableColor() {
	messageColors = map[messageLevel]string{
		fatal: "",
		error: "",
		warn:  "",
		info:  "",
		log:   "",
		debug: "",
	}
	colorReset = ""
}

// GetLevel returns the current log level as integer
func GetLevel() int {
	return int(loggerLevel)
}

// GetMessageLevel shows environment setting for logger level.
func GetMessageLevel() string {
	return fmt.Sprintf("SCIF_MESSAGELEVEL=%d", loggerLevel)
}

// Writer returns an object to pass to external logging utilities. if level <= -1.
func Writer() io.Writer {
	if loggerLevel <= -1 {
		// returning this ignores output.
		return ioutil.Discard
	}
	return os.Stderr
}
