package stream

type Stream[E any] struct {
	iterator iterator[E]
}

func (s *Stream[E]) Filter(predicate func(E) bool) *Stream[E] {
	return &Stream[E]{
		iterator: filterIterator(s.iterator, predicate),
	}
}

func (s *Stream[E]) Limit(size int) *Stream[E] {
	return &Stream[E]{
		iterator: limitPeek(s.iterator, size),
	}
}

func (s *Stream[E]) Peek(consumer func(E)) *Stream[E] {
	return &Stream[E]{
		iterator: peekIterator(s.iterator, consumer),
	}
}

func (s *Stream[E]) Skip(count int) *Stream[E] {
	return &Stream[E]{
		iterator: skipIterator(s.iterator, count),
	}
}

func (s *Stream[E]) ForEach(consumer func(E)) {
	d := s.iterator.next()
	for d.signal != complete {
		if d.signal == consume {
			consumer(*d.data)
		}
		d = s.iterator.next()
	}
}

func (s *Stream[E]) AnyMatch(predicate func(E) bool) bool {
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

func (s *Stream[E]) AllMatch(predicate func(E) bool) bool {
	d := s.iterator.next()
	for d.signal != complete {
		if d.signal == consume {
			if !predicate(*d.data) {
				return false
			}
		}
		d = s.iterator.next()
	}
	return true
}

func (s *Stream[E]) FirstMatch(predicate func(E) bool) E {
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

func (s *Stream[E]) Count() int {
	count := 0
	d := s.iterator.next()
	for d.signal != complete {
		if d.signal == consume {
			count++
		}
		d = s.iterator.next()
	}
	return count
}

func (s *Stream[E]) FindFirst() E {
	d := s.iterator.next()
	for d.signal != complete {
		if d.signal == consume {
			return *d.data
		}
		d = s.iterator.next()
	}
	panic("not found")
}

func (s *Stream[E]) Collect() []E {
	e := make([]E, 0)
	d := s.iterator.next()
	for d.signal != complete {
		if d.signal == consume {
			e = append(e, *d.data)
		}
		d = s.iterator.next()
	}
	return e
}

func StreamOfSlice[E any](slice []E) *Stream[E] {
	return &Stream[E]{
		iterator: sliceIterator(slice),
	}
}

func StreamOf[E any](e ...E) *Stream[E] {
	return &Stream[E]{
		iterator: sliceIterator(e),
	}
}

func Range(start, end int) *Stream[int] {
	return &Stream[int]{
		iterator: rangeIterator(start, end),
	}
}
