package main

import (
	"fmt"
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
	commandOptions := utils.CommandOptions{
		Command:     "docker",
		Args:        []string{"start", dockerContainerName},
		GetOutput:   true,
		TargetDir:   "",
		PrintOutput: false,
	}
	_, err = utils.RunCommandWithOptions(commandOptions)
	if err != nil {
		shared.StartDockerDesktop()
		_, err = utils.RunCommandWithOptions(commandOptions)
		if err != nil {
			fmt.Println(`
			Please find the root folder of your docker installation and ENSURE you add in the following order or else there may be issue in getting the docker engine to start at all
			[ROOT_PATH]/
			[ROOT_PATH]/resources
			[ROOT_PATH]/resources/bin
			`)
		}
	}
}
