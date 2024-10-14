package main

import (
	"log"

	"github.com/fgazat/trc/cmd/create"
	"github.com/fgazat/trc/cmd/list"
	"github.com/fgazat/trc/cmd/root"
	"github.com/fgazat/trc/cmd/update"
	"github.com/spf13/cobra"
)

func registerSubcommands(cmd *cobra.Command) {
	cmd.AddCommand(create.Create())
	cmd.AddCommand(list.List())
	cmd.AddCommand(update.Update())
}

func main() {
	cmd := root.MakeCmd("trc", "Yandex Tracker CLI")
	registerSubcommands(cmd)
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
