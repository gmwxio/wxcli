## args example

<!--tmpl,code=go:cat main.go -->
``` go 
package main

import (
	"fmt"

	"github.com/wxio/wxcli"
)

type Config struct {
	Shark  string   `wxcli:"mode=arg"`
	Octopi []string `wxcli:"mode=arg,min=1"`
}

func main() {
	c := Config{}
	wxcli.New(&c).MustParse()
	fmt.Printf("%+v\n", c)
}
```
<!--/tmpl-->

```
$ eg-args foo bar
```

<!--tmpl,code=plain:go run main.go foo bar -->
``` plain 
{Shark:foo Octopi:[bar]}
```
<!--/tmpl-->

```
$ eg-args --help
```

<!--tmpl,code=plain:go build -o eg-args && ./eg-args --help ; rm eg-args -->
``` plain 

  Usage: eg-args [options] <shark> <octopus> [octopus] ...

  Options:
  --help, -h  display help

```
<!--/tmpl-->
