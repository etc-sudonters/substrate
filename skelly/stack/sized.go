package stack

import (
	"errors"
)

var (
	ErrStackEmpty = errors.New("stack empty")
	ErrStackFull  = errors.New("stack full")
)

func NewSized[T any](size int) Sized[T] {
	return Sized[T]{
		items: make([]T, size),
	}
}

type Sized[T any] struct {
	items []T
	ptr   int
}

func (this *Sized[T]) Reset() {
	this.ptr = 0
}

func (this *Sized[T]) Slice(start, count int) []T {
	return this.items[start : start+count]
}

func (this *Sized[T]) PopN(n int) {
	this.ptr -= n
}

func (this *Sized[T]) Top() T {
	if this.ptr < 1 {
		panic(ErrStackEmpty)
	}

	return this.items[this.ptr-1]
}

func (this *Sized[T]) Push(item T) {
	if this.ptr == len(this.items) {
		panic(ErrStackFull)
	}

	this.items[this.ptr] = item
	this.ptr++
}

func (this *Sized[T]) Pop() T {
	if this.ptr == 0 {
		panic(ErrStackEmpty)
	}
	this.ptr--
	return this.items[this.ptr]
}
