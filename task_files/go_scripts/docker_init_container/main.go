package main

import (
	"fmt"
	"main/shared"
	"os"
	"path/filepath"
	"github.com/windmillcode/go_cli_scripts/v5/utils"
)

func main() {

	shared.CDToWorkspaceRoot()
	workspaceRoot, err := os.Getwd()
	if err != nil {
		fmt.Println("there was an error while trying to receive the current dir")
	}

	initScript := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt:  []string{"docker init script to run relative to workspace root "},
			Default: utils.JoinAndConvertPathToOSFormat(".windmillcode","Local","docker_init_container.go"),
		},
	)
	initScriptArgsStruct := utils.TakeVariableArgs(
		utils.TakeVariableArgsStruct{},
	)
	initScriptArgs := fmt.Sprintf("%s %s", workspaceRoot, initScriptArgsStruct.InputString)
	initScriptLocation := filepath.Dir(initScript)
	utils.CDToLocation(initScriptLocation)
	initScript = filepath.Base(initScript)

	utils.RunCommand(shared.GetGoExecutable(), []string{"run", initScript, initScriptArgs})
}
