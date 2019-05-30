package main

import (
	"fmt"

	"github.com/wxio/wxcli"
)

func main() {
	type config struct {
		File  string `wxcli:"help=file to load"`
		Lines int    `wxcli:"help=number of lines to show"`
	}
	c := config{}
	wxcli.Parse(&c)
	fmt.Printf("%+v\n", c)
}
