package graph

import (
	"context"
	"errors"

	"github.com/etc-sudonters/substrate/skelly/bitset"
	"github.com/etc-sudonters/substrate/skelly/queue"
	"github.com/etc-sudonters/substrate/skelly/stack"
)

var (
	// given Node n, return all nodes it points at
	Successors SelectorFunc[Destination] = Directed.Successors
	// given Node n, return all nodes that point at it
	Predecessors SelectorFunc[Origination] = Directed.Predecessors
	// visitors should return this err to gracefully exit a walk
	ErrVisitTerminated = errors.New("visit terminated")
)

type (
	BreadthFirst[T Direction] struct {
		Visitor
		Selector[T]
	}

	DepthFirst[T Direction] struct {
		Visitor
		Selector[T]
	}

	Walker[T Direction] interface {
		Walk(context.Context, Directed, Node) error
	}

	// called with the current node and a context
	// visitors may gracefully end a walk by returning ErrVisitTerminated
	// all other errors terminate the walk and the error is returned to the user
	Visitor interface {
		Visit(context.Context, Node) error
	}

	// responsible for determining which nodes to explore next and in which direction to explore
	Selector[T Direction] interface {
		Select(Directed, Node) (bitset.Bitset64, error)
	}

	SelectorFunc[T Direction] func(Directed, Node) (bitset.Bitset64, error)
	VisitorFunc               func(context.Context, Node) error
)

func (s SelectorFunc[T]) Select(g Directed, n Node) (bitset.Bitset64, error) {
	return s(g, n)
}

func (v VisitorFunc) Visit(ctx context.Context, n Node) error {
	return v(ctx, n)
}

func (b BreadthFirst[T]) Walk(ctx context.Context, g Directed, r Node) error {
	q := &queue.Q[Node]{r}
	seen := bitset.New(bitset.Buckets(uint64(g.NodeCount())))
	bitset.Set(&seen, r)

	var node Node
	for len(*q) > 0 {
		if err := ctx.Err(); err != nil {
			if errors.Is(err, ErrVisitTerminated) {
				break
			}
			return err
		}

		node, _ = q.Pop()

		if err := b.Visitor.Visit(ctx, node); err != nil {
			if errors.Is(err, ErrVisitTerminated) {
				break
			}
			return err
		}

		neighbors, err := b.Selector.Select(g, node)
		if err != nil {
			return err
		}

		iter := bitset.Iter(neighbors)
		for iter.MoveNext() {
			neighbor := uint64(iter.Current())
			if !seen.IsSet(neighbor) {
				seen.Set(neighbor)
				q.Push(Node(neighbor))
			}

		}
	}

	return nil
}

func (d DepthFirst[T]) Walk(ctx context.Context, g Directed, r Node) error {
	s := &stack.S[Node]{r}
	seen := bitset.New(bitset.Buckets(uint64(g.NodeCount())))

	var node Node
	for len(*s) > 0 {
		if err := ctx.Err(); err != nil {
			return err
		}

		node, _ = s.Pop()

		if !seen.IsSet(uint64(node)) {
			if err := d.Visitor.Visit(ctx, Node(node)); err != nil {
				return err
			}

			seen.Set(uint64(node))
			neighbors, err := d.Selector.Select(g, Node(node))
			if err != nil {
				return err
			}

			iter := bitset.Iter(neighbors)
			for iter.MoveNext() {
				s.Push(Node(iter.Current()))
			}
		}
	}

	return nil
}
