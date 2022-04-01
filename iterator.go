package stream

const (
	consume int = iota
	complete
	pass
	terminate
)

type dataSignal[E any] struct {
	data   *E
	signal int
}

func newDataSignal[E any](e *E, s int) *dataSignal[E] {
	return &dataSignal[E]{
		data:   e,
		signal: s,
	}
}

type iterator[E any] interface {
	next() *dataSignal[E]
}

type sourceIterator[E any] struct {
	onNext func() *E
}

func (i *sourceIterator[E]) next() *dataSignal[E] {
	e := i.onNext()
	if e == nil {
		return newDataSignal[E](nil, complete)
	}
	return newDataSignal(e, consume)
}

type intermediateIterator[E any] struct {
	chainedIterator iterator[E]
	beforeNext      func() *dataSignal[E]
	afterNext       func(*dataSignal[E]) *dataSignal[E]
}

func (i *intermediateIterator[E]) next() *dataSignal[E] {
	if d := i.beforeNext(); d.signal == terminate {
		return newDataSignal[E](nil, complete)
	}
	return i.afterNext(i.chainedIterator.next())
}
