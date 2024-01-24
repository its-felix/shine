package shine

type Option[T any] interface {
	IsSome() bool
	IsNone() bool
	Get() (T, bool)
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

func NewOptionOf[T any](v T) Option[T] {
	return NewOptionFrom(v, !isNil(v))
}

func NewOptionFrom[T any](v T, ok bool) Option[T] {
	if ok {
		return NewSome(v)
	}

	return NewNone[T]()
}

func OptMap[T any, R any](o Option[T], fn func(v T) R) Option[R] {
	if v, ok := o.Get(); ok {
		return NewSome(fn(v))
	}

	return NewNone[R]()
}

func OptAndThen[T any, R any](o Option[T], fn func(v T) Option[R]) Option[R] {
	if v, ok := o.Get(); ok {
		return fn(v)
	}

	return NewNone[R]()
}
