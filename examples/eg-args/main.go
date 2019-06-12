package main

import (
	"fmt"

	"github.com/wxio/wxcli"
)

type Config struct {
	Shark  string   `wxcli:"mode=arg"`
	Octopi []string `wxcli:"mode=arg,min=1"`
}

func main() {
	c := Config{}
	wxcli.New(&c).MustParse()
	fmt.Printf("%+v\n", c)
}
