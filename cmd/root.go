package cmd

import (
	"fmt"
	"log"

	"github.com/cloudsteak/cs-cli/cmd/test"
	"github.com/cloudsteak/cs-cli/pkg/utils"
	"github.com/spf13/cobra"
)

var appVersion bool

var rootCmd = &cobra.Command{
	Use:           "cs-cli",
	Short:         "CLI used on cloud resources",
	Long:          `CLI for interacting with several hyperscalers`,
	SilenceErrors: false,
	SilenceUsage:  false,
	Run: func(cmd *cobra.Command, args []string) {
		if appVersion {
			log.Printf("Version number: %s\n", "0.0.1")
			fmt.Printf("Version number: %s\n", "0.0.1")
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		utils.SetLogFlags()
		log.Fatal(err)
	}
}

func init() {
	rootCmd.AddCommand(test.NewCmd())
	rootCmd.Flags().BoolVarP(&appVersion, "version", "v", false, "Get version number")
}
