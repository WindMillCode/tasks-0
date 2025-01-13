package main

import (
	"fmt"
	"main/shared"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
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

	pyenvPath, err := findCommandPath("pyenv")
	// pyenvSubCommands := utils.JoinAndConvertPathToOSFormat(filepath.Dir(pyenvPath),"..","libexec")
	if err != nil || !strings.Contains(pyenvPath, "pyenv") {
		pyenvPath = utils.GetInputFromStdin(
			utils.GetInputFromStdinStruct{
				Prompt: []string{"Enter the path to pyenv:"},
				ErrMsg: "A value is required.",
			},
		)
	}
	pyenvPath = utils.ConvertPathToOSFormat(pyenvPath)
	if _, err := os.Stat(pyenvPath); err != nil {
		fmt.Println("Invalid file path:", err)
		return
	}
	if  !strings.HasSuffix(pyenvPath, "pyenv") {
		if runtime.GOOS == "windows" {
			pyenvPath = strings.TrimSuffix(pyenvPath, ".ps1")
		} else{
			fmt.Println("This may not be the pyenv executable.",pyenvPath)
			return
		}
	}

	pythonVersionsPath := utils.JoinAndConvertPathToOSFormat(filepath.Dir(pyenvPath),"..","versions")

	// Prompt for new Python version
	newPythonVersion := utils.GetInputFromStdin(utils.GetInputFromStdinStruct{
		Prompt: []string{"Enter the new Python version to install:"},
	})

	cliInfo := utils.ShowMenuModel{
		Prompt: "Reinstall Python version",
		Choices:[]string{"YES","NO"},
	}
	reinstallPythonVersion := utils.ShowMenu(cliInfo,nil)

	var pythonVersions []string
	err = utils.TraverseDirectory(utils.TraverseDirectoryParams{
		RootDir: pythonVersionsPath,
		Predicate: func(path string, info os.FileInfo) {
			if info.IsDir() {
				semverRegex := regexp.MustCompile(`^(?:v)?(\d+\.\d+\.\d+)$`)
				matches := semverRegex.FindStringSubmatch(info.Name())
				if matches != nil {
					// Log matched version for debugging
					pythonVersions = append(pythonVersions, matches[1])
				}
			}
		},
	})
	if err != nil {
		fmt.Println("Error listing installed Python versions:", err)
		return
	}

	if utils.ArrayContainsAny(pythonVersions, []string{newPythonVersion}) && reinstallPythonVersion != "YES" {
		fmt.Println("Python version", newPythonVersion, "is already available on the system.")
	} else {

		if reinstallPythonVersion == "YES"{
			if err := os.RemoveAll(filepath.Join(pythonVersionsPath, newPythonVersion)); err != nil {
				fmt.Println("Error:", err)
			}
		}

		fmt.Println("Installing Python version", newPythonVersion, "...")

		// Install the new Python version using pyenv
		pyenvInstallCmd := utils.CommandOptions{
			Command: pyenvPath, // Use cscript.exe to run the .vbs script
			Args: []string{"install",newPythonVersion},
			GetOutput:   true,
			PrintOutput: true,
			ExitRegex: ":: [Info] :: completed!",
		}
		_,err := utils.RunCommandWithOptions(pyenvInstallCmd)
		if err != nil {
			fmt.Println("Error installing Python version:", err)
			return
		}
	}


	// Reinstall global Python packages
	fmt.Println("Reinstalling global Python packages...")
	globalPackagesCmd := utils.CommandOptions{
		Command:   pyenvPath,
		Args:      []string{"exec", "pip", "freeze"},
		GetOutput: true,
	}
	globalPackages, err := utils.RunCommandWithOptions(globalPackagesCmd)
	if err != nil {
		fmt.Println("Error retrieving global Python packages:", err)
		return
	}

	pyenvUseCmd := utils.CommandOptions{
		Command: pyenvPath,
		Args:    []string{"global", newPythonVersion},
	}
	utils.RunCommandWithOptions(pyenvUseCmd)


	packages := strings.Split(globalPackages, "\n")
	for _, pkg := range packages {
		pkgName := strings.Split(pkg, "==")[0]
		if pkgName != "" {
			installCmd := utils.CommandOptions{
				Command: pyenvPath,
				Args:    []string{"exec", "pip", "install", pkgName},
			}
			utils.RunCommandWithOptions(installCmd)
		}
	}

	fmt.Println("Global packages have been migrated to Python version", newPythonVersion)
}

// findCommandPath attempts to locate the specified command path in a cross-platform manner.
func findCommandPath(commandName string) (string, error) {
	var commandCheck string
	var commandArgs []string

	switch runtime.GOOS {
	case "windows":
		commandCheck = "powershell"
		commandArgs = []string{"-Command", fmt.Sprintf("Get-Command %s | Select-Object -ExpandProperty Definition", commandName)}
	case "linux", "darwin":
		commandCheck = "sh"
		commandArgs = []string{"-c", fmt.Sprintf("command -v %s || which %s", commandName, commandName)}
	default:
		return "", fmt.Errorf("unsupported platform")
	}

	output, err := utils.RunCommandWithOptions(utils.CommandOptions{
		Command:   commandCheck,
		Args:      commandArgs,
		GetOutput: true,
	})

	if err != nil {
		return "", fmt.Errorf("could not find %s: %w", commandName, err)
	}

	commandPath := strings.TrimSpace(output)
	if commandPath == "" {
		return "", fmt.Errorf("%s command not found", commandName)
	}

	return commandPath, nil
}
