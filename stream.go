package stream

type Stream[E any] interface {
	Filter(predicate func(*E) bool) Stream[E]
	Limit(size int) Stream[E]
	Peek(consumer func(*E)) Stream[E]
	// TerminalOp
	ForEach(consumer func(*E))
	AnyMatch(predicate func(*E) bool) bool
	FirstMatch(predicate func(*E) bool) *E
}

type iterator[E any] interface {
	next() *E
	done() bool
}

type sliceIterator[E any] struct {
	data   []*E
	index  int
	isDone bool
}

func (i *sliceIterator[E]) next() *E {
	if i.isDone {
		return nil
	}
	d := i.data[i.index]
	i.index++
	if len(i.data) <= i.index {
		i.isDone = true
	}
	return d
}

func (i *sliceIterator[E]) done() bool {
	return i.isDone
}

type proxyIterator[E any] struct {
	rootIterator iterator[E]
	hookedNext   func(e *E) *E
}

func (i *proxyIterator[E]) next() *E {
	e := i.rootIterator.next()
	if e != nil {
		return i.hookedNext(e)
	}
	return nil
}

func (i *proxyIterator[E]) done() bool {
	return i.rootIterator.done()
}

type streamImpl[E any] struct {
	iterator iterator[E]
}

func (s *streamImpl[E]) Filter(predicate func(*E) bool) Stream[E] {
	return &streamImpl[E]{
		iterator: &proxyIterator[E]{
			rootIterator: s.iterator,
			hookedNext: func(e *E) *E {
				if predicate(e) {
					return e
				}
				return nil
			},
		},
	}
}

func (s *streamImpl[E]) Limit(size int) Stream[E] {
	return nil
}

func (s *streamImpl[E]) Peek(consumer func(*E)) Stream[E] {
	return nil
}

func (s *streamImpl[E]) ForEach(consumer func(*E)) {

}

func (s *streamImpl[E]) AnyMatch(predicate func(*E) bool) bool {
	return false
}

func (s *streamImpl[E]) FirstMatch(predicate func(*E) bool) *E {
	return nil
}

func StreamOf[E any](data []*E) Stream[E] {
	return &streamImpl[E]{
		iterator: &sliceIterator[E]{
			data:   data,
			index:  0,
			isDone: false,
		},
	}
}
