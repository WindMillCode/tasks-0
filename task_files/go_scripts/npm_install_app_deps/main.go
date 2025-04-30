package main

import (
	"main/shared"
	"os"
	"regexp"
	"sync"

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
	projectsCLI := utils.TakeVariableArgs(
		utils.TakeVariableArgsStruct{
			Prompt:  "Provide the paths of all the applications where you want the actions to take place",
			Default: workspaceRoot,
		},
	)

	packageManager := shared.ChooseNodePackageManager()
	cliInfo := utils.ShowMenuModel{
		Other:   true,
		Prompt:  "Choose the node.js app",
		Choices: settings.ExtensionPack.NPMInstallAppDeps.AppLocations,
	}
	if cliInfo.Choices == nil {
		cliInfo.Choices = settings.ExtensionPack.NodeJSAppLocations
	}
	appLocation := utils.ShowMenu(cliInfo, nil)

	cliInfo = utils.ShowMenuModel{
		Prompt:  "reinstall?",
		Choices: []string{"true", "false"},
	}
	reinstall := utils.ShowMenu(cliInfo, nil)

	cliInfo = utils.ShowMenuModel{
		Prompt:  "force?",
		Choices: []string{"true", "false"},
	}
	force := utils.ShowMenu(cliInfo, nil)

	var wg sync.WaitGroup
	regex0 := regexp.MustCompile(" ")
	projectsList := regex0.Split(projectsCLI.InputString, -1)
	for _, project := range projectsList {
		app := utils.JoinAndConvertPathToOSFormat(project, appLocation)

		wg.Add(1)
		go func() {
			defer wg.Done()
			if reinstall == "true" {
				os.Remove(utils.JoinAndConvertPathToOSFormat(app, "package-lock.json"))
				os.Remove(utils.JoinAndConvertPathToOSFormat(app, "yarn.lock"))
				os.RemoveAll(utils.JoinAndConvertPathToOSFormat(app, ".yarn"))
				os.Remove(utils.JoinAndConvertPathToOSFormat(app, ".pnp.js"))
				os.Remove(utils.JoinAndConvertPathToOSFormat(app, ".pnp.cjs"))
				os.Remove(utils.JoinAndConvertPathToOSFormat(app, ".pnp.loader.mjs"))
				os.Remove(utils.JoinAndConvertPathToOSFormat(app, "pnpm-lock.yaml"))
				os.RemoveAll(utils.JoinAndConvertPathToOSFormat(app, "node_modules"))
			}
			command := []string{"install"}
			if packageManager == "npm" {
				command = append(command, "-s")
			}
			if force == "true" {
				command = append(command, "--force")
			}
			command = append(command, "--verbose")
			utils.RunCommandInSpecificDirectory(packageManager, command, app)
		}()
	}
	wg.Wait()

}
