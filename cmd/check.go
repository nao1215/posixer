package cmd

import (
	"os"
	"os/exec"

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
	commands := posix.Commands()
	data := [][]string{}
	for _, c := range commands {
		path, err := exec.LookPath(c.Name)
		if err != nil {
			if posix.IsBuiltInCommands(c) {
				data = append(data, []string{c.Name, c.Type, "installed", "in " + shell.CurrentShell()})
			} else {
				data = append(data, []string{c.Name, c.Type, "not installed", "-"})
			}
		} else {
			data = append(data, []string{c.Name, c.Type, "installed", path})
		}
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Type", "In your system", "Path"})
	table.SetAutoWrapText(false)

	for _, v := range data {
		table.Append(v)
	}
	table.Render()

	return 0
}
