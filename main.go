package main

import (
	"embed"
	"flag"
	"log"

	inkwell "git.teletypegames.org/engines/inkwell"

	"git.teletypegames.org/games/realworld/inc"
)

//go:embed assets
var assets embed.FS

func init() { inkwell.AssetFS = assets }

func main() {
	var o inc.Opts
	flag.StringVar(&o.Screen, "screen", "", "starting screen (empty: the alley)")
	flag.BoolVar(&o.Finale, "finale", false,
		"start with the finale, to try the Nokia-punk theme switch")
	flag.Parse()

	if err := inkwell.Run(inc.New(o)); err != nil {
		log.Fatal(err)
	}
}
