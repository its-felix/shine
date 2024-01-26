package shine

import (
	"errors"
	"io"
)

type Err[T any] struct {
	error
}

func (e Err[T]) IsOk() bool {
	return false
}

func (e Err[T]) IsOkAnd(fn func(v T) bool) bool {
	return false
}

func (e Err[T]) IsErr() bool {
	return true
}

func (e Err[T]) IsErrAnd(fn func(e error) bool) bool {
	return fn(e)
}

func (e Err[T]) Get() (T, error, bool) {
	var def T
	return def, e, false
}

func (e Err[T]) IfPresent(fn func(v T)) bool {
	return false
}

func (e Err[T]) UnwrapOr(def T) T {
	return def
}

func (e Err[T]) UnwrapOrDefault() T {
	var def T
	return def
}

func (e Err[T]) UnwrapOrElse(fn func(e error) T) T {
	return fn(e)
}

func (e Err[T]) And(other Result[T]) Result[T] {
	return e
}

func (e Err[T]) AndThen(fn func(v T) Result[T]) Result[T] {
	return e
}

func (e Err[T]) Map(fn func(v T) T) Result[T] {
	return e
}

func (e Err[T]) MapErr(fn func(e error) error) Result[T] {
	return NewErr[T](fn(e))
}

func (e Err[T]) MapOr(def T, fn func(v T) T) T {
	return def
}

func (e Err[T]) MapOrElse(fnOk func(v T) T, fnErr func(e error) T) T {
	return fnErr(e)
}

func (e Err[T]) Ok() Option[T] {
	return NewNone[T]()
}

func (e Err[T]) Err() Option[error] {
	return NewSome[error](e)
}

func (e Err[T]) Or(other Result[T]) Result[T] {
	return other
}

func (e Err[T]) OrElse(fn func(e error) Result[T]) Result[T] {
	return fn(e)
}

func (e Err[T]) Iter() <-chan T {
	ch := make(chan T)
	close(ch)

	return ch
}

func (e Err[T]) Close() error {
	if cl, ok := e.error.(io.Closer); ok {
		return cl.Close()
	}

	return nil
}

func (e Err[T]) Is(err error) bool {
	return errors.Is(e.error, err)
}

func (e Err[T]) As(v any) bool {
	// passed value as expected to be a pointer
	return errors.As(e.error, v)
}

func NewErr[T any, E error](err E) Err[T] {
	return Err[T]{err}
}

var _ Result[struct{}] = Err[struct{}]{}
