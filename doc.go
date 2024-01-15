// # Multiple Logs package (go_multi_log)
//
// Multiple Logs package "go_multi_log" is Go package
// that provides logging in multiple loggers (console, file and other) with log levels.
//
// # GitHub repository:
//
// https://github.com/takecontrolsoft/go_multi_log
//
// # Multiple Logger Types:
//
//   - [loggers.ConsoleLogger]
//   - [loggers.FileLogger]
//   - Custom logger
//
// # Get started
//
//   - Default log level is "Info", which means that only "Info", "Warning", "Error" and "Fatal" messages will be logged, but "Debug" and "Trace" messages will be skipped.
//   - The function "Fatal" logs the messages and then calls "panic". It can be used for fatal errors and it will ensure storing the log into a file before closing the application.
//
// # Imports
//
//	import (
//
//	"github.com/takecontrolsoft/go_multi_log/logger"
//	"github.com/takecontrolsoft/go_multi_log/logger/levels"
//	"github.com/takecontrolsoft/go_multi_log/logger/loggers"
//
//	)
//
// # Usage
//
// - Log messages:
//
//	logger.Info("Test info log message")
//	logger.Warning("Test warning log message")
//	logger.Error("Test error log message")
//	logger.Fatal("Test log fatal message") // The function Fatal logs the messages and then calls Panic.
//
// - Log error object:
//
//	err := callFunction()
//	if err!=nil{
//		logger.Error(err)
//	}
//
// - Log any object
//
//	type person struct{ Name string }
//	person := person{Name: "Michael"}
//	logger.Info(person)
//
// - Log formatted message
//
// Using formatting functions "DebugF", "TraceF", "InfoF", "WarningF", "ErrorF", "FatalF"
//
//	type person struct{ Name string }
//	type car struct{ Year string }
//	person := person{Name: "Michael"}
//	car := car{Year: "2020"}
//	logger.InfoF("Person: %v, Car: %v", person, car)
//
// - Change log level
//
// To change the default log level use "SetLevel(levels.All)".
// This will cause all the messages for levels greater or equal to the new level
// also to be logged. The level is changed for the whole application.
// Avoid changing the level inside Goroutine (go lightweight thread).
//
//	currentLevel:= logger.DefaultLogger().GetLevel()
//	logger.DefaultLogger().SetLevel(levels.All)
//	logger.Debug("Test debug log message")
//	logger.Trace("Test trace log message")
//
// - Stop and Start logging
//
// No messages will be logged after calling "Stop" function.
// Logging could be resumed with calling "Start" function.
// In this example only "Message 2" will be logged.
//
//	logger.DefaultLogger().Stop()
//	logger.Info("Message 1")
//	logger.DefaultLogger().Start()
//	logger.Info("Message 2")
//
// # Manage Multiple Loggers Types
//
// Use the following functions to register, unregister or get loggers by key. One default logger always exists and can not be unregistered, but can be stopped.
//
//	key:="new_logger_key" // where the key must be unique, because more than one instances of the same logger type can be registered.
//
//	logger.RegisterLogger(key, logger) // where "logger" implements "loggers.LoggerInterface" and can support different log levels and log destinations.
//
//	logger.GetLogger(key) // will return an instance of the logger by key.
//
//	logger.UnregisterLogger(key) // will delete a specific logger from the collection with registered loggers.
//
//	logger.DefaultLogger() // will return an instance of the default logger of type "ConsoleLogger"
//
// # Console logger
//
// - Default console log:
//
// A simple "ConsoleLogger" is set by default and it can be obtained from "logger.DefaultLogger()". This logger can not be unregistered, but it can be stopped and resumed.
//
//	logger.DefaultLogger().Stop()
//
// - Custom console log:
//
// Custom console log can be registered for example to log only Debug messages using new formatting string. See the example:
//
//	c := loggers.NewConsoleLogger(levels.Debug, "***debug:'%s'")
//	_, err := logger.RegisterLogger("debug_log_key", c)
//
// # File logger
//
// "FileLogger" can be added as an additional logger to prints the messages to files.
//
// - Use "NewFileLoggerDefault" to initialize the file logger with the default settings.
//
//   - Default [LogLevel] is levels.Info.
//
//   - Default [FileOptions] are used:
//
//   - Directory: current executable directory.
//
//   - FilePrefix:  "mLog".
//
//   - FileExtension:  ".log".
//
// Example:
//
//	f := loggers.NewFileLoggerDefault()
//	_, err := logger.RegisterLogger("file_logger_key", f)
//
// - Use "NewFileLogger" to initialize the file logger with the customized settings.
//
//	level := levels.Error
//	format := "***error:'%s'"
//
//	fileOptions := loggers.FileOptions{
//	    Directory:     "./",
//	    FilePrefix:    "mLog",
//	    FileExtension: ".txt",
//	}
//
//	f := loggers.NewFileLogger(level, format, fileOptions)
//	_, err := logger.RegisterLogger("txt_file_key", f)
//
// # Custom loggers
//
// Custom loggers implementations can be easily added by implementing
// the interface [loggers.LoggerInterface] or deriving the base class
// [loggers.LoggerType], which already implements most of the function.
//
// # Examples:
//
// https://github.com/takecontrolsoft/go_multi_log?tab=readme-ov-file#custom-logger
//
// # Take Control - software & infrastructure
//
// The package is created and maintained by "Take Control - software & infrastructure".
//
// Web site: https://takecontrolsoft.eu
package go_multi_log
