package shine

import (
	"errors"
	"io"
)

type Err[T any] [1]error

func (e Err[T]) UnwrapOr(def T) T {
	return def
}

func (e Err[T]) UnwrapOrDefault() T {
	var def T
	return def
}

func (e Err[T]) UnwrapOrElse(fn func(e error) T) T {
	return fn(e[0])
}

func (e Err[T]) AndThen(fn func(v T) Result[T]) Result[T] {
	return e
}

func (e Err[T]) Ok() Option[T] {
	return NewNone[T]()
}

func (e Err[T]) Err() Option[error] {
	return NewSome[error](e[0])
}

func (e Err[T]) Close() error {
	if cl, ok := e[0].(io.Closer); ok {
		return cl.Close()
	}

	return nil
}

func (Err[T]) result() {

}

func (e Err[T]) Unwrap() error {
	return e[0]
}

func (e Err[T]) Error() string {
	return e[0].Error()
}

func (e Err[T]) Is(err error) bool {
	return errors.Is(e[0], err)
}

func (e Err[T]) As(v any) bool {
	// passed value as expected to be a pointer
	return errors.As(e[0], v)
}

func NewErr[T any, E error](err E) Err[T] {
	return Err[T]{err}
}

var _ interface {
	Result[struct{}]
	error
} = Err[struct{}]{}
