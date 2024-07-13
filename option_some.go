package shine

type Some[T any] [1]T

func (s Some[T]) UnwrapOr(def T) T {
	return s[0]
}

func (s Some[T]) UnwrapOrDefault() T {
	return s[0]
}

func (s Some[T]) UnwrapOrElse(fn func() T) T {
	return s[0]
}

func (s Some[T]) OkOr(err error) Result[T] {
	return NewOk[T](s[0])
}

func (s Some[T]) OkOrElse(fn func() error) Result[T] {
	return NewOk[T](s[0])
}

func (s Some[T]) Filter(predicate func(v T) bool) Option[T] {
	if predicate(s[0]) {
		return s
	}

	return NewNone[T]()
}

func (s Some[T]) AndThen(fn func(v T) Option[T]) Option[T] {
	return fn(s[0])
}

func (s Some[T]) Xor(other Option[T]) Option[T] {
	if _, ok := other.(None[T]); ok {
		return s
	}

	return NewNone[T]()
}

func (Some[T]) option() {

}

func (s Some[T]) Value() T {
	return s[0]
}

func NewSome[T any](v T) Some[T] {
	return Some[T]{v}
}

var _ Option[struct{}] = NewSome(struct{}{})
