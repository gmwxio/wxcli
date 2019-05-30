module github.com/wxio/wxcli-mdtmpl

//replace github.com/wxio/wxcli => ../wxcli

go 1.12

require (
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/jpillora/md-tmpl v1.2.2
	github.com/wxio/wxcli v1.0.4
)

replace github.com/wxio/wxcli => ../
