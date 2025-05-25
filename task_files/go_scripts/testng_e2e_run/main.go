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
			NonInteractive: settings.ExtensionPack.ProcessIfDefaultIsPresent,
		},
	)
	workSpaceFolder, err := os.Getwd()
	if err != nil {
		return
	}


	cliInfo := utils.ShowMenuModel{
		Prompt:  "choose the suite location",
		Choices: settings.ExtensionPack.TestNGE2ECreatePage.AppLocations,
		Other:   true,
	}
	if cliInfo.Choices == nil {
		cliInfo.Choices = settings.ExtensionPack.JavaSuiteLocations
	}
	suiteLocation := utils.ShowMenu(cliInfo, nil)
	suiteLocation = utils.JoinAndConvertPathToOSFormat(workspaceRoot, suiteLocation)
	utils.CDToLocation(suiteLocation)

		testArgs := utils.GetTestNGArgs(
		utils.GetTestNGArgsStruct{
			WorkspaceFolder: workSpaceFolder,
			TestNGFolder:    suiteLocation,
		},
	)

	envVarContent, err := utils.ReadFile(testArgs.EnvVarsFile)
	if err != nil {
		return
	}
	err = utils.OverwriteFile(".env", envVarContent)
	if err != nil {
		return
	}

	utils.RunCommand("mvn", []string{
		"clean",
		"test",
		fmt.Sprintf("-DsuiteFile=%s", testArgs.SuiteFile),
		fmt.Sprintf("-DparamEnv=%s", testArgs.ParamEnv),
	},
	)
}
