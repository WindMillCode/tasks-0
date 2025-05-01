package main

import (
	"fmt"
	"main/shared"
	"os"
	"path/filepath"
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

	initScript := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt:  []string{"docker init script to run relative to workspace root "},
			Default: utils.JoinAndConvertPathToOSFormat(".windmillcode","Local","docker_init_sql_container.go"),
		},
	)
	initScriptArgsStruct := utils.TakeVariableArgs(
		utils.TakeVariableArgsStruct{},
	)
	initScriptArgs := fmt.Sprintf("%s %s", workspaceRoot, initScriptArgsStruct.InputString)
	initScriptLocation := filepath.Dir(initScript)
	utils.CDToLocation(initScriptLocation)
	initScript = filepath.Base(initScript)
	shared.StartDockerDesktop()
	utils.RunCommand("go", []string{"run", initScript, initScriptArgs})
}
