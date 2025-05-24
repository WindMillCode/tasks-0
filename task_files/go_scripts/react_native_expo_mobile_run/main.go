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

	cliInfo = utils.ShowMenuModel{
		Prompt: "clear cache",
		Choices:[]string{"TRUE","FALSE"},
	}
	clearCache := utils.ShowMenu(cliInfo,nil)

	cliInfo = utils.ShowMenuModel{
		Prompt: "select the ennvironment",
		Choices:[]string{"DEV","PREVIEW","PROD"},
		Other: true,
	}
	EXPO_MOBILE_ENV:= utils.ShowMenu(cliInfo,nil)

	shared.CDToWorkspaceRoot()
	workspaceRoot, err := os.Getwd()
	if err != nil {
		return
	}
	settings, err := utils.GetSettingsJSON(workspaceRoot)
	if err != nil {
		return
	}
	utils.SetGlobalVars(
		utils.SetGlobalVarsOptions{
			NonInteractive: settings.ExtensionPack.ProcessIfDefaultIsPresent,
		},
	)
	utils.CDToExpoApp()

	commandArgs := []string{"expo", "start"}
	if useTunnel == "TRUE" {
		commandArgs = append(commandArgs, "--tunnel")
	}

	if clearCache == "TRUE" {
		commandArgs = append(commandArgs, "--clear")
	}

	opts := utils.CommandOptions{
		Command:             "npx",
		Args:                commandArgs,
		GetOutput:           false,
		PrintOutputOnly:     true,
		EnvVars:             map[string]string{
			"EXPO_MOBILE_ENV":EXPO_MOBILE_ENV,
		},
	}
	utils.RunCommandWithOptions(opts)
}
