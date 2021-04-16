package cmd

import (
	"fmt"

	"github.com/lolPants/redts/cli/src/pkg/version"
	"github.com/spf13/cobra"
)

type versionRow struct {
	label string
	value string
}

var (
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print version information",
		Run: func(cmd *cobra.Command, args []string) {
			printVersionInfo()
		},
	}
)

func printVersionInfo() {
	versionRows := make([]versionRow, 0)
	addRow := func(l, v string) {
		row := versionRow{label: l, value: v}
		versionRows = append(versionRows, row)
	}

	addRow("Version", version.AppVersion())
	addRow("Git Hash", version.CommitHash())
	addRow("Go Version", version.GoVersion())

	var widest int
	for _, r := range versionRows {
		width := len(r.label) + 2
		if width > widest {
			widest = width
		}
	}

	for _, r := range versionRows {
		fmt.Printf("%*s %s\n", widest*-1, r.label, r.value)
	}
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
