package stage

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/system/theme"
	"realworld/inc/system/widget"
)

type settingsScreen struct{}

var Settings = StageManager.Register(&Stage{
	Name: "settings",
	Title: inkwell.Text{
		{Lang: "en", Text: "SETTINGS"},
		{Lang: "hu", Text: "BEÁLLÍTÁSOK"},
	},
})

var _ = widget.WidgetManager.Register(&inkwell.Menu{
	Name:   "settings_menu",
	When:   inkwell.OnStage(Settings),
	Sticky: true,
	Bounds: window(320, 140),
	Root: inkwell.MenuItem{
		Label: inkwell.Text{
			{Lang: "en", Text: "SETTINGS"},
			{Lang: "hu", Text: "BEÁLLÍTÁSOK"},
		},
		Expand: settingsScreen{}.expand,
	},
})

var _ = widget.WidgetManager.Register(escapeToMenu("settings_escape", Settings))

func (settingsScreen) expand(context *inkwell.Ctx) []inkwell.MenuItem {
	otherLanguage := "hu"
	if inkwell.Lang == "hu" {
		otherLanguage = "en"
	}
	otherTheme := theme.NokiaPunk
	if context.Game.Theme() == theme.NokiaPunk {
		otherTheme = theme.RealWorld
	}
	return []inkwell.MenuItem{
		{
			Label: inkwell.Text{
				{Lang: "en", Text: "Language"},
				{Lang: "hu", Text: "Nyelv"},
			},
			Value:    inkwell.FixedValue(languageNames[inkwell.Lang]),
			Do:       inkwell.UseLang(otherLanguage),
			KeepOpen: true,
		},
		{
			Label: inkwell.Text{
				{Lang: "en", Text: "Theme"},
				{Lang: "hu", Text: "Téma"},
			},
			Value:    inkwell.FixedValue(themeNames[context.Game.Theme()]),
			Do:       inkwell.UseTheme(otherTheme),
			KeepOpen: true,
		},
		backRow(),
	}
}

var languageNames = map[string]inkwell.Text{
	"en": {
		{Lang: "en", Text: "English"},
		{Lang: "hu", Text: "angol"},
	},
	"hu": {
		{Lang: "en", Text: "Hungarian"},
		{Lang: "hu", Text: "magyar"},
	},
}

var themeNames = map[*theme.Theme]inkwell.Text{
	theme.RealWorld: {
		{Lang: "en", Text: "Real World '93"},
		{Lang: "hu", Text: "Real World '93"},
	},
	theme.NokiaPunk: {
		{Lang: "en", Text: "Nokia-punk"},
		{Lang: "hu", Text: "Nokia-punk"},
	},
}
