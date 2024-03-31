package main

import (
	"github.com/windmillcode/go_cli_scripts/v4/utils"
	"main/shared"
)

func main() {

	shared.CDToWorkspaceRoot()
	utils.CDToFlutterApp()

	utils.RunCommand("flutter", []string{"channel", "stable"})
	utils.RunCommandAndGetOutput("flutter", []string{"upgrade"})
	// utils.RunCommandAndGetOutput("flutter", []string{"upgrade","--force"})
}
