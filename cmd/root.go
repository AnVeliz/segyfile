package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "segyfile",
	Short: "The app works with segy files",
	Long:  `The application is aimed to work with segy files`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.AddCommand(displayCmd)
}
