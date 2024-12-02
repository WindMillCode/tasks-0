package main

import (
	"main/shared"
	"os"

	"github.com/windmillcode/go_cli_scripts/v5/utils"
)

func main() {

	shared.CDToWorkspaceRoot()
	workspaceRoot, err := os.Getwd()
	settings, err := utils.GetSettingsJSON(workspaceRoot)
	if err != nil {
		return
	}
	utils.SetGlobalVars(
		utils.SetGlobalVarsOptions{
			NonInteractive :settings.ExtensionPack.ProcessIfDefaultIsPresent,
		},
	)
	utils.CDToFlutterApp()

	utils.RunCommand("dart", []string{"fix", "--apply"})
	utils.RunCommand("flutter", []string{"build", "appbundle"})
}
