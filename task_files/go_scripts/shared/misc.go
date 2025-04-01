package shared

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
	"sync"
	"time"
	"runtime"

	"github.com/windmillcode/go_cli_scripts/v6/utils"
)


func CDToWorkspaceRoot() {
	utils.CDToLocation(filepath.Join("..", "..", ".."))
}

func RebuildExecutables(proceed string, tasksJSON utils.VSCodeTasksTasksJSON, goScriptsDestDirPath string, beforeActionPredicate func()) {
	var rebuild string
	var cliInfo utils.ShowMenuModel
	if proceed == "TRUE" {
		rebuild = "TRUE"
	} else {
		cliInfo = utils.ShowMenuModel{
			Prompt:  "Do you want to rebuild the go programs into binary exectuables ",
			Choices: []string{"TRUE", "FALSE"},
		}
		rebuild = utils.ShowMenu(cliInfo, nil)
	}

	if rebuild == "TRUE" {

		cliInfo = utils.ShowMenuModel{
			Prompt:  "Build one by one or all at once",
			Choices: []string{"ALL_AT_ONCE", "ONE_BY_ONE"},
		}
		inSync := utils.ShowMenu(cliInfo, nil)
		beforeActionPredicate()

		if inSync == "ALL_AT_ONCE" {
			var wg sync.WaitGroup
			fmt.Print(len(tasksJSON.Tasks))
			for _, task := range tasksJSON.Tasks {
				wg.Add(1)

				pattern0 := ":"
				regex0 := regexp.MustCompile(pattern0)
				programLocation0 := regex0.Split(task.Label, -1)
				pattern1 := " "
				regex1 := regexp.MustCompile(pattern1)
				programLocation1 := regex1.Split(strings.Join(programLocation0, ""), -1)
				programLocation2 := strings.Join(programLocation1, "_")
				absProgramLocation := utils.JoinAndConvertPathToOSFormat(goScriptsDestDirPath, programLocation2)
				go func() {
					defer wg.Done()
					BuildGoCLIProgram(absProgramLocation)
				}()
			}
			wg.Wait()
		} else {
			fmt.Print(len(tasksJSON.Tasks))
			for _, task := range tasksJSON.Tasks {

				pattern0 := ":"
				regex0 := regexp.MustCompile(pattern0)
				programLocation0 := regex0.Split(task.Label, -1)
				pattern1 := " "
				regex1 := regexp.MustCompile(pattern1)
				programLocation1 := regex1.Split(strings.Join(programLocation0, ""), -1)
				programLocation2 := strings.Join(programLocation1, "_")
				absProgramLocation := utils.JoinAndConvertPathToOSFormat(goScriptsDestDirPath, programLocation2)
				BuildGoCLIProgram(absProgramLocation)

			}
		}

	} else {
		beforeActionPredicate()
	}
}

func BuildGoCLIProgram(programLocation string) {

	fmt.Printf("%s \n", programLocation)
	utils.RunCommandInSpecificDirectory("go", []string{"build", "main.go"}, programLocation)
	fmt.Printf("Finished building %s \n", programLocation)

}

// func StartDockerDesktop() (string, error) {
//     commandOptions := utils.CommandOptions{
//       Command:     "dockerd",
//       Args:        []string{},
//       GetOutput:   true,
//       PrintOutput: false,
// 			IsElevated: true,
//     }
//     result, err := utils.RunCommandWithOptions(commandOptions)
//     if err != nil {
//       return "", fmt.Errorf("failed to start Docker Desktop on macOS: %v", err)
//     }
//     return result, nil
// }

func waitForDockerEngine() error {
	const maxAttempts = 30
	const delay = 2 * time.Second

	for i := 0; i < maxAttempts; i++ {
		cmdOptions := utils.CommandOptions{
			Command:     "docker",
			Args:        []string{"info"},
			GetOutput:   false,
			PrintOutput: false,
		}
		_, err := utils.RunCommandWithOptions(cmdOptions)
		if err == nil {
			return nil
		}
		time.Sleep(delay)
	}
	return fmt.Errorf("docker engine did not become ready after %d attempts", maxAttempts)
}

