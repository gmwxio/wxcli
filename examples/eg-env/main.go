package main

import (
	"fmt"

	"github.com/wxio/wxcli"
)

type Config struct {
	Foo string `wxcli:"env=FOO"`
	Bar string `wxcli:"env"`
}

func main() {
	c := Config{}
	//NOTE: we could also use UseEnv(), which
	//adds 'env' to all fields.
	wxcli.New(&c).
		// UseEnv().
		MustParse()
	fmt.Printf("%+v\n", c)
}
