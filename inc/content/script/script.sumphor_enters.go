package script

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

var _ = ScriptManager.Register(Script{
	Name: inc.ScriptSumphorEnters,
	Actions: inkwell.Seq(
		inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "Someone in the doorway. It's him. He slips past me like I'm not here."},
			{Lang: "hu", Text: "Valaki az ajtóban. Ő az. Elsuhan mellettem, mintha itt sem lennék."},
		}),
		inkwell.Say(inc.Sumphor, inkwell.Text{
			{Lang: "en", Text: "Have to get him out. Have to help him. Maybe from out here I can still reach in and turn it, maybe from out here…"},
			{Lang: "hu", Text: "Ki kell hozni. Segíteni kell rajta. Talán innen kintről még benyúlhatok és elfordíthatom, talán innen kintről…"},
		}),
		inkwell.Say(inc.Dex, inkwell.Text{
			{Lang: "en", Text: "He thinks he can steer the simulation from the outside. File that away, Paul. We may need it, or we may need to fear it."},
			{Lang: "hu", Text: "Azt hiszi, kívülről irányíthatja a szimulációt. Jegyezd meg, Paul. Vagy szükségünk lesz rá, vagy félnünk kell tőle."},
		}),
		inkwell.SetFlag(inc.FlagSumphorSeen),
	),
})
