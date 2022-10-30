package shine

type Tuple[T1 any, T2 any] struct {
	V1 T1
	V2 T2
}

type Container[T any] interface {
	Expect(msg string) T
	Unwrap() T
	UnwrapOr(def T) T
	UnwrapOrDefault() T
	UnwrapOrElse(fnc func() T) T
	Iter() <-chan T
}
