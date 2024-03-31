package main

import (
	"github.com/windmillcode/go_cli_scripts/v4/utils"
	"main/shared"
)

func main() {

	shared.CDToWorkspaceRoot()
	utils.CDToAngularApp()

	utils.RunCommand("yarn", []string{"compodoc:build-and-serve"})
}
