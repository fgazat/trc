package main

import (
	"log"

	"github.com/fgazat/trc/client/tracker"
	"github.com/fgazat/trc/cmd/create"
	"github.com/fgazat/trc/cmd/list"
	"github.com/fgazat/trc/cmd/root"
	"github.com/fgazat/trc/cmd/update"
	"github.com/fgazat/trc/config"
	"github.com/spf13/cobra"
)

func addSubcommands(cmd *cobra.Command, cfg *config.Config, client *tracker.Client) {
	cmd.AddCommand(create.Create(cfg, client))
	cmd.AddCommand(list.List(cfg))
	cmd.AddCommand(update.Update(cfg))
}

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}
	trackerClient := tracker.New(cfg)
	cmd := root.MakeCmd("trc", "Yandex Tracker CLI", cfg)
	addSubcommands(cmd, cfg, trackerClient)
	if err = cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
