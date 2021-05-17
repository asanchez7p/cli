package commands

import (
	"github.com/spf13/cobra"
)

var cmdClone = &cobra.Command{
	Use:     "clone",
	Short:   "Clona uno o varios proyectos en local",
	Example: "sites clone",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
