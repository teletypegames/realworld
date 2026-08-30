package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	WidgetManager.Register(&inkwell.RadialVerbs{
		Name:    "verbs",
		Trigger: inkwell.MouseButtonRight,
		Radius:  58,
		Labels: map[string]string{
			"look": "Look", "use": "Use", "talk": "Talk", "take": "Take",
		},
	})
}
