package main

import (
	"fmt"
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
			NonInteractive: settings.ExtensionPack.ProcessIfDefaultIsPresent,
		},
	)
	utils.CDToReactNativeExpoApp()

	cliInfo := utils.ShowMenuModel{
		Prompt: "profile",
		Choices:[]string{"development","preview","production"},
	}
	myProfile := utils.ShowMenu(cliInfo,nil)

	cliInfo = utils.ShowMenuModel{
		Prompt: "platform",
		Choices:[]string{"android","ios"},
	}
	myPlatform := utils.ShowMenu(cliInfo,nil)

	cliInfo = utils.ShowMenuModel{
		Prompt: "build locally",
		Choices:[]string{"TRUE","FALSE"},
	}
	localBuild := utils.ShowMenu(cliInfo,nil)

	fileExt := "ipa"

	if myPlatform == "android" && localBuild == "TRUE" {

		cliInfo := utils.ShowMenuModel{
			Prompt: "android file extension",
			Choices:[]string{"apk","aab"},
		}
		fileExt = utils.ShowMenu(cliInfo,nil)
	}

	outputDir := ""
	if localBuild == "TRUE" {
		outputDir = utils.GetInputFromStdin(
			utils.GetInputFromStdinStruct{
				Prompt: []string{"The output dir"},
				Default: utils.JoinAndConvertPathToOSFormat(workspaceRoot,"misc",fmt.Sprintf("local-%s-%s.%s",myProfile,myPlatform,fileExt)),
			},
		)
	}

	commandArgs:= []string{"build", "--profile", myProfile, "--platform", myPlatform }
	if localBuild == "TRUE" {
		commandArgs = append(commandArgs, "--output", outputDir)
	}

	if localBuild == "TRUE" {
		commandArgs = append(commandArgs, "--local")
	}

	opts := utils.CommandOptions{
		Command: "eas",
		Args:    commandArgs,
		GetOutput:       false,
		PrintOutputOnly: true,
	}
	utils.RunCommandWithOptions(opts)
}
