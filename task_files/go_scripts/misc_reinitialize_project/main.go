package main

import (
	"encoding/json"
	"fmt"
	"log"
	"main/shared"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/windmillcode/go_cli_scripts/v6/utils"
)

var (
	PROJECT_NAME                      = utils.CreateStringObjectType{}
	ORGANIZATION_NAME                 = utils.CreateStringObjectType{}
	WEB_DRIVER_PATH                                      = ""
	VCS_PRIVATE_KEY                                      = ""
	WEB_SEO_DESCRIPTION                                  = ""
	WEB_SEO_KEYWORDS                                     = ""
	PROXY_URLS_0                                         = ""
	FLUTTER_ANDRIOID_GOOGLE_ADS_ID                       = ""
	FLUTTER_ANDROID_GOOGLE_APPLICATION_ID                = ""
	FLUTTER_IOS_GOOGLE_ADS_ID                            = ""
	FLUTTER_ANDROID_KEY_PROPERTIES_STORE_FILE            = ""
	FLUTTER_ANDROID_KEY_PROPERTIES_KEY_PASSWORD          = ""
	FLUTTER_ANDROID_KEY_PROPERTIES_STORE_PASSWORD        = ""
	FLUTTER_ANDROID_PREVIEW_KEY_PROPERTIES_KEY_PASSWORD  = ""
	FLUTTER_ANDROID_PREVIEW_KEY_PROPERTIES_STORE_FILE    = ""
	FLUTTER_IOS_GOOGLE_APPLICATION_ID = ""
	FLUTTER_IOS_FACEBOOK_APP_ID = ""
	FLUTTER_IOS_FACEBOOK_CLIENT_TOKEN = ""
	FLUTTER_IOS_FACEBOOK_CUSTOM_URL_SCHEME = ""
	FLUTTER_IOS_GOOGLE_OAUTH_URL_SCHEMES = []string{}
	CHROME_DRIVER_PATH = ""
	FIREFOX_DRIVER_PATH = ""
	OPERA_DRIVER_PATH = ""
	EDGE_DRIVER_PATH = ""
	windmillcodeSettings = utils.VSCodeSettings{}
)

type LaunchJSON struct {
	Version        string `json:"version"`
	Configurations []struct {
		Name                   string            `json:"name"`
		Type                   string            `json:"type"`
		Request                string            `json:"request"`
		Mode                   string            `json:"mode,omitempty"`
		Console                string            `json:"console,omitempty"`
		Program                string            `json:"program"`
		Args                   []string          `json:"args,omitempty"`
		RuntimeArgs            []string          `json:"runtimeArgs,omitempty"`
		Cwd                    string            `json:"cwd,omitempty"`
		InternalConsoleOptions string            `json:"internalConsoleOptions,omitempty"`
		Env                    map[string]string `json:"env,omitempty"`
		VmAdditionalArgs       []string          `json:"vmAdditionalArgs,omitempty"`
		ToolArgs               []string          `json:"toolArgs,omitempty"`
		DeviceID               string            `json:"deviceId,omitempty"`
		FlutterMode            string            `json:"flutterMode,omitempty"`
		CodeLens               *struct {
			For   []string `json:"for,omitempty"`
			Path  string   `json:"path,omitempty"`
			Title string   `json:"title,omitempty"`
		} `json:"codeLens,omitempty"`
		PreLaunchTask string `json:"preLaunchTask,omitempty"`
		PythonPath    string `json:"python,omitempty"`
		JustMyCode    *bool  `json:"justMyCode,omitempty"`
		Gevent        *bool  `json:"gevent,omitempty"`
	} `json:"configurations"`
	Compounds []interface{} `json:"compounds,omitempty"`
}

type SettingsJSONSQLToolsConnections struct {
	PgOptions struct {
		StatementTimeout int `json:"statement_timeout"`
		QueryTimeout     int `json:"query_timeout"`
	} `json:"pgOptions"`
	PreviewLimit int    `json:"previewLimit"`
	Server       string `json:"server"`
	Port         int    `json:"port"`
	Driver       string `json:"driver"`
	Group        string `json:"group"`
	Name         string `json:"name"`
	Database     string `json:"database"`
	Username     string `json:"username"`
	Password     string `json:"password"`
}

type SettingsJSON struct {
	SQLToolsConnections       []SettingsJSONSQLToolsConnections `json:"sqltools.connections,omitempty"`
	WindmillcodeExtensionPack utils.WindmillcodeExtensionPack   `json:"windmillcode-extension-pack-0"`
}

