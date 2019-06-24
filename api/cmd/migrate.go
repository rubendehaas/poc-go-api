package cmd

import (
	"app/database/migrations"

	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate, will setup your database",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		migrations.Run()
	},
}
