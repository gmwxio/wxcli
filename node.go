package wxcli

import (
	"reflect"
)

//node is the main class, it contains
//all parsing state for a single set of
//arguments
type node struct {
	err error
	//embed item since an node can also be an item
	item
	userCfgPath bool
	//help
	order                          []string
	templates                      map[string]string
	repo, author, version, summary string
	repoInfer, authorInfer         bool
	lineWidth                      int
	padAll                         bool
	padWidth                       int
	//pretend these are in the user struct :)
	internalWXCli struct {
		Help       bool
		Version    bool
		Install    bool
		Uninstall  bool
		ConfigPath string
	}
	complete bool
}

type subnode struct {
	// err error
	//embed item since an node can also be an item
	item
	parent interface{}
	//help
	// order                          []string
	// templates                      map[string]string
	// repo, author, version, summary string
	// repoInfer, authorInfer         bool
	// lineWidth                      int
	// padAll                         bool
	// padWidth                       int
	//pretend these are in the user struct :)
	internalWXCli struct {
		Help       bool
		Version    bool
		Install    bool
		Uninstall  bool
		ConfigPath string
	}
	// complete bool
}

func newNode(val reflect.Value) *node {
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
	//all new node's MUST be an addressable struct
	t := val.Type()
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		val = val.Elem()
	}
	if !val.CanAddr() || t.Kind() != reflect.Struct {
		n.errorf("must be an addressable to a struct")
		return n
	}
	n.item.val = val
	return n
}

func newSubnode(val reflect.Value) *subnode {
	n := &subnode{
		// parent: nil,
		//each cmd/cmd has its own set of names
		item: item{
			flagNames: map[string]bool{},
			cmds:      map[string]*subnode{},
			envNames:  map[string]bool{},
		},
		//these are only set at the root
	}
	//all new node's MUST be an addressable struct
	t := val.Type()
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		val = val.Elem()
	}
	if !val.CanAddr() || t.Kind() != reflect.Struct {
		panic("must be an addressable to a struct")
	}
	n.item.val = val
	return n
}
