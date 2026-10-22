package armilla

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

var stripReading = inkwell.Text{
	{Lang: "en", Text: "SLOT2: READING"},
	{Lang: "hu", Text: "2. FOGL: OLVAS"},
}

func Slot2(game *inkwell.Game) *inkwell.Item {
	name, isText := game.State.Var(inc.VarSlot2).(string)
	if !isText {
		return nil
	}
	loaded, found := game.ItemManager.Get(name)
	if !found {
		return nil
	}
	return loaded
}

func Insert(cassette *inkwell.Item) inkwell.Action {
	return inkwell.Seq(
		inkwell.SetVar(inc.VarSlot2, cassette.Name),
		inkwell.TakeAway(cassette),
		inkwell.SetVar(inc.VarArmillaStrip, stripReading),
	)
}

func StripIfEmpty(text inkwell.Text) inkwell.Action {
	return inkwell.SetVarIfEmpty(inc.VarArmillaStrip, text)
}
