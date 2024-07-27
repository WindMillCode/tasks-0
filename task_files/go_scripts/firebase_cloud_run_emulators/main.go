package main

import (
	"fmt"
	"main/shared"
	"os"

	"github.com/windmillcode/go_cli_scripts/v5/utils"
)

func getDomain(primary, fallback string) string {
	if primary != "" {
		return primary
	}
	return fallback
}

func main() {

	shared.CDToWorkspaceRoot()
	workspaceRoot, err := os.Getwd()
	if err != nil {
		return
	}
	settings, err := utils.GetSettingsJSON(workspaceRoot)
	if err != nil {
		return
	}
	utils.CDToFirebaseApp()
	shared.SetJavaEnvironment()
	// cliInfo := utils.ShowMenuModel{
	// 	Prompt: "Debug Mode",
	// 	Choices:[]string{"TRUE","FALSE"},
	// }
	// debugMode := utils.ShowMenu(cliInfo,nil)
	// if debugMode=="TRUE"{
	// 	os.Setenv("FIREBASE_DEBUG", "true")
	// }
	outputFile := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt: []string{"The output file to see the results"},
			Default: utils.ConvertPathToOSFormat(settings.ExtensionPack.FirebaseCloudRunEmulators.KillPortOutputFile),
		},
	)

	cliInfo := utils.ShowMenuModel{
		Prompt: "kill port dry run",
		Choices:[]string{"TRUE","FALSE"},
		Default: "FALSE",
	}
	dryRun := utils.ShowMenu(cliInfo,nil)
	jobConfig := settings.ExtensionPack.FirebaseCloudRunEmulators
	globalDomain := settings.ExtensionPack.FirebaseCloudRunEmulators.GlobalDomain
	ports := settings.ExtensionPack.Ports

	envVars := map[string]struct {
		Domain string
		Port   int
	}{
		"EMULATOR_HOST":      {getDomain(jobConfig.UIDomain0, globalDomain), ports.FirebaseEmulatorUI0},
		"AUTH_EMULATOR_HOST":      {getDomain(jobConfig.AuthDomain0, globalDomain), ports.FirebaseEmulatorAuth0},
		"STORAGE_EMULATOR_HOST":   {getDomain(jobConfig.StorageDomain0, globalDomain), ports.FirebaseEmulatorStorage0},
		"FIRESTORE_EMULATOR_HOST": {getDomain(jobConfig.FirestoreDomain0, globalDomain), ports.FirebaseEmulatorFirestore0},
		"DATABASE_EMULATOR_HOST":  {getDomain(jobConfig.DatabaseDomain0, globalDomain), ports.FirebaseEmulatorDatabase0},
		"HOSTING_EMULATOR_HOST":   {getDomain(jobConfig.HostingDomain0, globalDomain), ports.FirebaseEmulatorHosting0},
		"FUNCTIONS_EMULATOR_HOST": {getDomain(jobConfig.FunctionsDomain0, globalDomain), ports.FirebaseEmulatorFunctions0},
		"PUBSUB_EMULATOR_HOST":    {getDomain(jobConfig.PubSubDomain0, globalDomain), ports.FirebaseEmulatorPubSub0},
	}

	envMap := make(map[string]string)
	for env, config := range envVars {
		if config.Port != 0 {
			host:= fmt.Sprintf("%s:%s", config.Domain, utils.IntToStr(config.Port))
			envMap[env] = host
			envMap[fmt.Sprintf("FIREBASE_%s", env)] = host
		}
	}
	for key,val := range envMap{
		os.Setenv(key,val)
	}

	firebasePorts := ports.GetFirebasePorts()
	firebasePorts = append(firebasePorts,jobConfig.AdditonalPortsToKill...)
	firebaseInterfacePorts := make([]interface{}, len(firebasePorts))
	for i, port := range firebasePorts {
		firebaseInterfacePorts[i] = port
	}

	options := utils.KillPortsOptions{
		Ports :utils.ConvertToStringArray(firebaseInterfacePorts),
		ProgramName: "firebase",
		OutputFile: outputFile,
		OpenOutputFile: outputFile != "",
		DryRun: dryRun =="TRUE",
	}
	fmt.Println(options.Ports)
	utils.KillPorts(options)

	commandOptions := utils.CommandOptions{
		Command: "npx",
		Args:     []string{"firebase", "emulators:start", "--import=devData", "--export-on-exit"},
		GetOutput:   true,
		PrintOutput: true,
		// PrintOutputOnly: true,
		EnvVars: envMap,
	}
	utils.RunCommandWithOptions(commandOptions)


}
