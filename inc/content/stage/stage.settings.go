package stage

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/system/theme"
	"realworld/inc/system/widget"
)

var Settings = StageManager.Register(&Stage{
	Name: "settings",
	Title: Text{
		{Lang: "en", Text: "SETTINGS"},
		{Lang: "hu", Text: "BEÁLLÍTÁSOK"},
	},
})

var _ = widget.WidgetManager.Register(&inkwell.Menu{
	Name:   "settings_menu",
	When:   inkwell.OnStage(Settings),
	Sticky: true,
	Bounds: window(320, 140),
	Root: MenuItem{
		Label: Text{
			{Lang: "en", Text: "SETTINGS"},
			{Lang: "hu", Text: "BEÁLLÍTÁSOK"},
		},
		Expand: settingsExpand,
	},
})

var _ = widget.WidgetManager.Register(escapeToMenu("settings_escape", Settings))

func settingsExpand(context *inkwell.Ctx) []MenuItem {
	game := context.Game.(*inkwell.Game)
	otherLanguage := "hu"
	if inkwell.Lang == "hu" {
		otherLanguage = "en"
	}
	otherTheme := theme.NokiaPunk
	if game.Theme() == theme.NokiaPunk {
		otherTheme = theme.RealWorld
	}
	return []MenuItem{
		{
			Label: Text{
				{Lang: "en", Text: "Language"},
				{Lang: "hu", Text: "Nyelv"},
			},
			Value:    inkwell.FixedValue(languageNames[inkwell.Lang]),
			Do:       inkwell.UseLang(otherLanguage),
			KeepOpen: true,
		},
		{
			Label: Text{
				{Lang: "en", Text: "Theme"},
				{Lang: "hu", Text: "Téma"},
			},
			Value:    inkwell.FixedValue(themeNames[game.Theme()]),
			Do:       inkwell.UseTheme(otherTheme),
			KeepOpen: true,
		},
		{
			Label: Text{
				{Lang: "en", Text: "Voices"},
				{Lang: "hu", Text: "Szinkron"},
			},
			Value:    inkwell.FixedValue(switchNames[game.VoiceOn()]),
			Do:       inkwell.SetVoice(!game.VoiceOn()),
			KeepOpen: true,
		},
		backRow(),
	}
}

var languageNames = map[string]Text{
	"en": {
		{Lang: "en", Text: "English"},
		{Lang: "hu", Text: "angol"},
	},
	"hu": {
		{Lang: "en", Text: "Hungarian"},
		{Lang: "hu", Text: "magyar"},
	},
}

var switchNames = map[bool]Text{
	true: {
		{Lang: "en", Text: "on"},
		{Lang: "hu", Text: "be"},
	},
	false: {
		{Lang: "en", Text: "off"},
		{Lang: "hu", Text: "ki"},
	},
}

var themeNames = map[*theme.Theme]Text{
	theme.RealWorld: {
		{Lang: "en", Text: "Real World '93"},
		{Lang: "hu", Text: "Real World '93"},
	},
	theme.NokiaPunk: {
		{Lang: "en", Text: "Nokia-punk"},
		{Lang: "hu", Text: "Nokia-punk"},
	},
}
