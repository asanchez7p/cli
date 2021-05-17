package commands

import (
	"os"

	"github.com/spf13/cobra"
)

const autocompleteHelp = `To load completion run

. <(sites autocomplete)

To configure your bash shell to load completions for each session add to your bashrc

# ~/.bashrc or ~/.profile
. <(sites autocomplete)
`

var cmdAutocomplete = &cobra.Command{
	Use:   "autocomplete",
	Short: "Generates bash completion scripts",
	Long:  autocompleteHelp,
	Run: func(cmd *cobra.Command, args []string) {
		CmdRoot.GenBashCompletion(os.Stdout)
	},
}
