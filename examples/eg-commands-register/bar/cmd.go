package bar

import (
	"fmt"

	"github.com/wxio/wxcli"
)

type cmd struct {
	Zip string
	Zop string
}

func New() wxcli.SubWXCli {
	c := cmd{}
	//default name for a subcommand is its package name ("bar")
	return wxcli.NewSub(&c)
}

func (b *cmd) Run() error {
	fmt.Printf("bar: %+v\n", b)
	return nil
}
