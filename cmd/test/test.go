package test

import (
	"fmt"
	"log"

	"github.com/cloudsteak/cs-cli/pkg/utils"
	"github.com/spf13/cobra"
)

type options struct {
	sanity bool
}

func NewCmd() *cobra.Command {
	o := &options{}
	cmd := &cobra.Command{
		Use:   "test",
		Short: "Test Cli",
		Long:  `Command for test cli`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(o)
		},
	}

	cmd.Flags().BoolVarP(&o.sanity, "sanity", "s", false, "Sanity test")

	return cmd
}

func run(o *options) error {
	utils.SetLogFlags()
	if o.sanity {
		log.Println("Sanity test is running...")
		fmt.Println("Sanity test is running...")
	} else {
		log.Println("Sanity test is not required")
		fmt.Println("Sanity test is not required")
	}
	return nil
}
