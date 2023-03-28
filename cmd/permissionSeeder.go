package cmd

import (
	"github.com/krobus00/auth-service/internal/bootstrap"
	"github.com/spf13/cobra"
)

// permissionSeederCmd represents the server command.
var permissionSeederCmd = &cobra.Command{
	Use:   "permission-seeder",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		bootstrap.StartPermissionSeeder()
	},
}

func init() {
	rootCmd.AddCommand(permissionSeederCmd)
}
