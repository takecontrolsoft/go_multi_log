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

// Package github.com/takecontrolsoft/go_multi_log/logger
// provides logging in multiple loggers (console, file and other)
// It logs messages, objects and errors in different levels:
// debug, trace, info, warning, error, fatal.
// Multiple loggers could be registered.
// loggers.ConsoleLogger and loggers.FileLogger are provided by this package.
// Custom loggers could be implemented using the loggers.LoggerInterface.
//
// (Logging a fatal message will close the application.)
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

type fnLog func(logger loggers.LoggerInterface, level int, arg any)
type fnLogF func(logger loggers.LoggerInterface, format string, level int, args ...interface{})

func _log(logger loggers.LoggerInterface, level int, arg any) {
	logger.Log(level, arg)
}

func _logF(logger loggers.LoggerInterface, format string, level int, args ...interface{}) {
	logger.LogF(level, format, args...)
}

func logAll(fn fnLog, level int, arg any) {
	mLogger = getMultiLog()
	for _, logger := range mLogger.registered_loggers {
		fn(logger, level, arg)
	}
	if level == levels.FatalLevel {
		panic(arg)
	}
}

func logFAll(fn fnLogF, format string, level int, args ...interface{}) {
	mLogger = getMultiLog()
	for _, logger := range mLogger.registered_loggers {
		fn(logger, format, level, args...)
	}
}

func RegisterLogger(key string, logger loggers.LoggerInterface) (loggers.LoggerInterface, error) {
	if len(key) == 0 {
		return logger, errors.Errorf("Empty key is not allowed for registering loggers.").Err
	}
	mLogger = getMultiLog()
	mLogger.registered_loggers[key] = logger
	return logger, nil
}

func UnregisterLogger(key string) (loggers.LoggerInterface, error) {
	mLogger = getMultiLog()
	logger := mLogger.registered_loggers[key]
	if logger == nil {
		return nil, errors.Errorf("A logger for given key does not exists.").Err
	}
	delete(mLogger.registered_loggers, key)
	return logger, nil
}

func GetLogger(key string) loggers.LoggerInterface {
	mLogger = getMultiLog()
	logger := mLogger.registered_loggers[key]
	return logger
}

func DefaultLogger() loggers.LoggerInterface {
	return GetLogger("")
}

func Debug(arg any) {
	logAll(_log, levels.DebugLevel, arg)
}

func Trace(arg any) {
	logAll(_log, levels.TraceLevel, arg)
}

func Info(arg any) {
	logAll(_log, levels.InfoLevel, arg)
}

func Warning(arg any) {
	logAll(_log, levels.WarningLevel, arg)
}

func Error(arg any) {
	logAll(_log, levels.ErrorLevel, arg)
}

func Fatal(arg any) {
	logAll(_log, levels.FatalLevel, arg)
}

func DebugF(format string, args ...interface{}) {
	logFAll(_logF, format, levels.DebugLevel, args...)
}

func TraceF(format string, args ...interface{}) {
	logFAll(_logF, format, levels.TraceLevel, args...)
}

func InfoF(format string, args ...interface{}) {
	logFAll(_logF, format, levels.InfoLevel, args...)
}

func WarningF(format string, args ...interface{}) {
	logFAll(_logF, format, levels.WarningLevel, args...)
}

func ErrorF(format string, args ...interface{}) {
	logFAll(_logF, format, levels.ErrorLevel, args...)
}

func FatalF(format string, args ...interface{}) {
	logFAll(_logF, format, levels.FatalLevel, args...)
}
