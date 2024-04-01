package main

import (
	"encoding/json"
	"fmt"
	"main/shared"
	"os"
	"regexp"
	"strings"

	"github.com/windmillcode/go_cli_scripts/v4/utils"
)

func main() {
	workSpaceFolder := os.Args[1]
	extensionFolder := os.Args[2]
	tasksJsonRelativeFilePath := os.Args[3]

	settings, err := utils.GetSettingsJSON(workSpaceFolder)
	if err != nil {
		return
	}

	// cliInfo := utils.ShowMenuModel{
	// 	Prompt:  "This will delete your vscode/tasks.json in your workspace folder. If you don't have a .vscode/tasks.json or you have not used this command before, it is safe to choose TRUE. Otherwise, consult with your manager before continuing",
	// 	Choices: []string{"TRUE", "FALSE"},
	// }
	// proceed := utils.ShowMenu(cliInfo, nil)
	proceed := "TRUE"

	cliInfo := utils.ShowMenuModel{
		Prompt:  "delete dest dir to ensure proper update (if updates are not taking place choose YES)",
		Choices: []string{"YES", "NO"},
		Default: "YES",
	}
	deleteDestDir := utils.ShowMenu(cliInfo, nil)

	cliInfo = utils.ShowMenuModel{
		Prompt:  "Run Tasks Via Interpreted or Complied Mode",
		Choices: []string{"COMPLIED", "INTERPRETED"},
		Default: "COMPLIED",
	}
	runMode := utils.ShowMenu(cliInfo, nil)

	// TODO implement lataer
	// cliInfo = utils.ShowMenuModel{
	// 	Prompt:  "use default user (if unsure select NO)",
	// 	Choices: []string{"NO", "YES"},
	// 	Default: "NO",
	// }
	// customUserIsPresent := utils.ShowMenu(cliInfo, nil)

	tasksJsonFilePath := utils.JoinAndConvertPathToOSFormat(extensionFolder, tasksJsonRelativeFilePath)
	content, err, shouldReturn := shared.CreateTasksJson(tasksJsonFilePath, false)
	if err != nil {
		fmt.Println("Error creating tasks json file", err)
		return
	}
	if shouldReturn {
		return
	}

	var tasksJSON shared.TasksJSON
	cleanJSON, err := utils.RemoveComments(content)
	if err != nil {
		fmt.Println("Error removing comments:", err)
		return
	}
	err = json.Unmarshal([]byte(cleanJSON), &tasksJSON)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	fmt.Println("\n\n\n If you see an unexpected end of input one of the array in your JSON has a ',' for its last item. this is not valid json remove that LAST comma\n\n\n")

	goExecutable := shared.GetGoExecutable()
	goScriptsSourceDirPath := utils.JoinAndConvertPathToOSFormat(extensionFolder, "task_files/go_scripts")
	goScriptsDestDirPath := utils.JoinAndConvertPathToOSFormat(workSpaceFolder, ".windmillcode/go_scripts")

	if proceed == "TRUE" {

		for index, task := range tasksJSON.Tasks {

			pattern0 := ":"
			regex0 := regexp.MustCompile(pattern0)
			programLocation0 := regex0.Split(task.Label, -1)
			pattern1 := " "
			regex1 := regexp.MustCompile(pattern1)
			programLocation1 := regex1.Split(strings.Join(programLocation0, ""), -1)
			programLocation2 := strings.Join(programLocation1, "_")
			programLocation3 := ".windmillcode//go_scripts//" + programLocation2
			taskExecutable := ".//main"
			if runMode == "INTERPRETED" || programLocation2 == "tasks_update_workspace_without_extension" {
				taskExecutable = fmt.Sprintf("%s %s", goExecutable, "run main.go")
			}
			linuxCommand0 := "cd " + programLocation3 + " ; " + taskExecutable
			windowsCommand0 := "cd " + strings.Replace(programLocation3, "//", "\\", -1) + " ; " + strings.Replace(taskExecutable+".exe", "//", "\\", -1)
			tasksJSON.Tasks[index].Windows.Command = windowsCommand0
			tasksJSON.Tasks[index].Osx.Command = linuxCommand0
			tasksJSON.Tasks[index].Linux.Command = linuxCommand0
			tasksJSON.Tasks[index].Linux.Options = shared.CommandOptions{
				Shell: shared.ShellOptions{
					Executable: "bash",
					Args:       []string{"-ic"},
				},
			}
			tasksJSON.Tasks[index].Metadata.Name = "windmillcode"
			if utils.ArrayContainsAny([]string{task.Label}, settings.ExtensionPack.TasksToRunOnFolderOpen) {
				tasksJSON.Tasks[index].RunOptions.RunOn = "folderOpen"
			}
		}
		workspaceTasksJSONFilePath := utils.JoinAndConvertPathToOSFormat(workSpaceFolder, "/.vscode/tasks.json")
		content, err, shouldReturn := shared.CreateTasksJson(workspaceTasksJSONFilePath, false)
		if err != nil {
			fmt.Println("Error creating tasks json file", err)
			return
		}
		if shouldReturn {
			return
		}
		var previousTasksJSON map[string]json.RawMessage
		previousCleanJSON, err := utils.RemoveComments(content)
		err = json.Unmarshal([]byte(previousCleanJSON), &previousTasksJSON)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			return
		}

		err = json.Unmarshal([]byte(previousCleanJSON), &previousTasksJSON)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			return
		}

		var previousTasks []json.RawMessage
		err = json.Unmarshal(previousTasksJSON["tasks"], &previousTasks)
		if err != nil {
			fmt.Println("Error unmarshalling tasks:", err)
			return
		}
		for i, taskRaw := range previousTasks {
			// Unmarshal the task into a map to access its properties
			var task map[string]interface{}
			if err := json.Unmarshal(taskRaw, &task); err != nil {
				fmt.Println("Error unmarshalling task:", err)
				continue
			}

			// Check if the task's label is in the TasksToRunOnFolderOpen list
			label, ok := task["label"].(string)
			if !ok {
				fmt.Println("Error: label is not a string")
				continue
			}
			var runOnValue = ""
			if utils.ArrayContainsAny([]string{label}, settings.ExtensionPack.TasksToRunOnFolderOpen) {
				runOnValue = "folderOpen"
			}
			runOptions, exists := task["runOptions"].(map[string]interface{})
			if !exists {
				runOptions = make(map[string]interface{})
				task["runOptions"] = runOptions
			}
			runOptions["runOn"] = runOnValue

			// Marshal the task back into json.RawMessage
			modifiedTaskRaw, err := json.Marshal(task)
			if err != nil {
				fmt.Println("Error marshalling task:", err)
				continue
			}

			// Update the previousTasks slice with the modified task
			previousTasks[i] = modifiedTaskRaw
		}

		currentTasksRaw := turnToDynamicJSONArray(tasksJSON.Tasks)

		var previousInputs []json.RawMessage
		err = json.Unmarshal(previousTasksJSON["inputs"], &previousInputs)
		currentInputsRaw := turnToDynamicJSONArray(tasksJSON.Inputs)
		var newTasksJSON shared.DynamicTasksJSON
		newTasksJSON.Version = tasksJSON.Version
		newTasksJSON.Tasks = append(filterJSONForOwnItems(previousTasks), currentTasksRaw...)
		newTasksJSON.Inputs = append(filterJSONForOwnItems(previousInputs), currentInputsRaw...)

		// marker
		tasksJSONData, err := json.MarshalIndent(newTasksJSON, "", "  ")
		if err != nil {
			fmt.Println("Error marshalling JSON:", err)
			return
		}

		workspaceTasksJSONFile, err := os.OpenFile(workspaceTasksJSONFilePath, os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer workspaceTasksJSONFile.Close()
		_, err = workspaceTasksJSONFile.Write(tasksJSONData)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	shared.RebuildExecutables(proceed, tasksJSON, goScriptsDestDirPath, goExecutable, preActions(deleteDestDir, goScriptsSourceDirPath, goScriptsDestDirPath))
}

