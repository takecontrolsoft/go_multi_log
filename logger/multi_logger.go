/* Copyright 2023 Take Control - Software & Infrastructure

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

// # Multi Logs GitHub repository:
//
// https://github.com/takecontrolsoft/go_multi_log
//
// # Package "logger"
//
// This package provides functions for logging
// in multiple loggers (console, file and other).
//
// It logs messages, objects and errors in different log levels:
// "Debug", "Trace", "Info", "Warning", "Error" and "Fatal".
//
// More than one loggers could be registered at the same time.
//
// This package provides implementations of [loggers.ConsoleLogger]
// and [loggers.FileLogger].
//
// Custom loggers could be also implemented using the [loggers.LoggerInterface].
//
// # Take Control - software & infrastructure
//
// The package is created and maintained by "Take Control - software & infrastructure".
//
// Web site: https://takecontrolsoft.eu
package logger

import (
	"sync"

	"github.com/go-errors/errors"
	"github.com/takecontrolsoft/go_multi_log/logger/levels"
	"github.com/takecontrolsoft/go_multi_log/logger/loggers"
)

var lock = &sync.Mutex{}

type multiLog struct {
	registered_loggers map[string]loggers.LoggerInterface
}

var mLogger *multiLog

func getMultiLog() *multiLog {
	if mLogger == nil {
		lock.Lock()
		defer lock.Unlock()
		if mLogger == nil {
			mLogger = &multiLog{
				registered_loggers: map[string]loggers.LoggerInterface{
					"": loggers.NewConsoleLoggerDefault(),
				},
			}
		}
	}

	return mLogger
}

type fnLog func(logger loggers.LoggerInterface, level levels.LogLevel, arg any)
type fnLogF func(logger loggers.LoggerInterface, format string, level levels.LogLevel, args ...interface{})

func _log(logger loggers.LoggerInterface, level levels.LogLevel, arg any) {
	logger.Log(level, arg)
}

func _logF(logger loggers.LoggerInterface, format string, level levels.LogLevel, args ...interface{}) {
	logger.LogF(level, format, args...)
}

func logAll(fn fnLog, level levels.LogLevel, arg any) {
	mLogger = getMultiLog()
	for _, logger := range mLogger.registered_loggers {
		fn(logger, level, arg)
	}
	if level == levels.Fatal {
		panic(arg)
	}
}

func logFAll(fn fnLogF, format string, level levels.LogLevel, args ...interface{}) {
	mLogger = getMultiLog()
	for _, logger := range mLogger.registered_loggers {
		fn(logger, format, level, args...)
	}
}

// Register an instance of an additional logger
// that implements [loggers.LoggerInterface].
func RegisterLogger(key string, logger loggers.LoggerInterface) error {
	if len(key) == 0 {
		return errors.Errorf("Empty key is not allowed for registering loggers.").Err
	}
	mLogger = getMultiLog()
	mLogger.registered_loggers[key] = logger
	return nil
}

// Unregister an instance of logger by key.
func UnregisterLogger(key string) error {
	mLogger = getMultiLog()
	logger := mLogger.registered_loggers[key]
	if logger == nil {
		return errors.Errorf("A logger for given key does not exists.").Err
	}
	delete(mLogger.registered_loggers, key)
	return nil
}

// Return a registered logger instance by key.
func GetLogger(key string) loggers.LoggerInterface {
	mLogger = getMultiLog()
	logger := mLogger.registered_loggers[key]
	return logger
}

// Return the default instance of [loggers.ConsoleLogger].
func DefaultLogger() loggers.LoggerInterface {
	return GetLogger("")
}

// Log object in Debug level.
func Debug(arg any) {
	logAll(_log, levels.Debug, arg)
}

// Log object in Trace level.
func Trace(arg any) {
	logAll(_log, levels.Trace, arg)
}

// Log object in Info level.
func Info(arg any) {
	logAll(_log, levels.Info, arg)
}

// Log object in Warning level.
func Warning(arg any) {
	logAll(_log, levels.Warning, arg)
}

// Log object in Error level.
func Error(arg any) {
	logAll(_log, levels.Error, arg)
}

// Log object in Fatal level and call Panic to exit.
func Fatal(arg any) {
	logAll(_log, levels.Fatal, arg)
}

// Log objects using format string in Debug level.
func DebugF(format string, args ...interface{}) {
	logFAll(_logF, format, levels.Debug, args...)
}

// Log objects using format string in Trace level.
func TraceF(format string, args ...interface{}) {
	logFAll(_logF, format, levels.Trace, args...)
}

// Log objects using format string in Info level.
func InfoF(format string, args ...interface{}) {
	logFAll(_logF, format, levels.Info, args...)
}

// Log objects using format string in Warning level.
func WarningF(format string, args ...interface{}) {
	logFAll(_logF, format, levels.Warning, args...)
}

// Log objects using format string in Error level.
func ErrorF(format string, args ...interface{}) {
	logFAll(_logF, format, levels.Error, args...)
}

// Log objects using format string in Fatal level and call Panic to exit.
func FatalF(format string, args ...interface{}) {
	logFAll(_logF, format, levels.Fatal, args...)
}
