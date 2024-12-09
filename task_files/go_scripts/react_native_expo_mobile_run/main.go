package main

import (
	"main/shared"
	"os"
	"github.com/windmillcode/go_cli_scripts/v6/utils"
)

func main() {

	cliInfo := utils.ShowMenuModel{
		Prompt: "use tunnel",
		Choices:[]string{"TRUE","FALSE"},
	}
	useTunnel := utils.ShowMenu(cliInfo,nil)

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
	utils.CDToReactNativeExpoApp()

	commandArgs := []string{"run", "start"}
	if useTunnel == "TRUE" {
		commandArgs = append(commandArgs, "--tunnel")
	}

	opts := utils.CommandOptions{
		Command: "npm",
		Args:    commandArgs,
		GetOutput:       false,
		PrintOutputOnly: true,
	}
	utils.RunCommandWithOptions(opts)
}
