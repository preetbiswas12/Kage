package main

import (
	"github.com/preetbiswas12/Kage/cmd"
	"github.com/preetbiswas12/Kage/config"
	"github.com/preetbiswas12/Kage/log"
	"github.com/samber/lo"
)

func main() {
	lo.Must0(config.Setup())
	lo.Must0(log.Setup())
	cmd.Execute()
}
