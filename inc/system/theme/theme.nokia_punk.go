package theme

import (
	"realworld/inc"
	"realworld/inc/util"
)

var (
	Green          = util.RGB(0x3BE86B)
	GreenLow       = util.RGB(0x1F7A39)
	nokiaPunkBlack = util.RGB(0x030603)
)

var _ = ThemeManager.Register(Theme{
	Name:                    inc.NokiaPunk,
	PanelBG:                 nokiaPunkBlack,
	StatusText:              Green,
	FlashText:               util.RGB(0xD8F06A),
	VerbButtonBG:            util.RGB(0x08140A),
	VerbButtonSelectedBG:    GreenLow,
	VerbButtonText:          Green,
	InventorySlotBG:         util.RGB(0x08140A),
	InventorySlotSelectedBG: GreenLow,
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
	HotspotOutline:          GreenLow,
	SceneBackdrop:           nokiaPunkBlack,
	TopBarBG:                util.RGB(0x061006),
	TopBarText:              GreenLow,
	TopBarAccent:            Green,
	ChatLogBG:               util.RGB(0x040A04),
	ChatLogPrompt:           GreenLow,
	ChatLogResponse:         Green,
	ChatLogSystem:           util.RGB(0x156030),
	CharacterPanelBG:        util.RGB(0x040A04),
	CharacterPanelBorder:    Green,
	CharacterPanelTitle:     util.RGB(0xB4FFC8),
})
