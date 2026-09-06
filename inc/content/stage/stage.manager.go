package stage

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"github.com/hajimehoshi/ebiten/v2"

	"realworld/inc/system/widget"
)

type (
	Stage    = inkwell.Stage
	Text     = inkwell.Text
	MenuItem = inkwell.MenuItem
	Asset    = inkwell.Asset
)

var StageManager = inkwell.NewManager[*Stage]()

func window(width, height float64) inkwell.Rectangle {
	return inkwell.Rect(
		(widget.ScreenWidth-width)/2,
		(widget.ScreenHeight-height)/2,
		width,
		height,
	)
}

var splashTransition = inkwell.Fade(0.5)

const splashHold = 2.0

var back = Text{
	{Lang: "en", Text: "Back"},
	{Lang: "hu", Text: "Vissza"},
}

type fixed bool

func (value fixed) Eval(*inkwell.Ctx) bool { return bool(value) }

var slotName = Text{
	{Lang: "en", Text: "Slot %d"},
	{Lang: "hu", Text: "%d. hely"},
}

var slotEmpty = Text{
	{Lang: "en", Text: "empty"},
	{Lang: "hu", Text: "üres"},
}

func slotValue(info inkwell.SlotInfo) Text {
	if !info.Exists {
		return slotEmpty
	}
	return Text{
		{Lang: "en", Text: info.SavedAt.Format("2006-01-02 15:04")},
		{Lang: "hu", Text: info.SavedAt.Format("2006-01-02 15:04")},
	}
}

func backRow() MenuItem {
	return MenuItem{
		Label: back,
		Do:    inkwell.GoToStage(MainMenu),
	}
}

func escapeToMenu(name string, stage *Stage) *inkwell.KeyBinding {
	return &inkwell.KeyBinding{
		Name: name,
		When: inkwell.OnStage(stage),
		Key:  ebiten.KeyEscape,
		Do:   inkwell.GoToStage(MainMenu),
	}
}
