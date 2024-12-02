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
			NonInteractive: settings.ExtensionPack.ProcessIfDefaultIsPresent,
		},
	)
	utils.CDToWxtApp()


	opts := utils.CommandOptions{
		Command: "npx",
		Args:    []string{"wxt", "build", "-b", "safari", "--mv3"},
		GetOutput:       false,
		PrintOutputOnly: true,
	}
	utils.RunCommandWithOptions(opts)

	commandOptions := utils.CommandOptions{
		Command:     "xcrun",
		Args:        []string{"safari-web-extension-converter", ".output/safari-mv3", "--bundle-identifier", settings.ExtensionPack.WxtBuildSafari.BundleIdentifier},
		GetOutput:       false,
		PrintOutputOnly: true,
	}
	utils.RunCommandWithOptions(commandOptions)
}
