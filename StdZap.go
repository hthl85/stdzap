package stdzap

import (
	zapconf "github.com/hthl85/conf-zap"
	"go.uber.org/zap"
)

// StdZap struct
type StdZap struct {
	logger *zap.Logger
}

// NewStdZap creates new std zap
func NewStdZap(zapConf *zapconf.ZapConf) (std *StdZap, err error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	logger.Sugar().Infof("init std zap logger\n")

	return &StdZap{
		logger: logger,
	}, nil
}

// Sync sync logger
func (std *StdZap) Sync() error {
	std.logger.Sugar().Infof("close std zap logger\n")
	return std.logger.Sync()
}

// Infof logs info
func (std *StdZap) Infof(msg string, args ...interface{}) {
	std.logger.Sugar().Infof(msg, args)
}

// Error logs error
func (std *StdZap) Error(err error) {
	std.logger.Sugar().Error(err)
}

// Errorf logs error info
func (std *StdZap) Errorf(msg string, args ...interface{}) {
	std.logger.Sugar().Errorf(msg, args)
}
