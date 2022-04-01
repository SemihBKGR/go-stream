package stream

func sliceIterator[E any](slice []E) iterator[E] {
	return &sourceIterator[E]{
		onNext: func(slice []E, index int) func() *E {
			return func() *E {
				if len(slice) <= index {
					return nil
				}
				e := &slice[index]
				index++
				return e
			}
		}(slice, 0),
	}
}

func filterIterator[E any](chainedIterator iterator[E], predicate func(E) bool) iterator[E] {
	return &intermediateIterator[E]{
		chainedIterator: chainedIterator,
		beforeNext:      func() *dataSignal[E] { return newDataSignal[E](nil, pass) },
		afterNext: func(d *dataSignal[E]) *dataSignal[E] {
			if d.signal == consume {
				if predicate(*d.data) {
					return d
				}
				return newDataSignal(d.data, pass)
			}
			return d
		},
	}
}

func forEachIterator[E any](chainedIterator iterator[E], consumer func(E)) iterator[E] {
	return &terminalIterator[E]{
		chainedIterator: chainedIterator,
		afterNext: func(d *dataSignal[E]) *dataSignal[E] {
			if d.signal == consume {
				consumer(*d.data)
			}
			return d
		},
	}
}
