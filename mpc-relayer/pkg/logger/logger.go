package logger

import (
	"fmt"
	"os"
	"strings"

	"github.com/qredo/fusionchain/mpc-relayer/pkg/common"
	"github.com/sirupsen/logrus"
	lo "github.com/sirupsen/logrus"
)

type Format string
type Level string

const (
	timeFormat = "2006-01-02T15:04:05"

	Debug Level = "debug"
	Info  Level = "info"
	Warn  Level = "warn"
	Error Level = "error"
	Fatal Level = "fatal"
	Panic Level = "panic"

	JSON  Format = "json"
	Plain Format = "plain"
)

// actionFormatter duplicate the log message inside a field action
// this is useful for log analysis, has datadog doesn't give access to message field
type actionFormatter struct {
	original logrus.Formatter
}

type ddErr struct {
	Message any `json:"message"`
}

// errorReMapper remap the field error value inside struct
// error fields conflicts with error.kind on datadog
type errorReMapper struct {
	logrus.Formatter
}

// NewLogger initializes a new (logrus) Logger instance
// Supported log types are: text, json, none
func NewLogger(logLevel Level, logFormat Format, Tofile bool, service string) (*logrus.Entry, error) {
	if err := IsValidLevel(logLevel, service); err != nil {
		return nil, err
	}
	if err := IsValidFormat(logFormat, service); err != nil {
		return nil, err
	}
	debugWarnMsg := fmt.Sprintf("%s RUNNING IN DEBUG MODE. DO NOT RUN IN PRODUCTION ENVIRONMENT", strings.ToUpper(service))
	// Set logging to json
	l := NewFormattedLogger(logLevel, logFormat, debugWarnMsg)

	// OPTIONAL OUTPUT TO .log file (Default false)
	if Tofile {
		f, err := os.OpenFile(fmt.Sprintf("%s.log", service), os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			return nil, fmt.Errorf("error opening file: %v", err)
		}

		// assign it to the standard logger
		l.SetOutput(f)
		l.Infof("Logging to file: %v", fmt.Sprintf("%s.log", service))
	}

	logger := lo.NewEntry(l)
	logger.Level = l.Level
	logger = logger.WithFields(lo.Fields{
		"serviceName": service,
		"version":     common.FullVersion,
	})
	return logger, nil
}

// return a formatted Logger object (log format is JSON by defulat)
func NewFormattedLogger(logLevel Level, logFormat Format, debugWarnMsg string) *logrus.Logger {
	l := logrus.New()
	var formatter logrus.Formatter
	// Select log Format
	switch logFormat {
	case Plain:
		formatter = &logrus.TextFormatter{
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyMsg: "message",
			},
			TimestampFormat: timeFormat,
		}
	default:
		formatter = &errorReMapper{&actionFormatter{&logrus.JSONFormatter{
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyMsg: "message",
			},
			TimestampFormat: timeFormat,
		}}}
	}

	l.SetFormatter(formatter)
	level := LogLevelFromString(logLevel)
	l.Level = level
	if logLevel == Debug {
		l.Warn(debugWarnMsg)
	}
	return l
}

func (formatter *actionFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	//entry.Data["action"] = entry.Message
	return formatter.original.Format(entry)
}

func (formatter *errorReMapper) Format(entry *logrus.Entry) ([]byte, error) {
	errorMessage := entry.Data["error"]
	if errorMessage != nil {
		entry.Data["error"] = ddErr{
			Message: errorMessage,
		}
	}
	return formatter.Formatter.Format(entry)
}

// Validation

func IsValidLevel(w Level, service string) error {
	switch w {
	case Debug, Info, Warn, Error, Fatal, Panic:
		return nil
	default:
		return fmt.Errorf("invalid %s log level input", service)
	}
}

func IsValidFormat(w Format, service string) error {
	switch w {
	case JSON, Plain:
		return nil
	default:
		return fmt.Errorf("invalid %s log format input", service)
	}
}

// logLevelFromString returns log level (Int) from string input
// Log Level (error, info, debug) from string input
func LogLevelFromString(logLevel Level) logrus.Level {
	switch logLevel {
	case Panic:
		return logrus.PanicLevel
	case Fatal:
		return logrus.FatalLevel
	case Error:
		return logrus.ErrorLevel
	case Warn:
		return logrus.WarnLevel
	case Info:
		return logrus.InfoLevel
	case Debug:
		return logrus.DebugLevel
	default:
		return logrus.InfoLevel
	}
}
