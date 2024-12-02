package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/windmillcode/go_cli_scripts/v6/utils"
)

type deploymentConfig struct {
	env           string
	deployRules   string
	deployIndexes string
}

func main() {
	os.Setenv("NO_UPDATE_NOTIFIER", "true")
	envs := []string{
		"preview",
		"prod",
	}

	var deployments []deploymentConfig

	for _, v := range envs {
		fmt.Println("Deploying firestore for", v)

		cliInfo := utils.ShowMenuModel{
			Prompt:  "deploy rules?",
			Choices: []string{"TRUE", "FALSE"},
		}
		deployRules := utils.ShowMenu(cliInfo, nil)

		cliInfo = utils.ShowMenuModel{
			Prompt:  "deploy indexes?",
			Choices: []string{"TRUE", "FALSE"},
		}
		deployIndexes := utils.ShowMenu(cliInfo, nil)

		deployments = append(deployments, deploymentConfig{
			env:           v,
			deployRules:   deployRules,
			deployIndexes: deployIndexes,
		})
	}

	var wg sync.WaitGroup
	for _, config := range deployments {
		wg.Add(1)
		go func(cfg deploymentConfig) {
			defer wg.Done()
			deployFirestore(cfg.env, cfg.deployRules, cfg.deployIndexes)
		}(config)
	}
	wg.Wait()
}
func deployFirestore(env, deployRules,deployIndexes string,) {

	// CLI ARGS
	targetEnviromnent, _ := utils.CreateStringObject(env, "")
	firebaseProjectId := "[PROJECT_NAME]"
	//
	if targetEnviromnent.Orig == "preview" {
		firebaseProjectId += "-preview"
	}

	currentFolder, _ := os.Getwd()
	workSpaceRoot := utils.JoinAndConvertPathToOSFormat(currentFolder, "..", "..")


	fbDest := utils.JoinAndConvertPathToOSFormat(workSpaceRoot, "apps", "cloud", "FirebaseApp")



	firebaseTarget :=""
	if deployRules == "TRUE" && deployIndexes == "TRUE" {
		firebaseTarget = "firestore"
	} else if deployRules == "TRUE" {
		firebaseTarget = "firestore:rules"
	} else if deployIndexes == "TRUE" {
		firebaseTarget = "firestore:indexes"
	}


	if firebaseTarget != "" {
		deployOptions := utils.CommandOptions{
			Command:         "npx",
			Args:            []string{"firebase", "deploy", "--only", firebaseTarget,"--project", firebaseProjectId},
			TargetDir:       fbDest,
			PrintOutputOnly: true,
		}
		utils.RunCommandWithOptions(deployOptions)
	}


	if err := os.Chdir(workSpaceRoot); err != nil {
		fmt.Println("Error changing back to workspace root:", err)
	}
}
