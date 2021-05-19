package docs

import (
	"cobra/cmd"
	"log"

	"github.com/spf13/cobra/doc"
)

func genRest() {
	err := doc.GenReSTTree(cmd.RootCmd, "./docs")
	if err != nil {
		log.Fatal(err)
	}
}
