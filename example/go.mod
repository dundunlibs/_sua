module github.com/dundunlabs/sua/example

go 1.20

replace github.com/dundunlabs/sua => ../

replace github.com/dundunlabs/sua/migr => ../migr

replace github.com/dundunlabs/sua/cli => ../cli

require (
	github.com/dundunlabs/sua v0.0.0-00010101000000-000000000000
	github.com/dundunlabs/sua/cli v0.0.0-00010101000000-000000000000
	github.com/dundunlabs/sua/migr v0.0.0-00010101000000-000000000000
	github.com/lib/pq v1.10.9
)

require (
	github.com/cpuguy83/go-md2man/v2 v2.0.2 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/urfave/cli/v3 v3.0.0-alpha3 // indirect
	github.com/xrash/smetrics v0.0.0-20201216005158-039620a65673 // indirect
)
