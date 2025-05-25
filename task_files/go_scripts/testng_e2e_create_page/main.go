package main

import (
	"fmt"
	"main/shared"
	"os"
	"strings"

	"github.com/iancoleman/strcase"
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
	cliInfo := utils.ShowMenuModel{
		Prompt:  "choose the suite location",
		Choices: settings.ExtensionPack.TestNGE2ECreatePage.AppLocations,
		Other:   true,
	}
	if cliInfo.Choices == nil {
		cliInfo.Choices = settings.ExtensionPack.JavaSuiteLocations
	}
	suiteLocation := utils.ShowMenu(cliInfo, nil)
	suiteLocation = utils.JoinAndConvertPathToOSFormat(workspaceRoot, suiteLocation)

	pageFolder := utils.JoinAndConvertPathToOSFormat(suiteLocation, "src", "main", "java", "pages")
	testFolder := utils.JoinAndConvertPathToOSFormat(suiteLocation, "src", "test", "java", "e2e")
	pageName := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt: []string{"The name of the screen (type in snake case)"},
			ErrMsg: "You must provide a page name",
		},
	)
	utils.CDToLocation(pageFolder)
	myPrefix := strcase.ToCamel(pageName)
	myDir := strings.ToLower(myPrefix)
	err = os.Mkdir(myDir, 0755)
	if err != nil {
		fmt.Printf("Error: ", err.Error())
	}
	myAct := utils.JoinAndConvertPathToOSFormat(pageFolder, myDir, myPrefix+"ActController.java")
	myPage := utils.JoinAndConvertPathToOSFormat(pageFolder, myDir, myPrefix+"Page.java")
	myVerify := utils.JoinAndConvertPathToOSFormat(pageFolder, myDir, myPrefix+"VerifyController.java")
	utils.CopyFile(utils.JoinAndConvertPathToOSFormat(".", "template", "TemplateActController.java"), myAct)
	utils.CopyFile(utils.JoinAndConvertPathToOSFormat(".", "template", "TemplatePage.java"), myPage)
	utils.CopyFile(utils.JoinAndConvertPathToOSFormat(".", "template", "TemplateVerifyController.java"), myVerify)

	utils.CDToLocation(testFolder)
	myTest := utils.JoinAndConvertPathToOSFormat(testFolder, myPrefix+"Test.java")
	utils.CopyFile(utils.JoinAndConvertPathToOSFormat(".", "TemplateTest.java"), myTest)
	for _, v := range []string{myAct, myPage, myVerify, myTest} {

		fileString, err := utils.ReadFile(v)
		if err != nil {
			return
		}
		fileString = strings.ReplaceAll(fileString, "Template", myPrefix)
		fileString = strings.ReplaceAll(fileString, "template", myDir)
		utils.OverwriteFile(v, fileString)
	}
}
