package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/windmillcode/go_cli_scripts/v6/utils"
)

func main() {

	utils.CDToLocation(strings.TrimSpace(os.Args[1]))
	workspaceRoot, err := os.Getwd()
	if err != nil {
		return
	}
	settings, err := utils.GetSettingsJSON(workspaceRoot)
	if err != nil {
		return
	}

	dbPort := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt:  []string{"the port to the db on your local computer"},
			Default:  strconv.Itoa(settings.ExtensionPack.Ports.Postgres0),
		},
	)
	cliInfo := utils.ShowMenuModel{
		Prompt:  "use history or restore from dump",
		Choices: []string{"DUMP", "HISTORY"},
		Default: "DUMP",
	}
	historyOrDump := utils.ShowMenu(cliInfo, nil)
	var restoreAction string
	var schemaOnly string
	if historyOrDump == "DUMP" {
		cliInfo := utils.ShowMenuModel{
			Prompt:  "SQL OR BINARY RESTORE",
			Choices: []string{ "SQL","BINARY"},
			Default: "SQL",
		}
		restoreAction = utils.ShowMenu(cliInfo, nil)
		cliInfo = utils.ShowMenuModel{
			Prompt:  "schema only",
			Choices: []string{"FALSE", "TRUE"},
			Default: "FALSE",
		}
		schemaOnly = utils.ShowMenu(cliInfo, nil)

	}
	init_sql_path := utils.JoinAndConvertPathToOSFormat(workspaceRoot, "apps", "database", "postgres", "init_db_conn.sql")
	setup_pg_cron_path := utils.JoinAndConvertPathToOSFormat(workspaceRoot, "apps", "database", "postgres", "setup_pg_cron.sh")
	give_postgres_user_acces_to_pg_cron_path := utils.JoinAndConvertPathToOSFormat(workspaceRoot, "apps", "database", "postgres", "give_postgres_user_acces_to_pg_cron.sql")
	schema_entries := utils.JoinAndConvertPathToOSFormat(workspaceRoot, "apps", "database", "postgres", "schema_entries")
	restore_db_script := utils.JoinAndConvertPathToOSFormat(workspaceRoot, ".windmillcode", "Local", "restore_db.sh")
	restore_dump_file := utils.JoinAndConvertPathToOSFormat(workspaceRoot, ".windmillcode", "Local", "backup_file.dump")
	restore_sql_file := utils.JoinAndConvertPathToOSFormat(workspaceRoot, ".windmillcode", "Local", "backup_file.sql")
	dockerPort := dbPort + ":5432"
	utils.RunCommand("docker", []string{"stop", settings.ExtensionPack.SQLDockerContainerName})
	utils.RunCommand("docker", []string{"rm", settings.ExtensionPack.SQLDockerContainerName})
	// utils.RunCommand("docker",[]string{ "run", "--name", settings.ExtensionPack.SQLDockerContainerName, "-e", "POSTGRES_PASSWORD=mysecretpassword","-e", "PGPASSWORD=mysecretpassword", "-p", dockerPort, "-d", "windmillcode/postgres-with-pg-cron-debian:16.1.0000"})
	// utils.RunCommand("docker",[]string{ "run", "--name", settings.ExtensionPack.SQLDockerContainerName, "-e", "POSTGRES_PASSWORD=mysecretpassword","-e", "PGPASSWORD=mysecretpassword", "-p", dockerPort, "-d", "ramazanpolat/postgres_cron:11"})
	utils.RunCommand("docker", []string{"run", "--name", settings.ExtensionPack.SQLDockerContainerName, "-e", "POSTGRES_PASSWORD=mysecretpassword", "-e", "PGPASSWORD=mysecretpassword", "-p", dockerPort, "-d", "postgres:latest"})
	utils.RunCommand("docker", []string{"cp", init_sql_path, settings.ExtensionPack.SQLDockerContainerName + ":/root/"})
	utils.RunCommand("docker", []string{"cp", give_postgres_user_acces_to_pg_cron_path, settings.ExtensionPack.SQLDockerContainerName + ":/root/"})
	utils.RunCommand("docker", []string{"cp", setup_pg_cron_path, settings.ExtensionPack.SQLDockerContainerName + ":/usr/bin/"})
	// utils.RunCommand("docker",[]string{"cp",schema_entries, settings.ExtensionPack.SQLDockerContainerName+":/root/"})
	time.Sleep(10 * time.Second)
	utils.RunCommand("docker", []string{"exec",
		"--workdir", "/root",
		settings.ExtensionPack.SQLDockerContainerName,
		"psql",
		"-U", "postgres",
		"-d", "postgres", "-a", "-f",
		"init_db_conn.sql"},
	)

	// just have this run becuase it not capable of restoring the db on its own
	// if historyOrDump == "HISTORY" {
		if true {
		pattern := `\d{2}-\d{1}-\d{2}_\d{2}-\d{2}-\d{2}`
		utils.ProcessFoldersMatchingPattern(schema_entries, pattern, func(path string) {
			updatePath := utils.JoinAndConvertPathToOSFormat(path, "update.sql")
			fmt.Println(updatePath)
			utils.RunCommand("docker", []string{"cp", updatePath, settings.ExtensionPack.SQLDockerContainerName + ":/root/"})
			utils.RunCommand("docker", []string{"exec", "--workdir", "/root", settings.ExtensionPack.SQLDockerContainerName, "psql", "-U", "postgres", "-d", settings.ExtensionPack.DatabaseName, "-a",  "-f", "update.sql"})
		})

	}
	if historyOrDump == "DUMP" {
		utils.RunCommand("docker", []string{"cp", restore_db_script, settings.ExtensionPack.SQLDockerContainerName + ":/usr/bin/"})
		utils.RunCommand("docker", []string{"cp", restore_dump_file, settings.ExtensionPack.SQLDockerContainerName + ":/root/"})
		utils.RunCommand("docker", []string{"cp", restore_sql_file, settings.ExtensionPack.SQLDockerContainerName + ":/root/"})
		utils.RunCommand("docker", []string{"exec", "--workdir", "/",
			settings.ExtensionPack.SQLDockerContainerName,
			"chmod", "+x", "/usr/bin/restore_db.sh"},
		)
		if restoreAction == "SQL" {
			utils.RunCommand("docker", []string{"exec", "--workdir", "/",
				settings.ExtensionPack.SQLDockerContainerName,
				"bash", "-c", fmt.Sprintf("/usr/bin/restore_db.sh %s %s %s", "SQL", schemaOnly, "/root/backup_file.sql")},
			)
		} else {
			utils.RunCommand("docker", []string{"exec", "--workdir", "/",
				settings.ExtensionPack.SQLDockerContainerName,
				"bash", "-c", fmt.Sprintf("/usr/bin/restore_db.sh %s %s %s", "BINARY", schemaOnly, "/root/backup_file.dump")},
			)
		}
	}

	utils.RunCommand("docker", []string{"exec", "--workdir", "/",
		settings.ExtensionPack.SQLDockerContainerName,
		"ls", "-l", "/usr/bin/setup_pg_cron.sh"},
	)
	utils.RunCommand("docker", []string{"exec", "--workdir", "/",
		settings.ExtensionPack.SQLDockerContainerName,
		"chmod", "+x", "/usr/bin/setup_pg_cron.sh"},
	)
	utils.RunCommand("docker", []string{"exec", "--workdir", "/",
		settings.ExtensionPack.SQLDockerContainerName,
		"bash", "-c", "/usr/bin/setup_pg_cron.sh"},
	)

	utils.RunCommand("docker", []string{"exec", "--workdir", "/root",
		settings.ExtensionPack.SQLDockerContainerName,
		"psql",
		"-U", "postgres",
		"-d", "postgres", "-a", "-f",
		"give_postgres_user_acces_to_pg_cron.sql"},
	)

	utils.RunCommand("docker", []string{"stop", settings.ExtensionPack.SQLDockerContainerName})
	utils.RunCommand("docker", []string{"start", settings.ExtensionPack.SQLDockerContainerName})

}
