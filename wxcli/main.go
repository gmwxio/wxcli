package main

import (
	"fmt"
	"os"

	"github.com/wxio/wxcli/wxcli/internal/genmd"

	"github.com/wxio/wxcli"
	initPkg "github.com/wxio/wxcli/wxcli/internal/init"
)

var (
	version = "dev"
	date    = "na"
	commit  = "na"
)

type root struct {
	help string
}

func main() {
	r := root{}
	o := wxcli.New(&r).
		EmbedGlobalFlagSet().
		Complete().
		Version(version).
		AddCommand(initPkg.New()).
		AddCommand(genmd.Register()).
		MustParse()
	r.help = o.Help()
	err := o.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "run error %v\n", err)
		os.Exit(2)
	}
}

func (r *root) Run() {
	fmt.Printf("version: %s\ndate: %s\ncommit: %s\n", version, date, commit)
	fmt.Println(r.help)
}
