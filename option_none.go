package shine

type None[T any] struct{}

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

func (n None[T]) OkOr(err error) Result[T] {
	return NewErr[T](err)
}

func (n None[T]) OkOrElse(fn func() error) Result[T] {
	return NewErr[T](fn())
}

func (n None[T]) Filter(predicate func(v T) bool) Option[T] {
	return n
}

func (n None[T]) AndThen(fn func(v T) Option[T]) Option[T] {
	return n
}

func (n None[T]) Xor(other Option[T]) Option[T] {
	if other, ok := other.(Some[T]); ok {
		return other
	}

	return n
}

func (None[T]) option() {

}

func NewNone[T any]() None[T] {
	return None[T]{}
}

var _ Option[struct{}] = NewNone[struct{}]()
