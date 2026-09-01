package stage

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/system/widget"
)

type settingsScreen struct{}

var _ = StageManager.Register(Stage{
	Name: inc.StageSettings,
	Title: inkwell.Text{
		{Lang: "en", Text: "SETTINGS"},
		{Lang: "hu", Text: "BEÁLLÍTÁSOK"},
	},
})

var _ = widget.WidgetManager.Register(&inkwell.Menu{
	Name:   "settings_menu",
	When:   inkwell.OnStage(inc.StageSettings),
	Sticky: true,
	Bounds: window(320, 140),
	Root: inkwell.MenuItem{
		Label: inkwell.Text{
			{Lang: "en", Text: "SETTINGS"},
			{Lang: "hu", Text: "BEÁLLÍTÁSOK"},
		},
		Expand: settingsScreen{}.rows,
	},
})

var _ = widget.WidgetManager.Register(escapeToMenu("settings_escape", inc.StageSettings))

func (settings settingsScreen) rows(context *inkwell.Ctx) []inkwell.MenuItem {
	return []inkwell.MenuItem{
		{
			Label: inkwell.Text{
				{Lang: "en", Text: "Language"},
				{Lang: "hu", Text: "Nyelv"},
			},
			Value:    inkwell.FixedValue(languageNames[inkwell.Lang]),
			Do:       inkwell.UseLang(settings.otherLanguage()),
			KeepOpen: true,
		},
		{
			Label: inkwell.Text{
				{Lang: "en", Text: "Theme"},
				{Lang: "hu", Text: "Téma"},
			},
			Value:    inkwell.FixedValue(themeNames[context.Game.Theme().Name]),
			Do:       inkwell.UseTheme(settings.otherTheme(context)),
			KeepOpen: true,
		},
		backRow(),
	}
}

func (settingsScreen) otherLanguage() string {
	if inkwell.Lang == "hu" {
		return "en"
	}
	return "hu"
}

func (settingsScreen) otherTheme(context *inkwell.Ctx) string {
	if context.Game.Theme().Name == inc.NokiaPunk {
		return inc.RealWorld
	}
	return inc.NokiaPunk
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

var themeNames = map[string]inkwell.Text{
	inc.RealWorld: {
		{Lang: "en", Text: "Real World '93"},
		{Lang: "hu", Text: "Real World '93"},
	},
	inc.NokiaPunk: {
		{Lang: "en", Text: "Nokia-punk"},
		{Lang: "hu", Text: "Nokia-punk"},
	},
}
