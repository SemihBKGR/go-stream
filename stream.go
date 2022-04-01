package stream

type Stream[E any] interface {
	Filter(predicate func(E) bool) Stream[E]
	//Limit(size int) Stream[E]
	//Peek(consumer func(E)) Stream[E]
	ForEach(consumer func(E))
	//AnyMatch(predicate func(E) bool) bool
	//FirstMatch(predicate func(E) bool) E
}

type stream[E any] struct {
	iterator iterator[E]
}

func (s *stream[E]) Filter(predicate func(E) bool) Stream[E] {
	return &stream[E]{
		iterator: filterIterator(s.iterator, predicate),
	}
}

func (s *stream[E]) ForEach(consumer func(E)) {
	startStream(&stream[E]{
		iterator: forEachIterator(s.iterator, consumer),
	})
}

func startStream[E any](s *stream[E]) {
	ds := s.iterator.next()
	for ds.signal == consume {
		ds = s.iterator.next()
	}
}

func StreamOf[E any](slice []E) Stream[E] {
	return &stream[E]{
		iterator: sliceIterator(slice),
	}
}
