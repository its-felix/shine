package shine

type Option[T any] interface {
	UnwrapOr(def T) T
	UnwrapOrDefault() T
	UnwrapOrElse(fn func() T) T
	OkOr(err error) Result[T]
	OkOrElse(fn func() error) Result[T]
	Filter(predicate func(v T) bool) Option[T]
	AndThen(fn func(v T) Option[T]) Option[T]
	Xor(other Option[T]) Option[T]
	option()
}

func NewOption[T any](v T) Option[T] {
	return NewOptionFrom(v, !isNil(v))
}

func NewOptionFrom[T any](v T, ok bool) Option[T] {
	if ok {
		return NewSome(v)
	}

	return NewNone[T]()
}

func NewOptionFromMap[K comparable, V any](m map[K]V, k K) Option[V] {
	v, ok := m[k]
	return NewOptionFrom(v, ok)
}
