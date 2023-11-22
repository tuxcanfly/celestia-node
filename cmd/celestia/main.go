package main

import (
	"context"
	"os"

	"github.com/spf13/cobra"

	"github.com/celestiaorg/celestia-node/cmd/celestia/bridge"
	"github.com/celestiaorg/celestia-node/cmd/celestia/full"
	"github.com/celestiaorg/celestia-node/cmd/celestia/light"
)

func init() {
	bridgeCmd := bridge.NewCommand(bridge.WithSubcommands())
	lightCmd := light.NewCommand(bridge.WithSubcommands())
	fullCmd := full.NewCommand(bridge.WithSubcommands())
	rootCmd.AddCommand(
		bridgeCmd,
		lightCmd,
		fullCmd,
		versionCmd,
	)
	rootCmd.SetHelpCommand(&cobra.Command{})
}

func main() {
	err := run()
	if err != nil {
		os.Exit(1)
	}
}

func run() error {
	return rootCmd.ExecuteContext(context.Background())
}

var rootCmd = &cobra.Command{
	Use: "celestia [  bridge  ||  full ||  light  ] [subcommand]",
	Short: `
	    ____      __          __  _
	  / ____/__  / /__  _____/ /_(_)___ _
	 / /   / _ \/ / _ \/ ___/ __/ / __  /
	/ /___/  __/ /  __(__  ) /_/ / /_/ /
	\____/\___/_/\___/____/\__/_/\__,_/
	`,
	Args: cobra.NoArgs,
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: false,
	},
}
