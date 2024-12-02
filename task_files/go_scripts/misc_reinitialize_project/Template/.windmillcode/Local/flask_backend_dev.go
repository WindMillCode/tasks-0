package main

import (
	"fmt"
	"os"

	"github.com/windmillcode/go_cli_scripts/v6/utils"
)

func main() {
	// scriptLocation := os.Args[1]
	workspaceRoot := os.Args[2]
	settings, err := utils.GetSettingsJSON(workspaceRoot)
	if err != nil {
		return
	}
	dockerCommand := utils.CommandOptions{
		Command:         "docker",
		Args:            []string{"start", settings.ExtensionPack.SQLDockerContainerName},
		PrintOutput:     false,
		PrintOutputOnly: false,
	}
	utils.RunCommandWithOptions(dockerCommand)


	envVars := fmt.Sprintf(`
	<ENVVARS>
	YOUR_ENV_VAR_KEY=YOUR_ENV_VAR_VALUE
	</ENVVARS>
	`,
	)
	fmt.Println(envVars)

}
