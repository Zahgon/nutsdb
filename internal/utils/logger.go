package utils

import "log"

var (
	printLoggerInstance ILogger = defaultLogger()
)

type ILogger interface {
	// Printf formats according to a format specifier and writes to the logger.
	// Arguments are handled in the manner of fmt.Printf.
	Printf(string, ...any)
}

type defaultPrintLogger struct {
	l *log.Logger
}

func (dpl *defaultPrintLogger) Printf(fmt string, args ...any) { _ = "STUB: not implemented"; return }

func defaultLogger() ILogger { _ = "STUB: not implemented"; return *new(ILogger) }

func SetLogger(logger ILogger) { _ = "STUB: not implemented"; return }

func GetLogger() ILogger { _ = "STUB: not implemented"; return *new(ILogger) }
