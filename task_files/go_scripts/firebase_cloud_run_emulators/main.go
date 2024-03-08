package main

import (
	"main/shared"

	"github.com/windmillcode/go_cli_scripts/v4/utils"
)

func main() {

	utils.CDToWorkspaceRoot()
	utils.CDToFirebaseApp()


	packageManager := shared.ChooseNodePackageManager()

	utils.RunCommand(packageManager, []string{"run","cleanup"})
	utils.RunCommand("npx", []string{"firebase", "emulators:start", "--import=devData", "--export-on-exit"})
}
