package main

import (
	"embed"
	"flag"
	"log"

	inkwell "git.teletypegames.org/engines/inkwell"

	"git.teletypegames.org/games/realworld/internal/boot"
)

// The art ships inside the binary. js/wasm has no OS filesystem, so without
// this every background in the web build falls back to a placeholder; and
// `make binaries` packages the executable alone, with no assets directory
// beside it.
//
//go:embed assets
var assets embed.FS

func init() { inkwell.AssetFS = assets }

func main() {
	var o boot.Opts
	flag.StringVar(&o.Screen, "screen", "", "starting screen (empty: the alley)")
	flag.BoolVar(&o.Finale, "finale", false,
		"start with the finale, to try the Nokia-punk theme switch")
	flag.Parse()

	if err := inkwell.Run(boot.New(o)); err != nil {
		log.Fatal(err)
	}
}
