package stream

type Stream[E any] interface {
	Filter(predicate func(E) bool) Stream[E]
	Limit(size int) Stream[E]
	Peek(consumer func(E)) Stream[E]
	ForEach(consumer func(E))
	AnyMatch(predicate func(E) bool) bool
	FirstMatch(predicate func(E) bool) E
}

type stream[E any] struct {
	iterator iterator[E]
}

func (s *stream[E]) Filter(predicate func(E) bool) Stream[E] {
	return &stream[E]{
		iterator: filterIterator(s.iterator, predicate),
	}
}

func (s *stream[E]) Limit(size int) Stream[E] {
	return &stream[E]{
		iterator: limitPeek(s.iterator, size),
	}
}

func (s *stream[E]) Peek(consumer func(E)) Stream[E] {
	return &stream[E]{
		iterator: peekIterator(s.iterator, consumer),
	}
}

func (s *stream[E]) ForEach(consumer func(E)) {
	d := s.iterator.next()
	for d.signal != complete {
		if d.signal == consume {
			consumer(*d.data)
		}
		d = s.iterator.next()
	}
}

func (s *stream[E]) AnyMatch(predicate func(E) bool) bool {
	d := s.iterator.next()
	for d.signal != complete {
		if d.signal == consume {
			if predicate(*d.data) {
				return true
			}
		}
		d = s.iterator.next()
	}
	return false
}

func (s *stream[E]) FirstMatch(predicate func(E) bool) E {
	d := s.iterator.next()
	for d.signal != complete {
		if d.signal == consume {
			if predicate(*d.data) {
				return *d.data
			}
		}
		d = s.iterator.next()
	}
	panic("not any match")
}

func StreamOf[E any](slice []E) Stream[E] {
	return &stream[E]{
		iterator: sliceIterator(slice),
	}
}