func main() {
	shared.CDToWorkspaceRoot()
	workspaceRoot, err := os.Getwd()
	settings, err := utils.GetSettingsJSON(workspaceRoot)
	if err != nil {
		return
	}
	windmillcodeSettings = settings
	utils.SetGlobalVars(
		utils.SetGlobalVarsOptions{
			NonInteractive: settings.ExtensionPack.ProcessIfDefaultIsPresent,
		},
	)
	executablePath, err := os.Executable()
	if err != nil {
		fmt.Printf("Error retrieving executable path: %v\n", err)
		return
	}
	executableDir := filepath.Dir(executablePath)

	amountToAppendToPortNumberString := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt:  []string{"This the amount to append to the port numbers for all apps in this project so you dont have port conflicts. check all projects on your computer to find a number that is not used in any of your projects with this template. Or keep track of it yourself"},
			Default: strconv.Itoa(settings.ExtensionPack.MiscReinitializeProject.AmountToAppendToPortNumberString),
		},
	)

	var ProjectName = utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt:  []string{"The project name"},
			Default: settings.ExtensionPack.MiscReinitializeProject.ProjectName,
		},
	)
	if ProjectName == "" {
		fmt.Println("Project name cannot be empty")
		return
	}
	PROJECT_NAME, _ = utils.CreateStringObject(ProjectName, "")
	var OrganizationName = utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt:  []string{"The organization name"},
			Default: settings.ExtensionPack.MiscReinitializeProject.OrganizationName,
		},
	)
	if OrganizationName == "" {
		fmt.Println("Organization name cannot be empty")
		return
	}
	ORGANIZATION_NAME, _ = utils.CreateStringObject(OrganizationName, "")

	VCS_PRIVATE_KEY = utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt:  []string{"The VCS private key"},
			Default: settings.ExtensionPack.MiscReinitializeProject.VCSPrivateKey,
		},
	)
	WEB_SEO_DESCRIPTION = utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt:  []string{"The web sso description"},
			Default: settings.ExtensionPack.MiscReinitializeProject.WebSEODescription,
		},
	)
	WEB_SEO_KEYWORDS = utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt:  []string{"The web seo keywords"},
			Default: settings.ExtensionPack.MiscReinitializeProject.WebSEOKeywords,
		},
	)
  PROXY_URLS_0:= ""
  if len(settings.ExtensionPack.MiscReinitializeProject.ProxyURLs) == 0{
    PROXY_URLS_0 =""
  }
	PROXY_URLS_0 = utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt:  []string{"The proxy urls to use"},
			Default: PROXY_URLS_0,
		},
	)
	FLUTTER_ANDRIOID_GOOGLE_ADS_ID = utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt:  []string{"The flutter android google ads id"},
			Default: settings.ExtensionPack.MiscReinitializeProject.FlutterAndroidGoogleAdsID,
		},
	)
	FLUTTER_ANDROID_GOOGLE_APPLICATION_ID = utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt:  []string{"The flutter android google application id"},
			Default: settings.ExtensionPack.MiscReinitializeProject.FlutterAndroidGoogleApplicationID,
		},
	)
	FLUTTER_IOS_GOOGLE_ADS_ID = utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt:  []string{"The flutter ios google ads id"},
			Default: settings.ExtensionPack.MiscReinitializeProject.FlutterIosGoogleAdsID,
		},
	)
	FLUTTER_IOS_GOOGLE_APPLICATION_ID = utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt:  []string{"The flutter ios google application id"},
			Default: settings.ExtensionPack.MiscReinitializeProject.FlutterIosGoogleApplicationID,
		},
	)
	FLUTTER_IOS_FACEBOOK_APP_ID = utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt:  []string{"The flutter ios facebook app id"},
			Default: settings.ExtensionPack.MiscReinitializeProject.FlutterIosFacebookAppID,
		},
	)
	FLUTTER_IOS_FACEBOOK_CLIENT_TOKEN = utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt:  []string{"The flutter ios facebook client token"},
			Default: settings.ExtensionPack.MiscReinitializeProject.FlutterIosFacebookClientToken,
		},
	)
	FLUTTER_IOS_FACEBOOK_CUSTOM_URL_SCHEME = utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt:  []string{"The flutter ios facebook custom url scheme"},
			Default: settings.ExtensionPack.MiscReinitializeProject.FlutterIosFacebookCustomURLScheme,
		},
	)
	portsArgs := utils.TakeVariableArgsStruct{
		Prompt:  "The Google OAuth URL Schemes for this project",
		Default: strings.Join(settings.ExtensionPack.MiscReinitializeProject.FlutterIosGoogleOAuthURLSchemes," "),
	}
	flutterIosGoogleOAuthURLSchemes := utils.TakeVariableArgs(portsArgs)
	FLUTTER_IOS_GOOGLE_OAUTH_URL_SCHEMES = flutterIosGoogleOAuthURLSchemes.InputArray
	CHROME_DRIVER_PATH = utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt:  []string{"The chrome driver path"},
			Default: settings.ExtensionPack.MiscReinitializeProject.ChromeDriverPath,
		},
	)
	CHROME_DRIVER_PATH = utils.ConvertPathToOSFormat(CHROME_DRIVER_PATH)
	FIREFOX_DRIVER_PATH = utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt:  []string{"The firefox driver path"},
			Default: settings.ExtensionPack.MiscReinitializeProject.FirefoxDriverPath,
		},
	)
	FIREFOX_DRIVER_PATH = utils.ConvertPathToOSFormat(FIREFOX_DRIVER_PATH)
	OPERA_DRIVER_PATH = utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt:  []string{"The opera driver path"},
			Default: settings.ExtensionPack.MiscReinitializeProject.OperaDriverPath,
		},
	)
	OPERA_DRIVER_PATH = utils.ConvertPathToOSFormat(OPERA_DRIVER_PATH)
	EDGE_DRIVER_PATH = utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt:  []string{"The edge driver path"},
			Default: settings.ExtensionPack.MiscReinitializeProject.EdgeDriverPath,
		},
	)
	EDGE_DRIVER_PATH = utils.ConvertPathToOSFormat(EDGE_DRIVER_PATH)



	updateLaunchJSON(workspaceRoot)
	updateSettingsJSON(workspaceRoot, amountToAppendToPortNumberString)
	removeFilesFromDotVscodeFolder(workspaceRoot)
	removeFilesFromDotWindmillcodeSlashLocalFolder(workspaceRoot)
	updateSubPath(executableDir, workspaceRoot, []string{".windmillcode", "Local"}, false)
	removeDevFoldersInDocs(workspaceRoot)
	updateSubPath(executableDir, workspaceRoot, []string{"docs", "tasks_docs", "promote"}, true)
	updateSubPath(executableDir, workspaceRoot, []string{".devcontainer"}, true)
	updateSubPath(executableDir, workspaceRoot, []string{".env"}, true)

	removeDirectories(
		workspaceRoot,
		[][]string{
			{"apps", ".secrets"},
			{"apps", "cloud"},
			{"apps", "database"},
			{"apps", "devops"},
			{"apps", "extensions"},
		},
		[]string{},
	)
	removePathsForAngularApp(workspaceRoot)
	removePathsForFlaskApp(workspaceRoot)
	removePathsForFlutterApp(workspaceRoot)
	removePathsForSeleniumApp(workspaceRoot)
	updateSubPath(executableDir, workspaceRoot, []string{"apps"}, true)

}



