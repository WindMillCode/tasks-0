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
		Choices:[]string{"dev","preview","prod"},
		Other: true,
	}
	dotenvEnv:= utils.ShowMenu(cliInfo,nil)

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
			"NODE_ENV":dotenvEnv,
		},
	}
	utils.RunCommandWithOptions(opts)
}
