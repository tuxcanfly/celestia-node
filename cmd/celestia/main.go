package main

import (
	"context"
	"os"

	"github.com/celestiaorg/celestia-node/cmd/celestia/bridge"
	"github.com/celestiaorg/celestia-node/cmd/celestia/full"
	"github.com/celestiaorg/celestia-node/cmd/celestia/light"
	"github.com/spf13/cobra"
)

var bridgeCmd = bridge.NewBridgeCmd()
var lightCmd = light.NewLightCmd()
var fullCmd = full.NewFullCmd()

func init() {
	bridgeCmd.AddCommand(bridge.DefaultCommands(bridge.DefaultFlags)...)
	lightCmd.AddCommand(light.DefaultCommands(bridge.DefaultFlags)...)
	fullCmd.AddCommand(full.DefaultCommands(full.DefaultFlags)...)
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
