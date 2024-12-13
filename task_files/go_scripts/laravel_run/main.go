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
	utils.CDToLaravelApp()

	commandOptions := utils.CommandOptions{
		Command: "php",
		Args:    []string{"artisan","serve"},
		GetOutput:   false,
		PrintOutputOnly: true,
	}
	utils.RunCommandWithOptions(commandOptions)
}
