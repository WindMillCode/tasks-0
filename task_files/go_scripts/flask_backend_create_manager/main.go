package main

import (
	"fmt"
	"main/shared"
	"os"
	"regexp"
	"strings"

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
	templateLocation := utils.JoinAndConvertPathToOSFormat(scriptLocation, "template_manager")
	testTemplateLocation := utils.JoinAndConvertPathToOSFormat(scriptLocation, "test_template_manager")

	utils.CDToFlaskApp()
	targetApp, err := os.Getwd()
	if err != nil {
		return
	}
	cliInfo := utils.ShowMenuModel{
		Prompt: "where is the managers folder located",
		Choices: []string{
			utils.JoinAndConvertPathToOSFormat(".", "managers"),
		},
		Other: true,
	}
	managersLocation := utils.JoinAndConvertPathToOSFormat(targetApp, utils.ShowMenu(cliInfo, nil))
	cliInfo = utils.ShowMenuModel{
		Prompt: "where is the test managers folder located",
		Choices: []string{
			utils.JoinAndConvertPathToOSFormat(".", "unit_tests", "FlaskTesting", "managers"),
		},
		Other: true,
	}
	testManagersLocation := utils.JoinAndConvertPathToOSFormat(targetApp, utils.ShowMenu(cliInfo, nil))
	managerName := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt: []string{"What is the name of the manager "},
			ErrMsg: "You must provide the name of a manager",
		},
	)
	// managerName :="feg"
	managerNameString, err := utils.CreateStringObject(managerName, "")
	snakeCaseManagerName := managerNameString.Snakecase(false, "_manager")
	managersLocation = utils.JoinAndConvertPathToOSFormat(managersLocation, snakeCaseManagerName)
	testManagersLocation = utils.JoinAndConvertPathToOSFormat(testManagersLocation, snakeCaseManagerName)
	utils.CopyDir(templateLocation, managersLocation)
	utils.CopyDir(testTemplateLocation, testManagersLocation)
	managerFilePath := utils.JoinAndConvertPathToOSFormat(managersLocation, fmt.Sprintf("%s.py", snakeCaseManagerName))
	testManagerFilePath := utils.JoinAndConvertPathToOSFormat(testManagersLocation, fmt.Sprintf("test_%s.py", snakeCaseManagerName))
	os.Rename(
		utils.JoinAndConvertPathToOSFormat(managersLocation, "template_manager.py"),
		managerFilePath,
	)
	os.Rename(
		utils.JoinAndConvertPathToOSFormat(testManagersLocation, "test_template_manager.py"),
		testManagerFilePath,
	)

	for _, path := range []string{managerFilePath, testManagerFilePath} {
		fileString, err := utils.ReadFile(path)
		if err != nil {
			return
		}
		fileString = strings.ReplaceAll(fileString, "Template", managerNameString.Classify(false, ""))
		fileString = strings.ReplaceAll(fileString, "template", managerNameString.Snakecase(false, ""))
		utils.OverwriteFile(path, fileString)
	}

	updateConfigsFile(targetApp, managerNameString)
}

func updateConfigsFile(targetConfigs string, managerNameString utils.CreateStringObjectType) {
	configsFile := utils.JoinAndConvertPathToOSFormat(targetConfigs, "configs.py")
	contentBytes, err := os.ReadFile(configsFile)
	if err != nil {
			fmt.Printf("Error reading configs.py: %v\n", err)
			return
	}
	contentStr := string(contentBytes)
	contentStr = strings.ReplaceAll(contentStr, "\r", "")

	// Determine the indentation (tabs or spaces)
	indentation := "\t" // Default to tab
	// Try to detect indentation from existing @property methods
	propertyIndentRegex := regexp.MustCompile(`(?m)^(\s*)@property`)
	indentMatch := propertyIndentRegex.FindStringSubmatch(contentStr)
	if len(indentMatch) >= 2 {
			indentation = indentMatch[1]
	} else {
			// If no @property found, detect indentation from def statements
			defIndentRegex := regexp.MustCompile(`(?m)^(\s*)def `)
			indentMatch = defIndentRegex.FindStringSubmatch(contentStr)
			if len(indentMatch) >= 2 {
					indentation = indentMatch[1]
			} else {
					// Default to 4 spaces if indentation cannot be detected
					indentation = "    "
			}
	}
	indentation = strings.ReplaceAll(indentation, "\n", "")


	propertyRegex := regexp.MustCompile(`(?s)(@property\s+def\s+\w+\(self\):.*?return\s+self\._managers\["\w+"\])`)
	matches := propertyRegex.FindAllStringIndex(contentStr, -1)

	// Prepare the new code block with the correct indentation
	newCodeBlock := fmt.Sprintf("\n\n%s@property\n%sdef %s_manager(self):\n%sif not self._managers.get(\"%s_manager\"):\n%sfrom managers.%s_manager.%s_manager import %sManager\n%sself._managers[\"%s_manager\"] = %sManager()\n%sreturn self._managers[\"%s_manager\"]\n",
			indentation,
			indentation,
			managerNameString.Snakecase(false, ""),
			indentation+indentation,
			managerNameString.Snakecase(false, ""),
			indentation+indentation+indentation,
			managerNameString.Snakecase(false, ""),
			managerNameString.Snakecase(false, ""),
			managerNameString.Classify(false, ""),
			indentation+indentation+indentation,
			managerNameString.Snakecase(false, ""),
			managerNameString.Classify(false, ""),
			indentation+indentation,
			managerNameString.Snakecase(false, ""))

	if len(matches) == 0 {
			// No @property methods found, add to class
			classRegex := regexp.MustCompile(`(?s)(class\s+\w+\(.*?\):)`)
			classMatch := classRegex.FindStringIndex(contentStr)
			if len(classMatch) > 0 {
					insertPos := classMatch[1]
					updatedContent := contentStr[:insertPos] + "\n" + indentation + newCodeBlock + contentStr[insertPos:]
					err = os.WriteFile(configsFile, []byte(updatedContent), 0644)
					if err != nil {
							fmt.Printf("Error writing to configs.py: %v\n", err)
					}
			} else {
					fmt.Println("No class definition found in configs.py")
					return
			}
	} else {
			// Insert after the last @property method
			lastMatch := matches[len(matches)-1]
			insertPos := lastMatch[1]
			updatedContent := contentStr[:insertPos] + newCodeBlock + contentStr[insertPos:]
			err = os.WriteFile(configsFile, []byte(updatedContent), 0644)
			if err != nil {
					fmt.Printf("Error writing to configs.py: %v\n", err)
			}
	}
}


