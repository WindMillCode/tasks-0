package main

import (
	"fmt"
	"main/shared"
	"os"

	"github.com/windmillcode/go_cli_scripts/v6/utils"
)

func main() {

	shared.CDToWorkspaceRoot()
	workspaceRoot, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting workspace root: %v", err)
		return
	}
	settings, err := utils.GetSettingsJSON(workspaceRoot)
	if err != nil {
		return
	}
	utils.SetGlobalVars(
		utils.SetGlobalVarsOptions{
			NonInteractive: settings.ExtensionPack.ProcessIfDefaultIsPresent,
		},
	)
	utils.CDToExpoApp()
	expoRoot, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting react native root: %v", err)
		return
	}

	cliInfo := utils.ShowMenuModel{
		Prompt:  "profile",
		Choices: []string{"development", "preview", "production"},
	}
	myProfile := utils.ShowMenu(cliInfo, nil)

	cliInfo = utils.ShowMenuModel{
		Prompt:  "platform",
		Choices: []string{"android", "ios"},
	}
	myPlatform := utils.ShowMenu(cliInfo, nil)

	cliInfo = utils.ShowMenuModel{
		Prompt:  "build locally",
		Choices: []string{"TRUE", "FALSE"},
	}
	localBuild := utils.ShowMenu(cliInfo, nil)

	fileExt := "ipa"

	if myPlatform == "android" && localBuild == "TRUE" {

		cliInfo := utils.ShowMenuModel{
			Prompt:  "android file extension",
			Choices: []string{"apk", "aab"},
		}
		fileExt = utils.ShowMenu(cliInfo, nil)
	}

	outputDir := ""
	if localBuild == "TRUE" {
		outputDir = utils.GetInputFromStdin(
			utils.GetInputFromStdinStruct{
				Prompt:  []string{"The output dir"},
				Default: utils.JoinAndConvertPathToOSFormat(workspaceRoot, "misc", fmt.Sprintf("local-%s-%s.%s", myProfile, myPlatform, fileExt)),
			},
		)
	}

	cliInfo = utils.ShowMenuModel{
		Prompt:  "Run prebuild?",
		Choices: []string{"YES", "NO"},
	}
	runPrebuild := utils.ShowMenu(cliInfo, nil)
	prebuildArgs := []string{"expo", "prebuild", "--platform", myPlatform}

	if runPrebuild == "YES" {
		cliInfo := utils.ShowMenuModel{
			Prompt:  "run prebuild with clean",
			Choices: []string{"YES", "NO"},
		}
		withClean := utils.ShowMenu(cliInfo, nil)
		if withClean == "YES" {
			prebuildArgs = append(prebuildArgs, "--clean")
		}

	}

	jsEngine := ""
	cliInfo = utils.ShowMenuModel{
		Prompt:  "upload source maps",
		Choices: []string{"YES", "NO"},
		Default: "YES",
	}
	uploadSourceMaps := utils.ShowMenu(cliInfo, nil)
	sentryOrg := ""
	sentryProject := ""
	sentryRelease := ""
	if uploadSourceMaps == "YES" {
		cliInfo = utils.ShowMenuModel{
			Prompt:  "jsEngine (if unsure pick jsc)",
			Choices: []string{"jsc", "hermes"},
			Default: "jsc",
		}
		jsEngine = utils.ShowMenu(cliInfo, nil)

		sentryOrg = utils.GetInputFromStdin(
			utils.GetInputFromStdinStruct{
				Prompt:  []string{"The sentry organization name"},
				Default: settings.ExtensionPack.ReactNativeExpoMobileBuild.SentryOrg,
			},
		)

		sentryProject = utils.GetInputFromStdin(
			utils.GetInputFromStdinStruct{
				Prompt:  []string{"The sentry project name"},
				Default: settings.ExtensionPack.ReactNativeExpoMobileBuild.SentryProject,
			},
		)

		sentryRelease = utils.GetInputFromStdin(
			utils.GetInputFromStdinStruct{
				Prompt:  []string{"The sentry release name"},
				Default: settings.ExtensionPack.ReactNativeExpoMobileBuild.SentryRelease,
			},
		)
	}

	commandArgs := []string{"build", "--profile", myProfile, "--platform", myPlatform}
	if localBuild == "TRUE" {
		commandArgs = append(commandArgs, "--output", outputDir)
	}

	if localBuild == "TRUE" {
		commandArgs = append(commandArgs, "--local")
	}
	if runPrebuild == "YES" {
		runPrebuildOptions := utils.CommandOptions{
			Command:         "npx",
			Args:            prebuildArgs,
			GetOutput:       false,
			PrintOutputOnly: true,
		}
		utils.RunCommandWithOptions(runPrebuildOptions)
	}
	os.RemoveAll("~/.app-store/")

	opts := utils.CommandOptions{
		Command:         "eas",
		Args:            commandArgs,
		GetOutput:       false,
		PrintOutputOnly: true,
		PanicOnError:    true,
	}
	utils.RunCommandWithOptions(opts)

	if uploadSourceMaps == "YES" {
		bundleOutput := map[string]string{
			"android": "index.android.bundle",
			"ios":     "main.jsbundle",
		}[myPlatform]
		sourceMapOutput := map[string]string{
			"android": "index.android.bundle.packager.map",
			"ios":     "main.jsbundle.map",
		}[myPlatform]
		if jsEngine == "hermes" {

			opts = utils.CommandOptions{
				Command:      "npx",
				PanicOnError: true,
				Args: []string{
					"expo",
					"export:embed",
					"--entry-file",
					utils.JoinAndConvertPathToOSFormat(expoRoot, "node_modules/expo-router/entry.js"),
					"--platform", myPlatform,
					"--dev", "false",
					"--reset-cache",
					"--bundle-output", bundleOutput,
					"--sourcemap-output", sourceMapOutput,
					"--minify", "false",
				},
			}
			utils.RunCommandWithOptions(opts)

			newCommand := map[string]string{
				"android": "node_modules/react-native/sdks/hermesc/osx-bin/hermesc",
				"ios":     "ios/Pods/hermes-engine/destroot/bin/hermesc",
			}[myPlatform]

			out := map[string]string{
				"android": "index.android.bundle.hbc",
				"ios":     "main.jsbundle.hbc",
			}[myPlatform]
			opts = utils.CommandOptions{
				Command:      utils.ConvertPathToOSFormat(newCommand),
				PanicOnError: true,
				Args: []string{
					"-O", "-emit-binary",
					"-output-source-map",
					"-out", out,
					bundleOutput,
				},
			}

			if err := os.Remove(bundleOutput); err != nil && !os.IsNotExist(err) {
				// panic(err)
			}

			if err := os.Rename(out, bundleOutput); err != nil {
				panic(err)
			}
			packagerMap := map[string]string{
				"android": "index.android.bundle.packager.map",
				"ios":     "main.jsbundle.packager.map",
			}[myPlatform]

			if myPlatform == "ios" {
				mvFile(sourceMapOutput, packagerMap)
			}

			hbcMap := map[string]string{
				"android": "index.android.bundle.hbc.map",
				"ios":     "main.jsbundle.hbc.map",
			}[myPlatform]
			outputMap := map[string]string{
				"android": "index.android.bundle.map",
				"ios":     "main.jsbundle.map",
			}[myPlatform]

			opts = utils.CommandOptions{
				Command:      "node",
				PanicOnError: true,
				Args: []string{
					utils.ConvertPathToOSFormat("node_modules/@sentry/react-native/scripts/copy-debugid.js"),
					packagerMap,
					hbcMap,
					"-o", outputMap,
				},
			}

			utils.RunCommandWithOptions(opts)

			opts = utils.CommandOptions{
				Command:      "node",
				PanicOnError: true,
				Args: []string{
					utils.ConvertPathToOSFormat("node_modules/@sentry/react-native/scripts/copy-debugid.js"),
					packagerMap, outputMap,
				},
			}

			utils.RunCommandWithOptions(opts)

			rmFile(packagerMap)

			opts = utils.CommandOptions{
				Command:      "npx",
				PanicOnError: true,
				Args: []string{
					"sentry-cli", "sourcemaps", "upload",
					"--debug-id-reference",
					"--strip-prefix", expoRoot,
					bundleOutput, outputMap,
				},
			}

			utils.RunCommandWithOptions(opts)
		} else if jsEngine == "jsc" {

			opts = utils.CommandOptions{
				Command:      "npx",
				PanicOnError: true,
				Args: []string{
					"expo", "export:embed",
					"--entry-file", utils.JoinAndConvertPathToOSFormat(expoRoot, "node_modules/expo-router/entry.js"),
					"--platform", myPlatform,
					"--dev", "false",
					"--reset-cache",
					"--bundle-output", bundleOutput,
					"--sourcemap-output", sourceMapOutput,
					"--minify", "true",
				},
			}

			utils.RunCommandWithOptions(opts)

			opts = utils.CommandOptions{
				Command:      "npx",
				PanicOnError: true,
				Args: []string{
					"sentry-cli", "sourcemaps", "upload",
					"--org", sentryOrg,
					"--project", sentryProject,
					"--release", sentryRelease,
					"--strip-prefix",
					expoRoot,
					bundleOutput, sourceMapOutput,
				},
			}

			utils.RunCommandWithOptions(opts)
		}
	}
}

func mvFile(source, destination string) error {
	if err := os.Rename(source, destination); err != nil {
		return fmt.Errorf("failed to move file from %s to %s: %w", source, destination, err)
	}
	return nil
}

func rmFile(filePath string) error {
	if err := os.Remove(filePath); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to remove file %s: %w", filePath, err)
	}
	return nil
}
