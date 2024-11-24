# Gogger

Golang + Logger is a simple and efficient logging library for Go (Golang) that provides easy-to-use logging functionalities.

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
    log, err := gogger.NewLogger("logfile.log", gogger.INFO)
    if err != nil {
        panic(err)
    }

    log.Info().Str("TaskType", "mytask").Int("EventCount", len(numberOfEvens)).Msg("This is an info message")
    log.Warning.Str("TaskType", "mytask").Int("EventCount", len(numberOfEvens)).Msg("This is an warning message")
    log.Error.Str("TaskType", "mytask").Int("EventCount", len(numberOfEvens)).Msg("This is an error message")
    log.Debug.Str("TaskType", "mytask").Int("EventCount", len(numberOfEvens)).Msg("This is an debug message")
}
```

## Configuration

You can customize the logger by setting different options in the configuration file:

```go
func main() {
    config := gogger.Config{
        Output: "stdout",
        Format: "json",
        Level:  gogger.INFO,
    }

    log, err := gogger.NewLoggerWithConfig(config)
    if err != nil {
        panic(err)
    }
}
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request on GitHub.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgements

Thanks to the Go community for their support and contributions. ❤️
