<p align="center">
<img width="443" alt="logo" src="https://user-images.githubusercontent.com/633843/57529538-84a22780-7378-11e9-9235-312633dc125e.png"><br>
<b>A Go (golang) library for build command-line interfaces that just works</b><br><br>
<a href="https://godoc.org/github.com/wxio/wxcli#WXCli" rel="nofollow">
	<img src="https://camo.githubusercontent.com/42566bdba17f1a0c86c1a1de859d6ab70bde1457/68747470733a2f2f676f646f632e6f72672f6769746875622e636f6d2f6a70696c6c6f72612f6f7074733f7374617475732e737667" alt="GoDoc" data-canonical-src="https://godoc.org/github.com/wxio/wxcli?status.svg" style="max-width:100%;">
</a>
<a href="https://circleci.com/gh/wxio/wxcli" rel="nofollow">
	<img src="https://camo.githubusercontent.com/34202387888c6b05f640653a29bb1e204f5a9e19/68747470733a2f2f636972636c6563692e636f6d2f67682f6a70696c6c6f72612f6f7074732e7376673f7374796c653d736869656c6426636972636c652d746f6b656e3d36396566396336616330643863656263623335346262383563333737656365666637376266623162" alt="CircleCI" data-canonical-src="https://circleci.com/gh/wxio/wxcli.svg?style=shield&amp;circle-token=69ef9c6ac0d8cebcb354bb85c377eceff77bfb1b" style="max-width:100%;">
</a>
</p>

---

Creating command-line interfaces should be simple:

```go
package main

import (
	"log"

	"github.com/wxio/wxcli"
)

func main() {
	type config struct {
		File  string `wxcli:"help=file to load"`
		Lines int    `wxcli:"help=number of lines to show"`
	}
	c := config{}
	wxcli.Parse(&c)
	log.Printf("%+v", c)
}
```

```sh
$ go build -o my-prog
$ ./my-prog --help

  Usage: my-prog [options]

  Options:
  --file, -f   file to load
  --lines, -l  number of lines to show
  --help, -h   display help

```

```sh
$ ./my-prog -f foo.txt -l 42
{File:foo.txt Lines:42}
```

*Try it out https://play.golang.org/p/D0jWFwmxRgt*

### Features (with examples)

