package docs

import (
	"cobra/cmd"
	"log"

	"github.com/spf13/cobra/doc"
)

func genMarkdown() {
	err := doc.GenMarkdownTree(cmd.RootCmd, "./docs")
	if err != nil {
		log.Fatal(err)
	}
}
