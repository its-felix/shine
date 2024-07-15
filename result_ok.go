package shine

import "io"

type Ok[T any] [1]T

func (o Ok[T]) UnwrapOr(def T) T {
	return o[0]
}

func (o Ok[T]) UnwrapOrDefault() T {
	return o[0]
}

func (o Ok[T]) UnwrapOrElse(fn func(e error) T) T {
	return o[0]
}

func (o Ok[T]) AndThen(fn func(v T) Result[T]) Result[T] {
	return fn(o[0])
}

func (o Ok[T]) Ok() Option[T] {
	return NewSome[T](o[0])
}

func (o Ok[T]) Err() Option[error] {
	return NewNone[error]()
}

func (o Ok[T]) Close() error {
	if cl, ok := any(o[0]).(io.Closer); ok {
		return cl.Close()
	}

	return nil
}

func (Ok[T]) result() {

}

func (o Ok[T]) Value() T {
	return o[0]
}

func NewOk[T any](v T) Ok[T] {
	return Ok[T]{v}
}

var _ Result[struct{}] = (*Ok[struct{}])(nil)
