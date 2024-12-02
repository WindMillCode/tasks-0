package main

import (
	"encoding/json"
	"fmt"
	"main/shared"
	"os"
	"github.com/windmillcode/go_cli_scripts/v6/utils"
)

type FirebaseEmulatorEntry struct {
	Enabled bool `json:"enabled"`
	Host string `json:"host"`
	Port int    `json:"port"`
}
type FirebaseConfig struct {
	Hosting struct {
		Public       string   `json:"public"`
		Ignore       []string `json:"ignore"`
		CleanUrls    bool     `json:"cleanUrls"`
		TrailingSlash bool    `json:"trailingSlash"`
		Rewrites     []struct {
			Source      string `json:"source"`
			Destination string `json:"destination"`
		} `json:"rewrites"`
	} `json:"hosting"`
	Storage struct {
		Rules string `json:"rules"`
	} `json:"storage,omitempty"`
	Firestore struct {
		Rules   string `json:"rules,omitempty"`
		Indexes string `json:"indexes,omitempty"`
	} `json:"firestore,omitempty"`
	Emulators struct {
		Hosting           *FirebaseEmulatorEntry `json:"hosting,omitempty"`
		Auth              *FirebaseEmulatorEntry `json:"auth,omitempty"`
		Storage           *FirebaseEmulatorEntry `json:"storage,omitempty"`
		UI                *FirebaseEmulatorEntry `json:"ui,omitempty"`
		Firestore         *FirebaseEmulatorEntry `json:"firestore,omitempty"`
		Database          *FirebaseEmulatorEntry `json:"database,omitempty"`
		Functions         *FirebaseEmulatorEntry `json:"functions,omitempty"`
		PubSub            *FirebaseEmulatorEntry `json:"pubsub,omitempty"`
		SingleProjectMode bool                  `json:"singleProjectMode"`
	} `json:"emulators"`
}

func updateFirebaseConfig(
	configPath string, envVars map[string]struct {
	Domain string
	Port   int
}) error {
	// Read the existing Firebase config file
	content, err := utils.ReadFile(configPath)
	if err != nil {
		return err
	}

	var config FirebaseConfig
	err = json.Unmarshal([]byte(content), &config)
	if err != nil {
		return err
	}

  // Update the Firebase config with envVars values
  if envVars["HOSTING_EMULATOR_HOST"].Port != 0 {
    config.Emulators.Hosting.Port = envVars["HOSTING_EMULATOR_HOST"].Port
    config.Emulators.Hosting.Host = envVars["HOSTING_EMULATOR_HOST"].Domain
    config.Emulators.Hosting.Enabled = true
  }

  if envVars["AUTH_EMULATOR_HOST"].Port != 0 {
    config.Emulators.Auth.Port = envVars["AUTH_EMULATOR_HOST"].Port
    config.Emulators.Auth.Host = envVars["AUTH_EMULATOR_HOST"].Domain
    config.Emulators.Auth.Enabled = true
  }

  if envVars["STORAGE_EMULATOR_HOST"].Port != 0 {
    config.Emulators.Storage.Port = envVars["STORAGE_EMULATOR_HOST"].Port
    config.Emulators.Storage.Host = envVars["STORAGE_EMULATOR_HOST"].Domain
    config.Emulators.Storage.Enabled = true
  }

  if envVars["EMULATOR_HOST"].Port != 0 {
    config.Emulators.UI.Port = envVars["EMULATOR_HOST"].Port
    config.Emulators.UI.Host = envVars["EMULATOR_HOST"].Domain
    config.Emulators.UI.Enabled = true
  }

  if envVars["FIRESTORE_EMULATOR_HOST"].Port != 0 {
    config.Emulators.Firestore.Port = envVars["FIRESTORE_EMULATOR_HOST"].Port
    config.Emulators.Firestore.Host = envVars["FIRESTORE_EMULATOR_HOST"].Domain
    config.Emulators.Firestore.Enabled = true
  }

  if envVars["DATABASE_EMULATOR_HOST"].Port != 0 {
    config.Emulators.Database.Port = envVars["DATABASE_EMULATOR_HOST"].Port
    config.Emulators.Database.Host = envVars["DATABASE_EMULATOR_HOST"].Domain
    config.Emulators.Database.Enabled = true
  }

  if envVars["FUNCTIONS_EMULATOR_HOST"].Port != 0 {
    config.Emulators.Functions.Port = envVars["FUNCTIONS_EMULATOR_HOST"].Port
    config.Emulators.Functions.Host = envVars["FUNCTIONS_EMULATOR_HOST"].Domain
    config.Emulators.Functions.Enabled = true
  }

  if envVars["PUBSUB_EMULATOR_HOST"].Port != 0 {
    config.Emulators.PubSub.Port = envVars["PUBSUB_EMULATOR_HOST"].Port
    config.Emulators.PubSub.Host = envVars["PUBSUB_EMULATOR_HOST"].Domain
    config.Emulators.PubSub.Enabled = true
  }


	// Marshal the updated config back to JSON
	updatedContent, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	// Overwrite the Firebase config file with updated content
	err = utils.OverwriteFile(configPath, string(updatedContent))
	if err != nil {
		return err
	}

	return nil
}


