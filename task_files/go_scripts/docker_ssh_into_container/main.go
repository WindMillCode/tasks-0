package main

import (
	"os"

	"main/shared"

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
			NonInteractive :settings.ExtensionPack.ProcessIfDefaultIsPresent,
		},
	)
	dockerContainerName := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt:  []string{"the name of the container"},
			ErrMsg:  "you must provide a container to run",
			Default: settings.ExtensionPack.SQLDockerContainerName,
		},
	)
	cliInfo := utils.ShowMenuModel{
		Prompt:  "the command line shell",
		Default: "bash",
		Choices: []string{"sh", "bash", "dash", "zsh", "cmd", "fish", "ksh", "powershell"},
	}
	shell := utils.ShowMenu(cliInfo, nil)
	shared.StartDockerDesktop()
	utils.RunCommand("docker", []string{"exec", "-it", dockerContainerName, shell})
}
