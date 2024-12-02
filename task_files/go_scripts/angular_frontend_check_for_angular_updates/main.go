package main

import (
	"os"
	"regexp"
	"sync"

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
	projectsCLI := utils.TakeVariableArgs(
		utils.TakeVariableArgsStruct{
			Prompt:  "Provide the paths of all the projects where you want the actions to take place",
			Default: workspaceRoot,
		},
	)
	angularAppLocation := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt:  []string{"provide the relative path to the angular app (note : for every project  the relative path should be the same other wi)"},
			Default: "apps/frontend/AngularApp",
		},
	)

	var wg sync.WaitGroup
	regex0 := regexp.MustCompile(" ")
	projectsList := regex0.Split(projectsCLI.InputString, -1)
	for _, project := range projectsList {
		app := utils.JoinAndConvertPathToOSFormat(project, angularAppLocation)
		wg.Add(1)
		go func() {
			defer wg.Done()
			utils.RunCommandInSpecificDirectory("npx", []string{"ng", "update"}, app)
		}()
	}
	wg.Wait()
}
