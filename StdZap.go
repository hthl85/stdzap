package stdzap

import (
	zapconf "github.com/hthl85/conf-zap"
	"go.uber.org/zap"
	"log"
)

// StdZap struct
type StdZap struct {
	sugar *zap.SugaredLogger
}

// NewStdZap creates new stdzap
func NewStdZap(zapConf *zapconf.ZapConf) (std *StdZap, err error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	if logger != nil {
		defer func() {
			log.Print("[DEBUG] stdzap defer sync is executing")
			if deferErr := logger.Sync(); deferErr != nil {
				err = deferErr
			}
		}()
	}

	sugar := logger.Sugar()

	return &StdZap{
		sugar: sugar,
	}, nil
}

// Infof logs info
func (std *StdZap) Infof(msg string, args ...interface{}) {
	std.sugar.Infof(msg, args)
}

// Error logs error
func (std *StdZap) Error(err error) {
	std.sugar.Error(err)
}

// Errorf logs error info
func (std *StdZap) Errorf(msg string, args ...interface{}) {
	std.sugar.Errorf(msg, args)
}
