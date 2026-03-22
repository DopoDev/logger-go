package logger

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

type Level string

// Types of logs
const (
	DEBUG    Level = "DEBUG"
	INFO     Level = "INFO"
	SUCCESS  Level = "SUCCESS"
	WARN     Level = "WARN"
	ERROR    Level = "ERROR"
	FATAL    Level = "FATAL"
	SECURITY Level = "SECURITY"
)

// Colors with ANSI
const (
	reset   = "\033[0m"
	gray    = "\033[90m"
	blue    = "\033[34m"
	green   = "\033[32m"
	yellow  = "\033[33m"
	red     = "\033[31m"
	boldRed = "\033[1;31m"
	cyan    = "\033[36m"
)

// Estruct of logger
type Logger struct {
	mu          sync.Mutex
	showTime    bool
	timeFormat  string
	minLevel    Level
	enableColor bool
}

var defaultLogger = &Logger{
	showTime:    true,
	timeFormat:  "2006-01-02 15:04:05",
	minLevel:    DEBUG,
	enableColor: true,
}

// Global Configuration
func SetMinLeve(level Level) {
	defaultLogger.minLevel = level
}

func SetTimeFormat(format string) {
	defaultLogger.timeFormat = format
}

func SetColor(enable bool) {
	defaultLogger.enableColor = enable
}

func (l *Logger) shouldLog(level Level) bool {
	order := map[Level]int{
		DEBUG:    1,
		INFO:     2,
		SUCCESS:  2,
		SECURITY: 2,
		WARN:     3,
		ERROR:    4,
		FATAL:    5,
	}

	return order[level] >= order[l.minLevel]
}

func (l *Logger) getColor(level Level) string {
	switch level {
	case DEBUG:
		return gray
	case INFO:
		return blue
	case SUCCESS:
		return green
	case WARN:
		return yellow
	case ERROR:
		return red
	case FATAL:
		return boldRed
	case SECURITY:
		return cyan
	default:
		return reset
	}
}

func (l *Logger) log(level Level, format string, args ...any) {
	if !l.shouldLog(level) {
		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	message := format

	if len(args) > 0 {
		message = fmt.Sprintf(format, args...)
	}

	timestamp := ""
	if l.showTime {
		timestamp = fmt.Sprintf("[%s] ", time.Now().Format(l.timeFormat))
	}

	leveltext := fmt.Sprintf("[%s]", strings.ToUpper(string(level)))

	if l.enableColor {
		color := l.getColor(level)
		fmt.Printf("%s%s%s %s%s\n", gray, timestamp, reset, color+leveltext+reset, " "+message)
	} else {
		fmt.Printf("%s%s %s\n", timestamp, leveltext, message)
	}

	if level == FATAL {
		os.Exit(1)
	}
}

// Public methods
func Debug(format string, args ...any) {
	defaultLogger.log(DEBUG, format, args...)
}

func Info(format string, args ...any) {
	defaultLogger.log(INFO, format, args...)
}

func Success(format string, args ...any) {
	defaultLogger.log(SUCCESS, format, args...)
}

func Warn(format string, args ...any) {
	defaultLogger.log(WARN, format, args...)
}

func Error(format string, args ...any) {
	defaultLogger.log(ERROR, format, args...)
}

func Fatal(format string, args ...any) {
	defaultLogger.log(FATAL, format, args...)
}

func Security(format string, args ...any) {
	defaultLogger.log(SECURITY, format, args...)
}
