package log

import (
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type HttpLogInfo struct{}

type LogLevel int

const (
	Panic LogLevel = iota
	Error
	Warn
	Info
)

const (
	FieldMethod     = "method"
	FieldPath       = "path"
	FieldStatusCode = "statusCode"
	FieldClientIP   = "client_ip"
	FieldDurationMs = "duration_ms"
	FieldUserAgent  = "user_agent"
)

func HttpLog(c *gin.Context, level LogLevel, httpStatus int, message string) {
	startTimeValue, exists := c.Get("startTime")
	if !exists {
		log.Error("startTime not found in context")
		startTimeValue = time.Now() // чтобы избежать паники, если startTime не найден
	}

	startTime, ok := startTimeValue.(time.Time)
	if !ok {
		log.Error("startTime has incorrect type")
		startTime = time.Now()
	}

	durationMs := time.Since(startTime).Milliseconds()

	fields := log.Fields{
		FieldMethod:     c.Request.Method,
		FieldPath:       c.Request.URL.Path,
		FieldStatusCode: httpStatus,
		FieldClientIP:   c.ClientIP(),
		FieldDurationMs: durationMs,
		FieldUserAgent:  c.Request.UserAgent(),
	}

	switch level {
	case Panic:
		log.WithFields(fields).Panic(message)
	case Error:
		log.WithFields(fields).Error(message)
	case Warn:
		log.WithFields(fields).Warn(message)
	case Info:
		log.WithFields(fields).Info(message)
	default:
		log.WithFields(fields).Info(message)
	}
}
