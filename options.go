package progress_bar

import (
	"io"
	"time"
)

// Option ...
type Option func(pb *ProgressBar)

// Reconfigure ...
func (pb *ProgressBar) Reconfigure(options ...Option) {
	for _, option := range options {
		option(pb)
	}
}

// WithSymbol sets the symbol
func WithSymbol(symbol rune) Option {
	return func(pb *ProgressBar) {
		pb.symbol = symbol
	}
}

// WithWriter sets the writer
func WithWriter(writer io.Writer) Option {
	return func(pb *ProgressBar) {
		pb.writer = writer
	}
}

// WithDelay sets the delay in milliseconds
func WithDelay(delay time.Duration) Option {
	return func(pb *ProgressBar) {
		pb.delay = delay
	}
}
