package runtimekit

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

func NewRuntimeContext() (context.Context, context.CancelFunc) {
	return signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
}
