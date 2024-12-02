package main

import (
	"fmt"
	"os"
	"sync"
	"github.com/windmillcode/go_cli_scripts/v6/utils"
	"strings"

	"main/shared"
)
var(
	taskSettings = utils.AngularDeployToFirebaseStruct {}
	myGlobal = map[string]string{}
)
func main() {
	// COMMAND LINE OPTIONS
	//
	// Change directory to appSource and run yarn build:preview
	// Copy the built app to the angular build destination
	// Change directory to fbDest and deploy using Firebase
	// Remove the directories
	// Change back to the workspace root
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
	taskSettings = settings.ExtensionPack.AngularDeployToFirebase

	os.Setenv("NO_UPDATE_NOTIFIER", "true")

	envsInfo := utils.TakeVariableArgs(
		utils.TakeVariableArgsStruct{
			Prompt:  "Provide the environments that should be part of this deployment",
			Default: strings.Join(taskSettings.Environments, " "),
		},
	)
	cliInfo := utils.ShowMenuModel{
		Prompt: "run lint",
		Choices:[]string{"TRUE", "FALSE"},
		Default: func() string {
			if taskSettings.RunLint {
				return "TRUE"
			}
			return "FALSE"
		}(),
	}
	runLint := utils.ShowMenu(cliInfo,nil)
	cliInfo = utils.ShowMenuModel{
		Prompt: "run ssg script",
		Choices:[]string{"TRUE","FALSE"},
		Default: func() string {
			if taskSettings.RunSSGScript {
				return "TRUE"
			}
			return "FALSE"
		}(),
	}
	runSSGScript := utils.ShowMenu(cliInfo,nil)
	cliInfo = utils.ShowMenuModel{
		Prompt: "remove build directories",
		Choices:[]string{"TRUE","FALSE"},
		Default: func() string {
			if taskSettings.RemoveBuildDirectories {
				return "TRUE"
			}
			return "FALSE"
		}(),
	}
	removeBuildDir := utils.ShowMenu(cliInfo,nil)
	cliInfo = utils.ShowMenuModel{
		Prompt: "deploy to sentry",
		Choices:[]string{"TRUE","FALSE"},
		Default: func() string {
			if taskSettings.DeployToSentry {
				return "TRUE"
			}
			return "FALSE"
		}(),
	}
	myGlobal["deployToSentry"]  = utils.ShowMenu(cliInfo,nil)
	cliInfo = utils.ShowMenuModel{
		Prompt: "deploy to firebase",
		Choices:[]string{"TRUE","FALSE"},
		Default: func() string {
			if taskSettings.DeployToFirebase {
				return "TRUE"
			}
			return "FALSE"
		}(),
	}
	myGlobal["deployToFirebase"]  = utils.ShowMenu(cliInfo,nil)


	envs := envsInfo.InputArray
	var wg sync.WaitGroup
	for _, v := range envs {
		wg.Add(1)
		go func(val string) {
			defer wg.Done()
			buildAndDeploy(val,runLint,runSSGScript,removeBuildDir)
		}(v)
	}
	wg.Wait()

}

func buildAndDeploy(env,runLint,runSSGScript,removeBuildDir string) {


	// CLI ARGS
	targetEnviromnent, _ := utils.CreateStringObject(env, "")
	os.Setenv("SENTRY_ORG", taskSettings.SentryOrg)
	os.Setenv("SENTRY_PROJECT", taskSettings.SentryProject)
	os.Setenv("SENTRY_AUTH_TOKEN", taskSettings.SentryAuthToken)
	firebaseProjectId := taskSettings.FirebaseProjectId
	//
	if targetEnviromnent.Orig == "preview" {
		firebaseProjectId += "-preview"
	}

	workSpaceRoot, _ := os.Getwd()

	appSource := utils.JoinAndConvertPathToOSFormat(workSpaceRoot, "apps", "frontend", "AngularApp")
	appDistSource := utils.JoinAndConvertPathToOSFormat(workSpaceRoot, "apps", "frontend", "AngularApp", "dist", fmt.Sprintf("angular-app-%s", targetEnviromnent.Orig), "browser")
	angularBuildDestination := utils.JoinAndConvertPathToOSFormat(workSpaceRoot, "apps", "cloud", "FirebaseApp", "dist", fmt.Sprintf("angular-app-%s", targetEnviromnent.Orig))
	fbDest := utils.JoinAndConvertPathToOSFormat(workSpaceRoot, "apps", "cloud", "FirebaseApp")
	if runLint == "TRUE" {
		lintOptions := utils.CommandOptions{
			Command:              "npm",
			Args:                 []string{"run", "lint"},
			TargetDir:            appSource,
			GetOutput:       false,
			PrintOutputOnly: true,
		}
		utils.RunCommandWithOptions(lintOptions)
	}


	buildOptions := utils.CommandOptions{
		Command:           "npm",
		Args:              []string{"run", fmt.Sprintf("build:%s", targetEnviromnent.Orig)},
		TargetDir:         appSource,
		GetOutput:       false,
		PrintOutputOnly: true,
	}
	utils.RunCommandWithOptions(buildOptions)

	if runSSGScript == "TRUE" {
		ssgOptions := utils.CommandOptions{
			Command:         "npm",
			Args:            []string{"run", fmt.Sprintf("ssg:%s", targetEnviromnent.Orig)},
			TargetDir:       appSource,
			PrintOutputOnly: true,
		}
		utils.RunCommandWithOptions(ssgOptions)
	}

	if err := os.RemoveAll(angularBuildDestination); err != nil {
		fmt.Println("Error removing old build destination:", err)
		return
	}

	if err := utils.CopyDir(appDistSource, angularBuildDestination); err != nil {
		fmt.Println("Error copying from the dist source in the AngularApp to the destination in the FirebaseApp", err)
		return
	}

	if myGlobal["deployToSentry"] == "TRUE" {
		sourceMapInjectOptions := utils.CommandOptions{
			Command: "sentry-cli",
			Args:    []string{"sourcemaps", "inject", appDistSource},
			GetOutput:       false,
			PrintOutputOnly: true,
		}
		utils.RunCommandWithOptions(sourceMapInjectOptions)

		sourceMapUploadOptions := utils.CommandOptions{
			Command: "sentry-cli",
			Args:    []string{"sourcemaps", "upload", appDistSource, "--", "--env", fmt.Sprintf("Angular_%s", targetEnviromnent.Uppercase(false, ""))},
			GetOutput:       false,
			PrintOutputOnly: true,
		}
		utils.RunCommandWithOptions(sourceMapUploadOptions)
	}


	if myGlobal["deployToFirebase"] == "TRUE" {
		deployOptions := utils.CommandOptions{
			Command:         "npx",
			Args:            []string{"firebase", "deploy", "--only", "hosting", "--config", fmt.Sprintf("angular.firebase.%s.json", targetEnviromnent.Orig), "--project", firebaseProjectId},
			TargetDir:       fbDest,
			PrintOutputOnly: true,
		}
		utils.RunCommandWithOptions(deployOptions)
	}


	if err := os.RemoveAll(angularBuildDestination); err != nil {
		fmt.Println("Error removing firebase build destination:", err)
	}

	if removeBuildDir == "TRUE" {
		if err := os.RemoveAll(appDistSource); err != nil {
			fmt.Println("Error removing angular app dist source:", err)
		}
	}

	if err := os.Chdir(workSpaceRoot); err != nil {
		fmt.Println("Error changing back to workspace root:", err)
	}
}
