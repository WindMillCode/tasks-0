package main

import (
	"fmt"
	"main/shared"
	"os"
	"strings"
	"github.com/windmillcode/go_cli_scripts/v6/utils"
)

func main() {

	scriptLocation, err := os.Getwd()
	if err != nil {
		return
	}
	templateLocation := utils.JoinAndConvertPathToOSFormat(scriptLocation, "template")
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

	componentName := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt: []string{"The name you would like to give to the shared component"},
			ErrMsg: "You must provide a value",
		},
	)
	componentNameObject,_ := utils.CreateStringObject(componentName,"")

	utils.CDToExpoApp()
	expoApp, err := os.Getwd()
	if err != nil {
		return
	}
	providerLocation := utils.JoinAndConvertPathToOSFormat(
		expoApp,
		"components",
		componentNameObject.PascalCase(false,""),
	)
	utils.CopyDir(templateLocation, providerLocation)
	newSharedComponentPath := utils.JoinAndConvertPathToOSFormat(providerLocation, fmt.Sprintf("%s.tsx",componentNameObject.PascalCase(false,"")))
	os.Rename(
		utils.JoinAndConvertPathToOSFormat(providerLocation,"template.tsx"),
		newSharedComponentPath,
	)

	utils.TraverseDirectory(utils.TraverseDirectoryParams{
		RootDir: providerLocation,
		Predicate: func(srcPath string, info os.FileInfo) {
			fileString, err := utils.ReadFile(srcPath)
			if err != nil {
				return
			}
			fileString = strings.ReplaceAll(fileString, "WMLTemplate", componentNameObject.Capitalize(false,""))
			utils.OverwriteFile(srcPath, fileString)
		},
		Filter: func(srcPath string, info os.FileInfo) bool {
			return !info.IsDir()
		},
	})

}
