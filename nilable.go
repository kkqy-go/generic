package generic

type Nilable[T any] struct {
	presented bool
	value     T
}

// Nil returns true if the value is not presented.
func (p *Nilable[T]) Nil() bool {
	return !p.presented
}

// Value returns the value if it is presented, otherwise returns the zero value.
func (p *Nilable[T]) Value() T {
	return p.value
}

// ValueOr returns the value if it is presented, otherwise returns the default value.
func (p *Nilable[T]) ValueOr(defaultValue T) T {
	if p.presented {
		return p.value
	} else {
		return defaultValue
	}
}

// Presented returns true if the value is presented.
func (p *Nilable[T]) Presented() bool {
	return p.presented
}

// SetValue sets the value and presented to true.
func (p *Nilable[T]) SetValue(value T) {
	p.value = value
	p.presented = true
}

// Clear sets presented to false and value to zero value.
func (p *Nilable[T]) Clear() {
	p.presented = false
	var t T
	p.value = t
}
