package main

import (
	"./config"
	"./structs"
)

var (
	Core structs.CoreCfg
)

func main() {
	config.Init(&Core)
}
