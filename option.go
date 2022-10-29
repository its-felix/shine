package shine

type Option[T any] interface {
	Container[T, Option[T], Option[any]]
	IsSome() bool
	IsNone() bool
	OkOr(e error) Result[T]
	OkOrElse(fnc func() error) Result[T]
	Filter(fnc func(T) bool) Option[T]
	OrElse(fnc func() Option[T]) Option[T]
	ZipAny(other Option[any]) Option[Tuple[any, any]]
	ZipWithAny(other Option[any], fnc func(T, any) Tuple[any, any]) Option[Tuple[any, any]]

	//Zip(other Option[T]) Option[Tuple[T, T]]
	//ZipAny(other Option[any]) Option[Tuple[T, any]]
	//ZipWith(other Option[T], fnc func(T, T) Tuple[T, T]) Option[Tuple[T, T]]
	//ZipWithAny(other Option[any], fnc func(T, any) Tuple[T, any]) Option[Tuple[T, any]]
}

var __none = None[any]{}

func OptionMap[T any, E any](o Option[T], fnc func(T) E) Option[E] {
	var r Option[E]

	switch o.(type) {
	case Some[T]:
		r = NewSome(fnc(o.Unwrap()))
	default:
		r = NewNone[E]()
	}

	return r
}

func OptionZip[T any, E any](o1 Option[T], o2 Option[E]) Option[Tuple[T, E]] {
	return OptionZipWith(o1, o2, func(t T, e E) Tuple[T, E] {
		return Tuple[T, E]{o1.Unwrap(), o2.Unwrap()}
	})
}

func OptionZipWith[T any, E any, O any](o1 Option[T], o2 Option[E], zip func(T, E) O) Option[O] {
	if o1.IsSome() && o2.IsSome() {
		return NewSome(zip(o1.Unwrap(), o2.Unwrap()))
	} else {
		return NewNone[O]()
	}
}

func OptionAndThen[T any, E any](o Option[T], fnc func(T) Option[E]) Option[E] {
	var r Option[E]

	switch o.(type) {
	case Some[T]:
		r = fnc(o.Unwrap())
	default:
		r = NewNone[E]()
	}

	return r
}

type Some[T any] struct {
	v T
}

func (s Some[T]) IsSome() bool {
	return true
}

func (s Some[T]) IsNone() bool {
	return false
}

func (s Some[T]) Expect(msg string) T {
	return s.v
}

func (s Some[T]) Unwrap() T {
	return s.v
}

func (s Some[T]) UnwrapOr(def T) T {
	return s.v
}

func (s Some[T]) UnwrapOrDefault() T {
	return s.v
}

func (s Some[T]) UnwrapOrElse(fnc func() T) T {
	return s.v
}

func (s Some[T]) OkOr(e error) Result[T] {
	return NewOk(s.v)
}

func (s Some[T]) OkOrElse(fnc func() error) Result[T] {
	return NewOk(s.v)
}

func (s Some[T]) Filter(fnc func(T) bool) Option[T] {
	if fnc(s.v) {
		return s
	} else {
		return NewNone[T]()
	}
}

func (s Some[T]) Iter() <-chan T {
	ch := make(chan T)
	ch <- s.v
	close(ch)

	return ch
}

func (s Some[T]) Map(fnc func(T) T) Option[T] {
	return NewSome(fnc(s.v))
}

func (s Some[T]) MapAny(fnc func(T) any) Option[any] {
	return NewSome(fnc(s.v))
}

func (s Some[T]) AndThen(fnc func(T) Option[T]) Option[T] {
	return fnc(s.v)
}

func (s Some[T]) AndThenAny(fnc func(T) Option[any]) Option[any] {
	return fnc(s.v)
}

func (s Some[T]) OrElse(fnc func() Option[T]) Option[T] {
	return s
}

func (s Some[T]) ZipAny(other Option[any]) Option[Tuple[any, any]] {
	var r Option[Tuple[any, any]]

	switch other := other.(type) {
	case Some[any]:
		r = NewSome(Tuple[any, any]{s.v, other.v})
	default:
		r = NewNone[Tuple[any, any]]()
	}

	return r
}

func (s Some[T]) ZipWithAny(other Option[any], fnc func(T, any) Tuple[any, any]) Option[Tuple[any, any]] {
	var r Option[Tuple[any, any]]

	switch other := other.(type) {
	case Some[any]:
		r = NewSome(fnc(s.v, other.v))
	default:
		r = NewNone[Tuple[any, any]]()
	}

	return r
}

