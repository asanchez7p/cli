package commands

import (
	"github.com/spf13/cobra"
)

func init() {
	CmdRoot.AddCommand(cmdAutocomplete)
	CmdRoot.AddCommand(cmdClone)
}

var CmdRoot = &cobra.Command{
	Use: "sites",
}
