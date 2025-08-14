package logger

import (
	"log"
)

type LogMessage struct {
	Level   string
	Message string
}

type Logger struct {
	logCh chan LogMessage
}

func NewLogger() *Logger {
	logCh := make(chan LogMessage, 100)
	l := &Logger{logCh: logCh}
	go l.processLogs()
	return l
}

func (l *Logger) processLogs() {
	for msg := range l.logCh {
		log.Printf("[%s] %s\n", msg.Level, msg.Message)
	}
}

func (l *Logger) Log(level, message string) {
	l.logCh <- LogMessage{
		Level:   level,
		Message: message,
	}
}

func (l *Logger) Close() {
	close(l.logCh)
}
