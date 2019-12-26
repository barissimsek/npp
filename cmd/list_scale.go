package cmd

import (
	"github.com/barissimsek/npp/pkg/scale"
	"github.com/spf13/cobra"
)

// scaleCmd represents the scale command
var listScaleCmd = &cobra.Command{
	Use:   "scale",
	Short: "List all root notes and their major/minor scales",
	Long:  `List all root notes and their major/minor scales`,
	Run: func(cmd *cobra.Command, args []string) {
		scale.List()
	},
}

func init() {
	listCmd.AddCommand(listScaleCmd)
}
