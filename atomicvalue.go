package generic

import "sync/atomic"

type AtomicValue[T any] struct {
	v atomic.Value
}

func (p *AtomicValue[T]) Load() T {
	v := p.v.Load()
	if v == nil {
		var zero T
		return zero
	}
	return v.(T)
}
func (p *AtomicValue[T]) Store(val T) {
	p.v.Store(val)
}
func (p *AtomicValue[T]) Swap(new T) T {
	return p.v.Swap(new).(T)
}
func (p *AtomicValue[T]) CompareAndSwap(old, new T) bool {
	return p.v.CompareAndSwap(old, new)
}
