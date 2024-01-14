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
	"strings"

	"github.com/takecontrolsoft/go_multi_log/logger/levels"
)

type LoggerInterface interface {
	Log(level int, arg any)
	LogF(level int, format string, args ...interface{})
	GetLevel() int
	SetLevel(level int)
	Start()
	Stop()
}

type LoggerType struct {
	LoggerInterface
	Level  int
	Format string

	isStopped bool
}

func (logger *LoggerType) IsLogAllowed(level int) bool {
	return !logger.isStopped && level >= logger.Level
}

func (logger *LoggerType) GetLevel() int {
	return logger.Level
}

func (logger *LoggerType) SetLevel(level int) {
	logger.Level = level
}

func (logger *LoggerType) Start() {
	logger.isStopped = false
}

func (logger *LoggerType) Stop() {
	logger.isStopped = true
}

func (logger *LoggerType) multi_log(level int, arg any) {
	var f string
	if len(logger.Format) > 0 {
		f = logger.Format
	} else {
		f = fmt.Sprintf("%s: [%s]", strings.ToUpper(GetLogLevelName(level)), "%v")
	}
	logger.multi_logF(level, f, arg)
}

func (logger *LoggerType) multi_logF(level int, format string, args ...interface{}) {
	log.Printf(format, args...)
}

func GetLogLevelName(level int) string {
	switch logLevel := level; logLevel {
	case levels.DebugLevel:
		return "Debug"
	case levels.TraceLevel:
		return "Trace"
	case levels.InfoLevel:
		return "Info"
	case levels.WarningLevel:
		return "Warning"
	case levels.ErrorLevel:
		return "Error"
	case levels.FatalLevel:
		return "Fatal"
	default:
		return "Unknown"
	}
}