func StartDockerDesktop() (string, error) {
	switch runtime.GOOS {
	case "windows":
		commandOptions := utils.CommandOptions{
			Command:     "powershell.exe",
			Args:        []string{"-Command", "Start-Process 'Docker Desktop.exe'"},
			GetOutput:   false,
			TargetDir: "",
			PrintOutput: false,
		}
		result, err := utils.RunCommandWithOptions(commandOptions)
		if err != nil {
			return "", fmt.Errorf("failed to start Docker Desktop on Windows: %v", err)
		}
		if err := waitForDockerEngine(); err != nil {
			return result, fmt.Errorf("Docker Desktop started but engine not ready: %v", err)
		}
		return result, nil

	case "darwin": // macOS
		commandOptions := utils.CommandOptions{
			Command:     "open",
			Args:        []string{"/Applications/Docker.app"},
			GetOutput:   false,
			TargetDir: "",
			PrintOutput: false,
		}
		result, err := utils.RunCommandWithOptions(commandOptions)
		if err != nil {
			return "", fmt.Errorf("failed to start Docker Desktop on macOS: %v", err)
		}
		if err := waitForDockerEngine(); err != nil {
			return result, fmt.Errorf("Docker Desktop started but engine not ready: %v", err)
		}
		return result, nil

	case "linux":
		commandOptions := utils.CommandOptions{
			Command:     "systemctl",
			Args:        []string{"--user", "start", "docker-desktop"},
			GetOutput:   false,
			TargetDir: "",
			PrintOutput: false,
		}
		result, err := utils.RunCommandWithOptions(commandOptions)
		if err != nil {
			return "", fmt.Errorf("failed to start Docker Desktop on Linux: %v", err)
		}
		if err := waitForDockerEngine(); err != nil {
			return result, fmt.Errorf("Docker Desktop started but engine not ready: %v", err)
		}
		return result, nil

	default:
		return "", fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}
}


func SetFlaskAppPort() bool {
	scriptRoot, err := os.Getwd()
	if err != nil {
		return true
	}
	CDToWorkspaceRoot()
	workspaceRoot, err := os.Getwd()
	settings, err := utils.GetSettingsJSON(workspaceRoot)
	if err != nil {
		return true
	}
	utils.SetGlobalVars(
		utils.SetGlobalVarsOptions{
			NonInteractive :settings.ExtensionPack.ProcessIfDefaultIsPresent,
		},
	)

	var appPort string
	settingsAppPort := utils.IntToStr(settings.ExtensionPack.Ports.FlaskRun0)
	if settingsAppPort != "0" {
		appPort = settingsAppPort
	} else {
		appPort = utils.IntToStr(5000)
	}
	os.Setenv("FLASK_BACKEND_PORT", appPort)

	cliInfo := utils.ShowMenuModel{
		Prompt: "Kill the app running on the port before starting? (If you know what you are doing choose TRUE)",
		Choices:[]string{"FALSE","TRUE"},
		Default: "FALSE",
	}
	killPort := utils.ShowMenu(cliInfo,nil)
	if killPort =="TRUE"{
		utils.KillPorts(utils.KillPortsOptions  {
			Ports:          []string{"appPort"},
		})
	}


	utils.CDToLocation(scriptRoot)
	return false
}

