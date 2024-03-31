package main

import (
	"github.com/windmillcode/go_cli_scripts/v4/utils"
	"main/shared"
)

func main() {

	shared.CDToWorkspaceRoot()
	utils.CDToTestNGApp()

	opts := utils.CommandOptions{
		Command: "",
		Args:    []string{},
	}
	utils.RunCommandWithOptions(opts)
}
