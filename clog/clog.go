//
//This package implements a custom log writer to facilitate formatting
//
package clog

import (
	"fmt"
	"log"
	"strings"
	"time"
)

// Custom io.Writer
type logWriter struct {
}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(time.Now().Format("2006-01-02 15:04:05") + " " + string(bytes))
}

// Initialization
func init() {
	log.SetFlags(0)
	log.SetOutput(new(logWriter))
}

// Public functions
func Log(level string, event string) {
	log.Print("["+strings.ToUpper(level)+"]", " ", event)
}

func Logf(level string, format string, event ...interface{}) {
	tmp := fmt.Sprintf(format, event...)
	Log(level, tmp)
}
