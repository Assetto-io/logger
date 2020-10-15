# logger
logger engine based on zap framework

## Installation

```bash
go get github.com/assetto-io/logger
```

## Usage
In order to use the library you need to import the corresponding package:

```go
import "github.com/assetto-io/logger"
```

## Performing Logging calls
The ``logger`` module provides convenient methods that you can use to perform different logging calls.

### Info

```go
logger.Info("some kind of event")
```

### Error

```go
if err := whoops.MethodWithError; err != nil {
    logger.Error("some kind of event", err)
}
```

### Working with tags

```go
url := "http://test.com"
request := NewCustomRequest()

if err := whoops.PerformRequest(url, request); err != nil {
    logger.Error("some kind of event", err, logger.Field(url, request))
}
```
