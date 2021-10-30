package save

import (
	"github.com/spf13/cobra"
)

var saveFile string

var SaveCmd = &cobra.Command{
	Use:   "save",
	Short: "",
}

func init() {
	SaveCmd.PersistentFlags().StringP("save", "s", "", "Save file")

	SaveCmd.AddCommand(infoCmd)
}
