package inc

type Tape struct {
	Name     string
	Item     string
	Dialogue string
}

func (t Tape) GetName() string { return t.Name }

func IsTape(name string) bool { return TapeManager.Has(name) }

func TapeOf(item string) (Tape, bool) {
	for _, t := range TapeManager.All() {
		if t.Item != "" && t.Item == item {
			return t, true
		}
	}
	return Tape{}, false
}

func IsTapeItem(item string) bool {
	_, ok := TapeOf(item)
	return ok
}

func TapeDialogue(item string) string {
	if t, ok := TapeOf(item); ok && t.Dialogue != "" {
		return t.Dialogue
	}
	return DlgDex
}
