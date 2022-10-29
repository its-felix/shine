package shine

type Result[T any] interface {
	Container[T, Result[T], Result[any]]
	IsOk() bool
	IsErr() bool
	ExpectErr(msg string) error
	UnwrapErr() error
	UnwrapBoth() (T, error)
	OrElse(fnc func(error) Result[T]) Result[T]
	Err() Option[error]
	Ok() Option[T]
}

func ResultMap[T any, E any](r Result[T], fnc func(T) E) Result[E] {
	var ret Result[E]

	switch r.(type) {
	case Ok[T]:
		ret = NewOk[E](fnc(r.Unwrap()))
	default:
		ret = NewErr[E](r.UnwrapErr())
	}

	return ret
}

func ResultAndThen[T any, E any](r Result[T], fnc func(T) Result[E]) Result[E] {
	var ret Result[E]

	switch r.(type) {
	case Ok[T]:
		ret = fnc(r.Unwrap())
	default:
		ret = NewErr[E](r.UnwrapErr())
	}

	return ret
}

type Ok[T any] struct {
	v T
}

func (ok Ok[T]) IsOk() bool {
	return true
}

func (ok Ok[T]) IsErr() bool {
	return false
}

func (ok Ok[T]) Expect(msg string) T {
	return ok.v
}

func (ok Ok[T]) Unwrap() T {
	return ok.v
}

func (ok Ok[T]) UnwrapOr(def T) T {
	return ok.v
}

func (ok Ok[T]) UnwrapOrDefault() T {
	return ok.v
}

func (ok Ok[T]) UnwrapOrElse(fnc func() T) T {
	return ok.v
}

func (ok Ok[T]) ExpectErr(msg string) error {
	panic(msg)
}

func (ok Ok[T]) UnwrapErr() error {
	panic("UnwrapErr on Ok")
}

func (ok Ok[T]) UnwrapBoth() (T, error) {
	return ok.v, nil
}

func (ok Ok[T]) Err() Option[error] {
	return NewNone[error]()
}

func (ok Ok[T]) Ok() Option[T] {
	return NewSome(ok.v)
}

func (ok Ok[T]) Iter() <-chan T {
	ch := make(chan T)
	ch <- ok.v
	close(ch)

	return ch
}

func (ok Ok[T]) Map(fnc func(T) T) Result[T] {
	return NewOk(fnc(ok.v))
}

func (ok Ok[T]) MapAny(fnc func(T) any) Result[any] {
	return NewOk(fnc(ok.v))
}

func (ok Ok[T]) AndThen(fnc func(T) Result[T]) Result[T] {
	return fnc(ok.v)
}

func (ok Ok[T]) AndThenAny(fnc func(T) Result[any]) Result[any] {
	return fnc(ok.v)
}

func (ok Ok[T]) OrElse(fnc func(error) Result[T]) Result[T] {
	return ok
}

type Err[T any] struct {
	err error
}

func (e Err[T]) Error() string {
	return e.err.Error()
}

func (e Err[T]) IsOk() bool {
	return false
}

func (e Err[T]) IsErr() bool {
	return true
}

func (e Err[T]) Expect(msg string) T {
	panic(msg)
}

func (e Err[T]) Unwrap() T {
	panic("Unwrap on Err")
}

func (e Err[T]) UnwrapOr(def T) T {
	return def
}

func (e Err[T]) UnwrapOrDefault() T {
	var def T
	return def
}

func (e Err[T]) UnwrapOrElse(fnc func() T) T {
	return fnc()
}

func (e Err[T]) ExpectErr(msg string) error {
	return e.err
}

func (e Err[T]) UnwrapErr() error {
	return e.err
}

func (e Err[T]) UnwrapBoth() (T, error) {
	var def T
	return def, e.err
}

func (e Err[T]) Err() Option[error] {
	return NewSome(e.err)
}

func (e Err[T]) Ok() Option[T] {
	return NewNone[T]()
}

func (e Err[T]) Iter() <-chan T {
	ch := make(chan T)
	close(ch)

	return ch
}

func (e Err[T]) Map(fnc func(T) T) Result[T] {
	return e
}

func (e Err[T]) MapAny(fnc func(T) any) Result[any] {
	return Err[any](e)
}

func (e Err[T]) AndThen(fnc func(T) Result[T]) Result[T] {
	return e
}

func (e Err[T]) AndThenAny(fnc func(T) Result[any]) Result[any] {
	return Err[any](e)
}

func (e Err[T]) OrElse(fnc func(error) Result[T]) Result[T] {
	return fnc(e.err)
}

func NewResult[T any](v T, err error) Result[T] {
	if err != nil {
		return NewErr[T](err)
	} else {
		return NewOk(v)
	}
}

func NewResultWithV[T any](fnc func() T, err error) Result[T] {
	if err != nil {
		return NewErr[T](err)
	} else {
		return NewOk(fnc())
	}
}

func NewResultWithE[T any](v T, fnc func() error) Result[T] {
	return NewResult(v, fnc())
}

func NewOk[T any](v T) Result[T] {
	return Ok[T]{v}
}

func NewErr[T any](err error) Result[T] {
	return Err[T]{err}
}
