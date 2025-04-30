package main

import (
	"fmt"
	"main/shared"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"github.com/windmillcode/go_cli_scripts/v6/utils"
	"regexp"
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

	vfoxPath, err := findCommandPath("vfox")
	askForPath := false
	if err != nil {
		askForPath = true
	} else if runtime.GOOS == "windows" {
		if !strings.HasSuffix(vfoxPath, "vfox.exe") {
			askForPath = true
		}
	} else {
		if !strings.HasSuffix(vfoxPath, "vfox") {
			askForPath = true
		}
	}
	if askForPath == true{
		vfoxPath = utils.GetInputFromStdin(
			utils.GetInputFromStdinStruct{
				Prompt: []string{"The path to vfox"},
				ErrMsg: "a value is required",
			},
		)
	}
	vfoxPath = utils.ConvertPathToOSFormat(vfoxPath)
  _, err = os.Stat(vfoxPath)
  if err != nil {
    fmt.Println("not a valid file path:", err)
  }
	if runtime.GOOS == "windows" {
		if !strings.HasSuffix(vfoxPath, "vfox.exe") {
			fmt.Println("this may not be the vfox executable")
			return
		}
	} else {
		if !strings.HasSuffix(vfoxPath, "vfox") {
			fmt.Println("this may not be the vfox executable")
			return
		}
	}




	newNodeVersion := utils.GetInputFromStdin(utils.GetInputFromStdinStruct{
		Prompt: []string{"Enter the new Node.js version to install:"},
		ErrMsg: "You must provide a version",
	})


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


	var globalPackages map[string]interface{}
	utils.ParseJSONFromString(globalPackagesJson, &globalPackages)

	packageNames := []string{}
	if dependencies, exists := globalPackages["dependencies"].(map[string]interface{}); exists {
		for pkgName := range dependencies {
			packageNames = append(packageNames, pkgName)
		}
	}




	var nodeVersions []string
	var nodeVersionsResult string
	nodeVersionsCmd := utils.CommandOptions{
		Command:     "vfox",
		Args:        []string{"list","nodejs"},
		GetOutput:   true,
		TargetDir:   "",
		PrintOutput: false,
	}
	nodeVersionsResult, err = utils.RunCommandWithOptions(nodeVersionsCmd)
	var lines []string
	if err != nil {
		fmt.Println("Error listing installed Node.js versions:", err)
		lines = []string{}
	} else{
		lines = strings.Split(nodeVersionsResult, "\n")
	}


	var semverRegex = regexp.MustCompile(`\d+\.\d+\.\d+(-[0-9A-Za-z-.]*)?`)

	for _, line := range lines {
			if line == "" {
					continue
			}


			match := semverRegex.FindString(line)

			if match != "" {
					nodeVersions = append(nodeVersions, match)
			}
	}
	fmt.Println("Installed Node.js versions:", nodeVersions)

	vfoxNodeVersion := fmt.Sprintf("nodejs@%s", newNodeVersion)
	if utils.ArrayContainsAny(nodeVersions, []string{newNodeVersion}) {
		fmt.Println("Node.js version", newNodeVersion, "is already available on the system.")
	} else {

		vfoxInstallCmd := utils.CommandOptions{
			Command: "vfox",
			Args:    []string{"install", vfoxNodeVersion},
			PrintOutput: true,
		}
		utils.RunCommandWithOptions(vfoxInstallCmd)
	}


	vfoxUseCmd := utils.CommandOptions{
		Command: "vfox",
		Args:    []string{"use","--global", vfoxNodeVersion},
	}
	utils.RunCommandWithOptions(vfoxUseCmd)

	regularInstallArray := []string{}
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
							var npmLinkCmd utils.CommandOptions
							if runtime.GOOS == "windows" {
								npmLinkCmd = utils.CommandOptions{
									Command:   "powershell",
									Args:      []string{"-c","npm link"},
									PrintOutput: true,
									TargetDir: utils.ConvertPathToOSFormat(resolvedPath),
								}
							} else {
								npmLinkCmd = utils.CommandOptions{
									Command:   "npm",
									Args:      []string{"link"},
									PrintOutput: true,
									TargetDir: utils.ConvertPathToOSFormat(resolvedPath),
								}
							}
							utils.RunCommandWithOptions(npmLinkCmd)
						}

					}
				}else {
					regularInstallArray = append(regularInstallArray, name)
				}
			}
		}


		var npmInstallCmd utils.CommandOptions
		if runtime.GOOS == "windows" {
			npmInstallCmd = utils.CommandOptions{
				Command: "powershell",
				Args: append([]string{"-c","npm install -g"}, regularInstallArray...),
				PrintOutput: true,
			}
		} else {
			npmInstallCmd = utils.CommandOptions{
				Command: "npm",
				Args: append([]string{"install", "-g"}, regularInstallArray...),
				PrintOutput: true,
			}
		}
		utils.RunCommandWithOptions(npmInstallCmd)

	}

	newNodeVersionViaCheckCmd := utils.CommandOptions{
		Command:     "node",
		Args:        []string{"-v"},
		GetOutput:   true,
		TargetDir:   "",
		PrintOutput: false,
	}
	newNodeVersionViaCheck, err := utils.RunCommandWithOptions(newNodeVersionViaCheckCmd)
	fmt.Println("Global packages have been migrated to Node.js version", newNodeVersionViaCheck)

}



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
