package main

import (
	"fmt"
	"github.com/windmillcode/go_cli_scripts/v4/utils"
	"main/shared"
)

func main() {
	flaskAppFolder, err := shared.SetupEnvironmentToRunFlaskApp()
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


