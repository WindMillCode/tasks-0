package main

import (
	"os"
	"time"

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
	utils.CDToFlutterApp()

	targetPath := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt:  []string{"provide the path to a file or directory"},
			Default: "./test",
		},
	)
	targetPath = utils.JoinAndConvertPathToOSFormat(targetPath)
	commandOptions := utils.CommandOptions{
		Command:   "flutter",
		Args:      []string{"test", targetPath},
		GetOutput: true,
	}
	for {
		utils.RunCommandWithOptions(commandOptions)
		time.Sleep(60 * time.Second)
	}

}
