package main

import (
	"flag"
	"log"

	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/internal/boot"
)

func main() {
	var o boot.Opts
	flag.StringVar(&o.Scene, "scene", "", "starting scene (empty: the alley)")
	flag.BoolVar(&o.Finale, "finale", false,
		"start with the finale, to try the Nokia-punk theme switch")
	flag.Parse()

	if err := inkwell.Run(boot.New(o)); err != nil {
		log.Fatal(err)
	}
}
