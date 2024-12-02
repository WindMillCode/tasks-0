package main

import (
	"main/shared"
	"os"
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
	var appPort string
	settingsAppPort := utils.IntToStr(settings.ExtensionPack.Ports.AngularAnalyzer0)
	if settingsAppPort != "0" {
		appPort = settingsAppPort
	} else {
		appPort = utils.IntToStr(9001)
	}
	executable := shared.ChooseNodePackageManager()
	cliInfo := utils.ShowMenuModel{
		Prompt:  "Choose an option:",
		Choices: []string{"dev", "preview", "prod"},
	}

	envType := utils.ShowMenu(cliInfo, nil)

	cliInfo = utils.ShowMenuModel{
		Prompt: "Choose analyzer (please choose vite-bundle-visualizer unless you know to choose webpack-bundle-analyzer)",
		Choices:[]string{"vite-bundle-visualizer","webpack-bundle-analyzer"},
	}
	analyzer := utils.ShowMenu(cliInfo,nil)
	var portStringArray []string
	if analyzer == "webpack-bundle-analyzer" {
		portStringArray = []string{"--"," --port ",appPort}
	}
	utils.CDToAngularApp()
	if executable == "npm" || executable == "pnpm" {
		utils.RunCommand(executable, append([]string{"run", "analyze:" + envType + ":" + analyzer}, portStringArray...))
	} else if executable == "yarn" {
		utils.RunCommand(executable, append([]string{"analyze:" + envType + ":" + analyzer },portStringArray...))
	}
}
