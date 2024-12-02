package main

import (
	"fmt"
	"os"

	"github.com/windmillcode/go_cli_scripts/v6/utils"
	// "os"
	// "path/filepath"
)

func main() {
	// scriptLocation := os.Args[1]
	workspaceRoot := os.Args[2]


	envVars := fmt.Sprintf(`
	<ENVVARS>
	</ENVVARS>
	`,)
	fmt.Println(envVars)

}
