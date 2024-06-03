package main

import (
	"fmt"
	"sync"
)

type MyLogger struct {
	logLevel int
}

func (l *MyLogger) SetLogLevel(level int) {
	l.logLevel = level
}

func (l *MyLogger) Log(message string) {
	fmt.Println(message)
}

var logger *MyLogger

var once sync.Once

func GetLoggerInstance() *MyLogger {
	// This is not concurrency safe
	// if logger == nil {
	// 	fmt.Println("Creating Logger Instance")
	// 	logger = &MyLogger{}
	// }

	once.Do(func() {
		fmt.Println("Creating Logger Instance")
		logger = &MyLogger{}
	})

	fmt.Println("Returning Logger Instance")
	return logger
}