func removePathsForSeleniumApp(workspaceRoot string) {
	removeDirectories(
		workspaceRoot,
		[][]string{
			{"apps", "testing", "appium"},
			{"apps", "testing", "AppiumApp"},
			{"apps", "testing", "selenium"},
			{"apps", "testing", "SeleniumApp"},
		},
		[]string{},
	)

}

func removePathsForFlutterApp(workspaceRoot string) {
	removeDirectories(
		workspaceRoot,
		[][]string{
			{"apps", "mobile", "FlutterApp"},
		},
		[]string{},
	)

}

func removePathsForAngularApp(workspaceRoot string) {
	removeDirectories(
		workspaceRoot,
		[][]string{
			{"apps", "frontend", "AngularApp", "src", ".well-known"},
			{"apps", "frontend", "AngularApp", "src", "assets", "scripts"},
			{"apps", "frontend", "AngularApp", "src", "app", "layouts"},
			{"apps", "frontend", "AngularApp", "src", "app", "pages"},
			{"apps", "frontend", "AngularApp", "src", "app", "shared", "guards"},
			{"apps", "frontend", "AngularApp", "src", "app", "shared", "services"},
		},
		[]string{},
	)
	removeDirectories(
		workspaceRoot,
		[][]string{
			{"apps", "frontend", "AngularApp", "src", "assets", "i18n"},
		},
		[]string{},
	)
	removeDirectories(
		workspaceRoot,
		[][]string{
			{"apps", "frontend", "AngularApp", "src", "assets", "media"},
		},
		[]string{
			"app",
		},
	)
	removeDirectories(
		workspaceRoot,
		[][]string{
			{"apps", "frontend", "AngularApp", "src", "assets", "media", "app"},
		},
		[]string{
			"apple-app-store.svg",
			"apple-logo.png",
			"background-0.svg",
			"drawer-logo-bg.png",
			"facebook-logo.png",
			"github-logo.png",
			"google_play_store.svg",
			"google-logo.png",
			"twitter-logo.png",
		},
	)
	removeDirectories(
		workspaceRoot,
		[][]string{
			{"apps", "frontend", "AngularApp", "src", "app", "shared", "components"},
		},
		[]string{
			"confirm-dialog-zero",
			"logo-img-zero",
			"overlay-loading",
			"sample-cpnt",
			"threejs-background-zero",
		},
	)
}

