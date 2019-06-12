package main

import (
	"fmt"

	"github.com/wxio/wxcli"
)

type a struct {
	X string
}
type b struct {
	Y string
}
type c struct {
	Z string
}
type d struct {
	AA string
}

func main() {
	type config struct {
		File  string `wxcli:"help=file to load"`
		Lines int    `wxcli:"help=number of lines to show"`
	}
	root := config{}
	wxcli.NewCLI("eg-helloworld").
		Configure(
			wxcli.Config{
				Path: "File",
				Tag: wxcli.Field{
					Name: "file",
					Help: "help text",
					Mode: &wxcli.Field_ModeFlag{
						ModeFlag: &wxcli.Field_Flag{
							Short: "f",
						},
					},
				},
				Default: "abc",
				Predictor: func(args string) []string {
					return []string{"a", "b", "c"}
				},
			},
			wxcli.Config{
				Path: "Lines",
				Tag: wxcli.Field{
					Name: "line",
					Help: "help text",
					Mode: &wxcli.Field_ModeFlag{
						ModeFlag: &wxcli.Field_Flag{
							Short: "l",
						},
					},
				},
				Default: 5,
			},
			wxcli.Config{
				Path: "sub.X",
				Tag: wxcli.Field{
					Name: "x",
					Help: "help text",
					Mode: &wxcli.Field_ModeFlag{
						ModeFlag: &wxcli.Field_Flag{
							Short: "x",
						},
					},
				},
				Default: "abcd",
			},
		).
		AddCommand(wxcli.NewCmd("sub").MustPrepare(&a{}).
			AddSubcommand(wxcli.NewCmd("subsub").MustPrepare(&b{}).
				AddSubcommand(wxcli.NewCmd("subsubsub").MustPrepare(&c{})),
			),
		).
		AddCommand(wxcli.NewCmd("sub2").MustPrepare(&d{})).
		MustPrepare(&root).
		Parse().
		RunFatal()
	fmt.Printf("%+v\n", root)
}
