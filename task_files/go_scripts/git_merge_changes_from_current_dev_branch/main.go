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
	sourceBranch := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt:  []string{"the branch to merge changes from:"},
			Default: "dev",
		},
	)

	utils.RunCommand("git", []string{"checkout", sourceBranch})
	utils.RunCommand("git", []string{"pull", "origin", sourceBranch})
	utils.RunCommand("git", []string{"checkout", "-"})
	utils.RunCommand("git", []string{"merge", sourceBranch})
}
