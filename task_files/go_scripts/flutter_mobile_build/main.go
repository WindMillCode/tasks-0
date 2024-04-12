package main

import (
	"fmt"
	"main/shared"
	"os"
	"path/filepath"
	"github.com/windmillcode/go_cli_scripts/v5/utils"
)

func main() {

	shared.CDToWorkspaceRoot()
	workspaceRoot, err := os.Getwd()
	if err != nil {
		return
	}
	settings, err := utils.GetSettingsJSON(workspaceRoot)
	if err != nil {
		return
	}
	utils.CDToFlutterApp()
	flutterRoot, err := os.Getwd()
	if err != nil {
		return
	}


	dartDefineArray := []string{}
	for _,val := range settings.ExtensionPack.FlutterMobileBuild.EnvVars{
		dartDefineArray = append(dartDefineArray, "--dart-define",val )
	}

	utils.RunCommand("dart", []string{"fix", "--apply"})
	utils.RunCommand("flutter", append([]string{"build", "appbundle"}, dartDefineArray...))
	releaseLocation := utils.JoinAndConvertPathToOSFormat(flutterRoot,"build\\app\\outputs\\bundle\\release")
	debugSymbolsFolder := utils.JoinAndConvertPathToOSFormat(flutterRoot,"build\\app\\intermediates\\merged_native_libs\\release\\out\\lib")
	zipFileName := "lib.zip"
	if err := zipFolder(debugSymbolsFolder, zipFileName); err != nil {
		fmt.Println("Error zipping folder:", err)
		return
	}
	srcPath := utils.JoinAndConvertPathToOSFormat(debugSymbolsFolder,zipFileName)
	destPath := utils.JoinAndConvertPathToOSFormat(releaseLocation, zipFileName)
	if err := os.Rename(srcPath, destPath); err != nil {
		fmt.Println("Error moving zip file:", err)
		return
	}

	utils.CDToLocation(releaseLocation)
	utils.RunCommand("code", []string{"app-release.aab"})



}

func zipFolder(srcFolder, destZip string) error {
	// Ensure the destination zip file has the .zip extension
	if filepath.Ext(destZip) != ".zip" {
		destZip += ".zip"
	}

	// Construct the command arguments for 7z
	args := []string{"a", destZip, srcFolder + "\\*"} // a for add, include all files from srcFolder

	// Run the 7z command
	_, err := utils.RunCommandWithOptions(utils.CommandOptions{
		Command: "7z",
		Args:    args,
		TargetDir: srcFolder,
	})

	if err != nil {
		fmt.Printf("Failed to zip folder: %s\n", err)
		return err
	}

	fmt.Println("Folder zipped successfully:", destZip)
	return nil
}
