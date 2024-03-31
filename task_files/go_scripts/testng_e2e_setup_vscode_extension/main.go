package main

import (
	"os"

	"github.com/windmillcode/go_cli_scripts/v4/utils"
	"main/shared"
)

func main() {

	shared.CDToWorkspaceRoot()
	workSpaceFolder, err := os.Getwd()
	if err != nil {
		return
	}
	testArgs := utils.GetTestNGArgs(
		utils.GetTestNGArgsStruct{
			WorkspaceFolder: workSpaceFolder,
		},
	)

	utils.CDToTestNGApp()
	envVarContent, err := utils.ReadFile(testArgs.EnvVarsFile)
	if err != nil {
		return
	}
	err = utils.OverwriteFile(".env", envVarContent)
	if err != nil {
		return
	}
}
