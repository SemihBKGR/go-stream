package stream

type Stream[E any] interface {
	Filter(predicate func(E) bool) Stream[E]
	Limit(size int) Stream[E]
	Peek(consumer func(E)) Stream[E]
	ForEach(consumer func(E))
	AnyMatch(predicate func(E) bool) bool
	FirstMatch(predicate func(E) bool) E
}

type iterator[E any] interface {
	next() *E
	isDone() bool
}

type sliceIterator[E any] struct {
	data  []E
	index int
}

func (i *sliceIterator[E]) next() *E {
	d := i.data[i.index]
	i.index++
	return &d
}

func (i *sliceIterator[E]) isDone() bool {
	return len(i.data) <= i.index
}

type proxyIterator[E any] struct {
	rootIterator iterator[E]
	hookedNext   func(*E) *E
	done         bool
}

func (i *proxyIterator[E]) next() *E {
	e := i.rootIterator.next()
	if e != nil {
		return i.hookedNext(e)
	}
	return nil
}

func (i *proxyIterator[E]) isDone() bool {
	return i.done || i.rootIterator.isDone()
}

type streamImpl[E any] struct {
	iterator iterator[E]
}

func (s *streamImpl[E]) Filter(predicate func(E) bool) Stream[E] {
	return &streamImpl[E]{
		iterator: &proxyIterator[E]{
			rootIterator: s.iterator,
			hookedNext: func(e E) E {
				if predicate(e) {
					return e
				}
			},
		},
	}
}

func (s *streamImpl[E]) Limit(size int) Stream[E] {
	return &streamImpl[E]{
		iterator: &proxyIterator[E]{
			rootIterator: s.iterator,
			hookedNext: func(i int) func(e *E) *E {
				return func(e *E) *E {
					if i >= size {
						return nil
					} else {
						i++
						return e
					}
				}
			}(0),
		},
	}
}

func (s *streamImpl[E]) Peek(consumer func(E)) Stream[E] {
	return &streamImpl[E]{
		iterator: &proxyIterator[E]{
			rootIterator: s.iterator,
			hookedNext: func(e *E) *E {
				consumer(e)
				return nil
			},
		},
	}
}

func (s *streamImpl[E]) ForEach(consumer func(E)) {
	for !s.iterator.done() {
		e := s.iterator.next()
		if e != nil {
			consumer(e)
		}
	}
}

func (s *streamImpl[E]) AnyMatch(predicate func(E) bool) bool {
	for !s.iterator.done() {
		e := s.iterator.next()
		if e != nil && predicate(e) {
			return true
		}
	}
	return false
}

func (s *streamImpl[E]) FirstMatch(predicate func(E) bool) E {
	for !s.iterator.done() {
		e := s.iterator.next()
		if e != nil && predicate(e) {
			return e
		}
	}
	return nil
}

func StreamOf[E any](data []E) Stream[E] {
	return &streamImpl[E]{
		iterator: &sliceIterator[E]{
			data:   data,
			index:  0,
			isDone: false,
		},
	}
}
