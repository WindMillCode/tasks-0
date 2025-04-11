package main

import (
	"fmt"
	"os"
	"runtime"
	"main/shared"
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

	Choices := make([]string, len(settings.ExtensionPack.MiscTranslateJson.AppLocations))
	for i, x := range settings.ExtensionPack.MiscTranslateJson.AppLocations {
		Choices[i] = utils.JoinAndConvertPathToOSFormat(workspaceRoot, x)
	}

	cliInfo := utils.ShowMenuModel{
		Other:   true,
		Prompt:  "choose the location of the i18n folder",
		Choices: Choices,
	}

	i18nLocation := utils.ShowMenu(cliInfo, nil)
	openAIAPIKey := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt:  []string{"provide the open ai api key"},
			ErrMsg:  "an open ai key is required to translate the app",
			Default: settings.ExtensionPack.OpenAIAPIKey0,
		},
	)
	if settings.ExtensionPack.OpenAIAPIBase0 == "" {
		settings.ExtensionPack.OpenAIAPIBase0 = "https://api.openai.com/v1"
	}
	openAIBase := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt:  []string{"Provide the open ai url"},
			Default: settings.ExtensionPack.OpenAIAPIBase0,
		},
	)
	langCodes := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt:  []string{" Provide a list of lang codes to run \n translation script. \n Provide them in comma separated format according to the options below. \n Example: 'zh, es, hi, bn' \n It's best to do 4 at a time. \n Options: zh, es, hi, uk, ar, bn, ms, fr, de, sw, am"},
			ErrMsg:  "Lang codes are required",
			Default: settings.ExtensionPack.LangCodes0,
		},
	)

	cliInfo = utils.ShowMenuModel{
		Prompt:  "Randomly remove key value json to better translate unmodifed key sections",
		Choices: []string{"FALSE","TRUE"},
		Default: "FALSE",
	}
	random_remove_json := utils.ShowMenu(cliInfo, nil)

	os.Setenv("OPENAI_API_BASE", openAIBase)
	os.Setenv("OPENAI_API_KEY_0", openAIAPIKey)
	utils.CDToLocation(utils.JoinAndConvertPathToOSFormat(workspaceRoot, ".windmillcode", "go_scripts","misc_translate_json", "i18n_script_via_ai"))
	// pathSeparator := string(filepath.Separator)
	shared.SetPythonEnvironment(settings.ExtensionPack.PythonVersion0)
	i18nScriptLocation, _ := os.Getwd()
	switch os := runtime.GOOS; os {
	case "windows":
		sitePackages := utils.JoinAndConvertPathToOSFormat(i18nScriptLocation, "site-packages", "windows")
		// sitePackages = strings.Join([]string{".",sitePackages},pathSeparator)
		if !utils.FolderExists(sitePackages)  {

			utils.RunCommand("pip", []string{"install", "-r", "requirements.txt", "--target", sitePackages})
		}
	case "linux", "darwin":
		sitePackages := utils.JoinAndConvertPathToOSFormat(i18nScriptLocation, "site-packages", "linux")
		// sitePackages = strings.Join([]string{".",sitePackages},pathSeparator)
		if !utils.FolderExists(sitePackages)  {
			utils.RunCommand("pip", []string{"install", "-r", "requirements.txt", "--target", sitePackages})
		}

	default:
		fmt.Println("Unknown Operating System:", os)
	}
	myCommand := []string{"index.py", "-c", langCodes, "--location", i18nLocation, "--source-file", "en.json"}
	if random_remove_json == "TRUE" {
		myCommand = append(myCommand, []string{"-r", "TRUE"}...)
	}

	utils.RunCommand("python", myCommand)
}
