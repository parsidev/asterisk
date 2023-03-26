package agi

import (
	"time"

	"go.uber.org/zap"
)

type Options struct {
	Addr         string
	IdleTimeout  time.Duration
	MaxReadBytes int64
	Logger       *zap.Logger
}

type Option func(a *Options) error

func WIthIdleTimeout(timeout time.Duration) Option {
	return func(a *Options) error {
		a.IdleTimeout = timeout
		return nil
	}
}

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
