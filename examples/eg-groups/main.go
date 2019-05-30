package main

import (
	"fmt"

	"github.com/wxio/wxcli"
)

type Config struct {
	Fizz       string
	Buzz       bool
	Foo        `wxcli:"group=Foo"`
	Ping, Pong int `wxcli:"group=More"`
}

type Foo struct {
	Bar  int
	Bazz int
}

func main() {
	c := Config{}
	wxcli.Parse(&c)
	fmt.Printf("%+v\n", c)
}
