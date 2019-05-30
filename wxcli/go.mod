module github.com/wxio/wxcli/wxcli

go 1.12

require (
	github.com/jpillora/md-tmpl v1.2.2
	github.com/wxio/wxcli v0.0.0-00010101000000-000000000000
)

//replace github.com/wxio/wxcli => github.com/millergarym/wxcli v1.1.2
replace github.com/wxio/wxcli => ../
