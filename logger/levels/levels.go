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
// # Package "levels"
//
// This package provides constants for different log levels.
//
// Based on the levels the logged messages could be filtered.
//
// Multiple loggers could be configured to log messages from different levels.
//
// # Take Control - software & infrastructure
//
// The package is created and maintained by "Take Control - software & infrastructure".
//
// Web site: https://takecontrolsoft.eu
package levels

// LogLevel type represents the supported log levels
type LogLevel int

const (
	All     LogLevel = 0
	Debug   LogLevel = 1
	Trace   LogLevel = 2
	Info    LogLevel = 3
	Warning LogLevel = 4
	Error   LogLevel = 5
	Fatal   LogLevel = 6
)

// Converts to string the name of the LogLevel value
func (level LogLevel) String() string {
	switch level {
	case Debug:
		return "Debug"
	case Trace:
		return "Trace"
	case Info:
		return "Info"
	case Warning:
		return "Warning"
	case Error:
		return "Error"
	case Fatal:
		return "Fatal"
	default:
		return "Unknown"
	}
}