//func (s Some[T]) Zip(other Option[T]) Option[Tuple[T, T]] {
//	var r Option[Tuple[T, T]]
//
//	switch other := other.(type) {
//	case Some[T]:
//		r = NewSome(Tuple[T, T]{s.v, other.v})
//	default:
//		r = NewNone[Tuple[T, T]]()
//	}
//
//	return r
//}
//
//func (s Some[T]) ZipAny(other Option[any]) Option[Tuple[T, any]] {
//	var r Option[Tuple[T, any]]
//
//	switch other := other.(type) {
//	case Some[any]:
//		r = NewSome(Tuple[T, any]{s.v, other.v})
//	default:
//		r = NewNone[Tuple[T, any]]()
//	}
//
//	return r
//}
//
//func (s Some[T]) ZipWith(other Option[T], fnc func(T, T) Tuple[T, T]) Option[Tuple[T, T]] {
//	var r Option[Tuple[T, T]]
//
//	switch other := other.(type) {
//	case Some[T]:
//		r = NewSome(fnc(s.v, other.v))
//	default:
//		r = NewNone[Tuple[T, T]]()
//	}
//
//	return r
//}
//
//func (s Some[T]) ZipWithAny(other Option[any], fnc func(T, any) Tuple[T, any]) Option[Tuple[T, any]] {
//	var r Option[Tuple[T, any]]
//
//	switch other := other.(type) {
//	case Some[any]:
//		r = NewSome(fnc(s.v, other.v))
//	default:
//		r = NewNone[Tuple[T, any]]()
//	}
//
//	return r
//}

type None[T any] struct{}

func (n None[T]) IsSome() bool {
	return false
}

func (n None[T]) IsNone() bool {
	return true
}

func (n None[T]) Expect(msg string) T {
	panic(msg)
}

func (n None[T]) Unwrap() T {
	panic("Unwrap on None")
}

func (n None[T]) UnwrapOr(def T) T {
	return def
}

func (n None[T]) UnwrapOrDefault() T {
	var def T
	return def
}

func (n None[T]) UnwrapOrElse(fnc func() T) T {
	return fnc()
}

func (n None[T]) OkOr(e error) Result[T] {
	return NewErr[T](e)
}

func (n None[T]) OkOrElse(fnc func() error) Result[T] {
	return NewErr[T](fnc())
}

func (n None[T]) Filter(fnc func(T) bool) Option[T] {
	return n
}

func (n None[T]) Iter() <-chan T {
	ch := make(chan T)
	close(ch)

	return ch
}

func (n None[T]) Map(fnc func(T) T) Option[T] {
	return n
}

func (n None[T]) MapAny(fnc func(T) any) Option[any] {
	return None[any](n)
}

func (n None[T]) AndThen(fnc func(T) Option[T]) Option[T] {
	return n
}

func (n None[T]) AndThenAny(fnc func(T) Option[any]) Option[any] {
	return None[any](n)
}

func (n None[T]) OrElse(fnc func() Option[T]) Option[T] {
	return fnc()
}

func (n None[T]) ZipAny(other Option[any]) Option[Tuple[any, any]] {
	return NewNone[Tuple[any, any]]()
}

func (n None[T]) ZipWithAny(other Option[any], fnc func(T, any) Tuple[any, any]) Option[Tuple[any, any]] {
	return NewNone[Tuple[any, any]]()
}

//func (n None[T]) Zip(other Option[T]) Option[Tuple[T, T]] {
//	return NewNone[Tuple[T, T]]()
//}
//
//func (n None[T]) ZipAny(other Option[any]) Option[Tuple[T, any]] {
//	return NewNone[Tuple[T, any]]()
//}
//
//func (n None[T]) ZipWith(other Option[T], fnc func(T, T) Tuple[T, T]) Option[Tuple[T, T]] {
//	return NewNone[Tuple[T, T]]()
//}
//
//func (n None[T]) ZipWithAny(other Option[any], fnc func(T, any) Tuple[T, any]) Option[Tuple[T, any]] {
//	return NewNone[Tuple[T, any]]()
//}

func NewOption[T any](v *T) Option[T] {
	if v == nil {
		return NewNone[T]()
	} else {
		return NewSome(*v)
	}
}

func NewSome[T any](v T) Option[T] {
	return Some[T]{v}
}

func NewNone[T any]() Option[T] {
	return None[T](__none)
}
