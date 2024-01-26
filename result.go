package shine

type Result[T any] interface {
	IsOk() bool
	IsOkAnd(fn func(v T) bool) bool
	IsErr() bool
	IsErrAnd(fn func(v error) bool) bool
	Get() (T, error, bool)
	IfPresent(fn func(v T)) bool
	UnwrapOr(def T) T
	UnwrapOrDefault() T
	UnwrapOrElse(fn func(e error) T) T
	And(other Result[T]) Result[T]
	AndThen(fn func(v T) Result[T]) Result[T]
	Map(fn func(v T) T) Result[T]
	MapErr(fn func(e error) error) Result[T]
	MapOr(def T, fn func(v T) T) T
	MapOrElse(fnOk func(v T) T, fnErr func(e error) T) T
	Ok() Option[T]
	Err() Option[error]
	Or(other Result[T]) Result[T]
	OrElse(fn func(e error) Result[T]) Result[T]
	Iter() <-chan T
	Close() error
}

func NewResult[T any](v T, err error) Result[T] {
	if err != nil && !isNil(err) {
		return NewErr[T](err)
	}

	return NewOk[T](v)
}

func ResMap[T any, R any](r Result[T], fn func(v T) R) Result[R] {
	if v, err, ok := r.Get(); ok {
		return NewOk[R](fn(v))
	} else {
		return NewErr[R](err)
	}
}

func ResAndThen[T any, R any](r Result[T], fn func(v T) Result[R]) Result[R] {
	if v, err, ok := r.Get(); ok {
		return fn(v)
	} else {
		return NewErr[R](err)
	}
}
