package logger

import (
	"github.com/ssengalanto/biscuit/pkg/constants"
	"github.com/ssengalanto/biscuit/pkg/interfaces"
	"github.com/ssengalanto/biscuit/pkg/logger/zap"
)

type buildProvider interface {
	logType() string
	build(env string) (interfaces.Logger, error)
}

// zapLog - buildProvider for zap logger.
type zapLog struct{}

func (z zapLog) logType() string {
	return constants.ZapLogType
}

func (z zapLog) build(env string) (interfaces.Logger, error) {
	logger, err := zap.New(env)
	if err != nil {
		return nil, err
	}

	return logger, nil
}

// getBuildProviders returns a slice of buildProvider.
func getBuildProviders() []buildProvider {
	return []buildProvider{
		zapLog{},
	}
}
