package main

import "fmt"

type Foo struct {
	Verbosity int
}

type option func(*Foo) interface{}

func (f *Foo) Option(opts ...option) []interface{} {
	var previous []interface{}
	for _, opt := range opts {
		previous = append(previous, opt(f))
	}
	return previous
}

func Verbosity(v int) option {
	return func(f *Foo) interface{} {
		previous := f.Verbosity
		f.Verbosity = v
		return previous
	}
}

func main() {
	foo := &Foo{1}
	v := foo.Option(Verbosity(2))
	fmt.Println("previous verbosity: ", v[0])
	fmt.Println("verbosity: ", foo.Verbosity)
}