func removePathsForFlaskApp(workspaceRoot string) {
	removeDirectories(
		workspaceRoot,
		[][]string{
			{"apps", "backend", "FlaskApp", "endpoints"},
			{"apps", "backend", "FlaskApp", "unit_tests", "FlaskTesting", "handlers"},
		},
		[]string{
			"healthcheck_endpoint.py",
			"test_healthcheck_endpoint.py",
		},
	)
	removeDirectories(
		workspaceRoot,
		[][]string{
			{"apps", "backend", "FlaskApp", "handlers"},
			{"apps", "backend", "FlaskApp", "unit_tests", "FlaskTesting", "handlers"},
		},
		[]string{
			"test_healthcheck_handler.py",
			"healthcheck_handler.py",
		},
	)

	removeDirectories(
		workspaceRoot,
		[][]string{
			{"apps", "backend", "FlaskApp", "managers"},
			{"apps", "backend", "FlaskApp", "unit_tests", "FlaskTesting", "managers"},
		},
		[]string{
			"cache_manager",
			"email_manager",
			"firebase_manager",
			"google_cloud_platform_manager",
			"hashicorp_vault_manager",
			"limiter_manager",
			"logger_manager",
			"one_signal_manager",
			"postgresql_manager",
			"reportlab_manager",
			"retry_manager",
			"sentry_manager",
			"socketio_manager",
		},
	)

	utilsFiles := []string{
		"api_msg_format.py",
		"constants.py",
		"crypto_utils.py",
		"dict_utils.py",
		"env_vars.py",
		"iterable_utils.py",
		"local_deps.py",
		"my_util.py",
		"print_if_dev.py",
		"string_utils.py",
		"url_utils.py",
	}

	var allUtilsFiles []string
	for _, file := range utilsFiles {
		allUtilsFiles = append(allUtilsFiles, file)
		if strings.HasSuffix(file, ".py") {
			testFile := "test_" + file
			allUtilsFiles = append(allUtilsFiles, testFile)
		}
	}
	removeDirectories(
		workspaceRoot,
		[][]string{
			{"apps", "backend", "FlaskApp", "utils"},
			{"apps", "backend", "FlaskApp", "unit_tests", "FlaskTesting", "utils"},
		},
		append([]string{
			"exceptions",
			"wml_libs",
			"api_msg_format.py",
		}, allUtilsFiles...),
	)
}

func removeDirectories(root string, targetDirs [][]string, keepDirs []string) bool {
	for _, dir := range targetDirs {
		fullPath := utils.JoinAndConvertPathToOSFormat(append([]string{root}, dir...)...)
		err := utils.TraverseDirectory(utils.TraverseDirectoryParams{
			RootDir: fullPath,
			Predicate: func(srcPath string, info os.FileInfo) {
				if srcPath != fullPath {
					err := os.RemoveAll(srcPath)
					if err != nil {
						fmt.Printf("Error removing %s: %v\n", srcPath, err)
					}
				}
			},
			Filter: func(srcPath string, info os.FileInfo) bool {
				return !utils.ArrayContainsAny(keepDirs, strings.Split(srcPath, string(filepath.Separator)))
			},
		})
		if err != nil {
			fmt.Printf("Error traversing directory: %v\n", err)
		}
	}

	return false
}

