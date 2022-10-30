package shine

type Option[T any] struct {
	some bool
	v    T
}

// OMap invokes the given function to transform the value of this Option if it represents a Some and returns a new Some; returns None for None
func OMap[T any, E any](o Option[T], fnc func(T) E) Option[E] {
	if o.IsSome() {
		return NewSome(fnc(o.Unwrap()))
	} else {
		return NewNone[E]()
	}
}

// Zip returns a new Option with both given Options values if both are Some; returns None if any of the two is None
func Zip[T any, E any](o1 Option[T], o2 Option[E]) Option[Tuple[T, E]] {
	return ZipWith(o1, o2, func(t T, e E) Tuple[T, E] {
		return Tuple[T, E]{o1.Unwrap(), o2.Unwrap()}
	})
}

// ZipWith invokes the given function to form a value for a new Option if both are Some; returns None if any of the two is None
func ZipWith[T any, E any, O any](o1 Option[T], o2 Option[E], zip func(T, E) O) Option[O] {
	if o1.IsSome() && o2.IsSome() {
		return NewSome(zip(o1.Unwrap(), o2.Unwrap()))
	} else {
		return NewNone[O]()
	}
}

// OAndThen invokes the given function to form a new Option if the given Option is Some; returns None for None
func OAndThen[T any, E any](o Option[T], fnc func(T) Option[E]) Option[E] {
	if o.IsSome() {
		return fnc(o.Unwrap())
	} else {
		return NewNone[E]()
	}
}

// IsSome returns true if this Option represents a Some
func (o Option[T]) IsSome() bool {
	return o.some
}

// IsNone returns true if this Option represents a None
func (o Option[T]) IsNone() bool {
	return !o.some
}

// Expect returns this Option underlying value if it represents a Some; panics with the given message otherwise
func (o Option[T]) Expect(msg string) T {
	if !o.some {
		panic(msg)
	}

	return o.v
}

// Unwrap returns this Option underlying value if it represents a Some; panics with a generic message otherwise
func (o Option[T]) Unwrap() T {
	if !o.some {
		panic("Unwrap on None")
	}

	return o.v
}

// UnwrapOr returns this Option underlying value if it represents a Some; returns the given default value otherwise
func (o Option[T]) UnwrapOr(def T) T {
	if o.some {
		return o.v
	} else {
		return def
	}
}

// UnwrapOrDefault returns this Option underlying value if it represents a Some; returns the default value for this type otherwise
func (o Option[T]) UnwrapOrDefault() T {
	if o.some {
		return o.v
	} else {
		var def T
		return def
	}
}

// UnwrapOrElse returns this Option underlying value if it represents a Some; returns the result of invoking the given function otherwise
func (o Option[T]) UnwrapOrElse(fnc func() T) T {
	if o.some {
		return o.v
	} else {
		return fnc()
	}
}

// OkOr returns this Option underlying value as a Result (Ok) if it represents a Some; returns a Result (Err) with the given error otherwise
func (o Option[T]) OkOr(e error) Result[T] {
	if o.some {
		return NewOk[T](o.v)
	} else {
		return NewErr[T](e)
	}
}

// OkOrElse returns this Option underlying value as a Result (Ok) if it represents a Some; returns a Result (Err) with the error returned by the given function otherwise
func (o Option[T]) OkOrElse(fnc func() error) Result[T] {
	if o.some {
		return NewOk[T](o.v)
	} else {
		return NewErr[T](fnc())
	}
}

// Filter returns this Option unmodified if it represents a Some and the given function returns true given the value; returns Option (None) otherwise
func (o Option[T]) Filter(fnc func(T) bool) Option[T] {
	if o.some && fnc(o.v) {
		return o
	} else {
		return NewNone[T]()
	}
}

// Iter returns a channel with this Option underlying value if it represents a Some; returns an empty channel otherwise
func (o Option[T]) Iter() <-chan T {
	ch := make(chan T, 1)

	if o.some {
		ch <- o.v
	}

	close(ch)

	return ch
}

// Map invokes the given function to transform the value of this Option is it's Some and returns a new Some; returns None for None
func (o Option[T]) Map(fnc func(T) T) Option[T] {
	if o.some {
		return NewSome(fnc(o.v))
	} else {
		return NewNone[T]()
	}
}

// MapAny invokes the given function to transform the value of this Option is it's Some and returns a new Some; returns None for None
func (o Option[T]) MapAny(fnc func(T) any) Option[any] {
	if o.some {
		return NewSome(fnc(o.v))
	} else {
		return NewNone[any]()
	}
}

// AndThen invokes the given function to form a new Option if the given Option is Some; returns None for None
func (o Option[T]) AndThen(fnc func(T) Option[T]) Option[T] {
	if o.some {
		return fnc(o.v)
	} else {
		return NewNone[T]()
	}
}

// AndThenAny invokes the given function to form a new Option if the given Option is Some; returns None for None
func (o Option[T]) AndThenAny(fnc func(T) Option[any]) Option[any] {
	if o.some {
		return fnc(o.v)
	} else {
		return NewNone[any]()
	}
}

// OrElse returns this Option unmodified if it represents a Some; returns the result of invoking the given function otherwise
func (o Option[T]) OrElse(fnc func() Option[T]) Option[T] {
	if o.some {
		return o
	} else {
		return fnc()
	}
}

// ZipAny returns a new Option with both given Options values if both are Some; returns None if any of the two is None
func (o Option[T]) ZipAny(other Option[any]) Option[[2]any] {
	if o.some {
		return NewSome([2]any{o.v, other.v})
	} else {
		return NewNone[[2]any]()
	}
}

// ZipWithAny invokes the given function to form a value for a new Option if both are Some; returns None if any of the two is None
func (o Option[T]) ZipWithAny(other Option[any], fnc func(T, any) [2]any) Option[[2]any] {
	if o.some {
		return NewSome(fnc(o.v, other.v))
	} else {
		return NewNone[[2]any]()
	}
}

// NewOption returns an Option (Some) with the given pointers value if it's not nil; returns Option (None) otherwise
func NewOption[T any](v *T) Option[T] {
	if v == nil {
		return NewNone[T]()
	} else {
		return NewSome(*v)
	}
}

// NewSome returns an Option (Some) with the given value
func NewSome[T any](v T) Option[T] {
	return Option[T]{
		some: true,
		v:    v,
	}
}

// NewNone returns an Option (None)
func NewNone[T any]() Option[T] {
	return Option[T]{}
}
