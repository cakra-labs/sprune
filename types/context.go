package types

import (
	"context"
	"time"

	cometbftlog "github.com/cometbft/cometbft/libs/log"
)

type Context struct {
	baseCtx context.Context
	logger  cometbftlog.Logger
}

// Read-only accessors
func (c Context) Context() context.Context   { return c.baseCtx }
func (c Context) Logger() cometbftlog.Logger { return c.logger }

func (c Context) Deadline() (deadline time.Time, ok bool) {
	return c.baseCtx.Deadline()
}

func (c Context) Done() <-chan struct{} {
	return c.baseCtx.Done()
}

func (c Context) Err() error {
	return c.baseCtx.Err()
}

// Create a new context
func NewContext(
	logger cometbftlog.Logger,
) Context {
	return Context{
		baseCtx: context.Background(),
		logger:  logger,
	}
}

// WithContext returns a Context with an updated context.Context.
func (c Context) WithContext(ctx context.Context) Context {
	c.baseCtx = ctx
	return c
}

// WithLogger returns a Context with an updated logger.
func (c Context) WithLogger(logger cometbftlog.Logger) Context {
	c.logger = logger
	return c
}

func (c Context) WithValue(key, value interface{}) Context {
	c.baseCtx = context.WithValue(c.baseCtx, key, value)
	return c
}

func (c Context) WithTimeout(timeout time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(c.baseCtx, timeout)
}

func (c Context) Value(key interface{}) interface{} {
	return c.baseCtx.Value(key)
}
