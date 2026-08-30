package theme

import (
	"git.teletypegames.org/games/realworld/inc"
)

var (
	Ink     = RGB(0xDCD5C4)
	InkDim  = RGB(0x8A806B)
	Amber   = RGB(0xE0A33E)
	AmberLo = RGB(0xA8701A)
	black   = RGB(0x14120E)
	panel   = RGB(0x0D0B08)
	sceneBG = RGB(0x241F18)
)

func init() {
	Manager.Register(Theme{
		Name:                    inc.RealWorld,
		PanelBG:                 black,
		StatusText:              Ink,
		FlashText:               Amber,
		VerbButtonBG:            RGB(0x1E1A14),
		VerbButtonSelectedBG:    AmberLo,
		VerbButtonText:          Ink,
		InventorySlotBG:         RGB(0x1E1A14),
		InventorySlotSelectedBG: AmberLo,
		SpeechBubbleBG:          RGBA(0x14120E, 0xDC),
		SpeechDefaultText:       Ink,
		DialogBG:                RGBA(0x0D0B08, 0xF0),
		DialogBorder:            AmberLo,
		DialogChoiceBG:          RGB(0x1E1A14),
		DialogChoiceHover:       Amber,
		DialogSpeaker:           Amber,
		DialogText:              Ink,
		EndCardBG:               RGBA(0x000000, 0xFA),
		EndCardText:             Ink,
		CursorColor:             Amber,
		HotspotOutline:          RGB(0x7A6A44),
		SceneBackdrop:           sceneBG,
		TopBarBG:                panel,
		TopBarText:              InkDim,
		TopBarAccent:            Amber,
		ChatLogBG:               panel,
		ChatLogPrompt:           InkDim,
		ChatLogResponse:         Amber,
		ChatLogSystem:           RGB(0x6B6353),
		CharacterPanelBG:        panel,
		CharacterPanelBorder:    AmberLo,
		CharacterPanelTitle:     Amber,
	})
}
