package cmd

import (
	"os"

	"github.com/nao1215/posixer/internal/posix"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "list",
	Short: "Displays the list of POSIX commands",
	Long: `Displays the list of POSIX commands.

list subcommand outputs a list of POSIX commands in a table,
indicating whether they are built-in or required, optional.`,
	Run: func(cmd *cobra.Command, args []string) {
		os.Exit(list(cmd, args))
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}

func list(cmd *cobra.Command, args []string) int {

	cmds := posix.Commands()
	data := [][]string{}
	for _, c := range cmds {
		data = append(data, []string{c.Name, c.Type, c.Desc})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Type", "Description"})
	table.SetAutoWrapText(false)

	for _, v := range data {
		table.Append(v)
	}
	table.Render()

	return 0
}
