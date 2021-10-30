package save

import (
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/l-ross/ficsit-toolkit/save"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Print save file info",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("%q\n", saveFile)
		f, err := os.Open(saveFile)
		if err != nil {
			return err
		}
		defer f.Close()

		h, err := save.ParseHeader(f)
		if err != nil {
			return err
		}

		spew.Dump(h)

		return nil
	},
}
