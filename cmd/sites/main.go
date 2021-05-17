package main

import (
	log "github.com/sirupsen/logrus"
	"libs.altipla.consulting/errors"

	"github.com/altipla-sites/cli/internal/commands"
)

func main() {
	if err := commands.CmdRoot.Execute(); err != nil {
		log.Fatal(errors.Stack(err))
	}
}
