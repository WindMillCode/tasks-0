package main

import (
	"main/shared"
	"os"

	"github.com/windmillcode/go_cli_scripts/v6/utils"
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
			NonInteractive: settings.ExtensionPack.ProcessIfDefaultIsPresent,
		},
	)
	utils.CDToSeleniumApp()

	commandOptions := utils.CommandOptions{
		Command: "",
		Args:    []string{},
		GetOutput:   false,
		PrintOutputOnly: true,
	}
	utils.RunCommandWithOptions(commandOptions)
}
