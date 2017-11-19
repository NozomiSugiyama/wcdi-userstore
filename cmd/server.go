package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wcdi/wcdi-userstore/server"
)

func init() {
	RootCmd.AddCommand(newServerCommand())
}

func newServerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "server command",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmd.AddCommand(
		newStartServerCmd(),
	)

	return cmd
}

func newStartServerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "server start !",
		RunE:  runStartServerCmd,
	}

	return cmd
}

func runStartServerCmd(cmd *cobra.Command, args []string) error {
	return server.StartServer()
}
