package main

import (
	_ "github.com/acoderup/goserver/mmo"

	"github.com/acoderup/goserver/core"
	"github.com/acoderup/goserver/core/module"
)

func main() {
	defer core.ClosePackages()
	core.LoadPackages("config.json")

	waiter := module.Start()
	waiter.Wait("main")
}
