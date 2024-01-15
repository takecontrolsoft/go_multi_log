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

// # Multiple Logs GitHub repository:
//
// https://github.com/takecontrolsoft/go_multi_log
//
// # Package "loggers"
//
// This Package provides different implementations for loggers,
// which are possible to be registered.
//
// The package supports:
//   - [loggers.ConsoleLogger], which logs the messages to the console
//   - [loggers.FileLogger], which logs the messages to files separated by goroutines.
//
// The common interface [loggers.LoggerInterface]
// makes it possible this package to be extended by implementing
// additional custom loggers for logging in json, xml and other formats,
// as well as sending the logs to external services.
//
// # Take Control - software & infrastructure
//
// The package is created and maintained by "Take Control - software & infrastructure".
//
// Web site: https://takecontrolsoft.eu
package loggers

import (
	"fmt"
	"log"
	"strings"

	"github.com/takecontrolsoft/go_multi_log/logger/levels"
)

// [LoggerInterface] describes all the methods required to be implemented by the loggers.
// Each logger, which implements this interface could be registered
// to log messages and objects.
type LoggerInterface interface {
	Log(level levels.LogLevel, arg any)
	LogF(level levels.LogLevel, format string, args ...interface{})
	GetLevel() levels.LogLevel
	SetLevel(level levels.LogLevel)
	Start()
	Stop()
}

// [LoggerType] provides base implementation of [loggers.LoggerInterface]
// and can be reused when extending the package with adding new
// loggers implementations.
type LoggerType struct {
	LoggerInterface
	Level  levels.LogLevel
	Format string

	isStopped bool
}

// Reports if the log message will be printed based on the
// log level. The log level can be set for a specific logger
// using [loggers.LoggerType.SetLevel].
// [loggers.LoggerType.IsLogAllowed] returns false if the logger
// is stopped using [loggers.LoggerType.Stop] function.
func (logger *LoggerType) IsLogAllowed(level levels.LogLevel) bool {
	return !logger.isStopped && level >= logger.Level
}

// Reports the log level for this logger.
func (logger *LoggerType) GetLevel() levels.LogLevel {
	return logger.Level
}

// Sets the log level for this logger.
func (logger *LoggerType) SetLevel(level levels.LogLevel) {
	logger.Level = level
}

// Resumes printing logs by this logger.
func (logger *LoggerType) Start() {
	logger.isStopped = false
}

// Stops printing logs by this logger.
func (logger *LoggerType) Stop() {
	logger.isStopped = true
}

func (logger *LoggerType) multi_log(level levels.LogLevel, arg any) {
	var f string
	if len(logger.Format) > 0 {
		f = logger.Format
	} else {
		f = fmt.Sprintf("%s: [%s]", strings.ToUpper(level.String()), "%v")
	}
	logger.multi_logF(level, f, arg)
}

func (logger *LoggerType) multi_logF(level levels.LogLevel, format string, args ...interface{}) {
	log.Printf(format, args...)
}
