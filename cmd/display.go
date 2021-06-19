package cmd

import (
	"fmt"

	"github.com/AnVeliz/segyfile/internal/segyparser"
	"github.com/spf13/cobra"
)

const (
	fileFlagName        = "file"
	displayTypeFlagName = "displaytype"
)

var (
	displayCmd = &cobra.Command{
		Use:       "display",
		Short:     "Display information from segy file",
		Long:      `Display some information from a segy file`,
		ValidArgs: []string{"everything"},
		Args:      cobra.ExactValidArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			file, err := cmd.Flags().GetString(fileFlagName)
			if err != nil {
				fmt.Println("Error with flags parsing. Flag file should be specified properly")
				return
			}
			displayType, err := cmd.Flags().GetString(displayTypeFlagName)
			if err != nil {
				fmt.Println("Error with flags parsing. Flag displaytype should be specified properly")
				return
			}
			fmt.Println("display called ", file, displayType)
			segyparser.Parse(file, segyparser.ParseEverything)
		},
	}
)

func init() {
	displayCmd.PersistentFlags().String(fileFlagName, "", "segy file name")
	displayCmd.MarkPersistentFlagRequired(fileFlagName)

	displayCmd.Flags().String(displayTypeFlagName, "", "Display type")
	displayCmd.MarkPersistentFlagRequired(displayTypeFlagName)
}
