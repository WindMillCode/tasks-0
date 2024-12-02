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
			NonInteractive :settings.ExtensionPack.ProcessIfDefaultIsPresent,
		},
	)
	utils.CDToWxtApp()
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v", err)
		return
	}
	os.RemoveAll(
		utils.JoinAndConvertPathToOSFormat(
			currentDir,
			"node_modules",
			".vite",
		),
	)
	opts := utils.CommandOptions{
		Command: "npm",
		Args:    []string{"run","dev"},
		GetOutput:       false,
		PrintOutputOnly: true,
	}
	utils.RunCommandWithOptions(opts)


}
