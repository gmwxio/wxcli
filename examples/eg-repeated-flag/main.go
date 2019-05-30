package main

import (
	"fmt"

	"github.com/wxio/wxcli"
)

func main() {
	type config struct {
		Files []string `wxcli:"help=a set of files to show"`
	}
	c := config{}
	wxcli.Parse(&c)
	fmt.Printf("%+v\n", c)
}
