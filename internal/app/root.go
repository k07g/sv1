package app

import "github.com/spf13/cobra"

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "app",
		Short: "CLI tool for app",
		Long:  `App is a CLI tool`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.HelpFunc()(cmd, args)
			}
		},
	}
	return cmd
}
