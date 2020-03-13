# GoLog Library for Logging in Go Applications

This is a tiny package which helps to log and print messages

## Installation

As a library

```shell
go get github.com/hop-/golog
```

## Usage

This library allows to log into file as well as to standard outputs

Code example:
```go
import log "github.com/hop-/golog"

func main() {
    log.InitFile("log.txt")

    log.Info("This application will log everithing to log.txt as well as to standard outputs")
    log.Warning("Some warining")

    if err != nil {
        log.Errorf("Some error happend: '%s'", err.Error())
    }

    log.Fatal("Some fatal error happend")
    // application will exit on fatal
}
```
