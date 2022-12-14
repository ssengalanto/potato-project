package zap

import (
	"fmt"
	"strings"

	"github.com/ssengalanto/biscuit/pkg/interfaces"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

type Logger struct {
	zap *zap.Logger
}

// New creates a new Logger instance.
func New(env string) (*Logger, error) {
	log, err := buildZapLogger(env)
	if err != nil {
		return nil, err
	}

	return &Logger{
		zap: log,
	}, nil
}

// Info logs a message at info level.
func (l *Logger) Info(msg string, fields interfaces.Fields) {
	zapFields := mapToZapFields(fields)
	l.zap.Info(msg, zapFields...)
}

// Error logs a message at error level.
func (l *Logger) Error(msg string, fields interfaces.Fields) {
	zapFields := mapToZapFields(fields)
	l.zap.Error(msg, zapFields...)
}

// Debug logs a message at debug level.
func (l *Logger) Debug(msg string, fields interfaces.Fields) {
	zapFields := mapToZapFields(fields)
	l.zap.Debug(msg, zapFields...)
}

// Warn logs a message at warn level.
func (l *Logger) Warn(msg string, fields interfaces.Fields) {
	zapFields := mapToZapFields(fields)
	l.zap.Warn(msg, zapFields...)
}

// Fatal logs a message at fatal level.
func (l *Logger) Fatal(msg string, fields interfaces.Fields) {
	zapFields := mapToZapFields(fields)
	l.zap.Fatal(msg, zapFields...)
}

// Panic logs a message at panic level.
func (l *Logger) Panic(msg string, fields interfaces.Fields) {
	zapFields := mapToZapFields(fields)
	l.zap.Panic(msg, zapFields...)
}

// mapToZapFields maps the logger fields to zap fields.
func mapToZapFields(fields map[string]any) []zap.Field {
	var zapFields []zap.Field
	for k, v := range fields {
		zapFields = append(zapFields, zap.Any(k, v))
	}

	return zapFields
}

// buildZapLogger builds a new zap.Logger for specific environment with predefined configuration.
func buildZapLogger(env string) (*zap.Logger, error) {
	var log *zap.Logger
	var err error

	buildProviders := getBuildProviders()
	lastIdx := len(buildProviders) - 1
	for i, provider := range buildProviders {
		matched := provider.env() == strings.ToLower(env)
		outOfScope := i == lastIdx && !matched

		if outOfScope {
			return nil,
				fmt.Errorf("%w: invalid env with value of `%s`, must be one of the ff: `development`, `testing`, `production`",
					ErrZapInitializationFailed, env)
		}

		if !matched {
			continue
		}

		log, err = provider.build()
		if err != nil {
			return nil, fmt.Errorf("%w: %s",
				ErrZapInitializationFailed, err)
		}
		break
	}
	return log, nil
}

// NewTestInstance creates a new Logger instance for testing.
// Use only for testing purposes.
func NewTestInstance(level zapcore.Level) (Logger, *observer.ObservedLogs) {
	observedZapCore, observedLogs := observer.New(level)
	observedLogger := zap.New(observedZapCore)

	return Logger{
		zap: observedLogger,
	}, observedLogs
}
