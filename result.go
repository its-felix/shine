package shine

type Result[T any] struct {
	v T
	e error
}

// RMap invokes the given function to transform the value of this Result if it represents Ok and returns a new Result (Ok); returns the Result (Err) with the same error otherwise
func RMap[T any, E any](r Result[T], fnc func(T) E) Result[E] {
	if r.IsOk() {
		return NewOk(fnc(r.Unwrap()))
	} else {
		return NewErr[E](r.UnwrapErr())
	}
}

// RAndThen invokes the given function to form a new Result if the given Result is Ok; returns the Result (Err) with the same error otherwise
func RAndThen[T any, E any](r Result[T], fnc func(T) Result[E]) Result[E] {
	if r.IsOk() {
		return fnc(r.Unwrap())
	} else {
		return NewErr[E](r.UnwrapErr())
	}
}

// IsOk returns true if this Result represents Ok; returns false otherwise
func (r Result[T]) IsOk() bool {
	return r.e == nil
}

// IsErr returns true if this Result represents Err; returns false otherwise
func (r Result[T]) IsErr() bool {
	return r.e != nil
}

// Expect returns this Result underlying value if it represents Ok; panics with the given message otherwise
func (r Result[T]) Expect(msg string) T {
	if r.e != nil {
		panic(msg)
	}

	return r.v
}

// Unwrap returns this Result underlying value if it represents Ok; panics with a generic message otherwise
func (r Result[T]) Unwrap() T {
	if r.e != nil {
		panic("Unwrap on Err")
	}

	return r.v
}

// UnwrapOr returns this Result underlying value if it represents Ok; returns the given default value otherwise
func (r Result[T]) UnwrapOr(def T) T {
	if r.e == nil {
		return r.v
	} else {
		return def
	}
}

// UnwrapOrDefault returns this Result underlying value if it represents Ok; returns the default value for this type otherwise
func (r Result[T]) UnwrapOrDefault() T {
	if r.e == nil {
		return r.v
	} else {
		var def T
		return def
	}
}

// UnwrapOrElse returns this Result underlying value if it represents Ok; returns the result of invoking the given function otherwise
func (r Result[T]) UnwrapOrElse(fnc func() T) T {
	if r.e == nil {
		return r.v
	} else {
		return fnc()
	}
}

// ExpectErr returns this Result underlying error if it represents Err; panics with the given message otherwise
func (r Result[T]) ExpectErr(msg string) error {
	if r.e == nil {
		panic(msg)
	}

	return r.e
}

// UnwrapErr returns this Result underlying error if it represents Err; panics with a generic message otherwise
func (r Result[T]) UnwrapErr() error {
	if r.e == nil {
		panic("UnwrapErr on Ok")
	}

	return r.e
}

// UnwrapBoth returns (value, nil) if this Result represents Ok; returns (<default value for T>, err) otherwise
func (r Result[T]) UnwrapBoth() (T, error) {
	return r.v, r.e
}

// Err returns this Result underlying error as Option (Some) if it represents Err; returns Option (None) otherwise
func (r Result[T]) Err() Option[error] {
	if r.e == nil {
		return NewNone[error]()
	} else {
		return NewSome(r.e)
	}
}

// Ok returns this Result underlying value as Option (Some) if it represents Ok; returns Option (None) otherwise
func (r Result[T]) Ok() Option[T] {
	if r.e == nil {
		return NewNone[T]()
	} else {
		return NewSome(r.v)
	}
}

// Iter returns a channel with this Result underlying value if it represents Ok; returns an empty channel otherwise
func (r Result[T]) Iter() <-chan T {
	ch := make(chan T, 1)

	if r.e == nil {
		ch <- r.v
	}

	close(ch)

	return ch
}

// Map invokes the given function to transform the value of this Result if it represents Ok and returns a new Result (Ok); returns the Result (Err) with the same error otherwise
func (r Result[T]) Map(fnc func(T) T) Result[T] {
	if r.e == nil {
		return NewOk[T](fnc(r.v))
	} else {
		return r
	}
}

// MapAny invokes the given function to transform the value of this Result if it represents Ok and returns a new Result (Ok); returns the Result (Err) with the same error otherwise
func (r Result[T]) MapAny(fnc func(T) any) Result[any] {
	if r.e == nil {
		return NewOk[any](fnc(r.v))
	} else {
		return NewErr[any](r.e)
	}
}

// AndThen invokes the given function to form a new Result if the given Result is Ok; returns the Result (Err) with the same error otherwise
func (r Result[T]) AndThen(fnc func(T) Result[T]) Result[T] {
	if r.e == nil {
		return fnc(r.v)
	} else {
		return r
	}
}

// AndThenAny invokes the given function to form a new Result if the given Result is Ok; returns the Result (Err) with the same error otherwise
func (r Result[T]) AndThenAny(fnc func(T) Result[any]) Result[any] {
	if r.e == nil {
		return fnc(r.v)
	} else {
		return NewErr[any](r.e)
	}
}

// OrElse returns this Result unmodified if it represents Ok; returns the result of invoking the given function with the underlying error otherwise
func (r Result[T]) OrElse(fnc func(error) Result[T]) Result[T] {
	if r.e == nil {
		return r
	} else {
		return fnc(r.e)
	}
}

// NewResult returns a Result (Ok) with the given value if the given error is nil; returns a Result (Err) with the given error otherwise
func NewResult[T any](v T, err error) Result[T] {
	if err != nil {
		return NewErr[T](err)
	} else {
		return NewOk(v)
	}
}

// NewResult2 see NewResult
func NewResult2(v1 any, v2 any, err error) Result[[]any] {
	if err != nil {
		return NewErr[[]any](err)
	} else {
		return NewOk([]any{v1, v2})
	}
}

// NewResult3 see NewResult
func NewResult3(v1 any, v2 any, v3 any, err error) Result[[]any] {
	if err != nil {
		return NewErr[[]any](err)
	} else {
		return NewOk([]any{v1, v2, v2})
	}
}

// NewResult4 see NewResult
func NewResult4(v1 any, v2 any, v3 any, v4 any, err error) Result[[]any] {
	if err != nil {
		return NewErr[[]any](err)
	} else {
		return NewOk([]any{v1, v2, v3, v4})
	}
}

// NewResult5 see NewResult
func NewResult5(v1 any, v2 any, v3 any, v4 any, v5 any, err error) Result[[]any] {
	if err != nil {
		return NewErr[[]any](err)
	} else {
		return NewOk([]any{v1, v2, v3, v4, v5})
	}
}

// NewResultVararg returns a Result (Err) if the last argument is an error and is not nil; otherwise returns Result (Ok) with all arguments (if the last arguments is not an error) or with all except the last argument (if it's an error)
func NewResultVararg(v ...any) Result[[]any] {
	l := len(v)

	if l == 0 {
		return NewOk([]any{})
	}

	lastArg := v[l-1]
	switch lastArg := lastArg.(type) {
	case error:
		if lastArg == nil {
			return NewOk(v[:l-1])
		} else {
			return NewErr[[]any](lastArg)
		}
	default:
		return NewOk(v)
	}
}

// NewOk returns a Result (Ok) with the given value
func NewOk[T any](v T) Result[T] {
	return Result[T]{
		v: v,
	}
}

// NewErr returns a Result (Err) with the given error
func NewErr[T any](err error) Result[T] {
	return Result[T]{
		e: err,
	}
}
