package armilla

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/system"
)

var stripReading = inkwell.Text{
	{Lang: "en", Text: "SLOT2: READING"},
	{Lang: "hu", Text: "2. FOGL: OLVAS"},
}

func Slot2() string {
	if value, isText := system.System.Game().State.Var(inc.VarSlot2).(string); isText {
		return value
	}
	return ""
}

func Insert(item string) inkwell.Action {
	return inkwell.Seq(
		inkwell.SetVar(inc.VarSlot2, item),
		inkwell.TakeAway(item),
		inkwell.SetVar(inc.VarArmillaStrip, stripReading),
	)
}

func StripIfEmpty(text inkwell.Text) inkwell.Action {
	return inkwell.SetVarIfEmpty(inc.VarArmillaStrip, text)
}
