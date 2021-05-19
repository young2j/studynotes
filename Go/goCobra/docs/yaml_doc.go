package docs

import (
	"cobra/cmd"
	"log"

	"github.com/spf13/cobra/doc"
)

func genYAML() {
	err := doc.GenYamlTree(cmd.RootCmd, "./docs")
	if err != nil {
		log.Fatal(err)
	}
}