func SetupEnvironmentToRunFlaskApp(env string) (string, error) {
	// Change to workspace root and get the current directory
	CDToWorkspaceRoot()
	workspaceFolder, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("error getting current directory: %w", err)
	}

	// Get settings from the workspace folder
	settings, err := utils.GetSettingsJSON(workspaceFolder)
	if err != nil {
		return "", fmt.Errorf("error getting settings JSON: %w", err)
	}

	// Change to Flask app directory and get the current directory
	utils.CDToFlaskApp()
	flaskAppFolder, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("error changing to Flask app directory: %w", err)
	}

	// Determine the helper script based on whether it's running in Docker
	helperScript := settings.ExtensionPack.FlaskBackendDevHelperScript
	if env == "test" {
		helperScript = settings.ExtensionPack.FlaskBackendTestHelperScript
	}
	if utils.IsRunningInDocker() {
		helperScript = strings.Replace(helperScript, "dev", "docker_dev", 1)
	}

	// Prompt for the location of the environment variables file
	cliInfo := utils.ShowMenuModel{
		Prompt:  "Where are the env vars located",
		Choices: []string{utils.JoinAndConvertPathToOSFormat(workspaceFolder, helperScript)},
		Other:   true,
	}
	envVarsFile := utils.ShowMenu(cliInfo, nil)

	// Get the Python version input
	SetPythonEnvironment(settings.ExtensionPack.PythonVersion0)

	// Execute the helper script to set environment variables
	utils.CDToLocation(workspaceFolder)
	envVarCommandOptions := utils.CommandOptions{
		Command:     "go",
		Args:        []string{"run", envVarsFile, filepath.Dir(utils.JoinAndConvertPathToOSFormat(envVarsFile)), workspaceFolder},
		GetOutput:   true,
		TargetDir:   filepath.Dir(utils.JoinAndConvertPathToOSFormat(envVarsFile)),
		PrintOutput: false,
	}
	envVars, err := utils.RunCommandWithOptions(envVarCommandOptions)
	if err != nil {
		return "", fmt.Errorf("error running command for env vars: %w", err)
	}

	// Parse and set the environment variables
	re := regexp.MustCompile(`(?s)<ENVVARS>(.*?)<\/ENVVARS>`)
	matches := re.FindStringSubmatch(envVars)
	envVarsArray := strings.Split(matches[1], ",")
	for _, x := range envVarsArray {
		indexOfEqual := strings.Index(x, "=")
		if indexOfEqual == -1 {
			continue
		}
		key := strings.TrimSpace(x[:indexOfEqual])
		value := strings.TrimSpace(x[indexOfEqual+1:])
		os.Setenv(key, value)
	}

	return flaskAppFolder, nil
}

func SetPythonEnvironment(myDefault string){

	pythonVersion := utils.GetInputFromStdin(utils.GetInputFromStdinStruct{
		Prompt:  []string{"Provide a Python version for pyenv to use type 'Skip' to stay with the existing version"},
		Default: myDefault,
	})
	if  !slices.Contains([]string{"", "Skip"}, pythonVersion){
		_, err := exec.LookPath("pyenv")
		if err == nil{
			utils.RunCommand("pyenv", []string{"global", pythonVersion})
		} else{
			fmt.Println("pyenv seems not to be install in your system please install or set manually")
		}
	}

}
func SetNodeJSEnvironment(myDefault string){
	nodeJSVersion := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt: []string{"provide the nodejs version is type 'Skip' to stay with the existing version"},
			Default: myDefault,
		},
	)
	_, err := exec.LookPath("nvm")
	if err == nil{
		if(nodeJSVersion != "Skip"){
			utils.RunCommand("nvm",[]string{"use",nodeJSVersion})
		}
	} else{
		fmt.Println("nvm seems not to be installed in your system please install or set manually")
	}
}

func SetJavaEnvironment(settings utils.VSCodeSettings){

	_, err := exec.LookPath("jvms")
	if err == nil{

	} else{
		fmt.Println("jvms seems not to be installed in your system please install or set manually")
		return;
	}
	output, _ := utils.RunCommandWithOptions(utils.CommandOptions{
		Command:   "jvms",
		Args:      []string{"ls"},
		GetOutput: true,
	})
	// Split the output into lines
	lines := strings.Split(output, "\n")
	var jdks []string

	// Iterate over the lines and extract JDK names
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if len(trimmedLine) > 0 && !strings.HasPrefix(trimmedLine, "Installed jdk") {
			// Remove the leading index and asterisk if present
			parts := strings.Split(trimmedLine, ")")
			if len(parts) > 1 {
				jdk := strings.TrimSpace(parts[1])
				// Remove the leading asterisk if present
				jdk = strings.TrimPrefix(jdk, "*")
				jdks = append(jdks, strings.TrimSpace(jdk))
			}
		}
	}
	cliInfo := utils.ShowMenuModel{
		Prompt: "select the java version to use",
		Choices:append([]string{"Skip"}, jdks...),
		Default: settings.ExtensionPack.JavaVersion0,
	}
	javaVersion := utils.ShowMenu(cliInfo,nil)
	if(javaVersion != "Skip"){
		utils.RunElevatedCommand("jvms",[]string{"switch",javaVersion})
	}

}




func ChooseNodePackageManager() string {
	cliInfo := utils.ShowMenuModel{
		Prompt:  "choose the package manager",
		Choices: []string{"npm", "yarn", "pnpm"},
		Default: "npm",
	}
	exectuable := utils.ShowMenu(cliInfo, nil)
	return exectuable
}


