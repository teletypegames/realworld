package tape

import (
	"image/color"

	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/content/dialog"
)

type Tape struct {
	Name   string
	Item   string
	Label  inkwell.Text
	Color  color.Color
	Dialog inkwell.Dialog
}

var TapeManager = inkwell.NewManager[Tape]()

var Speakers = inkwell.NewManager[inkwell.Speaker]()

func (tape Tape) GetName() string { return tape.Name }

func (tape Tape) talks() bool { return len(tape.Dialog.Nodes) > 0 }

func init() {
	TapeManager.Each(func(tape Tape) {
		Speakers.Register(inkwell.Speaker{
			Name:  tape.Name,
			Label: tape.Label,
			Voice: inkwell.VoiceLog,
			Color: tape.Color,
		})
		if !tape.talks() {
			return
		}
		conversation := tape.Dialog
		conversation.Name = tape.Name
		dialog.DialogManager.Register(conversation)
	})
}

func Of(item string) (Tape, bool) {
	for _, candidate := range TapeManager.All() {
		if candidate.Item != "" && candidate.Item == item {
			return candidate, true
		}
	}
	return Tape{}, false
}

func Dialog(item string) string {
	if loaded, found := Of(item); found && loaded.talks() {
		return loaded.Name
	}
	return inc.Dex
}
