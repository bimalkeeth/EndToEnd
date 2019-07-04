package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
)

type Level int

const (
	LevelDebug Level = iota
	LevelInfo
	LevelError
)

const defaultLogLevel = LevelInfo

type Logger struct {
	mu     sync.Mutex
	prefix string
	Level  Level
	w      io.Writer
	buf    bytes.Buffer
}

func New(w io.Writer, prefix string) *Logger {
	return &Logger{w: w, prefix: prefix, Level: defaultLogLevel}
}

var Console = New(os.Stderr, "")

func (l *Logger) Debug(v ...interface{}) {
	l.WriteEntry(LevelDebug, fmt.Sprintln(v...))
}

func (l *Logger) Info(v ...interface{}) {
	if LevelInfo < l.Level {
		return
	}
	l.WriteEntry(LevelInfo, fmt.Sprintln(v...))
}

func (l *Logger) Error(v ...interface{}) {
	if LevelError < l.Level {
		return
	}
	l.WriteEntry(LevelError, fmt.Sprintln(v...))
}

func (l *Logger) WriteEntry(level Level, sprintln string) {
	_, _ = l.w.Write([]byte(sprintln))

}
func (l *Logger) SetLevel(lvl Level) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.Level = lvl
}
func (l *Logger) GetLevel() Level {
	return l.Level
}

func main() {

	Console.Info("Hello")
	Console.Debug("Hello", "Debugger")
	Console.Error("Error")
	Console.SetLevel(LevelError)
	Console.Info("Hello")
	Console.Debug("Hello", "Debugger")
	Console.Error("Error")

}
