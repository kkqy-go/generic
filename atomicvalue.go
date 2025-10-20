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
    // Guard against swapping before initialization: if no value has been stored yet,
    // initialize with the new value and return the zero value of T.
    old := p.v.Load()
    if old == nil {
        p.v.Store(new)
        var zero T
        return zero
    }
    return p.v.Swap(new).(T)
}
func (p *AtomicValue[T]) CompareAndSwap(old, new T) bool {
    return p.v.CompareAndSwap(old, new)
}
