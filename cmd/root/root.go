package root

import (
	"github.com/spf13/cobra"
)

func MakeCmd(name, description string) *cobra.Command {
	cmd := &cobra.Command{
		Use:  name,
		Long: description,
	}
	return cmd
}
