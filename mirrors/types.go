package mirrors

import "reflect"

func T[T any]() reflect.Type {
	return reflect.TypeFor[T]()
}

func TypeOf[T any]() reflect.Type {
	return reflect.TypeFor[T]()
}

func Empty[T any]() T {
	var t T
	return t
}
