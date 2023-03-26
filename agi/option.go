package agi

import (
	"go.uber.org/zap"
)

type Options struct {
	Addr         string
	MaxReadBytes int64
	Logger       *zap.Logger
}

type Option func(a *Options) error

func WithMaxReadBytes(readBytes int64) Option {
	return func(a *Options) error {
		a.MaxReadBytes = readBytes
		return nil
	}
}

func WithLogger(logger *zap.Logger) Option {
	return func(a *Options) error {
		a.Logger = logger
		return nil
	}
}
