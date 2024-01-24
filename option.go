package shine

type Option[T any] interface {
	IsSome() bool
	IsNone() bool
	Expect(panicV any) T
	Unwrap() T
	UnwrapOr(def T) T
	UnwrapOrDefault() T
	UnwrapOrElse(fn func() T) T
	OkOr(err error) Result[T, error]
	OkOrElse(fn func() error) Result[T, error]
	Filter(predicate func(v T) bool) Option[T]
	Map(fn func(v T) T) Option[T]
	AndThen(fn func(v T) Option[T]) Option[T]
	OrElse(fn func() Option[T]) Option[T]
	Xor(other Option[T]) Option[T]
	Iter() <-chan T
}

func NewOption[T any](v *T) Option[T] {
	if v == nil {
		return NewNone[T]()
	}

	return NewSome(*v)
}
