package teapot

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"runtime/debug"
	"sync"
	"time"
)

// Level
type Level int

const (
	// DEBUG
	DEBUG Level = iota
	// INFO
	INFO
	// ERROR
	ERROR
)

func (l Level) string() string {
	switch l {
	case INFO:
		return "INFO"
	case ERROR:
		return "ERROR"
	default:
		return "DEBUG"
	}
}

// Logger
type Logger struct {
	mtx    sync.Mutex
	writer io.Writer
	pool   sync.Pool

	prefix string
	lvl    Level
}

// New
func New() *Logger {
	return &Logger{
		mtx:    sync.Mutex{},
		writer: os.Stdout,
		pool: sync.Pool{
			New: func() any { return new(bytes.Buffer) },
		},
		prefix: "teapot",
		lvl:    DEBUG,
	}
}

// SetOutput
func (l *Logger) SetOutput(w io.Writer) {
	l.writer = w
}

// SetLevel
func (l *Logger) SetLevel(lvl Level) {
	l.lvl = lvl
}

// SetPrefix
func (l *Logger) SetPrefix(s string) {
	l.prefix = s
}

// Debug
func (l *Logger) Debug(format string, args ...any) {
	l.logMsg(DEBUG, format, args...)
}

// Info
func (l *Logger) Info(format string, args ...any) {
	l.logMsg(INFO, format, args...)
}

// Error
func (l *Logger) Error(format string, args ...any) {
	l.logMsg(ERROR, format, args...)
}

func (l *Logger) logMsg(lvl Level, format string, args ...any) {
	if lvl < l.lvl {
		return
	}

	l.mtx.Lock()
	defer l.mtx.Unlock()

	buf := l.pool.Get().(*bytes.Buffer)
	buf.Reset()

	buf.WriteString("{\"timestamp\":\"")
	buf.WriteString(time.Now().Format(time.RFC3339Nano))
	buf.WriteString("\", \"service\":\"")
	buf.WriteString(l.prefix)
	buf.WriteString("\", \"level\":\"")
	buf.WriteString(lvl.string())
	buf.WriteString("\", \"message\":\"")

	fmt.Fprintf(buf, format, args...)

	if lvl == ERROR {
		stack := debug.Stack()
		buf.WriteString("\", \"stack_trace\":\n")
		buf.Write(stack)
	}

	buf.WriteString("\"}\n")

	buf.WriteTo(l.writer)

	l.pool.Put(buf)
}
