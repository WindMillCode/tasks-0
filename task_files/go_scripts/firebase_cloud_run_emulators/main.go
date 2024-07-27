package main

import (
	"main/shared"
	// "os"
	"github.com/windmillcode/go_cli_scripts/v5/utils"
)

func main() {

	shared.CDToWorkspaceRoot()
	// workspaceRoot, err := os.Getwd()
	// if err != nil {
	// 	return
	// }
	// settings, err := utils.GetSettingsJSON(workspaceRoot)
	// if err != nil {
	// 	return
	// }
	utils.CDToFirebaseApp()
	packageManager := shared.ChooseNodePackageManager()
	shared.SetJavaEnvironment()
	// cliInfo := utils.ShowMenuModel{
	// 	Prompt: "Debug Mode",
	// 	Choices:[]string{"TRUE","FALSE"},
	// }
	// debugMode := utils.ShowMenu(cliInfo,nil)
	// if debugMode=="TRUE"{
	// 	os.Setenv("FIREBASE_DEBUG", "true")
	// }



	// os.Setenv("FIREBASE_AUTH_EMULATOR_PORT", shared.IntToStr(settings.ExtensionPack.Ports.FirebaseEmulatorAuth0))
	// os.Setenv("FIRESTORE_EMULATOR_PORT", shared.IntToStr(settings.ExtensionPack.Ports.FirebaseEmulatorAuth0))
	// os.Setenv("FIREBASE_DATABASE_EMULATOR_PORT", shared.IntToStr(settings.ExtensionPack.Ports.FirebaseEmulatorAuth0))
	// os.Setenv("FIREBASE_STORAGE_EMULATOR_PORT", shared.IntToStr(settings.ExtensionPack.Ports.FirebaseEmulatorStorage0))
	// os.Setenv("FIREBASE_HOSTING_EMULATOR_PORT", shared.IntToStr(settings.ExtensionPack.Ports.FirebaseEmulatorAuth0))
	// os.Setenv("FIREBASE_FUNCTIONS_EMULATOR_PORT", shared.IntToStr(settings.ExtensionPack.Ports.FirebaseEmulatorAuth0))
	// os.Setenv("FIREBASE_PUBSUB_EMULATOR_PORT", shared.IntToStr(settings.ExtensionPack.Ports.FirebaseEmulatorAuth0))
	utils.RunCommand(packageManager, []string{"run", "cleanup"})
	utils.RunCommand("npx", []string{"firebase", "emulators:start", "--import=devData", "--export-on-exit"})
}
