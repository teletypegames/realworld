// Package theme holds the game's two colour themes and its colour tokens.
//
// In canon the machines of this world are matte black with a green or amber
// character display, and in the Real World finale "the Nokia-punk texture
// breaks in" — with inkwell's runtime theme switching that beat is a single
// line (see world.UseTheme and script.awakeningFinale).
//
// The colour rule: tapes speak in amber (green in Nokia-punk), the world and
// Paul speak in bone white. If something is amber, a tape said it.
package theme

import (
	"image/color"

	inkwell "git.teletypegames.org/engines/inkwell"
)

// The two theme names. Runtime switching (inkwell UseTheme) refers to these.
const (
	RealWorld = "realworld-93"
	NokiaPunk = "nokia-punk"
)

// RGB builds an opaque colour from a hex literal.
func RGB(hex uint32) color.Color {
	return color.RGBA{R: uint8(hex >> 16), G: uint8(hex >> 8), B: uint8(hex), A: 0xff}
}

// RGBA builds a colour with the given alpha from a hex literal.
func RGBA(hex uint32, a uint8) color.Color {
	return color.RGBA{R: uint8(hex >> 16), G: uint8(hex >> 8), B: uint8(hex), A: a}
}

// realworld-93 — 95% of the game. Muted earth tones, black HUD, amber tape
// channel. The Neumatronic layer is not visible yet.
// The realworld-93 palette. Ink and Amber are the two sides of the two-channel
// colour rule: the world speaks in bone white, tapes speak in amber.
var (
	Ink     = RGB(0xDCD5C4) // bone white — the voice of the world and of Paul
	InkDim  = RGB(0x8A806B)
	Amber   = RGB(0xE0A33E) // tape voice
	AmberLo = RGB(0xA8701A)
	black   = RGB(0x14120E)
	panel   = RGB(0x0D0B08)
	sceneBG = RGB(0x241F18)
)

func realWorldTheme() inkwell.Theme {
	return inkwell.Theme{
		Name:       RealWorld,
		PanelBG:    black,
		StatusText: Ink,
		FlashText:  Amber,

		VerbButtonBG:         RGB(0x1E1A14),
		VerbButtonSelectedBG: AmberLo,
		VerbButtonText:       Ink,

		InventorySlotBG:         RGB(0x1E1A14),
		InventorySlotSelectedBG: AmberLo,

		SpeechBubbleBG:    RGBA(0x14120E, 0xDC),
		SpeechDefaultText: Ink,

		DialogBG:          RGBA(0x0D0B08, 0xF0),
		DialogBorder:      AmberLo,
		DialogChoiceBG:    RGB(0x1E1A14),
		DialogChoiceHover: Amber,
		DialogSpeaker:     Amber,
		DialogText:        Ink,

		EndCardBG:   RGBA(0x000000, 0xFA),
		EndCardText: Ink,

		CursorColor:    Amber,
		HotspotOutline: RGB(0x7A6A44),
		SceneBackdrop:  sceneBG,

		TopBarBG:     panel,
		TopBarText:   InkDim,
		TopBarAccent: Amber,

		ChatLogBG:       panel,
		ChatLogPrompt:   InkDim, // your own actions — quiet
		ChatLogResponse: Amber,  // tape voice
		ChatLogSystem:   RGB(0x6B6353),

		CharacterPanelBG:     panel,
		CharacterPanelBorder: AmberLo,
		CharacterPanelTitle:  Amber,
	}
}

// nokia-punk — at peak decay, on BBS/ECHO terminals, and in the finale.
// Matte black and phosphor green: a tuned variant of the engine's
// terminal-green preset.
// The nokia-punk palette — phosphor green on black glass.
var (
	Green   = RGB(0x3BE86B)
	GreenLo = RGB(0x1F7A39)
	npBlack = RGB(0x030603)
)

func nokiaPunkTheme() inkwell.Theme {
	return inkwell.Theme{
		Name:       NokiaPunk,
		PanelBG:    npBlack,
		StatusText: Green,
		FlashText:  RGB(0xD8F06A),

		VerbButtonBG:         RGB(0x08140A),
		VerbButtonSelectedBG: GreenLo,
		VerbButtonText:       Green,

		InventorySlotBG:         RGB(0x08140A),
		InventorySlotSelectedBG: GreenLo,

		SpeechBubbleBG:    RGBA(0x030603, 0xDC),
		SpeechDefaultText: Green,

		DialogBG:          RGBA(0x030603, 0xF0),
		DialogBorder:      Green,
		DialogChoiceBG:    RGB(0x08140A),
		DialogChoiceHover: RGB(0xB4FFC8),
		DialogSpeaker:     RGB(0xB4FFC8),
		DialogText:        Green,

		EndCardBG:   RGBA(0x000000, 0xFA),
		EndCardText: Green,

		CursorColor:    Green,
		HotspotOutline: GreenLo,
		SceneBackdrop:  npBlack,

		TopBarBG:     RGB(0x061006),
		TopBarText:   GreenLo,
		TopBarAccent: Green,

		ChatLogBG:       RGB(0x040A04),
		ChatLogPrompt:   GreenLo,
		ChatLogResponse: Green,
		ChatLogSystem:   RGB(0x156030),

		CharacterPanelBG:     RGB(0x040A04),
		CharacterPanelBorder: Green,
		CharacterPanelTitle:  RGB(0xB4FFC8),
	}
}

// Register adds both themes to the game.
func Register(g *inkwell.Game) {
	g.ThemeManager.Register(realWorldTheme())
	g.ThemeManager.Register(nokiaPunkTheme())
}
