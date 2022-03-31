package stream

const (
	consume int = iota
	complete
)

type dataSignal[E any] struct {
	data   *E
	signal int
}

func newDataSignal[E any](e *E, s int) dataSignal[E] {
	return dataSignal[E]{
		data:   e,
		signal: s,
	}
}

type iterator[E any] interface {
	next() dataSignal[E]
}

type sourceIterator[E any] interface {
	iterator[E]
	hasNext() bool
}

type sliceIterator[E any] struct {
	data  []E
	index int
}

func (i *sliceIterator[E]) next() dataSignal[E] {
	if !i.hasNext() {
		return newDataSignal[E](nil, complete)
	}
	e := &i.data[i.index]
	i.index++
	return newDataSignal[E](e, consume)
}

func (i *sliceIterator[E]) hasNext() bool {
	return i.index <= len(i.data)
}

func newSliceIterator[E any](data []E) iterator[E] {
	return &sliceIterator[E]{
		data:  data,
		index: 0,
	}
}

type intermediateIterator[E any] struct {
}
