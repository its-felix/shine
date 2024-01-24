package shine

import "errors"

type Err[T any, E error] struct {
	err E
}

func (e Err[T, E]) IsOk() bool {
	return false
}

func (e Err[T, E]) IsOkAnd(fn func(v T) bool) bool {
	return false
}

func (e Err[T, E]) IsErr() bool {
	return true
}

func (e Err[T, E]) IsErrAnd(fn func(e E) bool) bool {
	return fn(e.err)
}

func (e Err[T, E]) Get() (T, E, bool) {
	var def T
	return def, e.err, false
}

func (e Err[T, E]) Expect(panicV any) T {
	panic(panicV)
}

func (e Err[T, E]) ExpectErr(panicV any) E {
	return e.err
}

func (e Err[T, E]) Unwrap() T {
	panic("Unwrap on Err")
}

func (e Err[T, E]) UnwrapErr() E {
	return e.err
}

func (e Err[T, E]) UnwrapOr(def T) T {
	return def
}

func (e Err[T, E]) UnwrapOrDefault() T {
	var def T
	return def
}

func (e Err[T, E]) UnwrapOrElse(fn func() T) T {
	return fn()
}

func (e Err[T, E]) And(other Result[T, E]) Result[T, E] {
	return e
}

func (e Err[T, E]) AndThen(fn func(v T) Result[T, E]) Result[T, E] {
	return e
}

func (e Err[T, E]) Map(fn func(v T) T) Result[T, E] {
	return e
}

func (e Err[T, E]) MapErr(fn func(e E) E) Result[T, E] {
	return NewErr[T](fn(e.err))
}

func (e Err[T, E]) MapOr(def T, fn func(v T) T) T {
	return def
}

func (e Err[T, E]) MapOrElse(fnOk func(v T) T, fnErr func(e E) T) T {
	return fnErr(e.err)
}

func (e Err[T, E]) Ok() Option[T] {
	return NewNone[T]()
}

func (e Err[T, E]) Err() Option[E] {
	return NewSome[E](e.err)
}

func (e Err[T, E]) Or(other Result[T, E]) Result[T, E] {
	return other
}

func (e Err[T, E]) OrElse(fn func(e E) Result[T, E]) Result[T, E] {
	return fn(e.err)
}

func (e Err[T, E]) Iter() <-chan T {
	ch := make(chan T)
	close(ch)

	return ch
}

func NewErr[T any, E error](err E) Err[T, E] {
	return Err[T, E]{err}
}

var _ Result[struct{}, error] = NewErr[struct{}](errors.ErrUnsupported)
