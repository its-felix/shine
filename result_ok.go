package shine

import "io"

type Ok[T any] struct {
	v T
}

func (o Ok[T]) IsOk() bool {
	return true
}

func (o Ok[T]) IsOkAnd(fn func(v T) bool) bool {
	return fn(o.v)
}

func (o Ok[T]) IsErr() bool {
	return false
}

func (o Ok[T]) IsErrAnd(fn func(e error) bool) bool {
	return false
}

func (o Ok[T]) Get() (T, error, bool) {
	return o.v, nil, true
}

func (o Ok[T]) IfPresent(fn func(v T)) bool {
	fn(o.v)
	return true
}

func (o Ok[T]) UnwrapOr(def T) T {
	return o.v
}

func (o Ok[T]) UnwrapOrDefault() T {
	return o.v
}

func (o Ok[T]) UnwrapOrElse(fn func(e error) T) T {
	return o.v
}

func (o Ok[T]) And(other Result[T]) Result[T] {
	return other
}

func (o Ok[T]) AndThen(fn func(v T) Result[T]) Result[T] {
	return fn(o.v)
}

func (o Ok[T]) Map(fn func(v T) T) Result[T] {
	return NewOk[T](fn(o.v))
}

func (o Ok[T]) MapErr(fn func(e error) error) Result[T] {
	return o
}

func (o Ok[T]) MapOr(def T, fn func(v T) T) T {
	return fn(o.v)
}

func (o Ok[T]) MapOrElse(fnOk func(v T) T, fnErr func(e error) T) T {
	return fnOk(o.v)
}

func (o Ok[T]) Ok() Option[T] {
	return NewSome[T](o.v)
}

func (o Ok[T]) Err() Option[error] {
	return NewNone[error]()
}

func (o Ok[T]) Or(other Result[T]) Result[T] {
	return o
}

func (o Ok[T]) OrElse(fn func(e error) Result[T]) Result[T] {
	return o
}

func (o Ok[T]) Iter() <-chan T {
	ch := make(chan T, 1)
	ch <- o.v
	close(ch)

	return ch
}

func (o Ok[T]) Close() error {
	if cl, ok := any(o.v).(io.Closer); ok {
		return cl.Close()
	}

	return nil
}

func (o Ok[T]) Value() T {
	return o.v
}

func NewOk[T any](v T) Ok[T] {
	return Ok[T]{v}
}

var _ Result[struct{}] = Ok[struct{}]{}
