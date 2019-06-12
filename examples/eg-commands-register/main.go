package main

import (
	"github.com/wxio/wxcli"
	"github.com/wxio/wxcli-examples/eg-commands-register/foo"
)

type cmd struct{}

func main() {
	c := cmd{}
	//default name for the root command (package main) is the binary name
	wxcli.New(&c).
		AddCommand(foo.New()).
		MustParse().
		RunFatal()
}
