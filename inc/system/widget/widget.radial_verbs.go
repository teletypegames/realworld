package widget

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	WidgetManager.Register(&inkwell.RadialVerbs{
		Name:    "verbs",
		Trigger: inkwell.MouseButtonRight,
		Radius:  58,
		Labels: map[string]inkwell.Text{
			"look": {
				{Lang: "en", Text: "Look"},
				{Lang: "hu", Text: "Nézd"},
			},
			"use": {
				{Lang: "en", Text: "Use"},
				{Lang: "hu", Text: "Húzd"},
			},
			"talk": {
				{Lang: "en", Text: "Talk"},
				{Lang: "hu", Text: "Szólj"},
			},
			"take": {
				{Lang: "en", Text: "Take"},
				{Lang: "hu", Text: "Vedd"},
			},
		},
	})
}
