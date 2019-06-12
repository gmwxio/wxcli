package wxcli

import (
	"flag"
	"fmt"
	"reflect"
)

//errorf to be stored until parse-time
func (n *node) errorf(format string, args ...interface{}) error {
	err := authorError(fmt.Sprintf(format, args...))
	//only store the first error
	if n.err == nil {
		n.err = err
	}
	return err
}

func (n *node) Name(name string) WXCli {
	n.name = name
	return n
}

func (n *node) Configure(cfgs ...Config) Commander {
	n.cfgs = cfgs
	return n
}

func (n *node) Tags(tags ...Tag) Defaulter {
	n.tags = tags
	return n
}
func (n *node) Completions(completes ...Completion) Commander {
	n.completions = completes
	return n
}
func (n *subnode) Completions(completes ...Completion) SubCommander {
	n.completions = completes
	return n
}
func (n *subnode) Tags(tags ...Tag) SubDefaulter {
	n.tags = tags
	return n
}

func (n *node) Defaults(defaults ...Default) Completers {
	n.defaults = defaults
	return n
}
func (n *subnode) Defaults(defaults ...Default) SubCompleters {
	n.defaults = defaults
	return n
}

// func (n *node) AddCommand(cmd WXCommand) Commander {
// 	sub, ok := cmd.(*node)
// 	if !ok {
// 		panic("another implementation of wxcli???")
// 	}
// 	if _, exists := n.cmds[sub.name]; exists {
// 		panic(n.errorf("cannot add command, '%s' already exists", sub.name))
// 	}
// 	sub.parent = n
// 	n.cmds[sub.name] = sub
// 	return n
// }

func (n *node) Stuff(config interface{}) (WXCli, error) {
	val := reflect.ValueOf(config)
	//all new node's MUST be an addressable struct
	t := val.Type()
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		val = val.Elem()
	}
	if !val.CanAddr() || t.Kind() != reflect.Struct {
		return nil, n.errorf("must be an addressable to a struct")
		// return n
	}
	n.item.val = val
	return n, nil
}
func (n *subnode) Stuff(config interface{}) (SubCommander, error) {
	val := reflect.ValueOf(config)
	//all new node's MUST be an addressable struct
	t := val.Type()
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		val = val.Elem()
	}
	if !val.CanAddr() || t.Kind() != reflect.Struct {
		return nil, fmt.Errorf("must be an addressable to a struct")
		// return n
	}
	n.item.val = val
	return n, nil
}

func (n *node) MustStuff(config interface{}) WXCli {
	if _, err := n.Stuff(config); err != nil {
		panic(err)
	}
	return n
}
func (n *subnode) MustStuff(config interface{}) SubCommander {
	if _, err := n.Stuff(config); err != nil {
		panic(err)
	}
	return n
}

//Name sets the name of the program
func (n *node) SubName(name string) WXCommand {
	n.name = name
	return n
}

//Version sets the version of the program
//and renders the 'version' template in the help text
func (n *node) Version(version string) WXCli {
	n.version = version
	return n
}

func (n *node) Summary(summary string) WXCli {
	n.summary = summary
	return n
}

//Summary sets the text summary of the program,
//which, by default, is inserted below the usage text
func (n *node) SubSummary(summary string) WXCommand {
	n.summary = summary
	return n
}

//Repo sets the repository link of the program
//and renders the 'repo' template in the help text
func (n *node) Repo(repo string) WXCli {
	n.repo = repo
	return n
}

func (n *node) PkgRepo() WXCli {
	n.repoInfer = true
	return n
}

func (n *node) Author(author string) WXCli {
	n.author = author
	return n
}

//PkgRepo infers the repository link of the program
//from the package import path of the struct (So note,
//this will not work for 'main' packages)
func (n *node) PkgAuthor() WXCli {
	n.authorInfer = true
	return n
}

//Set the padding width, which defines the amount padding
//when rendering help text (defaults to 72)
func (n *node) SetPadWidth(p int) WXCli {
	n.padWidth = p
	return n
}

func (n *node) DisablePadAll() WXCli {
	n.padAll = false
	return n
}

func (n *node) SetLineWidth(l int) WXCli {
	n.lineWidth = l
	return n
}

func (n *node) ConfigPath(path string) WXCli {
	n.internalWXCli.ConfigPath = path
	return n
}

func (n *node) UserConfigPath() WXCli {
	n.userCfgPath = true
	return n
}

func (n *node) UseEnv() WXCli {
	n.useEnv = true
	return n
}

//DocBefore inserts a text block before the specified template
func (n *node) DocBefore(target, newID, template string) WXCli {
	return n.docOffset(0, target, newID, template)
}

//DocAfter inserts a text block after the specified template
func (n *node) DocAfter(target, newID, template string) WXCli {
	return n.docOffset(1, target, newID, template)
}

//DecSet replaces the specified template
func (n *node) DocSet(id, template string) WXCli {
	if _, ok := DefaultTemplates[id]; !ok {
		if _, ok := n.templates[id]; !ok {
			n.errorf("template does not exist: %s", id)
			return n
		}
	}
	n.templates[id] = template
	return n
}

func (n *node) docOffset(offset int, target, newID, template string) *node {
	if _, ok := n.templates[newID]; ok {
		n.errorf("new template already exists: %s", newID)
		return n
	}
	for i, id := range n.order {
		if id == target {
			n.templates[newID] = template
			index := i + offset
			rest := []string{newID}
			if index < len(n.order) {
				rest = append(rest, n.order[index:]...)
			}
			n.order = append(n.order[:index], rest...)
			return n
		}
	}
	n.errorf("target template not found: %s", target)
	return n
}

func (n *node) EmbedFlagSet(fs *flag.FlagSet) WXCli {
	n.flagsets = append(n.flagsets, fs)
	return n
}

func (n *node) EmbedGlobalFlagSet() WXCli {
	return n.EmbedFlagSet(flag.CommandLine)
}

func (n *node) Call(fn func(o WXCli)) WXCli {
	fn(n)
	return n
}

func (n *item) flagGroup(name string) *itemGroup {
	//NOTE: the default group is the empty string
	//get existing group
	for _, g := range n.flagGroups {
		if g.name == name {
			return g
		}
	}
	//otherwise, create and append
	g := &itemGroup{name: name}
	n.flagGroups = append(n.flagGroups, g)
	return g
}

func (n *item) flags() []*item {
	flags := []*item{}
	for _, g := range n.flagGroups {
		flags = append(flags, g.flags...)
	}
	return flags
}

type authorError string

func (e authorError) Error() string {
	return string(e)
}

type exitError string

func (e exitError) Error() string {
	return string(e)
}
