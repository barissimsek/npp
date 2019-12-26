package cmd

import (
	"github.com/barissimsek/npp/pkg/chord"
	"github.com/spf13/cobra"
)

// chordCmd represents the chord command
var listChordCmd = &cobra.Command{
	Use:   "chord",
	Short: "List all chords and its triads",
	Long:  `List all chords and its triads`,
	Run: func(cmd *cobra.Command, args []string) {
		chord.List()
	},
}

func init() {
	listCmd.AddCommand(listChordCmd)
}
