package theme

import (
	"git.teletypegames.org/games/realworld/inc"
)

var (
	Green   = RGB(0x3BE86B)
	GreenLo = RGB(0x1F7A39)
	npBlack = RGB(0x030603)
)

func init() {
	Manager.Register(Theme{
		Name:                    inc.NokiaPunk,
		PanelBG:                 npBlack,
		StatusText:              Green,
		FlashText:               RGB(0xD8F06A),
		VerbButtonBG:            RGB(0x08140A),
		VerbButtonSelectedBG:    GreenLo,
		VerbButtonText:          Green,
		InventorySlotBG:         RGB(0x08140A),
		InventorySlotSelectedBG: GreenLo,
		SpeechBubbleBG:          RGBA(0x030603, 0xDC),
		SpeechDefaultText:       Green,
		DialogBG:                RGBA(0x030603, 0xF0),
		DialogBorder:            Green,
		DialogChoiceBG:          RGB(0x08140A),
		DialogChoiceHover:       RGB(0xB4FFC8),
		DialogSpeaker:           RGB(0xB4FFC8),
		DialogText:              Green,
		EndCardBG:               RGBA(0x000000, 0xFA),
		EndCardText:             Green,
		CursorColor:             Green,
		HotspotOutline:          GreenLo,
		SceneBackdrop:           npBlack,
		TopBarBG:                RGB(0x061006),
		TopBarText:              GreenLo,
		TopBarAccent:            Green,
		ChatLogBG:               RGB(0x040A04),
		ChatLogPrompt:           GreenLo,
		ChatLogResponse:         Green,
		ChatLogSystem:           RGB(0x156030),
		CharacterPanelBG:        RGB(0x040A04),
		CharacterPanelBorder:    Green,
		CharacterPanelTitle:     RGB(0xB4FFC8),
	})
}
