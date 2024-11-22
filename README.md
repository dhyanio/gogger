# Gogger

Golang + Logger, is a simple and efficient logging library for Go (Golang) that provides easy-to-use logging functionalities.

## Features

- Simple API for logging at different levels (Info, Warning, Error, Debug)
- Customizable log format
- Support for logging to different outputs (console, file, etc.)
- Lightweight and efficient

## Installation

To install Gogger, use `go get`:

```sh
go get github.com/dhyanio/gogger
```

## Usage

Here's a basic example of how to use Gogger:

```go
package main

import (
    "github.com/dhyanio/gogger"
)

func main() {
    logger, err := gogger.NewLogger("logfile.log", gogger.DebugLevel)
    if err != nil {
        panic(err)
    }

    logger.Info("This is an info message")
    logger.Warning("This is a warning message")
    logger.Error("This is an error message")
    logger.Debug("This is a debug message")
}
```

## Configuration

You can customize the logger by setting different options in the configuration file:

```go
package main

import (
    "github.com/dhyanio/gogger"
)

func main() {
    config := gogger.Config{
        Output: "stdout",
        Format: "json",
        Level:  gogger.DebugLevel,
    }

    logger, err := gogger.NewLoggerWithConfig(config)
    if err != nil {
        panic(err)
    }

    logger.Info("This is an info message")
    logger.Warning("This is a warning message")
    logger.Error("This is an error message")
    logger.Debug("This is a debug message")
}
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request on GitHub.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgements

Thanks to the Go community for their support and contributions.
