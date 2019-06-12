package wxcli

import (
	"flag"
	"reflect"
	"strings"
)

//WXCli is a single configuration command instance. It represents a node
//in a tree of commands. Use the AddCommand method to add subcommands (child nodes)
//to this command instance.
type WXCli interface {
	// //Name of the command. For the root command, Name defaults to the executable's
	// //base name. For subcommands, Name defaults to the package name, unless its the
	// //main package, then it defaults to the struct name.
	// Name(name string) WXCli
	//Version of the command. Commonly set using a package main variable at compile
	//time using ldflags (for example, go build -ldflags -X main.version=42).
	Version(version string) WXCli
	//ConfigPath is a path to a JSON file to use as defaults. This is useful in
	//global paths like /etc/my-prog.json. For a user-specified path. Use the
	//UserConfigPath method.
	ConfigPath(path string) WXCli
	//UserConfigPath is the same as ConfigPath however an extra flag (--config-path)
	//is added to this WXCli instance to give the user control of the filepath.
	//Configuration unmarshalling occurs after flag parsing.
	UserConfigPath() WXCli
	//UseEnv enables the default environment variables on all fields. This is
	//equivalent to adding the wxcli tag "env" on all flag fields.
	UseEnv() WXCli
	//Complete enables auto-completion for this command. When enabled, two extra
	//flags are added (--install and --uninstall) which can be used to install
	//a dynamic shell (bash, zsh, fish) completion for this command. Internally,
	//this adds a stub file which runs the Go binary to auto-complete its own
	//command-line interface. Note, the absolute path returned from os.Executable()
	//is used to reference to the Go binary.
	Complete() WXCli
	//EmbedFlagSet embeds the given pkg/flag.FlagSet into
	//this WXCli instance. Placing the flags defined in the FlagSet
	//along side the configuration struct flags.
	EmbedFlagSet(*flag.FlagSet) WXCli
	//EmbedGlobalFlagSet embeds the global pkg/flag.CommandLine
	//FlagSet variable into this WXCli instance.
	EmbedGlobalFlagSet() WXCli

	//Summary adds a short sentence below the usage text
	Summary(summary string) WXCli
	//Repo sets the source repository of the program and is displayed
	//at the bottom of the help text.
	Repo(repo string) WXCli
	//Author sets the author of the program and is displayed
	//at the bottom of the help text.
	Author(author string) WXCli
	//PkgRepo automatically sets Repo using the struct's package path.
	//This does not work for types defined in the main package.
	PkgRepo() WXCli
	//PkgAuthor automatically sets Author using the struct's package path.
	//This does not work for types defined in the main package.
	PkgAuthor() WXCli
	//DocSet replaces an existing template.
	DocSet(id, template string) WXCli
	//DocBefore inserts a new template before an existing template.
	DocBefore(existingID, newID, template string) WXCli
	//DocAfter inserts a new template after an existing template.
	DocAfter(existingID, newID, template string) WXCli
	//DisablePadAll removes the padding from the help text.
	DisablePadAll() WXCli
	//SetPadWidth alters the padding to specific number of spaces.
	//By default, pad width is 2.
	SetPadWidth(padding int) WXCli
	//SetLineWidth alters the maximum number of characters in a
	//line (excluding padding). By default, line width is 96.
	SetLineWidth(width int) WXCli

	//Parse uses os.Args to parse the internal FlagSet and
	//returns a ParsedWXCli instance.
	Parse() (ParsedWXCli, error)
	MustParse() ParsedWXCli
	//ParseArgs uses a given set of args to to parse the
	//current flags and args. Assumes the executed program is
	//the first arg.
	ParseArgs(args []string) (ParsedWXCli, error)
}

type WXCommand interface {
	// //Name of the command. For the root command, Name defaults to the executable's
	// //base name. For subcommands, Name defaults to the package name, unless its the
	// //main package, then it defaults to the struct name.
	// SubName(name string) WXCommand
	//Summary adds an arbitrarily long string to below the usage text
	// SubSummary(summary string) WXCommand
	//AddCommand adds another WXCli instance as a subcommand.
	// AddSubcommand(WXCommand) WXCommand
}

type ParsedWXCli interface {
	//Help returns the final help text
	Help() string
	//IsRunnable returns whether the matched command has a Run method
	IsRunnable() bool
	//Run assumes the matched command is runnable and executes its Run method.
	//The target Run method must be 'Run() error' or 'Run()'
	Run() error
	//RunFatal assumes the matched command is runnable and executes its Run method.
	//However, any error will be printed, followed by an exit(1).
	RunFatal()
}

type Config struct {
	Path      string
	Tag       Field
	Default   interface{}
	Predictor Complete
}

func NewCLI(name string) Configurer {
	n := &node{
		// parent: nil,
		//each cmd/cmd has its own set of names
		item: item{
			flagNames: map[string]bool{},
			cmds:      map[string]*subnode{},
			envNames:  map[string]bool{},
		},
		//these are only set at the root
		order:     defaultOrder(),
		templates: map[string]string{},
		//public defaults
		lineWidth: 96,
		padAll:    true,
		padWidth:  2,
	}
	n.name = name
	return n
}

func NewCmd(name string) SubConfigurer {
	n := &subnode{
		parent: nil,
		item: item{
			//each cmd/cmd has its own set of names
			flagNames: map[string]bool{},
			// envNames:  map[string]bool{},
			cmds: map[string]*subnode{},
		},
	}
	n.name = name
	return n
}

type Configurer interface {
	Commander
	Configure(cfgs ...Config) Commander
}
type SubConfigurer interface {
	SubCommander
	Configure(cfgs ...Config) SubCommander
}

type Commander interface {
	Preparitory
	AddCommand(WXCommand) Commander
}
type SubCommander interface {
	SubPreparitory
	AddSubcommand(WXCommand) SubCommander
}

type Preparitory interface {
	Prepare(config interface{}) (WXCli, error)
	MustPrepare(config interface{}) WXCli
}
type SubPreparitory interface {
	Prepare(config interface{}) (SubCommander, error)
	MustPrepare(config interface{}) SubCommander
}

//New creates a new WXCli instance using the given configuration
//struct pointer.
func New(config interface{}) WXCli {
	return newNode(reflect.ValueOf(config))
}

//New creates a new WXCommand instance using the given configuration
//struct pointer.
func NewSub(config interface{}) WXCommand {
	sub := newNode(reflect.ValueOf(config))
	//default name should be package name,
	//unless its in the main packagWXCommand
	//the default becomes the struct name
	structType := sub.item.val.Type()
	pkgPath := structType.PkgPath()
	if sub.name == "" && pkgPath != "main" && pkgPath != "" {
		parts := strings.Split(pkgPath, "/")
		sub.name = parts[len(parts)-1]
	}
	structName := structType.Name()
	if sub.name == "" && structName != "" {
		sub.name = camel2dash(structName)
	}
	return sub
}

//Parse is shorthand for
//  wxcli.New(config).MustParse()
func Parse(config interface{}) ParsedWXCli {
	return New(config).MustParse()
}

//Setter is any type which can be set from a string.
//This includes flag.Value.
type Setter interface {
	Set(string) error
}
