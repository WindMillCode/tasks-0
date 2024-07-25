package main

import (
	"log"
	"main/shared"
	"os"

	"github.com/windmillcode/go_cli_scripts/v5/utils"
)

func main() {

	shared.CDToWorkspaceRoot()
	utils.CDToShopifyApp()
	shopifyFolder, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting shopify directory: %v", err)
		return;
	}
	files, err := utils.GetItemsInFolder(shopifyFolder)
	if err != nil {
		log.Fatalf("Error getting shopify items: %v", err)
		return
	}

	var apps []string
	for _, file := range files {
		itemPath := shopifyFolder + "/" + file
		fileInfo, err := os.Stat(itemPath)
		if err != nil {
			log.Printf("Error getting info for %s: %v", itemPath, err)
			continue
		}
		if fileInfo.IsDir() {
			apps = append(apps, file)
		}
	}
	cliInfo := utils.ShowMenuModel{
		Prompt: "select the shopify app",
		Choices:apps,
	}
	targetApp := utils.ShowMenu(cliInfo,nil)
	utils.CDToLocation(targetApp)


	opts := utils.CommandOptions{
		Command: "npm",
		Args:    []string{"run","dev"},
		PrintOutput: true,
	}
	utils.RunCommandWithOptions(opts)
}
