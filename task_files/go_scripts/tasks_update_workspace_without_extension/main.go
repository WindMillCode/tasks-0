package main

import (
	"encoding/json"
	"fmt"
	"main/shared"
	"os"

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



	tasksJsonFilePath := utils.JoinAndConvertPathToOSFormat(workspaceRoot, ".vscode/tasks.json")

	content, err, shouldReturn := utils.CreateTasksJson(tasksJsonFilePath, false)
	if shouldReturn {
		return
	}

	var tasksJSON utils.VSCodeTasksTasksJSON
	err = json.Unmarshal(content, &tasksJSON)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	goScriptsDestDirPath := utils.JoinAndConvertPathToOSFormat(workspaceRoot, ".windmillcode/go_scripts")

	shared.RebuildExecutables("FALSE", tasksJSON, goScriptsDestDirPath, func() {})

}
