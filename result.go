package shine

import "reflect"

type Result[T any, E error] interface {
	IsOk() bool
	IsOkAnd(fn func(v T) bool) bool
	IsErr() bool
	IsErrAnd(fn func(v E) bool) bool
	Expect(panicV any) T
	ExpectErr(panicV any) E
	Unwrap() T
	UnwrapErr() E
	UnwrapOr(def T) T
	UnwrapOrDefault() T
	UnwrapOrElse(fn func() T) T
	And(other Result[T, E]) Result[T, E]
	AndThen(fn func(v T) Result[T, E]) Result[T, E]
	Map(fn func(v T) T) Result[T, E]
	MapErr(fn func(e E) E) Result[T, E]
	MapOr(def T, fn func(v T) T) T
	MapOrElse(fnOk func(v T) T, fnErr func(e E) T) T
	Ok() Option[T]
	Err() Option[E]
	Or(other Result[T, E]) Result[T, E]
	OrElse(fn func(e E) Result[T, E]) Result[T, E]
	Iter() <-chan T
}

func NewResult[T any, E error](v T, err E) Result[T, E] {
	errV := reflect.ValueOf(err)
	switch errV.Kind() {
	case reflect.Ptr, reflect.Interface, reflect.Slice, reflect.Map, reflect.Chan, reflect.Func:
		if !errV.IsNil() {
			return NewErr[T](err)
		}

	default:
	}

	return NewOk[T, E](v)
}
