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
	utils.CDToTestNGApp()

	opts := utils.CommandOptions{
		Command: "",
		Args:    []string{},
	}
	utils.RunCommandWithOptions(opts)
}
