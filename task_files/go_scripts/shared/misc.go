package shared

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"

	"github.com/windmillcode/go_cli_scripts/v4/utils"
)

type ShellOptions struct {
	Executable string   `json:"executable"`
	Args       []string `json:"args"`
}

type CommandOptions struct {
	Shell ShellOptions `json:"shell"`
}

type Task struct {
	Label   string `json:"label"`
	Type    string `json:"type"`
	Windows struct {
		Command string `json:"command"`
	} `json:"windows"`
	Linux struct {
    Command string       `json:"command"`
    Options CommandOptions `json:"options"`
	} `json:"linux"`
	Osx struct {
		Command string `json:"command"`
	} `json:"osx"`
	RunOptions struct {
		RunOn         string `json:"runOn"`
		InstanceLimit int    `json:"instanceLimit"`
	} `json:"runOptions"`
}

type Input struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Default     string `json:"default"`
	Type        string `json:"type"`
}

type TasksJSON struct {
	Version string  `json:"version"`
	Tasks   []Task  `json:"tasks"`
	Inputs  []Input `json:"inputs"`
}

func RebuildExecutables(proceed string, tasksJSON TasksJSON, goScriptsDestDirPath string, goExecutable string) {
	var rebuild string
	var cliInfo utils.ShowMenuModel
	if proceed == "TRUE" {
		rebuild = "TRUE"
	} else {
		cliInfo = utils.ShowMenuModel{
			Prompt:  "Do you want to rebuild the go programs into binary exectuables ",
			Choices: []string{"TRUE", "FALSE"},
		}
		rebuild = utils.ShowMenu(cliInfo, nil)
	}

	if rebuild == "TRUE" {

		cliInfo = utils.ShowMenuModel{
			Prompt:  "Build one by one or all at once",
			Choices: []string{"ALL_AT_ONCE", "ONE_BY_ONE"},
		}
		inSync := utils.ShowMenu(cliInfo, nil)

		if inSync == "ALL_AT_ONCE" {
			var wg sync.WaitGroup
			fmt.Print(len(tasksJSON.Tasks))
			for _, task := range tasksJSON.Tasks {
				wg.Add(1)

				pattern0 := ":"
				regex0 := regexp.MustCompile(pattern0)
				programLocation0 := regex0.Split(task.Label, -1)
				pattern1 := " "
				regex1 := regexp.MustCompile(pattern1)
				programLocation1 := regex1.Split(strings.Join(programLocation0, ""), -1)
				programLocation2 := strings.Join(programLocation1, "_")
				absProgramLocation := utils.JoinAndConvertPathToOSFormat(goScriptsDestDirPath, programLocation2)
				go func() {
					defer wg.Done()
					BuildGoCLIProgram(absProgramLocation, goExecutable)
				}()
			}
			wg.Wait()
		} else{
			fmt.Print(len(tasksJSON.Tasks))
			for _, task := range tasksJSON.Tasks {

				pattern0 := ":"
				regex0 := regexp.MustCompile(pattern0)
				programLocation0 := regex0.Split(task.Label, -1)
				pattern1 := " "
				regex1 := regexp.MustCompile(pattern1)
				programLocation1 := regex1.Split(strings.Join(programLocation0, ""), -1)
				programLocation2 := strings.Join(programLocation1, "_")
				absProgramLocation := utils.JoinAndConvertPathToOSFormat(goScriptsDestDirPath, programLocation2)
				BuildGoCLIProgram(absProgramLocation, goExecutable)

			}
		}

	}
}

func BuildGoCLIProgram(programLocation string, goExecutable string) {

	fmt.Printf("%s \n", programLocation)
	utils.RunCommandInSpecificDirectory(goExecutable, []string{"build", "main.go"}, programLocation)
	fmt.Printf("Finished building %s \n", programLocation)

}

func CreateTasksJson(tasksJsonFilePath string, triedCreateOnError bool) ([]byte, error, bool) {
	content, err := os.ReadFile(tasksJsonFilePath)
	if err != nil {
		if triedCreateOnError {
			return nil, err, true
		}

		// If the file doesn't exist, create it.
		_, createErr := os.Create(tasksJsonFilePath)
		if createErr != nil {
			return nil, createErr, true
		}

		// Recursively attempt to read the file after creating it.
		return CreateTasksJson(tasksJsonFilePath, true)
	}

	return content, nil, false
}

func SetupEnvironmentToRunFlaskApp() (string, error) {
	utils.CDToWorkspaceRoot()
	workspaceFolder, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("error getting current directory: %w", err)
	}
	settings, err := utils.GetSettingsJSON(workspaceFolder)
	if err != nil {
		return "", fmt.Errorf("error getting settings JSON: %w", err)
	}
	utils.CDToFlaskApp()
	flaskAppFolder, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("error changing to Flask app directory: %w", err)
	}

	helperScript := settings.ExtensionPack.FlaskBackendDevHelperScript
	if utils.IsRunningInDocker() {
		helperScript = strings.Replace(helperScript, "dev", "docker_dev", 1)
	}

	cliInfo := utils.ShowMenuModel{
		Prompt: "Where are the env vars located",
		Choices: []string{
			utils.JoinAndConvertPathToOSFormat(workspaceFolder, helperScript),
		},
		Other: true,
	}
	envVarsFile := utils.ShowMenu(cliInfo, nil)

	pythonVersion := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt:  []string{"Provide a Python version for pyenv to use"},
			Default: settings.ExtensionPack.PythonVersion0,
		},
	)

	if pythonVersion != "" {
		utils.RunCommand("pyenv", []string{"global", pythonVersion})
	}

	utils.CDToLocation(workspaceFolder)
	envVarCommandOptions := utils.CommandOptions{
		Command:     GetGoExecutable(),
		Args:        []string{"run", envVarsFile, filepath.Dir(utils.JoinAndConvertPathToOSFormat(envVarsFile)), workspaceFolder},
		GetOutput:   true,
		TargetDir:   filepath.Dir(utils.JoinAndConvertPathToOSFormat(envVarsFile)),
		PrintOutput: false,
	}
	envVars, err := utils.RunCommandWithOptions(envVarCommandOptions)
	if err != nil {
		return "", fmt.Errorf("error running command for env vars: %w", err)
	}

	envVarsArray := strings.Split(envVars, ",")
	for _, x := range envVarsArray {
		indexOfEqual := strings.Index(x, "=")
		if indexOfEqual == -1 {
			continue
		}
		key := x[:indexOfEqual]
		value := x[indexOfEqual+1:]
		key = strings.TrimSpace(key)
		value = strings.TrimSpace(value)
		os.Setenv(key, value)
	}

	return flaskAppFolder, nil
}


func GetGoExecutable() string {
	cliInfo := utils.ShowMenuModel{
		Prompt:  "choose the executable to use (try with windmillcode_go first if not then use go)",
		Choices: []string{"go", "windmillcode_go"},
		Default: "windmillcode_go",
	}
	goExecutable := utils.ShowMenu(cliInfo, nil)
	return goExecutable
}
