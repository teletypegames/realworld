package tape

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

type Tape struct {
	Name     string
	Item     string
	Dialogue string
}

var TapeManager = inkwell.NewManager[Tape]()

func (tape Tape) GetName() string { return tape.Name }

func Is(name string) bool { return TapeManager.Has(name) }

func Of(item string) (Tape, bool) {
	for _, candidate := range TapeManager.All() {
		if candidate.Item != "" && candidate.Item == item {
			return candidate, true
		}
	}
	return Tape{}, false
}

func IsItem(item string) bool {
	_, found := Of(item)
	return found
}

func Dialogue(item string) string {
	if loaded, found := Of(item); found && loaded.Dialogue != "" {
		return loaded.Dialogue
	}
	return inc.DialogDex
}
