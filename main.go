package main

import (
	"log"
	"log/slog"

	"github.com/fgazat/trc/cmd/create"
	"github.com/fgazat/trc/cmd/list"
	"github.com/fgazat/trc/cmd/root"
	"github.com/fgazat/trc/cmd/update"
	"github.com/fgazat/trc/config"
	"github.com/fgazat/trc/internal/client/tracker"
	"github.com/spf13/cobra"
)

func addSubcommands(cmd *cobra.Command, cfg *config.Config, client *tracker.Client) {
	cmd.AddCommand(create.Create(cfg, client))
	cmd.AddCommand(list.List(cfg, client))
	cmd.AddCommand(update.Update(cfg))
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()
	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}
	initLogger(cfg.Debug)

	trackerClient := tracker.New(cfg)
	cmd := root.MakeCmd("trc", "Yandex Tracker CLI", cfg)
	addSubcommands(cmd, cfg, trackerClient)
	if err = cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func initLogger(debug bool) {
	if debug {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}
	log.SetFlags(0)
	log.SetPrefix("")
	// tbd: better logging
}
