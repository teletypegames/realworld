package tape

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/content/dialog"
	"realworld/inc/content/item"
	"realworld/inc/content/speaker"
)

type Tape struct {
	Name     string
	Item     *item.Item
	Speaker  *speaker.Speaker
	Dialog   *dialog.Dialog
	OnInsert inkwell.Action
}

var TapeManager = inkwell.NewManager[*Tape]()

func (tape Tape) GetName() string { return tape.Name }

func Of(loaded *item.Item) (*Tape, bool) {
	if loaded == nil {
		return nil, false
	}
	for _, candidate := range TapeManager.All() {
		if candidate.Item == loaded {
			return candidate, true
		}
	}
	return nil, false
}

func DialogOf(loaded *item.Item) *dialog.Dialog {
	if card, found := Of(loaded); found && card.Dialog != nil {
		return card.Dialog
	}
	return dialog.Dex
}
