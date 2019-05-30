package main

import (
	"fmt"

	"github.com/wxio/wxcli"
)

type Config struct {
	//register commands by including them in the parent struct
	Foo  `wxcli:"mode=cmd,help=This text also becomes commands summary text"`
	*Bar `wxcli:"mode=cmd,help=command two of two"`
}

func main() {
	c := Config{}
	wxcli.Parse(&c).Run()
}

type Foo struct {
	Ping string
	Pong string
}

func (f *Foo) Run() error {
	fmt.Printf("foo: %+v\n", f)
	return nil
}

type Bar struct {
	Zip string
	Zap string
}

func (b *Bar) Run() error {
	fmt.Printf("bar: %+v\n", b)
	return nil
}
