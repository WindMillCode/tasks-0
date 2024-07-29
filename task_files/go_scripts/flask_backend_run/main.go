package main

import (
	"fmt"
	"main/shared"
	"os"

	"github.com/windmillcode/go_cli_scripts/v5/utils"
)

func main() {

	scriptRoot, err := os.Getwd()
	if err != nil {
		return
	}
	shared.CDToWorkspaceRoot()
	workspaceRoot, err := os.Getwd()
	if err != nil {
		return
	}
	settings, err := utils.GetSettingsJSON(workspaceRoot)
	if err != nil {
		return
	}

	var appPort string
	settingsAppPort := utils.IntToStr(settings.ExtensionPack.Ports.FlaskRun0)
	if settingsAppPort != "0" {
		appPort = settingsAppPort
	} else {
		appPort = utils.IntToStr(5000)
	}
	os.Setenv("FLASK_BACKEND_PORT",appPort)

	utils.CDToLocation(scriptRoot)
	flaskAppFolder, err := shared.SetupEnvironmentToRunFlaskApp("dev")
	if err != nil {
		fmt.Println("Error during setup:", err)
		return
	}

	// Remaining part of the original script that uses flaskAppFolder
	for {
		utils.CDToLocation(flaskAppFolder)
		runOptions := utils.CommandOptions{
			Command:         "python",
			Args:            []string{"app.py", "--reloader_type", "stat"},
			GetOutput:       false,
			PrintOutputOnly: true,
		}
		utils.RunCommandWithOptions(runOptions)
	}
}
