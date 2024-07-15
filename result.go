package shine

import "io"

type Result[T any] interface {
	io.Closer
	UnwrapOr(def T) T
	UnwrapOrDefault() T
	UnwrapOrElse(fn func(e error) T) T
	AndThen(fn func(v T) Result[T]) Result[T]
	Ok() Option[T]
	Err() Option[error]
	result()
}

func NewResult[T any](v T, err error) Result[T] {
	if err != nil && !isNil(err) {
		return NewErr[T](err)
	}

	return NewOk[T](v)
}
