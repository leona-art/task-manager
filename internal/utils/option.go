package utils

type Option[T any] struct {
	value T
	none  bool
}

func (s Option[T]) IsNone() bool {
	return s.none
}
func (s Option[T]) IsSome() bool {
	return !s.none
}

func (s Option[T]) Get() (value T, ok bool) {
	if s.none {
		return value, false
	}
	return s.value, true
}

func (s Option[T]) Unwrap() T {
	if s.none {
		panic("called `Option.Unwrap()` on a `None` value")
	}
	return s.value
}

func Some[T any](value T) Option[T] {
	return Option[T]{value: value, none: false}
}
func None[T any]() Option[T] {
	return Option[T]{none: true}
}
func From[T any](value T, ok bool) Option[T] {
	if ok {
		return Some(value)
	}
	return None[T]()
}
