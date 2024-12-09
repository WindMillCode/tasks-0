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
	utils.CDToReactNativeExpoApp()

	cliInfo := utils.ShowMenuModel{
		Prompt: "profile",
		Choices:[]string{"development","staging","production"},
	}
	myProfile := utils.ShowMenu(cliInfo,nil)

	cliInfo = utils.ShowMenuModel{
		Prompt: "platform",
		Choices:[]string{"android","ios"},
	}
	myPlatform := utils.ShowMenu(cliInfo,nil)

	cliInfo = utils.ShowMenuModel{
		Prompt: "build locally",
		Choices:[]string{"TRUE","FALSE"},
	}
	localBuild := utils.ShowMenu(cliInfo,nil)
	commandArgs:= []string{"build", "--profile", myProfile, "--platform", myPlatform}

	if localBuild == "TRUE" {
		commandArgs = append(commandArgs, "--local")
	}

	opts := utils.CommandOptions{
		Command: "eas",
		Args:    commandArgs,
		GetOutput:       false,
		PrintOutputOnly: true,
	}
	utils.RunCommandWithOptions(opts)
}
