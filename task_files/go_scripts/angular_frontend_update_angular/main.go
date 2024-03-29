package main

import (
	"fmt"
	"main/shared"
	"os"
	"regexp"
	"strings"
	"sync"

	"github.com/windmillcode/go_cli_scripts/v4/utils"
)

func main() {

	utils.CDToWorkspaceRoot()
	workspaceRoot, err := os.Getwd()
	if err != nil {
		fmt.Println("there was an error while trying to receive the current dir")
	}
	projectsCLI := utils.TakeVariableArgs(
		utils.TakeVariableArgsStruct{
			Prompt:  "Provide the paths of all the projects where you want the actions to take place",
			Default: workspaceRoot,
		},
	)
	angularAppLocation := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt:  []string{"provide the relative path to the angular app (note : for every project  the relative path should be the same other )"},
			Default: "apps/frontend/AngularApp",
		},
	)

	var wg sync.WaitGroup
	regex0 := regexp.MustCompile(" ")
	projectsList := regex0.Split(projectsCLI.InputString, -1)
	for _, project := range projectsList {
		rootProject := project
		AngularApp := utils.JoinAndConvertPathToOSFormat(project, angularAppLocation)
		wg.Add(1)
		go func() {
			defer wg.Done()
			updateAngular(rootProject, AngularApp)
		}()
	}
	wg.Wait()

}

func updateAngular(project string, angularApp string) {
	utils.RunCommandInSpecificDirectory("git", []string{"add", "."}, project)
	utils.RunCommandInSpecificDirectory("git", []string{"commit", "-m", "[CHECKPOINT] before upgrading to next angular version"}, project)
	inputText := utils.RunCommandInSpecifcDirectoryAndGetOutput("npx", []string{"ng", "update"}, angularApp)
	inputLines := strings.Split(inputText, "\n")
	packagesToUpdate := []string{}
	for _, line := range inputLines {
		if strings.Contains(line, "ng update @") {
			packagesToUpdate = append(packagesToUpdate, line)
		}
	}
	updateCommand := "ng update"
	for _, pkg := range packagesToUpdate {
		packageGroup := strings.TrimSpace(strings.Split(pkg, "->")[0])
		packageName := strings.TrimSpace(strings.Split(packageGroup, " ")[0])
		updateCommand += " " + packageName
	}
	utils.RunCommandInSpecificDirectory("npx", strings.Split(updateCommand, " "), angularApp)
	cliInfo := utils.ShowMenuModel{
		Prompt:  "auto update additional packages",
		Choices: []string{"TRUE", "FALSE"},
	}
	deps := []string{"@windmillcode/angular-wml-components-base", "@rxweb/reactive-form-validators", "@fortawesome/fontawesome-free", "@compodoc/compodoc", "@sentry/angular-ivy", "@sentry/tracing"}
	devDeps := []string{"@faker-js/faker", "@windmillcode/angular-templates", "webpack-bundle-analyzer", "browserify"}
	addtl := utils.ShowMenu(cliInfo, nil)
	if addtl == "TRUE" {

		packageManager := shared.ChooseNodePackageManager()
		if packageManager == "yarn" {
			utils.RunCommandInSpecificDirectory("yarn", append([]string{"upgrade"}, deps...), angularApp)
			utils.RunCommandInSpecificDirectory("yarn", append([]string{"upgrade", "--dev"}, devDeps...), angularApp)
		} else if (packageManager == "pnpm") {
			utils.RunCommandInSpecificDirectory("pnpm", append([]string{"update"}, deps...), angularApp)
		} else {
			utils.RunCommandInSpecificDirectory("npm", append([]string{"update"}, deps...), angularApp)
			utils.RunCommandInSpecificDirectory("npm", append([]string{"update", "--include=dev"}, devDeps...), angularApp)
		}
	}
}
