# ChangeLog

All notable changes to the "Windmillcode Tasks Zero" extension pack will be documented in this file.

* Version updates will be based on vscode relesases
on every vscode update a new version will be release

* the software version extends the vscode patch version by 3 zeros giving us
1000 possible updates before there is an update to vscode and extends back to zero

* you would have to check the CHANGELOG for any breaking, (major), minor or patched updates which will be denoted respectively



## [1.85.1000] - 12-27-2023
* Extension made available to the public ready for use

## [1.85.1001] - 12-27-2023
* [FIX] fixed a bug with flask backend create endpoint


## [1.85.1002] - 12-27-2023
* [UPDATE] added a feature where you can view coverage info at localhost:8003 for angular_frontend_test
* [UPDATE] added a feature where you can view coverage info at localhost:8004 for flask_backend_test


## [1.85.1003] - 12-27-2023
* [PATCH] - fix bug in flask_backned_run flask_backend_test and docker_init_container trying
based on an underlying command fn from windmillcode/go_scripts package under investigation
* [UPDATE]  seperated coverage http-server to its on task flask_backend_view_coverage_info from flask_backend_test


## [1.85.1004] - 1-2-2024
* [UPDATE] - configured angular frontend run and flask backend run so a developer wont have to toggle between developer and docker development in the settings

## [1.85.1005] - 1-5-2024
* [UPDATE] - made an an update to tasks update workspace by providing additional prompts that will address devices with less capable peformance components

## [1.85.1006] - 1-5-2024
* update tasks to ensure for linux that the bash shell is used and ran in interactive mode in an attempt to source the .bashrc with important features such as the $PATH

## [1.85.1007] - 1-5-2024
* [PATCH] patch an issue with an underlying library

## [1.85.1008] - 1-6-2024
* [UPDATE] updated undelying go libraries so now output that returns a value also prints to the console

## [1.85.1009] - 1-6-2024
* [FIX] ensured the go prorams build into go executables w/o error


## [1.85.1010] - 1-7-2024
8 [UPDATE] ensured the sql make db update entry starts with year so recent db snapshots appear lower in the file explorer than older ones

## [1.85.1011] - 1-11-2024
* [PATCH] - updated internal go code

## [1.85.1012] - 1-13-2024
* [FIX] - update flask_backend_create_manager to conform to the testing location for managers

## [1.85.1013] - 1-13-2024
* [PATCH] - updated sql_make_db_schema_update_entry so that the correct timestamp format comes out

## [1.85.1014] - 1-15-2024
* [PATCH] - corrected an issue with flask_backend_create_manager


