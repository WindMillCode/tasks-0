package main

import (
	"fmt"
	"main/shared"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

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
			NonInteractive: settings.ExtensionPack.ProcessIfDefaultIsPresent,
		},
	)



	cliInfo := utils.ShowMenuModel{
		Prompt: "Handle special packages",
		Choices:[]string{"YES","NO"},
		Default: "YES",
	}
	handleSpecialPackages := utils.ShowMenu(cliInfo,nil)


	// Prompt for new Node.js version
	newNodeVersion := utils.GetInputFromStdin(utils.GetInputFromStdinStruct{
		Prompt: []string{"Enter the new Node.js version to install:"},
	})
	// newNodeVersion := "20.18.0"

	currentNodeVersionCmd := utils.CommandOptions{
		Command:   "node",
		Args:      []string{"-v"},
		GetOutput: true,
	}
	currentNodeVersion, err := utils.RunCommandWithOptions(currentNodeVersionCmd)
	if err != nil {
		fmt.Println("Error getting current Node.js version:", err)
		return
	}
	currentNodeVersion = strings.TrimSpace(currentNodeVersion)

	executablePath, _ := os.Executable()
	scriptDir := filepath.Dir(executablePath)
	listGlobalPackagesCmd := utils.CommandOptions{
		Command:   "npm",
		Args:      []string{"ls", "-g", "--depth=0", "--long", "--json"},
		GetOutput: true,
		TargetDir: scriptDir,
	}
	globalPackagesJson, err := utils.RunCommandWithOptions(listGlobalPackagesCmd)
	if err != nil {
		fmt.Println("Error listing global packages:", err)
		return
	}

	// Parse the JSON output to extract package names
	var globalPackages map[string]interface{}
	utils.ParseJSONFromString(globalPackagesJson, &globalPackages)

	packageNames := []string{}
	if dependencies, exists := globalPackages["dependencies"].(map[string]interface{}); exists {
		for pkgName := range dependencies {
			packageNames = append(packageNames, pkgName)
		}
	}

	nvmPath, err := findCommandPath("nvm")
	if err != nil {
		fmt.Println("Error finding NVM path:", err)
		return
	}

	nodeVersionsPath := utils.JoinAndConvertPathToOSFormat(filepath.Dir(nvmPath))
	if runtime.GOOS != "windows" {
		nodeVersionsPath = utils.JoinAndConvertPathToOSFormat(filepath.Dir(nvmPath), "..", "versions", "node")
	}
	var nodeVersions []string
	err = utils.TraverseDirectory(utils.TraverseDirectoryParams{
		RootDir: nodeVersionsPath,
		Predicate: func(path string, info os.FileInfo) {
			if info.IsDir() {
				semverRegex := regexp.MustCompile(`^v(\d+\.\d+\.\d+)$`)
				matches := semverRegex.FindStringSubmatch(info.Name())
				if matches != nil && matches[1] == newNodeVersion {
					nodeVersions = append(nodeVersions, matches[1])
				}
			}
		},
	})
	if err != nil {
		fmt.Println("Error listing installed Node.js versions:", err)
		return
	}

	if utils.ArrayContainsAny(nodeVersions, []string{newNodeVersion}) {
		fmt.Println("Node.js version", newNodeVersion, "is already available on the system.")
	} else {
		fmt.Println("Installing Node.js version", newNodeVersion, "...")

		// Install the new Node.js version using nvm
		nvmInstallCmd := utils.CommandOptions{
			Command: nvmPath,
			Args:    []string{"install", newNodeVersion},
		}
		utils.RunCommandWithOptions(nvmInstallCmd)
	}


	nvmUseCmd := utils.CommandOptions{
		Command: "nvm",
		Args:    []string{ "use", newNodeVersion},
	}
	utils.RunCommandWithOptions(nvmUseCmd)
	// Reinstall global packages

	if dependencies, ok := globalPackages["dependencies"].(map[string]interface{}); ok {
		for name, pkg := range dependencies {
			fmt.Println(name)
			if pkgMap, ok := pkg.(map[string]interface{}); ok {

				if handleSpecialPackages =="YES" && utils.ArrayContainsAny([]string{name}, []string{"pnpm","corepack"}) {
					if name == "pnpm" {
						corePackEnableCommand := utils.CommandOptions{
							Command: "corepack",
							Args:    []string{"enable"},
						}
						utils.RunCommandWithOptions(corePackEnableCommand)
					}
				} else if resolved, ok := pkgMap["resolved"].(string); ok && resolved != "" {
					if linkAbsPath, ok := pkgMap["path"].(string); ok && linkAbsPath != "" {

						resolvedPath, err := os.Readlink(linkAbsPath)
						if err != nil {
							fmt.Printf("Error resolving symlink: %v", err)
						} else{
							npmLinkCmd := utils.CommandOptions{
								Command:   "npm",
								Args:      []string{"link"},
								TargetDir: utils.ConvertPathToOSFormat(resolvedPath),
							}
							utils.RunCommandWithOptions(npmLinkCmd)
						}

					}
				}else {
					npmInstallCmd := utils.CommandOptions{
						Command: "npm",
						Args:    []string{"install", "-g", name},
					}
					utils.RunCommandWithOptions(npmInstallCmd)
				}
			}

		}

	}

	fmt.Println("Global packages have been migrated to Node.js version", newNodeVersion)

}

// findCommandPath attempts to locate the specified command path in a cross-platform manner.

func findCommandPath(commandName string) (string, error) {
	var commandCheck string
	var commandArgs []string

	switch runtime.GOOS {
	case "windows":
		// En Windows, usamos `Get-Command` en PowerShell para localizar el comando
		commandCheck = "powershell"
		commandArgs = []string{"-Command", fmt.Sprintf("Get-Command %s | Select-Object -ExpandProperty Definition", commandName)}
	case "linux":
		// En sistemas Unix-like, usamos `command -v` o `which`
		commandCheck = "sh"
		commandArgs = []string{"-c", fmt.Sprintf("command -v %s || which %s", commandName, commandName)}

	// TODO cant get mac os to work
	// case "darwin":
	// 	commandCheck = "zsh"
	// 	commandArgs = []string{"-c",fmt.Sprintf("(type -a %s | tail -n 1 | awk '{print $NF}')",commandName)}

	default:
		return "", fmt.Errorf("unsupported platform")
	}

	// Usar la función de utils para ejecutar el comando
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