func removeDevFoldersInDocs(workspaceRoot string) bool {
	err := os.RemoveAll(
		utils.JoinAndConvertPathToOSFormat(
			workspaceRoot, "snippets",
		),
	)
	if err != nil {
		fmt.Printf("Error removing snippets folder: %v\n", err)
		return true
	}
	docs := []string{
		utils.JoinAndConvertPathToOSFormat(workspaceRoot, "docs", "tasks_docs"),
		utils.JoinAndConvertPathToOSFormat(workspaceRoot, "docs", "app_docs"),
		utils.JoinAndConvertPathToOSFormat(workspaceRoot, "issues"),
		utils.JoinAndConvertPathToOSFormat(workspaceRoot, "misc"),
	}
	keep := []string{
		"template", "system_design",
	}
	for _, doc := range docs {
		err = utils.TraverseDirectory(utils.TraverseDirectoryParams{
			RootDir: doc,
			Predicate: func(srcPath string, info os.FileInfo) {
				if srcPath != doc {
					err := os.RemoveAll(srcPath)
					if err != nil {

						fmt.Println(fmt.Errorf("Error removing %s: %v\n", srcPath, err))
					}
				}
			},
			Filter: func(srcPath string, info os.FileInfo) bool {
				return !utils.ArrayContainsAny(keep, strings.Split(srcPath, string(filepath.Separator)))
			},
		})
		if err != nil {
			fmt.Printf("Error traversing directory: %v\n", err)
		}
	}
	return false
}

