package main

import (
	"fmt"
	"main/shared"
	"os"
	"strconv"
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

	portNumber := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt: []string{"The port for the app"},
			Default: func() string {
				if settings.ExtensionPack.Ports.LaravelRun0 == 0 {
					return "8000"
				}
				return strconv.Itoa(settings.ExtensionPack.Ports.LaravelRun0)
			}(),
		},
	)


	utils.CDToLaravelApp()




	commandOptions := utils.CommandOptions{
		Command: "php",
		Args:    []string{"artisan","serve",fmt.Sprintf("--port=%s",portNumber)},
		GetOutput:   false,
		PrintOutputOnly: true,
	}
	utils.RunCommandWithOptions(commandOptions)
}
