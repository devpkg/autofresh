package logger

import (
	"fmt"
	"log"
)

type LogWriter struct {
	logger *log.Logger
}

func (lw LogWriter) Write(p []byte) (n int, err error) {
	fmt.Print(string(p))
	return len(p), nil
}