## [1.85.1015] - 1-15-2024
* [UPDATE] added misc run proxy to run x amount of proxies on your running apps
the windmillcode-extension-pack-0.proxyURLs is a space seperated string for multiple entries
it optionally runs a diode tunnel for each proxy making the proxy public on the www learn more [here](https://support.diode.io/article/ss32engxlq-publish-your-local-webserver)

* [UPDATE] added customUserIsPresent option to
tasks_update_workspace_with_latest_tasks so you dont have to manully hit enter for the Windmillcode user

## [1.85.1016] - 1-15-2024
* [UPDATE] -made internal changes
* [UPDATE]- flask backend env does not print output to the console anymore

## [1.85.1017] - 1-15-2024
* [BREAKING CHANGE] -removed flask backend view coverage info
* [UPDATE] flask backend test supports both unit testing and coverage info
* [UPDATE] can leverage tasksToRunOnFolderOpen with the labels of the tasks you want to automatically run on folder open


## [1.85.1018] -1-21-2024
* [UPDATE] let user know what port angular app coverage info is running on
* [FIX] fixed issue in flask backend run where you could not see output info


## [1.85.2002] -1-25-2024
* [FIX] - fixed feature where reload functionality is broken on some OS systems for flask backend run
* [PATCH] -python install specific packages now updates the requirements files accordingly but the script that was written to remove local packages was written by chat-gpt and may not cover all edge cases use with caution

## [1.85.2003] -1-25-2024
* [FIX] - fixed a bug with python install specific packages

## [1.85.2004] -1-26-2024
* [PATCH] - fixed an issue with flask backend run

## [1.85.2005] -1-28-2024
* [PATCH] - updated env var extraction logic for flask backend run and flask backend test

## [1.85.2006] -1-30-2024
* [FIX] -fixed issue with python_install_specific_packages
install and uninstall  correctly remove duplicates from requirements.txt
instal and uninstall removes the package itself along with its dist-info folder

## [1.85.2007] -1-31-2024
* [UPDATE] added flask backend run with reloader for python application running gevent
* [UPDATE] python install specific packagges has a force option for uninstall
* [UPDATE] python install specific packagges has an upgrade option for upgreade
* [UPDATE] update workspace without extension does not build the go executables anymore easier to build the tasks

## [1.85.2008] -1-31-2024
* [UPDATE] ensured tasks to run on folder open works correctly

## [1.86.1000] -2-11-2024
* [UPDATE] - corrected error message from python install specific packages

## [1.86.1001] -2-12-2024
* [PATCH] - fixed issue with template for flask backend create manager

## [1.86.1010] -2-13-2024
* [MAJOR UPDATE] -fixed angular run translate i18n script does an amazing job at translating to dest_json

## [1.86.1011] -2-14-2024
* [UPDATE] - update all scripts to give the option of which go executable they want to use

## [1.86.1012] -2-14-2024
* [PATCH] - a few scripts were missing the go executable option

## [1.86.1013] -2-19-2024
* [PATCH] - patch update work with latest tasks to use go instead of windmillcode_go to install everything

## [1.87.0] - 2-29-2024
* [UPDATE] -angular frontend run now supports concurrently with scss so scss with expensive computations only run once

## [1.87.1] - 3-1-2024
* [BREAKING CHANGE] - sql_make_db_schema_update_entry now supports the new format for create schema entries
it will take the path at "apps", "database", databaseToBackup, "schema_entries" and copy whatevr is in the template folder
in that folder to make a new entry
* [UPDATE] angular_frontend_test now tries to look in karma.conf.js for the coverage info and extract the subdir path
from coverageReporter.dir if it cant find it it will run the test regardless and if it cant find karma.conf.js will run the test regardless

## [1.87.1000] - 3-8-2024
* [PATCH] Implemented the shared `ChooseNodePackageManager` function to standardize package manager selection across various Go scripts, replacing redundant code blocks.
* [PATCH] Removed manual package manager prompts in `angular_frontend_analyze/main.go`, `angular_frontend_update_angular/main.go`, `firebase_cloud_run_emulators/main.go`, `npm_install_app_deps/main.go`, and `npm_install_specific_packages/main.go`, now utilizing the shared `ChooseNodePackageManager`.
* [PATCH] Updated the `npm_install_app_deps/main.go` script to handle dependency reinstallation and support `pnpm` along with existing package managers.
* [PATCH] Modified `npm_install_specific_packages/main.go` to use the shared package manager selection and enhanced package installation/uninstallation logic.
* [PATCH] Refined `flask_backend_test/main.go` to utilize a shared setup function, simplifying the test environment setup process.
* [PATCH] Updated `go.mod` to adjust the order of required modules for proper dependency management.
* [PATCH] Introduced changes in `tasks_update_workspace_with_latest_tasks/main.go` and `tasks_update_workspace_without_extension/main.go` to improve the tasks update process with new prompts and conditional logic.

## [1.87.1001] - 3-9-2024
* [UPDATE] Modified `SetupEnvironmentToRunFlaskApp` in `shared/misc.go` to accept an `env` string parameter for different environment setups.
* [UPDATE] Updated calls to `SetupEnvironmentToRunFlaskApp` in `main.go` files within `flask_backend_run`, `flask_backend_run_with_reloader`, and `flask_backend_test` directories to pass environment strings.
* [BUG] Fixed an inconsistency in setting Python version in `SetupEnvironmentToRunFlaskApp` by moving the comment and ensuring the Python version is set after environment variables are set.
* [FIX] Removed the file `old.go` in the `flask_backend_run_with_reloader` directory as part of cleanup or refactoring.
* [PATCH] Adjusted helper script determination logic in `SetupEnvironmentToRunFlaskApp` to select the correct script based on the environment.

## [1.87.1002] - 3-9-2024
* [UPDATE] - Updated go.work.sum and task_files/go_scripts/go.sum to include new dependencies on github.com/ghodss/yaml and gopkg.in/yaml.v2.
* [UPDATE] - Upgraded github.com/windmillcode/go_cli_scripts/v4 to v4.5.0 in task_files/go_scripts/go.mod and go.work.sum.
* [REMOVE] - Removed direct requirements on github.com/fsnotify/fsnotify and github.com/windmillcode/go_cli_scripts/v4 in task_files/go_scripts/go.mod, now marked as indirect dependencies.
* [CHANGE] - Binary difference detected in diff_output.txt without specific details due to the binary nature of the file change.

## [1.87.1003] - 3-10-2024
* [UPDATE] - Updated go.work.sum and task_files/go_scripts/go.mod to use version v4.5.1 of github.com/windmillcode/go_cli_scripts/v4.
* [REMOVE] - Removed references to github.com/ghodss/yaml and gopkg.in/yaml.v2 from go.work.sum and task_files/go_scripts/go.mod.
* [ADD] - Added github.com/tailscale/hujson to go.work.sum and task_files/go_scripts/go.sum as a new dependency.
* [UPDATE] - Updated go.sum to include new checksums for github.com/tailscale/hujson and github.com/windmillcode/go_cli_scripts/v4 at their updated versions.

## [1.87.1004] - 3-10-2024
* [UPDATE] - Formatted struct alignment in task_files/go_scripts/shared/misc.go for Linux command options.
* [UPDATE] - Removed whitespace and adjusted formatting in various functions in task_files/go_scripts/shared/misc.go.
* [REFACTOR] - Enhanced SetupEnvironmentToRunFlaskApp function in task_files/go_scripts/shared/misc.go to better extract and parse environment variables using regex.
* [UPDATE] - Moved the Python version setting logic in SetupEnvironmentToRunFlaskApp to after the environment variables are set.
* [FORMAT] - Adjusted spacing and formatting in various parts of task_files/go_scripts/shared/misc.go for better readability.



## [1.87.2000] - 3-11-2024
* [UPDATE] In `flutter_mobile_create_page/main.go`, updated file naming for template and riverpod provider paths to include `_page` in their names for consistency.
* [UPDATE] In `i18n_script_via_ai/requirements.txt`, added a new dependency `g4f[all]`.

## [1.87.2001] - 3-16-2024
* [UPDATE] Updated import statement in `template_page.dart` to use `template_page_riverpod_provider.dart` instead of `template_riverpod_provider.dart`.

## [1.87.2002] - 3-16-2024
* [UPDATE] Changed the location where the provider is stored in the Flutter application. The provider now resides in the 'lib/shared/widgets' directory instead of the 'lib/pages' directory, reflecting a move to a more appropriate location for shared widget resources.
* [FIX] Modified the path generation logic for new template and RiverPod provider files to align with the updated provider location, ensuring that these resources are correctly stored within the new directory structure.

## [1.87.2003] - 3-17-2024
* [UPDATE] Added new package imports for WMLColors, WMLFonts, WMLNav, and WMLSpacing in Flutter templates (template_page.dart, template_page.dart, template_widget.dart).
* [UPDATE] Integrated WMLColorsRiverpodProvider, WMLFontsRiverpodProvider, WMLNavRiverpodProvider, and WMLSpacingRiverpodProvider into the build context of WMLTemplateWidget in Flutter templates.
* [UPDATE] Added contextHeight and contextWidth variables to capture screen dimensions in the build context of WMLTemplateWidget in Flutter templates.

## [1.87.2004] - 3-27-2024
[UPDATE] Modified file naming in Go scripts for Flutter mobile to include '_riverpod_provider' in the filename, enhancing clarity and consistency in the provider naming convention. Changes made in 'flutter_mobile_create_riverpod_async_notifier_provider/main.go' and 'flutter_mobile_create_riverpod_provider/main.go'.

[PATCH] Adjusted the path construction for template copying in Go scripts, ensuring the correct file names are used during the generation of Riverpod providers in Flutter applications.

## [1.87.2005] - 3-28-2024
[UPDATE] Added a `copyWith` method to `WMLTemplateRiverpodProviderValue` class in various Dart template files to support more flexible state management in Flutter. Changes applied to `template_riverpod_provider.dart` files in `flutter_mobile_create_layout`, `flutter_mobile_create_page`, and `flutter_mobile_create_shared_widget` directories.

## [1.87.2006] - 3-28-2024
[FIX] ensured empty directoires are copied

## [1.87.2100] 03-31-2024

[BREAKING CHANGE]- ignore folder now becomes .windmillcode folder all projects need to rename ignore to .windmillcode
[UPDATE] Introduced a new `Metadata` struct in `misc.go`, refactoring the `Task` and `Input` structures for improved maintainability.
[PATCH] Added a `filterJSONForOwnItems` function in `main.go` to preserve user's original task JSON while merging with the extension's own tasks JSON, ensuring user data integrity during updates.
[BUG] Corrected data merging logic in `main.go` to prevent overwriting user's tasks with extension's tasks, ensuring user's original task configurations are maintained.


## [1.87.2101] 3-31-2024

[UPDATE] Refactored the RunOptions structure in misc.go, extracting it into a separate type for improved code clarity and reuse.
[PATCH] Adjusted the Windows command construction in tasks_update_workspace_with_latest_tasks/main.go to append ".exe" to the executable, ensuring correct execution path on Windows systems.
[COMPLEX MERGE] Enhanced task processing logic in tasks_update_workspace_with_latest_tasks/main.go to dynamically assign runOn settings based on task labels, facilitating more flexible task execution control.

## [1.87.2102] 3-31-2024 8:24:38 PM EST
[UPDATE] We made the 'runOptions' in 'tasks_update_workspace_with_latest_tasks/main.go' smarter. Now, it directly sets 'runOn' to 'folderOpen' when needed, no extra steps. It's like we gave it a shortcut, making things faster and the code cleaner.

## [1.87.2200] 4/11/2024 7:29:00 PM EST

- [UPDATE] Hey there, we just upgraded our utility scripts to use version 5 of `go_cli_scripts`! It's like updating your phone's OS to make sure everything runs smoothly.

- [CHECKPOINT] We've added a cool feature in the `flutter_mobile_build` script. Now, when you're building your app, you can pick an environment and provide a Sentry DSN. It's like telling your app where to send those sneaky bugs!

- [UPDATE] In `flutter_mobile_build`, there's now a neat step that bundles up your debug symbols into a zip file. It's like packing your suitcase before a trip, making sure you've got everything you need.

- [UPDATE] Oh, and when you're done building your Flutter app, the script will open the .aab file

## [1.88.0] [4/12/2024 2:59:00 PM EST]
[UPDATE] In flutter_mobile_build/main.go, we streamlined how environment variables are managed and how the Flutter app bundle is built. Plus, we changed the zipping part to use a simpler command, making the process quicker and less prone to errors.

## [1.88.1] 4/15/2024 10:12:36 PM EST


[COMPLEX MERGE] Major overhaul in flutter mobile build its now flutter mobile build deploy We’ve reworked how we handle deploying to the Play Store. Now, there’s a bunch of steps to get inputs from you (like key files and package names), and depending on your choices, we either zip up some stuff or kick off deployment right from the terminal. Plus, we’ve added multi-threading with sync.WaitGroup to manage tasks like opening directories or deploying packages concurrently. It’s a big mix of updates all in one place, so be sure to test this out.


[NEW FEATURE] We're now supporting automated APK and AAB deployments to Google Play! With functions like deployAppbundle and deployAPK, you can automate uploading directly to various Play Store tracks. Set up your credentials, specify your track, and let it roll out automatically. Just be ready with that keyFile and the right trackName.

[PATCH] Improved error handling across the board, especially with file operations and API interactions. We've tightened up the way we handle failures, like zipping folders or moving files, so expect more robust error messages instead of just crashing out.

## [1.88.2] 4/16/2024 12:42:17 PM EST

[UPDATE] Added options to run 'flutter clean' and './gradlew --refresh-dependencies' in `flutter_mobile_build_deploy/main.go`. Now developers can choose to clean up the build environment or refresh dependencies before deploying to the Play Store. Make sure to set these options to "TRUE" if needed!

[PATCH] Updated the `main()` function in `flutter_mobile_build_deploy/main.go` to include conditional commands based on new menu options. This change means you can better manage your build process directly from the menu prompts without manual tweaks.

## [1.88.1000] 4/24/2024 10:30 AM EST

[UPDATE] update angular_frontend run to remove .nx cache from application to clear cache

## [1.88.1010] 4/24/2024 1:10:00 PM EST

[UPDATE] Updated Go version from 1.21.0 to 1.22.2 in the `go.work` and Go module files. This might require developers to adjust their environment to the new version.

[PATCH] Updated the desired extension version for Go from 1.21.6 to 1.22.2 in `src/installGo.ts`. Developers need to ensure compatibility with the new version in their projects.

[UPDATE] Added a function call to set the Node.js environment based on a predefined version in `task_files/go_scripts/angular_frontend_run/main.go`. Developers using this script should verify that Node.js is configured correctly.

[PATCH] Updated dependency from `github.com/windmillcode/go_cli_scripts/v5` version 5.1.2 to 5.1.3 in `go.mod` and `go.sum`. Developers should update their local dependencies to avoid potential conflicts.

[UPDATE] Improved error handling in Python environment setup in `task_files/go_scripts/shared/misc.go`. Now checks if `pyenv` is installed before attempting to use it, providing a fallback message otherwise.

[NEW FEATURE] Added functionality to prompt for and use a specific Node.js version in `task_files/go_scripts/shared/misc.go`. This affects developers needing to switch Node.js versions frequently.

## [1.88.1012] [5/25/2024 6:57:00 PM EST]

**[UPDATE]**
- **File:** `task_files/go_scripts/firebase_cloud_run_emulators/main.go`
- **Change:** Commented out code for debug mode.
- **Detail:** Added commented code for enabling Firebase debug mode.

**[UPDATE]**
- **File:** `task_files/go_scripts/flutter_mobile_build_deploy/main.go`
- **Change:** Added options to remove unused imports and executed the commands concurrently.
- **Detail:** Added prompts to remove unused imports and used goroutines to run the command for both `lib` and `test` directories.

## [1.88.1014] [5/25/2024 4:32:10 PM EST]

[UPDATE] Added ignore rules for unused local variables and unused catch stack traces in multiple template files: `template_page.dart`, `template_riverpod_provider.dart`, `template_widget.dart`.

Files impacted:
- `flutter_mobile_create_layout/template/template_page.dart`
- `flutter_mobile_create_layout/template/template_riverpod_provider.dart`
- `flutter_mobile_create_page/template/template_page.dart`
- `flutter_mobile_create_page/template/template_riverpod_provider.dart`
- `flutter_mobile_create_riverpod_async_notifier_provider/template/template.dart`
- `flutter_mobile_create_riverpod_provider/template/template.dart`
- `flutter_mobile_create_shared_widget/template/template_riverpod_provider.dart`
- `flutter_mobile_create_shared_widget/template/template_widget.dart`

Changes:
- Ignore unused local variables
- Ignore unused catch stack traces

## [1.88.1015] [5/25/2024 4:40:33 PM EST]

[UPDATE] Renamed classes and variables for better clarity and consistency in multiple template files.

Files impacted:
- `flutter_mobile_create_layout/template/template_page.dart`
- `flutter_mobile_create_layout/template/template_riverpod_provider.dart`
- `flutter_mobile_create_page/template/template_page.dart`
- `flutter_mobile_create_page/template/template_riverpod_provider.dart`
- `flutter_mobile_create_shared_widget/template/template_riverpod_provider.dart`

Changes:
- Renamed `WMLTemplateWidget` to `WMLTemplateLayout`
- Renamed `WMLTemplateRiverpodProviderValue` to `WMLTemplateLayoutRiverpodProviderValue`
- Updated related variables and provider instances to reflect new names
- Adjusted references in the code to use new class and variable names

## [1.88.1016] [5/25/2024 4:40:50 PM EST]

[FIX] Renamed state class names to match widget names in template files: `template_page.dart`, `template_widget.dart`.

Files impacted:
- `flutter_mobile_create_layout/template/template_page.dart`
- `flutter_mobile_create_page/template/template_page.dart`
- `flutter_mobile_create_shared_widget/template/template_widget.dart`

Changes:
- Renamed `ConsumerState` class from `_WMLTemplateState` to `_WMLTemplateLayoutState` in `template_page.dart`.
- Renamed `ConsumerState` class from `_WMLTemplateState` to `_WMLTemplatePageState` in `template_page.dart`.
- Renamed `ConsumerState` class from `_WMLTemplateState` to `_WMLTemplateWidgetState` in `template_widget.dart`.

## [1.88.1017] [5/25/2024 4:45:10 PM EST]

[FIX] Updated provider names in template files: `template_page.dart`, `template_widget.dart`.

Files impacted:
- `flutter_mobile_create_layout/template/template_page.dart`
- `flutter_mobile_create_page/template/template_page.dart`
- `flutter_mobile_create_shared_widget/template/template_widget.dart`

Changes:
- Changed provider from `WMLTemplateRiverpodProvider` to `WMLTemplateLayoutRiverpodProvider` in `template_page.dart`.
- Changed provider from `WMLTemplateRiverpodProvider` to `WMLTemplatePageRiverpodProvider` in `template_page.dart`.
- Changed provider from `WMLTemplateRiverpodProvider` to `WMLTemplateWidgetRiverpodProvider` in `template_widget.dart`.


## [1.88.1018] [5/28/2024 02:15:30 PM EST]

[UPDATE] In `installGo.ts`, the `extensionDesiredVersion` is now `"1.22.3"` instead of `"1.22.2"`. Just a tiny version bump.

[UPDATE] In `main.go`, added `useForce` variable. It's set to `"TRUE"` by default and has a menu prompt to choose between `"TRUE"` and `"FALSE"`.

[FIX] In `main.go`, added `--force` to `updateArgs` if `useForce` is `"TRUE"`. This means the `npx` command will now use `--force` if chosen.

## [1.88.1020] [5/29/2024 12:05:00 PM EST]

[FIX] task_files/go_scripts/tasks_update_workspace_with_latest_tasks/main.go: Fixed error handling in main function to properly check for and handle errors when reading JSON files.

[PATCH] task_files/go_scripts/tasks_update_workspace_with_latest_tasks/main.go: Added error handling for JSON unmarshalling and file operations.

[PATCH] task_files/go_scripts/tasks_update_workspace_with_latest_tasks/main.go: Simplified preActions function for copying files.

[UPDATE] its not necessary to have a tasks.json or settings.json when intializing the extension in your workspace it will create it for you along with the upsert .gitingore so large executables dont end up in github.com or your remote VCS

## [1.88.1021] [6/4/2024 12:05:00 PM EST]

[UPDATE] pushing working to git remote allows you to specify any relative and absoluate paths
in settings.json
```ts
{
  ...
  "windmillcode-extension-pack-0":{
...
    "gitPushingWorkingToGitRemote":{
      "relativePaths":["."],
      "absolutePaths":["C:\\Users\\User\\My_Project"]
    }
  },
  ...
}
```

## [1.88.1022] [6/5/2024 11:42:00 AM EST]

[UPDATE] A new script `main.go` was added to the `android_replace_all_emulators` directory. This script helps you manage Android Virtual Devices (AVDs). You can list, delete, and create AVDs using this script.

[UPDATE] The `tasks.json` file was updated with a new task labeled "android: replace all emulators". This task runs a shell command to replace all Android emulators.

## [1.88.1023] [6/5/2024 10:20:34 AM EST]

**[UPDATE]**
- **File:** `android_replace_all_emulators`
- **Change:** Added `shared.SetJavaEnvironment()` call at the beginning of `main` function.
- **Impact:** Ensures Java environment is set before showing the main menu. Developers must ensure `jvms` is installed.

**[UPDATE]**
- **File:** `misc.go`
- **Function:** `SetJavaEnvironment`
- **Change:** Added new function `SetJavaEnvironment`.
- **Impact:** Developers can now set the Java environment using `jvms` command. They must ensure `jvms` is installed.

## [1.88.1024] [6/5/2024 7:20:34 PM EST]
* [UPDATE] allowed the use to choose jdk via jvms command if installed on the system
also ensured if jvms was not on system to handle appropriately

## [1.88.1025] 6/6/2024 10:32:45 AM EST


[PATCH] added SetJavaEnvironment function in main.go at task_files/go_scripts/flutter_mobile_build_deploy/main.go to set Java environment for building and deploying Flutter mobile apps

## [1.90.1000] [6/19/2024 10:15:00 AM EST]

[FIX]
fixed widget naming issues with  flutter_mobile_create_layout and flutter_mobile_create_page

[UPDATE]
added a feature to translate_json to remove-json-keys to help the script better fully translate the file


## [1.90.1001] [7/5/2024 12:15:30 PM EST]
* [FIX] fixed an issue with docker init container on mac os

## [1.91.1] [7/7/2024 10:45:12 AM EST]

* [UPDATE] for commands that ask to set the language software version you can say skip

## [1.91.2] 7/7/2024 10:23:45 AM EST

[UPDATE]
* notify user that they cany type Skip to not set python and nodejs version

[FIX]
- **Description:** Fixed Windows command to conditionally append `.exe` based on the task when updating the workspace without the extension.
