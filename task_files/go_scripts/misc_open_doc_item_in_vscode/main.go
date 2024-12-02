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
	cliInfo := utils.ShowMenuModel{
		Prompt: "Choose an option:",
		Choices: []string{
			utils.JoinAndConvertPathToOSFormat("docs", "tasks_docs"),
			utils.JoinAndConvertPathToOSFormat("docs", "app_docs"),
			"issues",
		},
	}
	docLocation := utils.ShowMenu(cliInfo, nil)
	docLocation = utils.JoinAndConvertPathToOSFormat(docLocation)
	entityNames, err := utils.GetItemsInFolder(docLocation)
	if err != nil {

		fmt.Println("Error retrieving file names please check the spelling of the provided/selected folder")
	}
	cliInfo = utils.ShowMenuModel{
		Prompt:  "Select the entity to open",
		Choices: entityNames,
		Other:   true,
	}
	targetName := utils.ShowMenu(cliInfo, nil)
	targetPath := utils.JoinAndConvertPathToOSFormat(docLocation, targetName)
	utils.RunCommand("code", []string{targetPath})
}
