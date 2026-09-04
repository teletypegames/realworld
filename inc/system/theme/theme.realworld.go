package theme

import (
	"realworld/inc/util"
)

var (
	Ink             = util.RGB(0xDCD5C4)
	InkDim          = util.RGB(0x8A806B)
	Amber           = util.RGB(0xE0A33E)
	AmberLow        = util.RGB(0xA8701A)
	AmberPale       = util.RGB(0xE8C86A)
	Bone            = util.RGB(0xB9A98A)
	black           = util.RGB(0x14120E)
	panel           = util.RGB(0x0D0B08)
	sceneBackground = util.RGB(0x241F18)
)

var RealWorld = ThemeManager.Register(&Theme{
	Name:                    "realworld-93",
	PanelBG:                 black,
	StatusText:              Ink,
	FlashText:               Amber,
	VerbButtonBG:            util.RGB(0x1E1A14),
	VerbButtonSelectedBG:    AmberLow,
	VerbButtonText:          Ink,
	InventorySlotBG:         util.RGB(0x1E1A14),
	InventorySlotSelectedBG: AmberLow,
	SpeechBubbleBG:          util.RGBA(0x14120E, 0xDC),
	SpeechDefaultText:       Ink,
	DialogBG:                util.RGBA(0x0D0B08, 0xF0),
	DialogBorder:            AmberLow,
	DialogChoiceBG:          util.RGB(0x1E1A14),
	DialogChoiceHover:       Amber,
	DialogSpeaker:           Amber,
	DialogText:              Ink,
	EndCardBG:               util.RGBA(0x000000, 0xFA),
	EndCardText:             Ink,
	CursorColor:             Amber,
	HotspotOutline:          util.RGB(0x7A6A44),
	SceneBackdrop:           sceneBackground,
	TopBarBG:                panel,
	TopBarText:              InkDim,
	TopBarAccent:            Amber,
	ChatLogBG:               panel,
	ChatLogPrompt:           InkDim,
	ChatLogResponse:         Amber,
	ChatLogSystem:           util.RGB(0x6B6353),
	CharacterPanelBG:        panel,
	CharacterPanelBorder:    AmberLow,
	CharacterPanelTitle:     Amber,
})
