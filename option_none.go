package shine

type None[T any] struct{}

func (n None[T]) IsSome() bool {
	return false
}

func (n None[T]) IsNone() bool {
	return true
}

func (n None[T]) Expect(panicV any) T {
	panic(panicV)
}

func (n None[T]) Unwrap() T {
	panic("Unwrap on None")
}

func (n None[T]) UnwrapOr(def T) T {
	return def
}

func (n None[T]) UnwrapOrDefault() T {
	var def T
	return def
}

func (n None[T]) UnwrapOrElse(fn func() T) T {
	return fn()
}

func (n None[T]) OkOr(err error) Result[T, error] {
	return NewErr[T](err)
}

func (n None[T]) OkOrElse(fn func() error) Result[T, error] {
	return NewErr[T](fn())
}

func (n None[T]) Filter(predicate func(v T) bool) Option[T] {
	return n
}

func (n None[T]) Map(fn func(v T) T) Option[T] {
	return n
}

func (n None[T]) AndThen(fn func(v T) Option[T]) Option[T] {
	return n
}

func (n None[T]) OrElse(fn func() Option[T]) Option[T] {
	return fn()
}

func (n None[T]) Xor(other Option[T]) Option[T] {
	if other, ok := other.(Some[T]); ok {
		return other
	}

	return n
}

func (n None[T]) Iter() <-chan T {
	ch := make(chan T)
	close(ch)

	return ch
}

func NewNone[T any]() None[T] {
	return None[T]{}
}

var _ Option[struct{}] = NewNone[struct{}]()
