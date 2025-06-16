package generic

type Nilable[T any] struct {
	presented bool
	value     T
}

func (p *Nilable[T]) Nil() bool {
	return !p.presented
}

func (p *Nilable[T]) Value() T {
	return p.value
}

func (p *Nilable[T]) Presented() bool {
	return p.presented
}

func (p *Nilable[T]) SetValue(value T) {
	p.value = value
	p.presented = true
}

func (p *Nilable[T]) Clear() {
	p.presented = false
	var t T
	p.value = t
}
