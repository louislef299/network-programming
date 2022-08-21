package ch03

import (
	"context"
	"io"
	"time"
)

func Pinger(ctx context.Context, w io.Writer, reset <-chan time.Duration) {
	var interval time.Duration
	select {
	case <-ctx.Done():
		return
	case interval = <-reset: // pulled initial interval off res
	}
}
