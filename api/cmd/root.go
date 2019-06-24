package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sankaku",
	Short: "Project root command",
	Long:  ``,
}

func init() {
	rootCmd.AddCommand(migrateCmd)
	rootCmd.AddCommand(seedCmd)
}

func Execute() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
