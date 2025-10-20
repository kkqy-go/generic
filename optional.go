package generic

import (
    "bytes"
    "encoding/json"
)

type Optional[T any] struct {
    value     T
    presented bool
}

// Value returns the value if it is presented, otherwise returns the zero value.
func (p *Optional[T]) Value() T {
    return p.value
}

// Ptr returns a pointer to the value if it is presented, otherwise returns nil.
func (p *Optional[T]) Ptr() *T {
    if p.presented {
        return &p.value
    }
    return nil
}

// ValueOr returns the value if it is presented, otherwise returns the default value.
func (p *Optional[T]) ValueOr(defaultValue T) T {
    if p.presented {
        return p.value
    }
    return defaultValue
}

// Presented returns true if the value is presented.
func (p *Optional[T]) Presented() bool {
    return p.presented
}

// NotPresented returns true if the value is not presented.
func (p *Optional[T]) NotPresented() bool {
    return !p.presented
}

// SetValue sets the value and presented to true.
func (p *Optional[T]) SetValue(value T) {
    p.value = value
    p.presented = true
}

// Clear sets presented to false and value to zero value.
func (p *Optional[T]) Clear() {
    p.presented = false
    var t T
    p.value = t
}

func (p Optional[T]) MarshalJSON() ([]byte, error) {
    if !p.presented {
        return []byte("null"), nil
    }
    return json.Marshal(p.value)
}

func (p *Optional[T]) UnmarshalJSON(data []byte) error {
    trimmed := bytes.TrimSpace(data)
    if bytes.Equal(trimmed, []byte("null")) {
        p.presented = false
        var zero T
        p.value = zero
        return nil
    }
    p.presented = true
    return json.Unmarshal(trimmed, &p.value)
}

func NewOptional[T any](value T) Optional[T] {
    return Optional[T]{value: value, presented: true}
}
