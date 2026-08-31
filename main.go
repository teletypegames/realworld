package main

import (
	"embed"
	"flag"
	"log"

	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/system/boot"
)

//go:embed assets
var assets embed.FS

func init() { inkwell.AssetFS = assets }

func main() {
	var o boot.Opts
	flag.StringVar(&o.Scene, "scene", "", "starting scene (empty: the alley)")
	flag.BoolVar(&o.Finale, "finale", false,
		"start with the finale, to try the Nokia-punk theme switch")
	flag.StringVar(&o.Lang, "lang", "en", "language the game speaks (en, hu)")
	flag.Parse()

	if err := inkwell.Run(boot.New(o)); err != nil {
		log.Fatal(err)
	}
}
