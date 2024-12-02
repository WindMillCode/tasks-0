package main

import (
	"fmt"
	"main/shared"

	"github.com/windmillcode/go_cli_scripts/v6/utils"
)

func main() {

	shouldReturn := shared.SetFlaskAppPort()
	if shouldReturn {
		return
	}
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
			Args:            []string{"reloader.py"},
			GetOutput:       false,
			PrintOutputOnly: true,
		}
		utils.RunCommandWithOptions(runOptions)
	}
}


