package docs

import (
	"cobra/cmd"
	"log"

	"github.com/spf13/cobra/doc"
)

func genManPage() {

	header := &doc.GenManHeader{
		Title: "MINE",
		Section: "3",
	}
	err := doc.GenManTree(cmd.RootCmd, header, "./docs")
	if err != nil {
		log.Fatal(err)
	}
}
