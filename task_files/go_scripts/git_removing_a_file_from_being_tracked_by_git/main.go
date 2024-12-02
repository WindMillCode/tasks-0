package main

import (
	"fmt"
	"os"
	"strings"

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
	cliInfo := utils.ShowMenuModel{
		Prompt:  "Is it a file or directory:",
		Choices: []string{"file", "directory"},
	}
	targetType := utils.ShowMenu(cliInfo, nil)

	targetName := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt: []string{"name of the item"},
		},
	)

	var args0 []string
	if targetType == "file" {

		argsString := fmt.Sprintf("rm --cached --sparse %s", targetName)
		args0 = strings.Split(argsString, " ")
	} else {

		argsString := fmt.Sprintf("rm -r --cached --sparse %s", targetName)
		args0 = strings.Split(argsString, " ")
	}

	utils.RunCommand("git", args0)
}
