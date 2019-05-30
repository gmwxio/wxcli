package main

import (
	"fmt"

	"github.com/wxio/wxcli"
)

type Config struct {
	Foo string
	Bar string `wxcli:"default=world"` //only changes help text
}

func main() {
	c := Config{Foo: "hello"}
	wxcli.Parse(&c)
	fmt.Println(c.Foo)
	fmt.Println(c.Bar)
}
