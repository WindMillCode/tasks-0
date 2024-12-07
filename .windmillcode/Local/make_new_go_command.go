package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sort"
	"github.com/windmillcode/go_cli_scripts/v6/utils"
)

func main() {

	destDir := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt: []string{"name of the command"},
			ErrMsg: "you must provide a name for the command",
		},
	)
	destDir = strings.TrimSpace(destDir)

	instanceLimit := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt: []string{"How many instances of this command can be run at once?"},
			Default: "1",
		},
	)
	instanceLimitNumber,err := strconv.Atoi(instanceLimit)
	if err != nil {
		fmt.Println("Error converting instance limit to int:", err)
		return
	} else if instanceLimitNumber < 1 {
		fmt.Println("Instance limit must be at least 1")
		return
	}


	utils.CDToLocation(
		utils.JoinAndConvertPathToOSFormat("..","..", "task_files", "go_scripts"),
	)

	err = utils.CopyDir("template", destDir)
	if err != nil {
		log.Fatalf("Failed to copy directory: %v", err)
	}

	utils.CDToLocation(
		utils.JoinAndConvertPathToOSFormat(".."),
	)

	cwd,_ := os.Getwd()

	tasksJsonFilePath := utils.JoinAndConvertPathToOSFormat(cwd,"tasks.json")

	content, err, _ := utils.CreateTasksJson(tasksJsonFilePath, false)
	cleanJSON, err := utils.CleanJSON(content)
	if err != nil {
		fmt.Println("Error cleaning JSON:", err)
		return
	}
	var tasksJSON utils.VSCodeTasksTasksJSON
	err = json.Unmarshal(cleanJSON, &tasksJSON)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	tasksJSON.Tasks = append(tasksJSON.Tasks, utils.VSCodeTasksTask{
		Label: strings.Join(strings.Split(destDir, "_"), " "),
		Type:"shell",
		RunOptions: utils.VSCodeTasksRunOptions{
			InstanceLimit: instanceLimitNumber,
		},
	})

  sort.Slice(tasksJSON.Tasks, func(i, j int) bool {
    return tasksJSON.Tasks[i].Label < tasksJSON.Tasks[j].Label
  })

	tasksJSONContent, err := json.MarshalIndent(tasksJSON, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling tasks.json:", err)
		return
	}

	err = os.WriteFile(tasksJsonFilePath, tasksJSONContent, 0644)
	if err != nil {
		fmt.Println("Error writing tasks.json file:", err)
		return
	}


}
