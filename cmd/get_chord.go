package cmd

import (
	"fmt"

	"github.com/barissimsek/npp/pkg/chord"

	"github.com/spf13/cobra"
)

// chordCmd represents the chord command
var chordCmd = &cobra.Command{
	Use:   "chord",
	Short: "Get triads for a given root note",
	Long:  `Get triads for a given root note`,
	Run: func(cmd *cobra.Command, args []string) {
		r, _ := cmd.Flags().GetString("root")
		t, _ := cmd.Flags().GetString("type")

		fmt.Println(chord.Get(t, r))
	},
}

func init() {
	getCmd.AddCommand(chordCmd)
	chordCmd.Flags().String("root", "C", "Root note")
	chordCmd.Flags().String("type", "maj", "Type. maj or minn")
}
