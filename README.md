
<img src="https://takecontrolsoft.eu/wp-content/uploads/2023/11/TakeControlTransparentGreenLogo-1.png" alt="Sync Device by Take Control - software & infrastructure" width="30%">

[![Web site](https://img.shields.io/badge/Web_site-takecontrolsoft.eu-pink)](https://takecontrolsoft.eu/)
[![Linked in](https://img.shields.io/badge/Linked_In-page-blue)](https://www.linkedin.com/company/take-control-si/)
[![Documentation](https://pkg.go.dev/badge/github.com/takecontrolsoft/go_multi_log.svg)](https://pkg.go.dev/github.com/takecontrolsoft/go_multi_log)
[![Release](https://img.shields.io/github/v/release/takecontrolsoft/go_multi_log.svg)](https://github.com/takecontrolsoft/go_multi_log/releases/latest)
[![License](https://img.shields.io/badge/License-Apache-purple)](https://www.apache.org/licenses/LICENSE-2.0)

# Multi Logs package (go_multi_log)
Multi Logs package "go_multi_log" is a Go package that provides logging with multiple loggers (console, file and custom loggers) with different log levels.

# [Multi Loggers Types](#multi-loggers-types)
* [Console logger](#console-logger) (defaults)
* [File logger](#file-logger)
* [Custom logger](#custom-logger)

# Get started
* Default log level is `Info`, which means that only `Info`, `Warning`, `Error` and `Fatal` messages will be logged, but `Debug` and `Trace` messages will be skipped.
* The function `Fatal` logs the messages and then calls `panic`. It can be used for fatal errors and it will ensure storing the log into a file before closing the application.

## Imports
```go
import (	
    "github.com/takecontrolsoft/go_multi_log/logger"
    "github.com/takecontrolsoft/go_multi_log/logger/levels"
    "github.com/takecontrolsoft/go_multi_log/logger/loggers"
)
````	
## Usage
### Log messages:

```go
    logger.Info("Test info log message")
    logger.Warning("Test warning log message")
    logger.Error("Test error log message")
    logger.Fatal("Test log fatal message")
```

### Log error object:
```go
err := callFunction()
if err!=nil{
    logger.Error(err)
}
```
### Log any object

```go	
type person struct{ Name string }
person := person{Name: "Michael"}
logger.Info(person)	
```

### Log formatted message
Using formatting functions `DebugF`, `TraceF`, `InfoF`, `WarningF`, `ErrorF`, `FatalF`

```go	
type person struct{ Name string }
type car struct{ Year string }

person := person{Name: "Michael"}
car := car{Year: "2020"}

logger.InfoF("Person: %v, Car: %v", person, car)	
```

### Change log level
To change the default log level use `SetLevel(levels.All)`. This will cause all the messages for levels greater or equal to the new level also to be logged. The level is changed for the whole application. Avoid changing the level inside Goroutine (go lightweight thread).
       
```go
currentLevel:= logger.DefaultLogger().GetLevel()
logger.DefaultLogger().SetLevel(levels.All)
logger.Debug("Test debug log message")
logger.Trace("Test trace log message")
```
### Stop and Start logging
No messages will be logged after calling `Stop` function.
Logging could be resumed with calling `Start` function.
In this example only "Message 2" will be logged.
```go
logger.DefaultLogger().Stop()
logger.Info("Message 1")
logger.DefaultLogger().Start()
logger.Info("Message 2")		
```

## Multi Loggers Types

### Manage loggers
Use the following functions to **register**, **unregister** or **get** loggers by key. One **default logger** always exists and can not be unregistered, but can be stopped.

```go
key:="new_logger_key" 
// where the key must be unique, because more than one instances of the same logger type can be registered.

logger.RegisterLogger(key, logger) 
// where `logger` implements `loggers.LoggerInterface` and can support different log levels and log destinations.

logger.GetLogger(key) 
// will return an instance of the logger by key.

logger.UnregisterLogger(key) 
// will delete a specific logger from the collection with registered loggers.

logger.DefaultLogger() 
// will return an instance of the default logger of type `ConsoleLogger`
```

### Console logger

#### Default console log:
A simple `ConsoleLogger` is set by default and it can be obtained from `logger.DefaultLogger()`. This logger can not be unregistered, but it can be stopped and resumed.
```go
logger.DefaultLogger().Stop()
```
#### Custom console log:
Custom console log can be registered for example to log only Debug messages using new formatting string. See the example:

```go
c := loggers.NewConsoleLogger(levels.Debug, "***debug:'%s'")
_, err := logger.RegisterLogger("debug_log_key", c)
```

### File logger
`FileLogger` can be added as an additional logger to prints the messages to files. 
#### Use `NewFileLoggerDefault` to initialize the file logger with the default settings.
   * Default LogLevel is levels.Info.
   * Default FileOptions are used:        
        * Directory: current executable directory.
        * FilePrefix:  "mLog".
        * FileExtension:  ".log".
```go
f := loggers.NewFileLoggerDefault()
_, err := logger.RegisterLogger("file_logger_key", f)
```

#### Use `NewFileLogger` to initialize the file logger with the customized settings.
```go
level := levels.Error
format := "***error:'%s'"
fileOptions := loggers.FileOptions{
    Directory:     "./",
    FilePrefix:    "tLog",
    FileExtension: ".txt",
}

f := loggers.NewFileLogger(level, format, fileOptions)
_, err := logger.RegisterLogger("txt_file_key", f)
	
```
### Custom logger
Custom loggers implementations can be easily added by implementing the interface `loggers.LoggerInterface` or deriving the base class `loggers.LoggerType`, which already implements most of the function. 
#### Implementation example: 
* Create a new file `json_logger.go`.
* Import `go_multi_log` package.
```go
import (
	"github.com/takecontrolsoft/go_multi_log/logger/levels"
	"github.com/takecontrolsoft/go_multi_log/logger/loggers"
)
```
* Create `JsonLogger` type.
```go
//json_logger.go

// Define new type
type JsonLogger struct {
	loggers.LoggerType
}

// Create function for initializing
func NewJsonLogger(level levels.LogLevel) *JsonLogger {
	return &JsonLogger{
		LoggerType: loggers.LoggerType{Level: level},
	}
}

// Implement Log function
func (logger *JsonLogger) Log(level levels.LogLevel, arg any) {
	if logger.IsLogAllowed(level) {
		// Define your format here
		format := fmt.Sprintf("{%s: %s},", strings.ToLower(level.String()), "%v")
		logger.LogF(level, format, arg)
	}
}

// Implement LogF function
func (logger *JsonLogger) LogF(level levels.LogLevel, format string, args ...interface{}) {
	if logger.IsLogAllowed(level) {
		// You can send the messages in Json format to an external service here
		fmt.Printf(format, args...)
	}
}

```
#### Usage example: 

```go
 err := logger.RegisterLogger("json_key", NewJsonLogger(levels.Info))
 // ... error check...
 logger.Error("Test log error message")
 logger.Info("Test log info message")
```

# Build source
* Go version 1.21 is required.
* Create and go to folder `go_multi_log`.
* run the following commands:
    * `git clone https://github.com/takecontrolsoft/go_multi_log.git`
    * `go build -v ./...`
    * `go test -v ./...`

# Contribute
See [CONTRIBUTING.md](CONTRIBUTING.md) for instructions about building the source.

# License
Multi Logs package ("go_multi_log") is published under the Apache License 2.0.

***The package is created and maintained by **["Take Control - software & infrastructure"](https://takecontrolsoft.eu/)*****

***The "Go" name and logo are trademarks owned by Google.***




