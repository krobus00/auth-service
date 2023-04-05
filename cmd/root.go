package cmd

import (
	"fmt"
	"os"

	"github.com/krobus00/auth-service/internal/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "auth-service",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func Init() {
	if err := config.LoadConfig(); err != nil {
		logrus.Fatalln(err.Error())
	}
	logrus.Info(fmt.Sprintf("starting %s:%s...", config.ServiceName(), config.ServiceVersion()))
}
