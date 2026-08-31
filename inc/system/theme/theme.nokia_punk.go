package theme

import (
	"realworld/inc"
	"realworld/inc/util"
)

var (
	Green   = util.RGB(0x3BE86B)
	GreenLo = util.RGB(0x1F7A39)
	npBlack = util.RGB(0x030603)
)

func init() {
	ThemeManager.Register(Theme{
		Name:                    inc.NokiaPunk,
		PanelBG:                 npBlack,
		StatusText:              Green,
		FlashText:               util.RGB(0xD8F06A),
		VerbButtonBG:            util.RGB(0x08140A),
		VerbButtonSelectedBG:    GreenLo,
		VerbButtonText:          Green,
		InventorySlotBG:         util.RGB(0x08140A),
		InventorySlotSelectedBG: GreenLo,
		SpeechBubbleBG:          util.RGBA(0x030603, 0xDC),
		SpeechDefaultText:       Green,
		DialogBG:                util.RGBA(0x030603, 0xF0),
		DialogBorder:            Green,
		DialogChoiceBG:          util.RGB(0x08140A),
		DialogChoiceHover:       util.RGB(0xB4FFC8),
		DialogSpeaker:           util.RGB(0xB4FFC8),
		DialogText:              Green,
		EndCardBG:               util.RGBA(0x000000, 0xFA),
		EndCardText:             Green,
		CursorColor:             Green,
		HotspotOutline:          GreenLo,
		SceneBackdrop:           npBlack,
		TopBarBG:                util.RGB(0x061006),
		TopBarText:              GreenLo,
		TopBarAccent:            Green,
		ChatLogBG:               util.RGB(0x040A04),
		ChatLogPrompt:           GreenLo,
		ChatLogResponse:         Green,
		ChatLogSystem:           util.RGB(0x156030),
		CharacterPanelBG:        util.RGB(0x040A04),
		CharacterPanelBorder:    Green,
		CharacterPanelTitle:     util.RGB(0xB4FFC8),
	})
}