func updateSubPath(executableDir, workspaceRoot string, targetRelativePath []string, overwrite bool) bool {
	srcDir := utils.JoinAndConvertPathToOSFormat(
		append([]string{executableDir, "Template"}, targetRelativePath...)...,
	)

	destDir := utils.JoinAndConvertPathToOSFormat(
		append([]string{workspaceRoot}, targetRelativePath...)...,
	)

	utils.TraverseDirectory(utils.TraverseDirectoryParams{
		RootDir: srcDir,
		Predicate: func(srcPath string, info os.FileInfo) {
			// Construct the relative path
			relPath := utils.RemovePathPrefix(srcPath, []string{srcDir})

			destPath := utils.JoinAndConvertPathToOSFormat(destDir, relPath)

			if info.IsDir() {
				// Create the directory in the destination if it doesn't exist
				if _, err := os.Stat(destPath); os.IsNotExist(err) {
					err = os.MkdirAll(destPath, os.ModePerm)
					if err != nil {
						fmt.Printf("Error creating directory %s: %v\n", destPath, err)
					} else {
						fmt.Printf("Created directory: %s\n", destPath)
					}
				}
			} else {
				// Ensure the parent directory exists
				parentDir := filepath.Dir(destPath)
				if _, err := os.Stat(parentDir); os.IsNotExist(err) {
					err = os.MkdirAll(parentDir, os.ModePerm)
					if err != nil {
						fmt.Printf("Error creating parent directory %s: %v\n", parentDir, err)
						return
					}
				}

				// Check if the file should be copied
				shouldCopy := overwrite
				if !overwrite {
					if _, err := os.Stat(destPath); os.IsNotExist(err) {
						shouldCopy = true
					}
				}

				if shouldCopy {
					// Copy the file
					err := utils.CopyFile(srcPath, destPath)
					if err != nil {
						fmt.Printf("Error copying file %s to %s: %v\n", srcPath, destPath, err)
					} else {
						// Process the copied file with replacements
						fileString, _ := utils.ReadFile(destPath)
						fileString = strings.ReplaceAll(fileString, "[FLUTTER_ANDROID_KEY_PROPERTIES_STORE_FILE]", FLUTTER_ANDROID_KEY_PROPERTIES_STORE_FILE)
						fileString = strings.ReplaceAll(fileString, "[FLUTTER_ANDROID_KEY_PROPERTIES_KEY_PASSWORD]", FLUTTER_ANDROID_KEY_PROPERTIES_KEY_PASSWORD)
						fileString = strings.ReplaceAll(fileString, "[FLUTTER_ANDROID_KEY_PROPERTIES_STORE_PASSWORD]", FLUTTER_ANDROID_KEY_PROPERTIES_STORE_PASSWORD)
						fileString = strings.ReplaceAll(fileString, "[FLUTTER_ANDROID_PREVIEW_KEY_PROPERTIES_KEY_PASSWORD]", FLUTTER_ANDROID_PREVIEW_KEY_PROPERTIES_KEY_PASSWORD)
						fileString = strings.ReplaceAll(fileString, "[FLUTTER_ANDROID_PREVIEW_KEY_PROPERTIES_STORE_FILE]", FLUTTER_ANDROID_PREVIEW_KEY_PROPERTIES_STORE_FILE)

						fileString = strings.ReplaceAll(fileString, "[FLUTTER_IOS_GOOGLE_APPLICATION_ID]", FLUTTER_IOS_GOOGLE_APPLICATION_ID)
						fileString = strings.ReplaceAll(fileString, "[FLUTTER_IOS_FACEBOOK_APP_ID]", FLUTTER_IOS_FACEBOOK_APP_ID)
						fileString = strings.ReplaceAll(fileString, "[FLUTTER_IOS_FACEBOOK_CLIENT_TOKEN]", FLUTTER_IOS_FACEBOOK_CLIENT_TOKEN)

						fileString = strings.ReplaceAll(fileString, "[FLUTTER_IOS_FACEBOOK_CUSTOM_URL_SCHEME]", FLUTTER_IOS_FACEBOOK_CUSTOM_URL_SCHEME)
						FLUTTER_IOS_GOOGLE_OAUTH_URL_SCHEMES_Value := ""
						for _, scheme := range FLUTTER_IOS_GOOGLE_OAUTH_URL_SCHEMES {
							FLUTTER_IOS_GOOGLE_OAUTH_URL_SCHEMES_Value += `
							<dict>
								<key>CFBundleTypeRole</key>
								<string>Editor</string>
								<key>CFBundleURLSchemes</key>
								<array>
									<string>` + scheme + `</string>
								</array>
							</dict>
							`
						}

						fileString = strings.ReplaceAll(fileString, "[FLUTTER_IOS_GOOGLE_OAUTH_URL_SCHEMES]", FLUTTER_IOS_GOOGLE_OAUTH_URL_SCHEMES_Value)

						fileString = strings.ReplaceAll(fileString, "[PROJECT_NAME]", PROJECT_NAME.Orig)
						fileString = strings.ReplaceAll(
							fileString,
							"[PROJECT_NAME_CAPITAL]",
							PROJECT_NAME.Capitalize(false, ""),
						)
						fileString = strings.ReplaceAll(fileString, "[ORGANIZATION_NAME]", ORGANIZATION_NAME.Orig)
						fileString = strings.ReplaceAll(fileString, "[ORGANIZATION_NAME_CAPITAL]", ORGANIZATION_NAME.Capitalize(false, ""))
						fileString = strings.ReplaceAll(fileString, "[WEB_SEO_DESCRIPTION]", WEB_SEO_DESCRIPTION)
						fileString = strings.ReplaceAll(fileString, "[WEB_SEO_KEYWORDS]", WEB_SEO_KEYWORDS)
						fileString = strings.ReplaceAll(fileString, "[WEBDRIVERS_PATH]", WEB_DRIVER_PATH)
						fileString = strings.ReplaceAll(fileString, "[VCS_PRIVATE_KEY]", VCS_PRIVATE_KEY)
						fileString = strings.ReplaceAll(fileString, "[PROXY_URLS_0]", PROXY_URLS_0)
						fileString = strings.ReplaceAll(fileString, "[FLUTTER_ANDRIOID_GOOGLE_ADS_ID]", FLUTTER_ANDRIOID_GOOGLE_ADS_ID)
						fileString = strings.ReplaceAll(fileString, "[FLUTTER_IOS_GOOGLE_ADS_ID]", FLUTTER_IOS_GOOGLE_ADS_ID)

						fileString = strings.ReplaceAll(fileString, "[CHROME_DRIVER_PATH]", CHROME_DRIVER_PATH)
						fileString = strings.ReplaceAll(fileString, "[FIREFOX_DRIVER_PATH]", FIREFOX_DRIVER_PATH)
						fileString = strings.ReplaceAll(fileString, "[OPERA_DRIVER_PATH]", OPERA_DRIVER_PATH)
						fileString = strings.ReplaceAll(fileString, "[EDGE_DRIVER_PATH]", EDGE_DRIVER_PATH)

						fileString = strings.ReplaceAll(fileString, "[Angular_Run_0]", strconv.Itoa(windmillcodeSettings.ExtensionPack.Ports.AngularRun0))
						fileString = strings.ReplaceAll(fileString, "[Flask_Run_0]", strconv.Itoa(windmillcodeSettings.ExtensionPack.Ports.FlaskRun0))
						fileString = strings.ReplaceAll(fileString, "[Postgres_0]", strconv.Itoa(windmillcodeSettings.ExtensionPack.Ports.Postgres0))
						fileString = strings.ReplaceAll(fileString, "[Firebase_Emulator_Auth_0]", strconv.Itoa(windmillcodeSettings.ExtensionPack.Ports.FirebaseEmulatorAuth0))



						err = os.WriteFile(destPath, []byte(fileString), 0644)
						if err != nil {
							fmt.Printf("Error writing file %s: %v\n", destPath, err)
						}
						fmt.Printf("Copied: %s -> %s\n", srcPath, destPath)
					}
				} else {
					fmt.Printf("Skipped (exists): %s\n", destPath)
				}
			}
		},
		Filter: func(path string, info os.FileInfo) bool {
			return true
		},
	})
	return false
}

