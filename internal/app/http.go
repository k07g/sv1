package app

import (
	"github.com/k07g/sv1/internal/pkg/infrastructure/http"
	"github.com/spf13/cobra"
)

func NewHTTPCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "http",
		Short: "CLI HTTP server",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.HelpFunc()(cmd, args)
			}
		},
	}
	cmd.AddCommand(&cobra.Command{
		Use:   "run",
		Short: "Run HTTP server",
		Run: func(cmd *cobra.Command, args []string) {
			http.Run()
		},
	})
	return cmd
}
