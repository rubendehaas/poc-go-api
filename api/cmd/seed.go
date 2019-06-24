package cmd

import (
	"app/database/seeds"

	"github.com/spf13/cobra"
)

var seedCmd = &cobra.Command{
	Use:   "api:seed",
	Short: "Seed, will provide sample data to your database",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		seeds.Run()
	},
}