func removeFilesFromDotWindmillcodeSlashLocalFolder(workspaceRoot string) bool {

	targetFolder := "Local"
	keep := []string{
		"minify_flask_app",
		"docker_dump_db.go",
		"docker_init_sql_container.go",
		"go.mod",
		"go.sum",
	}
	rootDir := utils.JoinAndConvertPathToOSFormat(workspaceRoot, ".windmillcode", targetFolder)
	err := utils.TraverseDirectory(utils.TraverseDirectoryParams{
		RootDir: rootDir,
		Predicate: func(path string, info os.FileInfo) {

			relPath := utils.RemovePathPrefix(
				path,
				strings.Split(rootDir, string(filepath.Separator)),
			)
			relPathArray := strings.Split(relPath, string(filepath.Separator))
			if !utils.ArrayContainsAny(keep, relPathArray) {
				os.Remove(path)
			}

		},
		Filter: func(path string, info os.FileInfo) bool {
			return filepath.Base(path) != targetFolder
		},
	})
	if err != nil {
		utils.LogErrorWithTraceBack("Error traversing directory", err)
	}
	return false
}

func removeFilesFromDotVscodeFolder(workspaceRoot string) bool {
	regexPattern := `(?i)^(launch|settings|tasks)\.json$`
	regex, err := regexp.Compile(regexPattern)
	if err != nil {
		utils.LogErrorWithTraceBack("Error compiling regex pattern", err)
		return true
	}
	err = utils.TraverseDirectory(utils.TraverseDirectoryParams{
		RootDir: utils.JoinAndConvertPathToOSFormat(workspaceRoot, ".vscode"),
		Predicate: func(path string, info os.FileInfo) {
			baseName := filepath.Base(path)
			if !regex.MatchString(baseName) && baseName != ".vscode" {
				err := os.Remove(path)
				if err != nil {
					utils.LogErrorWithTraceBack("Error removing file", err)
				} else {
					println("Removed item:", path)
				}
			}
		},
		Filter: func(path string, info os.FileInfo) bool {
			return true
		},
	})
	if err != nil {
		utils.LogErrorWithTraceBack("Error traversing directory", err)
	}
	return false
}

func structToRawMessage(data interface{}) (json.RawMessage, error) {
	// Marshal the data into a JSON byte slice
	bytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// Return the byte slice as json.RawMessage
	return json.RawMessage(bytes), nil
}

func updateLaunchJSON(workspaceRoot string) bool {
	var launchJSON LaunchJSON

	launchJSONPath := utils.JoinAndConvertPathToOSFormat(workspaceRoot, ".vscode", "launch.json")
	launchJSONContent, err := os.ReadFile(launchJSONPath)
	if err != nil {
		log.Fatalf("Failed to read launch.json file: %v", err)
	}
	if err != nil {
		fmt.Println(err)
		return true
	}
	cleanLaunchJSON, err := utils.CleanJSON(launchJSONContent)
	if err != nil {
		utils.LogErrorWithTraceBack("Error removing comments:", err)
		return true
	}
	err = json.Unmarshal([]byte(cleanLaunchJSON), &launchJSON)
	if err != nil {
		utils.LogErrorWithTraceBack("Error unmarshalling JSON:", err)
		return true
	}

	for i, config := range launchJSON.Configurations {
		if config.Type == "dart" {
			config.DeviceID = ""
			config.ToolArgs = []string{
				"--dart-define",
				"YOUR_FLUTTER_ENV_VAR=YOUR_FLUTTER_ENV_VAR_VALUE",
			}
		} else if config.Type == "debugpy" {
			config.Env = map[string]string{
				"FLASK_BACKEND_ENV":"DEV",
				"YOUR_PYTHON_ENV_VAR": "YOUR_PYTHON_ENV_VAR_VALUE",
			}
		}

		launchJSON.Configurations[i] = config
	}

	launchJSONContent, err = json.MarshalIndent(launchJSON, "", "  ")
	if err != nil {
		utils.LogErrorWithTraceBack("Error marshalling JSON for launch.json:", err)
		return true
	}

	err = os.WriteFile(launchJSONPath, launchJSONContent, 0644)
	if err != nil {
		utils.LogErrorWithTraceBack("Error writing launch.json file:", err)
		return true
	}
	return false
}

