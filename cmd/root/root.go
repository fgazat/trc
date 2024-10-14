package root

import (
	"github.com/fgazat/trc/config"
	"github.com/spf13/cobra"
)

func MakeCmd(name, description string, cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:  name,
		Long: description,
	}
	return cmd
}
