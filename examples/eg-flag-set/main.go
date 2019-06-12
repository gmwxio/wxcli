package main

import (
	"github.com/golang/glog"
	"github.com/wxio/wxcli"
)

func main() {
	wxcli.New(&app{}).
		EmbedGlobalFlagSet().
		Complete().
		SetLineWidth(90).
		MustParse().
		RunFatal()
}

type app struct {
}

func (a *app) Run() {
	glog.Infof("hello from app via glog\n")
}
