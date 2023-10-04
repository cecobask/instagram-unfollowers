package root

import (
	"github.com/cecobask/instagram-insights/cmd/followdata"
	"github.com/cecobask/instagram-insights/cmd/information"
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "instagram",
		Short: "Instagram Insights CLI",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
		SilenceUsage: true,
	}
	cmd.AddCommand(information.NewRootCommand())
	cmd.AddCommand(followdata.NewRootCommand())
	return cmd
}
