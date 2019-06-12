package main

import (
	"fmt"

	"github.com/wxio/wxcli"
)

type Config struct {
	Foo string `wxcli:"mode=arg,help=<foo> is a very important argument"`
	Bar string
}

func main() {
	c := Config{}
	wxcli.New(&c).MustParse()
	fmt.Printf("%+v\n", c)
}
