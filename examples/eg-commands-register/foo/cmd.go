package foo

import (
	"fmt"

	"github.com/wxio/wxcli"
	"github.com/wxio/wxcli-examples/eg-commands-register/bar"
)

func New() wxcli.SubWXCli {
	c := cmd{}
	//default name for a subcommand is its package name ("foo")
	return wxcli.NewSub(&c).SubAddCommand(bar.New())
}

type cmd struct {
	Ping string
	Pong string
}

func (f *cmd) Run() error {
	fmt.Printf("foo: %+v\n", f)
	return nil
}
