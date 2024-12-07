package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/windmillcode/go_cli_scripts/v6/utils"
)

func main() {

	destDir := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt: []string{"name of the command"},
			ErrMsg: "you must provide a name for the command",
		},
	)

	utils.CDToLocation(
		utils.JoinAndConvertPathToOSFormat("..","..", "task_files", "go_scripts"),
	)

	err := utils.CopyDir("template", destDir)
	if err != nil {
		log.Fatalf("Failed to copy directory: %v", err)
	}

	utils.CDToLocation(
		utils.JoinAndConvertPathToOSFormat("..",".."),
	)

	tasksJsonFilePath := utils.JoinAndConvertPathToOSFormat(".", ".vscode/tasks.json")

	content, err, _ := utils.CreateTasksJson(tasksJsonFilePath, false)
	var tasksJSON utils.VSCodeTasksTasksJSON
	err = json.Unmarshal(content, &tasksJSON)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

}
