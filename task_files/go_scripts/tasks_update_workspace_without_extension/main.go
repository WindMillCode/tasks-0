package main

import (
	"encoding/json"
	"fmt"
	"main/shared"
	"os"

	"github.com/windmillcode/go_cli_scripts/v5/utils"
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

	goExecutable := shared.GetGoExecutable()

	tasksJsonFilePath := utils.JoinAndConvertPathToOSFormat(workspaceRoot, ".vscode/tasks.json")

	content, err, shouldReturn := shared.CreateTasksJson(tasksJsonFilePath, false)
	if shouldReturn {
		return
	}

	var tasksJSON shared.TasksJSON
	err = json.Unmarshal(content, &tasksJSON)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	goScriptsDestDirPath := utils.JoinAndConvertPathToOSFormat(workspaceRoot, ".windmillcode/go_scripts")

	shared.RebuildExecutables("FALSE", tasksJSON, goScriptsDestDirPath, goExecutable, func() {})

}
