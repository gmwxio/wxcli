package main

import (
	"fmt"

	"github.com/wxio/wxcli"
)

type Config struct {
	Foo string
	Bar string
}

func main() {
	c := Config{}
	wxcli.New(&c).
		ConfigPath("config.json").
		MustParse()
	fmt.Println(c.Foo)
	fmt.Println(c.Bar)
}
