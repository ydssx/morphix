package concurrent

import (
	"context"

	"golang.org/x/sync/errgroup"
)

type Group struct {
	limit    int
	eg       *errgroup.Group
	ctx      context.Context
	errChan  chan struct{}
	fastFail bool
}

type Opt func(*Group)

// WithSemaphore limit the max actively goroutines to semaphore.
func WithSemaphore(semaphore int) Opt {
	return func(g *Group) { g.limit = semaphore }
}

// WithFastFail if set true, Run will return when first err return by function passed to Run,
// else Run will block util all functions finished.
func WithFastFail(fastFail bool) Opt {
	return func(g *Group) { g.fastFail = fastFail }
}

func NewGroup(ctx context.Context, opts ...Opt) *Group {
	eg, ctx := errgroup.WithContext(ctx)
	g := &Group{eg: eg, errChan: make(chan struct{}, 1), ctx: ctx}

	for _, opt := range opts {
		opt(g)
	}
	if g.limit > 0 {
		eg.SetLimit(g.limit)
	}
	return g
}

func (g *Group) Run(fs ...func() error) (err error) {
	for _, f := range fs {
		g.eg.Go(f)
	}
	if !g.fastFail {
		return g.eg.Wait()
	}

	go func() {
		err = g.eg.Wait()
		g.errChan <- struct{}{}
		close(g.errChan)
	}()

	select {
	case <-g.ctx.Done():
		return g.ctx.Err()
	case <-g.errChan:
	}
	return
}
