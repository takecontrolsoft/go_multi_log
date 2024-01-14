/* Copyright 2024 Take Control - Software & Infrastructure

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package loggers

import (
	"log"
	"os"

	"github.com/takecontrolsoft/go_multi_log/logger/levels"
)

// [ConsoleLogger] type represents the logger that print
// the messages to the standard output [os.Stdout].
type ConsoleLogger struct {
	LoggerType
}

// Returns an instance of [ConsoleLogger] with
// default log level "Info".
func NewConsoleLoggerDefault() *ConsoleLogger {
	return &ConsoleLogger{
		LoggerType: LoggerType{Level: levels.Info},
	}
}

// Returns an instance of [ConsoleLogger] with
// given log level and format string defined by the caller.
func NewConsoleLogger(level levels.LogLevel, format string) *ConsoleLogger {
	return &ConsoleLogger{
		LoggerType: LoggerType{Level: level, Format: format},
	}
}

// Prints the message or the object "arg" into the console.
// If there is no format set when initializing this [ConsoleLogger],
// a default format is used: {time} {log level}: [{message}]
func (logger *ConsoleLogger) Log(level levels.LogLevel, arg any) {
	if logger.IsLogAllowed(level) {
		log.SetOutput(os.Stdout)
		logger.multi_log(level, arg)
	}
}

// Prints one or more objects "args" into the console
// as a message formatted using the given format string
// by the caller.
func (logger *ConsoleLogger) LogF(level levels.LogLevel, format string, args ...interface{}) {
	if logger.IsLogAllowed(level) {
		log.SetOutput(os.Stdout)
		logger.multi_logF(level, format, args...)
	}
}
