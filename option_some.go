package shine

type Some[T any] struct {
	v T
}

func (s Some[T]) IsSome() bool {
	return true
}

func (s Some[T]) IsNone() bool {
	return false
}

func (s Some[T]) Expect(panicV any) T {
	return s.v
}

func (s Some[T]) Unwrap() T {
	return s.v
}

func (s Some[T]) UnwrapOr(def T) T {
	return s.v
}

func (s Some[T]) UnwrapOrDefault() T {
	return s.v
}

func (s Some[T]) UnwrapOrElse(fn func() T) T {
	return s.v
}

func (s Some[T]) OkOr(err error) Result[T, error] {
	return NewOk[T, error](s.v)
}

func (s Some[T]) OkOrElse(fn func() error) Result[T, error] {
	return NewOk[T, error](s.v)
}

func (s Some[T]) Filter(predicate func(v T) bool) Option[T] {
	if predicate(s.v) {
		return s
	}

	return NewNone[T]()
}

func (s Some[T]) Map(fn func(v T) T) Option[T] {
	return NewSome(fn(s.v))
}

func (s Some[T]) AndThen(fn func(v T) Option[T]) Option[T] {
	return fn(s.v)
}

func (s Some[T]) OrElse(fn func() Option[T]) Option[T] {
	return s
}

func (s Some[T]) Xor(other Option[T]) Option[T] {
	if _, ok := other.(None[T]); ok {
		return s
	}

	return NewNone[T]()
}

func (s Some[T]) Iter() <-chan T {
	ch := make(chan T, 1)
	ch <- s.v
	close(ch)

	return ch
}

func NewSome[T any](v T) Some[T] {
	return Some[T]{v}
}

var _ Option[struct{}] = NewSome(struct{}{})