func filterJSONForOwnItems(items []json.RawMessage) []json.RawMessage {
	var filteredItems []json.RawMessage
	for _, item := range items {
		var itemWithMetadata struct {
			Metadata shared.Metadata `json:"metadata"`
		}

		if err := json.Unmarshal(item, &itemWithMetadata); err != nil {
			fmt.Println("Error unmarshalling item:", err)
			continue
		}

		if itemWithMetadata.Metadata.Name != "windmillcode" {
			filteredItems = append(filteredItems, item)
		}
	}
	return filteredItems
}

func turnToDynamicJSONArray[T any](mySource []T) []json.RawMessage {
	var rawItems []json.RawMessage
	for _, item := range mySource {
		rawItem, err := json.Marshal(item)
		if err != nil {
			fmt.Println("Error marshalling:", err)
			continue
		}
		rawItems = append(rawItems, rawItem)
	}
	return rawItems
}

func preActions(deleteDestDir, goScriptsSourceDirPath, goScriptsDestDirPath string) func() {
	return func() {
		if deleteDestDir == "YES" {
			fmt.Println("Deleting Dest dir ...")
			if err := os.RemoveAll(goScriptsDestDirPath); err != nil {
				fmt.Println("Error:", err)
				return
			}
		}
		fmt.Println("Copying over files ...")
		utils.CopyDir(goScriptsSourceDirPath, goScriptsDestDirPath)
	}
}
