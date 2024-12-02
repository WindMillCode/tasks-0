package main

import (
	"fmt"
	"os"

	"github.com/windmillcode/go_cli_scripts/v6/utils"
)

func main() {

	extensionFolder, err := os.Getwd()
	if err != nil {
		return
	}
	utils.CDToLocation(utils.JoinAndConvertPathToOSFormat("..", ".."))
	workspaceRoot, err := os.Getwd()
	if err != nil {
		return
	}
	settings, err := utils.GetSettingsJSON(workspaceRoot)
	if err != nil {
		return
	}
	utils.CDToLocation(extensionFolder)
	cliInfo := utils.ShowMenuModel{
		Prompt:  "choose the file type",
		Choices: []string{"SQL", "BINARY"},
	}
	dumpFileType := utils.ShowMenu(cliInfo, nil)
	cliInfo = utils.ShowMenuModel{
		Prompt:  "schema only",
		Choices: []string{"TRUE", "FALSE"},
		Default: "FALSE",
	}
	schemaOnly := utils.ShowMenu(cliInfo, nil)

	backupFile := "backup_file.dump"
	if dumpFileType == "SQL" {
		backupFile = "backup_file.sql"
	}
	utils.RunCommand("docker", []string{"cp", "dump_db.sh", settings.ExtensionPack.SQLDockerContainerName + ":/usr/bin/"})
	utils.RunCommand("docker", []string{"exec", "--workdir", "/",
		settings.ExtensionPack.SQLDockerContainerName,
		"chmod", "+x", "/usr/bin/dump_db.sh"},
	)
	utils.RunCommand("docker", []string{"exec", "--workdir", "/",
		settings.ExtensionPack.SQLDockerContainerName,
		"bash", "-c", fmt.Sprintf("/usr/bin/dump_db.sh %s %s %s", dumpFileType,
			schemaOnly,
			backupFile),
	},
	)
	utils.RunCommand("docker", []string{"cp", settings.ExtensionPack.SQLDockerContainerName + ":/" + backupFile, "."})

}
