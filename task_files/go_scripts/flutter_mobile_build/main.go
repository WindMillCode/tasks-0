package main

import (
	"archive/zip"
	"fmt"
	"io"
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


	cliInfo := utils.ShowMenuModel{
		Prompt: "Choose an environment",
		Choices:settings.ExtensionPack.Environments,
		Default:"PROD",
	}
 	environment := utils.ShowMenu(cliInfo,nil)

	sentryDSN := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt: []string{"Provide the sentry dsn"},
			Default: settings.ExtensionPack.SentryDSN,
		},
	)
	utils.RunCommand("dart", []string{"fix", "--apply"})
	utils.RunCommand("flutter", []string{"build", "appbundle","--dart-define",fmt.Sprintf("FLUTTER_MOBILE_ENV=%s",environment),"--dart-define",fmt.Sprintf("SENTRY_DSN=%s",sentryDSN)})
	releaseLocation := utils.ConvertPathToOSFormat("build\\app\\outputs\\bundle\\release")
	debugSymbolsFolder := utils.ConvertPathToOSFormat("build\\app\\intermediates\\merged_native_libs\\release\\out\\lib")
	zipFileName := "debugSymbols.zip"
	if err := zipFolder(debugSymbolsFolder, zipFileName); err != nil {
		fmt.Println("Error zipping folder:", err)
		return
	}
	destPath := utils.JoinAndConvertPathToOSFormat(releaseLocation, zipFileName)
	if err := os.Rename(zipFileName, destPath); err != nil {
		fmt.Println("Error moving zip file:", err)
		return
	}

	utils.CDToLocation(releaseLocation)
	utils.RunCommand("code", []string{"app-release.aab"})



}

func zipFolder(srcFolder, destZip string) error {
	zipfile, err := os.Create(destZip)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	err = filepath.Walk(srcFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		relPath, err := filepath.Rel(srcFolder, path)
		if err != nil {
			return err
		}

		zipFile, err := archive.Create(relPath)
		if err != nil {
			return err
		}

		srcFile, err := os.Open(path)
		if err != nil {
			return err
		}
		defer srcFile.Close()

		_, err = io.Copy(zipFile, srcFile)
		return err
	})

	return err
}