func getDomain(primary, fallback string) string {
	if primary != "" {
		return primary
	}
	return fallback
}



func main() {

	scriptRoot, err := os.Getwd()
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
	utils.CDToFirebaseApp()
	shared.SetJavaEnvironment(settings)
	firebaseApp, err := os.Getwd()
	if err != nil {
		return
	}
	// cliInfo := utils.ShowMenuModel{
	// 	Prompt: "Debug Mode",
	// 	Choices:[]string{"TRUE","FALSE"},
	// }
	// debugMode := utils.ShowMenu(cliInfo,nil)
	// if debugMode=="TRUE"{
	// 	os.Setenv("FIREBASE_DEBUG", "true")
	// }

	outputFile :=""
	if settings.ExtensionPack.FirebaseCloudRunEmulators.KillPortOutputFileAcceptDefault == false{
		outputFile = utils.GetInputFromStdin(
			utils.GetInputFromStdinStruct{
				Prompt: []string{"The output file to see the results (Default will write kill-port.csv to script directory)"},
				Default: utils.ConvertPathToOSFormat(settings.ExtensionPack.FirebaseCloudRunEmulators.KillPortOutputFile),
			},
		)
	}

	cliInfo := utils.ShowMenuModel{
		Prompt: "Open Output File",
		Choices:[]string{"FALSE","TRUE"},
		Default: "FALSE",
	}
	openOutputFile := utils.ShowMenu(cliInfo,nil)

	cliInfo = utils.ShowMenuModel{
		Prompt: "kill ports",
		Choices:[]string{"TRUE","FALSE"},
		Default: "TRUE",
	}
	killPorts := utils.ShowMenu(cliInfo,nil)
	// dryRun := "FALSE"
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
	updateFirebaseConfig(utils.JoinAndConvertPathToOSFormat(firebaseApp,"firebase.json"),envVars)

	envMap := make(map[string]string)
	for env, config := range envVars {
		if config.Port != 0 {
			host:= fmt.Sprintf("%s:%s", config.Domain, utils.IntToStr(config.Port))
			envMap[env] = host
			envMap[fmt.Sprintf("FIREBASE_%s", env)] = host
		}
	}


	firebasePorts := ports.GetFirebasePorts()
	firebasePorts = append(firebasePorts,jobConfig.AdditonalPortsToKill...)
	firebaseInterfacePorts := make([]interface{}, len(firebasePorts))
	for i, port := range firebasePorts {
		if(port != 0){
			firebaseInterfacePorts[i] = port
		}
	}

	if outputFile == ""{
		outputFile = utils.JoinAndConvertPathToOSFormat(scriptRoot, "kill-port.csv")
	}
	options := utils.KillPortsOptions{
		Ports :utils.ConvertToStringArray(firebaseInterfacePorts),
		// ProgramNames: []string{"node", "java", "System Idle"},
		OutputFile: outputFile,
		OpenOutputFile: openOutputFile == "TRUE",
		DryRun: killPorts =="FALSE",
	}
	fmt.Println(options.Ports)
	utils.KillPorts(options)

	commandOptions := utils.CommandOptions{
		Command: "npx",
		Args:     []string{"firebase", "emulators:start", "--import=devData", "--export-on-exit"},
		GetOutput:   true,
		PrintOutput: true,
		EnvVars: envMap,
	}
	utils.RunCommandWithOptions(commandOptions)


}
