//Package wxcli defines a struct-tag based API for
//rapidly building command-line interfaces. For example:
//
//  package main
//
//  import (
//  	"log"
//  	"github.com/wxio/wxcli"
//  )
//
//  func main() {
//  	type config struct {
//  		File  string `wxcli:"help=file to load"`
//  		Lines int    `wxcli:"help=number of lines to show"`
//  	}
//  	c := config{}
//  	wxcli.Parse(&c)
//  	log.Printf("%+v", c)
//  }
//
//Build and run:
//
//  $ go build -o my-prog
//  $ ./my-prog --help
//
//    Usage: my-prog [options]
//
//    Options:
//    --file, -f   file to load
//    --lines, -l  number of lines to show
//    --help, -h   display help
//
//  $ ./my-prog -f foo.txt -l 42
//  {File:foo.txt Lines:42}
//
//See https://github.com/wxio/wxcli for more information and more examples.
package wxcli
