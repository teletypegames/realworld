//go:build steam

package main

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/engines/inkwell/steam"
)

var platformExtensions = []inkwell.Extension{
	&steam.Steam{
		AppID: 480,
		Languages: map[string]string{
			"english":   "en",
			"hungarian": "hu",
		},
		Cloud: []string{inkwell.SlotsKey},
	},
}
