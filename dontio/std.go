package dontio

import (
	"context"
	"fmt"
	"io"
	"os"
)

type ctxkey string

var stdkey ctxkey = "std"

type notInCtx string // what wasn't present

func (what notInCtx) Error() string {
	return fmt.Sprintf("%s was not found in context", string(what))
}

type Std struct {
	Out io.Writer
	Err io.Writer
	In  io.Reader
}

func AddStdToContext(ctx context.Context, s *Std) context.Context {
	return context.WithValue(ctx, stdkey, s)
}

func StdFromContext(ctx context.Context) (*Std, error) {
	v := ctx.Value(stdkey)
	if v == nil {
		return nil, notInCtx("stdio")
	}
	return v.(*Std), nil
}

func StdIo() Std {
	return Std{
		Out: os.Stdout,
		Err: os.Stderr,
		In:  os.Stdin,
	}
}