- Easy to use ([eg-helloworld](https://github.com/wxio/wxcli/tree/master/examples/eg-helloworld/))
- Promotes separation of CLI code and library code ([eg-app](https://github.com/wxio/wxcli/tree/master/examples/eg-app/))
- Automatically generated `--help` text via struct tags ([eg-help](https://github.com/wxio/wxcli/tree/master/examples/eg-help/))
- Default values by modifying the struct prior to `MustParse()` ([eg-defaults](https://github.com/wxio/wxcli/tree/master/examples/eg-defaults/))
- Default values from a JSON config file, unmarshalled via your config struct ([eg-config](https://github.com/wxio/wxcli/tree/master/examples/eg-config/))
- Default values from environment, defined by your field names ([eg-env](https://github.com/wxio/wxcli/tree/master/examples/eg-env/))
- Repeated flags using slices ([eg-repeated-flag](https://github.com/wxio/wxcli/tree/master/examples/eg-repeated-flag/))
- Group your flags in the help output ([eg-groups](https://github.com/wxio/wxcli/tree/master/examples/eg-groups/))
- Sub-commands by nesting structs ([eg-commands-inline](https://github.com/wxio/wxcli/tree/master/examples/eg-commands-inline/))
- Sub-commands by providing child `WXCli` ([eg-commands-main](https://github.com/wxio/wxcli/tree/master/examples/eg-commands-main/))
- Infers program name from executable name
- Infers command names from struct or package name
- Define custom flags types via `wxcli.Setter` or `flag.Value` ([eg-custom-flag](https://github.com/wxio/wxcli/tree/master/examples/eg-custom-flag/))
- Customizable help text by modifying the default templates ([eg-help](https://github.com/wxio/wxcli/tree/master/examples/eg-help/))
- Built-in shell auto-completion ([eg-complete](https://github.com/wxio/wxcli/tree/master/examples/eg-complete))

Find these examples and more in the [`examples`](https://github.com/wxio/wxcli/tree/master/examples) folder.

### Package API

See https://godoc.org/github.com/wxio/wxcli#WXCli

[![GoDoc](https://godoc.org/github.com/wxio/wxcli?status.svg)](https://godoc.org/github.com/wxio/wxcli)

### Struct Tag API

**wxcli** tries to set sane defaults so, for the most part, you'll get the desired behaviour by simply providing a configuration struct.

However, you can customise this behaviour by providing the `wxcli` struct
tag with a series of settings in the form of **`key=value`**:

```
`wxcli:"key=value,key=value,..."`
```

Where **`key`** must be one of:

- `-` (dash) - Like `json:"-"`, the dash character will cause wxcli to ignore the struct field. Unexported fields are always ignored.

- `name` - Name is used to display the field in the help text. By default, the flag name is infered by converting the struct field name to lowercase and adding dashes between words.

- `help` - The help text used to summaryribe the field. It will be displayed next to the flag name in the help output.

	*Note:* `help` can also be set as a stand-alone struct tag (i.e. `help:"my text goes here"`). You must use the stand-alone tag if you wish to use a comma `,` in your help string.

- `mode` - The **wxcli** mode assigned to the field. All fields will be given a `mode`. Where the `mode` **`value`** must be one of:

	* `flag` - The field will be treated as a flag: an optional, named, configurable field. Set using `./program --<flag-name> <flag-value>`. The struct field must be a [*flag-value*](#flag-values) type. `flag` is the default `mode` for any [*flag-value*](#flag-values).

	* `arg` - The field will be treated as an argument: a required, positional, unamed, configurable field. Set using `./program <argument-value>`. The struct field must be a [*flag-value*](#flag-values) type.

	* `embedded` - A special mode which causes the fields of struct to be used in the current struct. Useful if you want to split your command-line options across multiple files (default for `struct` fields). The struct field must be a `struct`. `embedded` is the default `mode` for a `struct`. *Tip* You can play group all fields together placing an `group` tag on the struct field.

	* `cmd` - A inline command, shorthand for `.AddCommmand(wxcli.New(&field))`, which also implies the struct field must be a `struct`.

	* `cmdname` - A special mode which will assume the name of the selected command. The struct field must be a `string`.

- `short` - One letter to be used a flag's "short" name. By default, the first letter of `name` will be used. It will remain unset if there is a duplicate short name. Only valid when `mode` is `flag`.

- `group` - The name of the flag group to store the field. Defining this field will create a new group of flags in the help text (will appear as "`<group>` options"). The default flag group is the empty string (which will appear as "Options"). Only valid when `mode` is `flag` or `embedded`.

- `env` - An environent variable to use as the field's **default** value. It can always be overridden by providing the appropriate flag. Only valid when `mode` is `flag`.

	For example, `wxcli:"env=FOO"`. It can also be infered using the field name with simply `wxcli:"env"`. You can enable inference on all flags with the `wxcli.WXCli` method `UseEnv()`.

- `min` `max` - A minimum or maximum length of a slice. Only valid when `mode` is `arg`, *and* the struct field is a slice.

#### flag-values:

In general an wxcli _flag-value_ type aims to be any type that can be get and set using a `string`. Currently, **wxcli** supports the following types:

- `string`
- `bool`
- `int`, `int8`, `int16`, `int32`, `int64`
- `uint`, `uint8`, `uint16`, `uint32`, `uint64`
- `float32`, `float64`
- [`wxcli.Setter`](https://godoc.org/github.com/wxio/wxcli#Setter)
	- *The interface `func Set(string) error`*
- [`flag.Value`](https://golang.org/pkg/flag/#Value)
	- *Is an `wxcli.Setter`*
- `time.Duration`
- `encoding.TextMarshaler`
	- *Includes `time.Time` and `net.IP`*
- `encoding.BinaryMarshaler`
	- *Includes `url.URL`*

In addition, `flag`s and `arg`s can also be a slice of any _flag-value_ type. Slices allow multiple flags/args. For example, a struct field flag `Foo []int` could be set with `--foo 1 --foo 2`, and would result in `[]int{1,2}`.

### Help text

By default, **wxcli** attempts to output well-formatted help text when the user provides the `--help` (`-h`) flag. The [examples](https://github.com/wxio/wxcli/tree/master/examples) repositories shows various combinations of this default help text, resulting from using various features above.

Modifications be made by customising the underlying [Go templates](https://golang.org/pkg/text/template/) found here [DefaultTemplates](https://godoc.org/github.com/wxio/wxcli#pkg-variables).

### Talk

I gave a talk on **wxcli** at the Go Meetup Sydney (golang-syd) on the 23rd of May, 2019. You can find the slides here https://github.com/wxio/wxcli-talk.

### Other projects

Other related projects which infer flags from struct tags but aren't as feature-complete:

- https://github.com/alexflint/go-arg
- https://github.com/jessevdk/go-flags

#### MIT License

Copyright © 2019 &lt;dev@wxio.com&gt;

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
'Software'), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED 'AS IS', WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
