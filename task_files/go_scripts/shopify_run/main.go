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
	utils.CDToShopifyApp()
	shopifyFolder, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting shopify directory: %v", err)
		return
	}
	files, err := utils.GetItemsInFolder(shopifyFolder)
	if err != nil {
		fmt.Printf("Error getting shopify items: %v", err)
		return
	}

	var apps []string
	for _, file := range files {
		itemPath := shopifyFolder + "/" + file
		fileInfo, err := os.Stat(itemPath)
		if err != nil {
			fmt.Printf("Error getting info for %s: %v", itemPath, err)
			continue
		}
		if fileInfo.IsDir() {
			apps = append(apps, file)
		}
	}
	cliInfo := utils.ShowMenuModel{
		Prompt:  "select the shopify app",
		Choices: apps,
		Default: settings.ExtensionPack.ShopifyRun.ProjectName,
	}
	targetApp := utils.ShowMenu(cliInfo, nil)
	utils.CDToLocation(targetApp)

	opts := utils.CommandOptions{
		Command:     "npm",
		Args:        []string{"run", "dev"},
		PrintOutput: true,
	}
	utils.RunCommandWithOptions(opts)
}
