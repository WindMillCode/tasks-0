package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/windmillcode/go_cli_scripts/v4/utils"
)

func main() {

	utils.CDToWorkspaceRoot()
	workspaceFolder, err := os.Getwd()
	if err != nil {
		fmt.Println("there was an error while trying to receive the current dir")
	}
	settings, err := utils.GetSettingsJSON(workspaceFolder)
	if err != nil {
		return
	}
	utils.CDToFlaskApp()
	flaskAppFolder, err := os.Getwd()
	if err != nil {
		return
	}
	envVarsFile := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt:  []string{"where are the env vars located"},
			Default: utils.JoinAndConvertPathToOSFormat(workspaceFolder, settings.ExtensionPack.FlaskBackendTestHelperScript),
		},
	)
	pythonVersion := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt:  []string{"provide a python version for pyenv to use"},
			Default: settings.ExtensionPack.PythonVersion0,
		},
	)
	if pythonVersion != "" {
		utils.RunCommand("pyenv", []string{"global", pythonVersion})
	}
	utils.CDToLocation(workspaceFolder)

	envVarCommandOptions := utils.CommandOptions{
		Command:      "windmillcode_go",
		Args:         []string{"run", envVarsFile, filepath.Dir(utils.JoinAndConvertPathToOSFormat(envVarsFile)), workspaceFolder},
		GetOutput:    true,
		TargetDir:     filepath.Dir(utils.JoinAndConvertPathToOSFormat(envVarsFile)),
	}
	envVars,err := utils.RunCommandWithOptions(envVarCommandOptions)
	if err != nil {
		return
	}
	envVarsArray := strings.Split(envVars, ",")


	for _, x := range envVarsArray {
		// Splitting only on the first "="
		keyPair := strings.SplitN(x, "=", 2)

		if len(keyPair) != 2 {
				fmt.Println("Invalid format in environment variable:", x)
				continue
		}

		key := strings.TrimSpace(keyPair[0])
		value := strings.TrimSpace(strings.ReplaceAll(keyPair[1], ",", ""))

		os.Setenv(key, value)
	}
	flaskAppUnitTestFolder := utils.JoinAndConvertPathToOSFormat(
		flaskAppFolder, "unit_tests",
	)
	utils.CDToLocation(flaskAppUnitTestFolder)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		port := 8004
		utils.CDToLocation(utils.JoinAndConvertPathToOSFormat("covhtml"))
		http.Handle("/", http.FileServer(http.Dir(".")))
		fmt.Println(fmt.Sprintf("Coverage info accesible at localhost:%d", port))
		err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
		if err != nil {
			fmt.Println("Error starting the server:", err)
		}
	}()


	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			utils.RunCommandWithOptions(
				utils.CommandOptions{
					Command: "python",
					Args: []string{"run_tests.py"},
					TargetDir:  flaskAppUnitTestFolder,
					PrintOutputOnly: true,
				},
			)
		}
	}()

	wg.Wait()

}


