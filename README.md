# go_multi_log
Go package go_multi_log that provides logging in multiple loggers (console, file and other) with log levels.


Default log level is `Info`, which means that only `Info`, `Warning`, `Error` and `Fatal` messages will be logged.
The function `Fatal` logs the messages and then calls `panic`. It can be used for fatal errors and it will ensure storing the log into a file before closing the application.

# Imports
```go
import (	
    "github.com/takecontrolsoft/go_multi_log/logger"
	"github.com/takecontrolsoft/go_multi_log/logger/levels"
	"github.com/takecontrolsoft/go_multi_log/logger/loggers"
)
````	
# Usage
## Log messages:

```go
    logger.Info("Test info log message")
    logger.Warning("Test warning log message")
    logger.Error("Test error log message")
    logger.Fatal("Test log fatal message") // The function Fatal logs the messages and then calls Panic.
```

## Log error object:
```go
err := callFunction()
logger.Error(err)
```
## Log any object

```go	
type person struct{ Name string }
person := person{Name: "Michael"}
logger.Info(person)	
```

## Log formatted message
Using formatting functions `DebugF`, `TraceF`, `InfoF`, `WarningF`, `ErrorF`, `FatalF`

```go	
type person struct{ Name string }
type car struct{ Year string }

person := person{Name: "Michael"}
car := car{Year: "2020"}

logger.InfoF("Person: %v, Car: %v", person, car)	
```

## Change log level
To change the default log level use `SetLevel(levels.All)`. This will cause all the messages for levels greater or equal to the new level also to be logged. The level is changed for the whole application. Avoid changing the level inside Goroutine (go lightweight thread).
       
```go
currentLevel:= logger.DefaultLogger().GetLevel()
logger.DefaultLogger().SetLevel(levels.All)
logger.Debug("Test debug log message")
logger.Trace("Test trace log message")
```
## Stop and Start logging
No messages will be logged after calling `Stop` function.
Logging could be resumed with calling `Start` function.
In this example only "Message 2" will be logged.
```go
logger.DefaultLogger().Stop()
logger.Info("Message 1")
logger.DefaultLogger().Start()
logger.Info("Message 2")		
```

# Logger types

## Console logger
## File logger
## Implement custom logger
