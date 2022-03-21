package cmd

import (
	"os"

	"github.com/nao1215/posixer/internal/posix"
	"github.com/nao1215/posixer/internal/shell"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check whether the POSIX command is installed on your system",
	Long:  `Check whether the POSIX command is installed on your system`,
	Run: func(cmd *cobra.Command, args []string) {
		os.Exit(check(cmd, args))
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}

func check(cmd *cobra.Command, args []string) int {
	cmds := posix.Commands()
	data := [][]string{}
	for _, c := range cmds {
		if posix.IsBuiltInCmds(c) {
			data = append(data, []string{c.Name, c.Type, "installed"})
		} else {
			if shell.ExistCmd(c.Name) {
				data = append(data, []string{c.Name, c.Type, "installed"})
			} else {
				data = append(data, []string{c.Name, c.Type, "not installed"})
			}
		}
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Type", "In your system"})
	table.SetAutoWrapText(false)

	for _, v := range data {
		table.Append(v)
	}
	table.Render()

	return 0
}
