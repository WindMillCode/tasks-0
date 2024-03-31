package main

import (
	"github.com/windmillcode/go_cli_scripts/v4/utils"
	"main/shared"
)

func main() {

	shared.CDToWorkspaceRoot()
	utils.CDToFlutterApp()

	utils.RunCommand("dart", []string{"fix", "--apply"})
	utils.RunCommand("flutter", []string{"build", "appbundle"})
}
