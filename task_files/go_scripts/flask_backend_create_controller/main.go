package main

import (
	"fmt"
	"main/shared"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"github.com/iancoleman/strcase"
	"github.com/windmillcode/go_cli_scripts/v6/utils"
)

func main() {

	scriptLocation, err := os.Getwd()
	if err != nil {
		return
	}
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
	templateLocation := utils.JoinAndConvertPathToOSFormat(scriptLocation, "template")
	templateEndpointFile := utils.JoinAndConvertPathToOSFormat(templateLocation, "template_endpoint.py")
	templateHandlerFile := utils.JoinAndConvertPathToOSFormat(templateLocation, "template_handler.py")
	testTemplateEndpointFile := utils.JoinAndConvertPathToOSFormat(templateLocation, "test_template_endpoint.py")
	testTemplateHandlerFile := utils.JoinAndConvertPathToOSFormat(templateLocation, "test_template_handler.py")

	utils.CDToFlaskApp()
	targetApp, err := os.Getwd()
	if err != nil {
		return
	}

	endpointsFolder := utils.JoinAndConvertPathToOSFormat(targetApp, "endpoints")
	handlersFolder := utils.JoinAndConvertPathToOSFormat(targetApp, "handlers")
	testEndpointsFolder := utils.JoinAndConvertPathToOSFormat(targetApp,"unit_tests","FlaskTesting",  "endpoints")
	testHandlersFolder := utils.JoinAndConvertPathToOSFormat(targetApp, "unit_tests","FlaskTesting", "handlers")

	targetName := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt: []string{"Please provide the name of the controller"},
			ErrMsg: "You must provide the name of the controller",
		},
	)
	urlPrefix := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt:  []string{"provide the url prefix for the controller"},
			Default: targetName,
		},
	)
	snakeCaseUrlPrefix := strcase.ToSnake(urlPrefix)
	snakeCaseTargetName := strcase.ToSnake(targetName)
	snakeCaseEndpointTargetName := strcase.ToSnake(targetName + "_endpoint")
	snakeCaseHandlersTargetName := strcase.ToSnake(targetName + "_handler")
	endpointsFile := utils.JoinAndConvertPathToOSFormat(endpointsFolder, fmt.Sprintf("%s.py", snakeCaseEndpointTargetName))
	handlersFile := utils.JoinAndConvertPathToOSFormat(handlersFolder, fmt.Sprintf("%s.py", snakeCaseHandlersTargetName))
	testEndpointsFile := utils.JoinAndConvertPathToOSFormat(testEndpointsFolder, fmt.Sprintf("%s.py", snakeCaseEndpointTargetName))
	testHandlersFile := utils.JoinAndConvertPathToOSFormat(testHandlersFolder, fmt.Sprintf("%s.py", snakeCaseHandlersTargetName))
	utils.CopyFile(templateEndpointFile, endpointsFile)
	utils.CopyFile(templateHandlerFile, handlersFile)
	utils.CopyFile(testTemplateEndpointFile,testEndpointsFile)
	utils.CopyFile(testTemplateHandlerFile,testHandlersFile)

	for _, path := range []string{endpointsFile, handlersFile,testEndpointsFile,testHandlersFile} {
		fileString, err := utils.ReadFile(path)
		if err != nil {
			return
		}
		fileString = strings.ReplaceAll(fileString, "wml_template_url_prefix", snakeCaseUrlPrefix)
		fileString = strings.ReplaceAll(fileString, "wml_template", snakeCaseTargetName)

		utils.OverwriteFile(path, fileString)
	}

	updateAppFile(targetApp,endpointsFile)

}

func updateAppFile(targetApp, endpointsFile string) {
	appFile := utils.JoinAndConvertPathToOSFormat(targetApp, "app.py")

	// Read app.py
	contentBytes, err := os.ReadFile(appFile)
	if err != nil {
			fmt.Printf("Error reading app.py: %v\n", err)
			return
	}

	// Convert content to string
	contentStr := string(contentBytes)

	// Get endpointName from endpointsFile
	endpointFileBase := filepath.Base(endpointsFile)
	endpointName := strings.TrimSuffix(endpointFileBase, filepath.Ext(endpointFileBase))
	// endpointName is something like 'new_endpoint'

	// New import line
	newImportLine := fmt.Sprintf("from endpoints.%s import %s", endpointName, endpointName)

	// New registration line
	newRegistrationLine := fmt.Sprintf("app.register_blueprint(%s)", endpointName)

	// Now, process the content line by line
	lines := strings.Split(contentStr, "\n")

	// We will keep track of where to insert
	lastImportIndex := -1
	lastRegistrationIndex := -1

	// Also, we can check if the import and registration are already present
	importAlreadyPresent := false
	registrationAlreadyPresent := false

	importRegex := regexp.MustCompile(`^from endpoints\..+ import .+$`)
	registrationRegex := regexp.MustCompile(`^app\.register_blueprint\(.+\)`)

	for i, line := range lines {
			lineTrimmed := strings.TrimSpace(line)
			if lineTrimmed == newImportLine {
					importAlreadyPresent = true
			}
			if lineTrimmed == newRegistrationLine {
					registrationAlreadyPresent = true
			}
			if importRegex.MatchString(line) {
					lastImportIndex = i
			}
			if registrationRegex.MatchString(line) {
					lastRegistrationIndex = i
			}
	}

	// Insert the new import if not already present
	if !importAlreadyPresent {
			if lastImportIndex != -1 {
					// Insert after lastImportIndex
					lines = append(lines[:lastImportIndex+1], append([]string{newImportLine}, lines[lastImportIndex+1:]...)...)
			} else {
					// No existing import, add at the top
					lines = append([]string{newImportLine}, lines...)
			}
	}

	// Similarly for registration
	if !registrationAlreadyPresent {
			if lastRegistrationIndex != -1 {
					// Insert after lastRegistrationIndex
					lines = append(lines[:lastRegistrationIndex+1], append([]string{newRegistrationLine}, lines[lastRegistrationIndex+1:]...)...)
			} else {
					// No existing registration, add at the end
					lines = append(lines, newRegistrationLine)
			}
	}

	// Now write back to app.py
	newContentStr := strings.Join(lines, "\n")
	err = os.WriteFile(appFile, []byte(newContentStr), 0644)
	if err != nil {
			fmt.Printf("Error writing to app.py: %v\n", err)
	}
}