func updateSettingsJSON(workspaceRoot string, amountToAppendToPortNumberString string) bool {
	amountToAppendToPortNumber, err := strconv.Atoi(amountToAppendToPortNumberString)
	if err != nil {
		fmt.Println(err)
	}
	var settingsJSON map[string]json.RawMessage
	var parseableSettingsJSON SettingsJSON

	settingsJSONPath := utils.JoinAndConvertPathToOSFormat(workspaceRoot, ".vscode", "settings.json")
	settingsJSONContent, err := os.ReadFile(settingsJSONPath)
	if err != nil {
		log.Fatalf("Failed to read settings.json file: %v", err)
	}
	if err != nil {
		fmt.Println(err)
		return true
	}
	cleanSettingsJSON, err := utils.CleanJSON(settingsJSONContent)
	if err != nil {
		utils.LogErrorWithTraceBack("Error removing comments:", err)
		return true
	}
	err = json.Unmarshal([]byte(cleanSettingsJSON), &settingsJSON)
	if err != nil {
		utils.LogErrorWithTraceBack("Error unmarshalling JSON:", err)
		return true
	}

	err = json.Unmarshal([]byte(cleanSettingsJSON), &parseableSettingsJSON)
	if err != nil {
		utils.LogErrorWithTraceBack("Error unmarshalling JSON:", err)
		return true
	}

	parseableSettingsJSON.SQLToolsConnections = []SettingsJSONSQLToolsConnections{}
	parseableSettingsJSON.WindmillcodeExtensionPack.TasksToRunOnFolderOpen = append(
		parseableSettingsJSON.WindmillcodeExtensionPack.TasksToRunOnFolderOpen,
		"angular frontend: run",
		"firebase cloud: run emulators",
		"misc: run proxies",
	)

	ports := &utils.WMLPorts{}
	basePorts := map[string]int{
		"AngularRun0":                10000,
		"AngularTest0":               10010,
		"AngularCoverageTest0":       10020,
		"AngularAnalyzer0":           10030,
		"AngularSSG0":                10040,
		"AngularSSG1":                10050,
		"AngularSSG2":                10060,
		"FlaskRun0":                  10070,
		"FlaskTest0":                 10080,
		"Postgres0":                  10090,
		"FirebaseEmulatorUI0":        10100,
		"FirebaseEmulatorAuth0":      10110,
		"FirebaseEmulatorStorage0":   10130,
		"FirebaseEmulatorHosting0":   10140,
		"FirebaseEmulatorFirestore0": 10150,
		"DiodeProxies0":              10160,
	}
	v := reflect.ValueOf(ports).Elem()
	for fieldName, basePort := range basePorts {
		field := v.FieldByName(fieldName)
		if field.IsValid() && field.CanSet() {
			field.SetInt(int64(basePort + amountToAppendToPortNumber))
		}
	}
	parseableSettingsJSON.WindmillcodeExtensionPack.Ports = *ports
	parseableSettingsJSON.WindmillcodeExtensionPack.ProcessIfDefaultIsPresent = utils.ProcessIfDefaultIsPresentStruct{
		Global: false,
	}
	parseableSettingsJSON.WindmillcodeExtensionPack.ShopifyRun = utils.ShopifyRunStruct{
		ProjectName: "base-folder-name-of-shopify-app",
	}
	parseableSettingsJSON.WindmillcodeExtensionPack.ProxyURLs = fmt.Sprintf("https://example.com:%d", parseableSettingsJSON.WindmillcodeExtensionPack.Ports.FlaskRun0)
	parseableSettingsJSON.WindmillcodeExtensionPack.PythonVersion0 = "Skip"
	parseableSettingsJSON.WindmillcodeExtensionPack.NodeJSVersion0 = "Skip"
	parseableSettingsJSON.WindmillcodeExtensionPack.JavaVersion0 = "Skip"
	parseableSettingsJSON.WindmillcodeExtensionPack.OpenAIAPIKey0 = ""

	settingsJSON["sqltools.connections"], _ = structToRawMessage(parseableSettingsJSON.SQLToolsConnections)
	settingsJSON["windmillcode-extension-pack-0"], _ = structToRawMessage(parseableSettingsJSON.WindmillcodeExtensionPack)

	// fmt.Println(string(settingsJSON["windmillcode-extension-pack-0"]))
	settingsJSONContent, err = json.MarshalIndent(settingsJSON, "", "  ")
	if err != nil {
		utils.LogErrorWithTraceBack("Error marshalling JSON for settings.json:", err)
		return true
	}

	err = os.WriteFile(settingsJSONPath, settingsJSONContent, 0644)
	if err != nil {
		utils.LogErrorWithTraceBack("Error writing settings.json file:", err)
		return true
	}
	return false
}
