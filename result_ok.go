package shine

type Ok[T any, E error] struct {
	v T
}

func (o Ok[T, E]) IsOk() bool {
	return true
}

func (o Ok[T, E]) IsOkAnd(fn func(v T) bool) bool {
	return fn(o.v)
}

func (o Ok[T, E]) IsErr() bool {
	return false
}

func (o Ok[T, E]) IsErrAnd(fn func(e E) bool) bool {
	return false
}

func (o Ok[T, E]) Get() (T, E) {
	var def E
	return o.v, def
}

func (o Ok[T, E]) Expect(panicV any) T {
	return o.v
}

func (o Ok[T, E]) ExpectErr(panicV any) E {
	panic(panicV)
}

func (o Ok[T, E]) Unwrap() T {
	return o.v
}

func (o Ok[T, E]) UnwrapErr() E {
	panic("UnwrapErr on Ok")
}

func (o Ok[T, E]) UnwrapOr(def T) T {
	return o.v
}

func (o Ok[T, E]) UnwrapOrDefault() T {
	return o.v
}

func (o Ok[T, E]) UnwrapOrElse(fn func() T) T {
	return o.v
}

func (o Ok[T, E]) And(other Result[T, E]) Result[T, E] {
	return other
}

func (o Ok[T, E]) AndThen(fn func(v T) Result[T, E]) Result[T, E] {
	return fn(o.v)
}

func (o Ok[T, E]) Map(fn func(v T) T) Result[T, E] {
	return NewOk[T, E](fn(o.v))
}

func (o Ok[T, E]) MapErr(fn func(e E) E) Result[T, E] {
	return o
}

func (o Ok[T, E]) MapOr(def T, fn func(v T) T) T {
	return fn(o.v)
}

func (o Ok[T, E]) MapOrElse(fnOk func(v T) T, fnErr func(e E) T) T {
	return fnOk(o.v)
}

func (o Ok[T, E]) Ok() Option[T] {
	return NewSome[T](o.v)
}

func (o Ok[T, E]) Err() Option[E] {
	return NewNone[E]()
}

func (o Ok[T, E]) Or(other Result[T, E]) Result[T, E] {
	return o
}

func (o Ok[T, E]) OrElse(fn func(e E) Result[T, E]) Result[T, E] {
	return o
}

func (o Ok[T, E]) Iter() <-chan T {
	ch := make(chan T, 1)
	ch <- o.v
	close(ch)

	return ch
}

func NewOk[T any, E error](v T) Ok[T, E] {
	return Ok[T, E]{v}
}

var _ Result[struct{}, error] = NewOk[struct{}, error](struct{}{})
