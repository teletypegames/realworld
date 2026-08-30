package tape

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
)

type Tape struct {
	Name     string
	Item     string
	Dialogue string
}

var Manager = inkwell.NewManager[Tape]()

func (t Tape) GetName() string { return t.Name }

func Is(name string) bool { return Manager.Has(name) }

func Of(item string) (Tape, bool) {
	for _, t := range Manager.All() {
		if t.Item != "" && t.Item == item {
			return t, true
		}
	}
	return Tape{}, false
}

func IsItem(item string) bool {
	_, ok := Of(item)
	return ok
}

func Dialogue(item string) string {
	if t, ok := Of(item); ok && t.Dialogue != "" {
		return t.Dialogue
	}
	return inc.DlgDex
}
