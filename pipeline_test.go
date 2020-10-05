package pipelineinternal

import (
	"context"
	"errors"
	"fmt"
	"testing"
)

func TestPipe(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	p1 := NewPipeNode()
	p2 := NewPipeNode()
	p3 := NewPipeNode()

	// work 1 frist work
	p1.SetNextNode(p2)
	p1.SetWork(&p1Worker{})

	// work 2
	p2.SetNextNode(p3)
	p2.SetWork(&p2Worker{})

	// work 3 last work
	p3.SetWork(&p3Worker{})

	if err := DefaultExec(ctx, p1); err != nil {
		t.Error(err)
	}
}

type p1Worker struct {
	Name string
}

func (p1 *p1Worker) Run(ctx context.Context) (context.Context, error) {
	fmt.Println("p1")
	p1.Name = "worker 1"
	ctx = context.WithValue(ctx, p1Worker{}, p1)
	return ctx, nil
}

type p2Worker struct{}

func (p2 *p2Worker) Run(ctx context.Context) (context.Context, error) {
	p1, ok := ctx.Value(p1Worker{}).(*p1Worker)
	if !ok {
		return context.TODO(), errors.New("p1 worker error")
	}

	fmt.Println("p2")
	fmt.Println(p1.Name)

	return ctx, nil
}

type p3Worker struct{}

func (p3 *p3Worker) Run(ctx context.Context) (context.Context, error) {
	fmt.Println("p3")
	return ctx, nil
}