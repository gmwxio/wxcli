package main

import (
	"fmt"

	"github.com/wxio/wxcli"
)

type a struct{}
type b struct{}
type c struct{}
type d struct{}

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
			}).
		// Tags(
		// 	wxcli.Tag{Path: "File", Tag: wxcli.Field{
		// 		Name: "file",
		// 		Help: "help text",
		// 		Mode: &wxcli.Field_ModeFlag{
		// 			ModeFlag: &wxcli.Field_Flag{
		// 				Short: "f",
		// 			},
		// 		},
		// 	}},
		// 	wxcli.Tag{Path: "Lines", Tag: wxcli.Field{
		// 		Name: "line",
		// 		Help: "help text",
		// 		Mode: &wxcli.Field_ModeFlag{
		// 			ModeFlag: &wxcli.Field_Flag{
		// 				Short: "l",
		// 			},
		// 		},
		// 	}},
		// ).
		// Defaults(
		// 	wxcli.Default{Path: "File", Value: "abc"},
		// 	wxcli.Default{Path: "Lines", Value: 5},
		// ).
		// Completions(
		// 	wxcli.Completion{Path: "File", Predictor: func(args string) []string {
		// 		return []string{"a", "b", "c"}
		// 	}},
		// ).
		AddCommand(wxcli.NewCmd("sub").MustStuff(&a{}).
			AddSubcommand(wxcli.NewCmd("subsub").MustStuff(&b{}).
				AddSubcommand(wxcli.NewCmd("subsubsub").MustStuff(&c{})),
			),
		).
		AddCommand(wxcli.NewCmd("sub2").MustStuff(&d{})).
		MustStuff(&root).
		Parse().
		RunFatal()
	fmt.Printf("%+v\n", root)
}
