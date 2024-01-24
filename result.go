package shine

type Result[T any, E error] interface {
	IsOk() bool
	IsOkAnd(fn func(v T) bool) bool
	IsErr() bool
	IsErrAnd(fn func(v E) bool) bool
	Get() (T, E, bool)
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
	if !isNil(err) {
		return NewErr[T](err)
	}

	return NewOk[T, E](v)
}

func ResMap[T any, E error, R any](r Result[T, E], fn func(v T) R) Result[R, E] {
	if v, err, ok := r.Get(); ok {
		return NewOk[R, E](fn(v))
	} else {
		return NewErr[R](err)
	}
}

func ResAndThen[T any, E error, R any](r Result[T, E], fn func(v T) Result[R, E]) Result[R, E] {
	if v, err, ok := r.Get(); ok {
		return fn(v)
	} else {
		return NewErr[R](err)
	}
}

func ResMapErr[T any, E error, R error](r Result[T, E], fn func(err E) R) Result[T, R] {
	if v, err, ok := r.Get(); ok {
		return NewOk[T, R](v)
	} else {
		return NewErr[T](fn(err))
	}
}

func ResOrElse[T any, E error, R error](r Result[T, E], fn func(e E) Result[T, R]) Result[T, R] {
	if v, err, ok := r.Get(); ok {
		return NewOk[T, R](v)
	} else {
		return fn(err)
	}
}
