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
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/takecontrolsoft/go_multi_log/logger/levels"
	"github.com/timandy/routine"
)

// A FileLogger is safe for concurrent use by multiple goroutines
type FileLogger struct {
	LoggerType
	FileOptions
}

// Represent a set of file options,
// which are used when the log file name is generated.
//   - Directory - an absolute or relative path to log files location where the process has write access.
//   - FilePrefix - should be a short string of symbols allowed for OS file names.
//   - FileExtension - should starts with ".".
type FileOptions struct {
	Directory, FilePrefix, FileExtension string
}

// Returns an instance of [FileLogger] with
// default log level "Info".
// Default [FileOptions] are used:
//   - Directory: current executable directory.
//   - FilePrefix:  "sync_server"
//   - FileExtension:  ".log"
func NewFileLoggerDefault() *FileLogger {
	return &FileLogger{
		LoggerType:  LoggerType{Level: levels.Info},
		FileOptions: FileOptions{FilePrefix: "mLog", FileExtension: ".log"},
	}
}

// Returns an instance of [FileLogger] with
// given log level and format string defined by the caller.
// [FileOptions] are not required, but could be used for setting the
// file name prefix and file extension as well as the path location,
// where the log files to be stored.
func NewFileLogger(level levels.LogLevel, format string, options FileOptions) *FileLogger {
	return &FileLogger{
		LoggerType:  LoggerType{Level: level, Format: format},
		FileOptions: options,
	}
}

// Prints the message or the object "arg" into files (named with goroutine id).
// If there is no format set when initializing this [FileLogger],
// a default format is used: {time} {log level}: [{message}]
func (logger *FileLogger) Log(level levels.LogLevel, arg any) {
	if logger.IsLogAllowed(level) {
		fLog := setFileLog(logger)
		defer fLog.Close()
		logger.multi_log(level, arg)
	}
}

// Prints one or more objects "args" into files (named with goroutine id)
// as a message formatted using the given format string by the caller.
func (logger *FileLogger) LogF(level levels.LogLevel, format string, args ...interface{}) {
	if logger.IsLogAllowed(level) {
		fLog := setFileLog(logger)
		defer fLog.Close()
		logger.multi_logF(level, format, args...)
	}
}

func setFileLog(logger *FileLogger) *os.File {
	goid := routine.Goid()
	fName := fmt.Sprintf("%s_%d_%d%s", logger.FilePrefix, os.Getpid(), goid, logger.FileExtension)
	logFile := filepath.Join(logger.Directory, fName)
	fLog, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	log.SetOutput(fLog)
	return fLog
}
