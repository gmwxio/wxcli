package main

import (
	"fmt"

	"github.com/wxio/wxcli"
)

type Config struct{}

func main() {
	wxcli.New(&Config{}).
		AddCommand(
			wxcli.NewSub(&Foo{}).
				SubAddCommand(
					wxcli.NewSub(&Bar{}),
				),
		).
		MustParse().
		Run()
}

type Foo struct {
	Ping string
	Pong string
}

func (f *Foo) Run() {
	fmt.Printf("foo: %+v\n", f)
}

type Bar struct {
	Zip string
	Zop string
}

func (b *Bar) Run() {
	fmt.Printf("bar: %+v\n", b)
}
