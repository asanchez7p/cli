package commands

import (
	"github.com/spf13/cobra"
)

func init() {
	CmdRoot.AddCommand(cmdAutocomplete)
	CmdRoot.AddCommand(cmdClone)
	cmdClone.Flags().StringVarP(&org, "org", "o", "", "Organizacion de Github (requerido)")
	cmdClone.MarkFlagRequired("org")
	cmdClone.Flags().StringVarP(&pattern, "pattern", "p", "", "Prefijo de los proyectos a clonar (requerido)")
	cmdClone.MarkFlagRequired("pattern")
}

var CmdRoot = &cobra.Command{
	Use: "sites",
}
