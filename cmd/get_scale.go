package cmd

import (
	"fmt"

	"github.com/barissimsek/npp/pkg/scale"

	"github.com/spf13/cobra"
)

// scaleCmd represents the scale command
var scaleCmd = &cobra.Command{
	Use:   "scale",
	Short: "Get scale for a given root note",
	Long:  `Get scale for a given root note`,
	Run: func(cmd *cobra.Command, args []string) {
		r, _ := cmd.Flags().GetString("root")
		t, _ := cmd.Flags().GetString("type")

		fmt.Println(scale.Get(t, r))
	},
}

func init() {
	getCmd.AddCommand(scaleCmd)
	scaleCmd.Flags().String("root", "C", "Root note")
	scaleCmd.Flags().String("type", "maj", "Type. maj or minn")
}
