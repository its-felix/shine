package shine

type Tuple[T1 any, T2 any] struct {
	v1 T1
	v2 T2
}

type ContainerBase[T any] interface {
	Expect(msg string) T
	Unwrap() T
	UnwrapOr(def T) T
	UnwrapOrDefault() T
	UnwrapOrElse(fnc func() T) T
	Iter() <-chan T
}

type ContainerExtensionSelf[T any, Self any] interface {
	Map(fnc func(T) T) Self
	AndThen(fnc func(T) Self) Self
}

type ContainerExtensionAny[T any, SelfAny any] interface {
	MapAny(fnc func(T) any) SelfAny
	AndThenAny(fnc func(T) SelfAny) SelfAny
}

type Container[T any, Self any, SelfAny any] interface {
	ContainerBase[T]
	ContainerExtensionSelf[T, Self]
	ContainerExtensionAny[T, SelfAny]
}
