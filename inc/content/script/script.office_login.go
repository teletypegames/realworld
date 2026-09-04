package script

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/content/character"
	"realworld/inc/content/speaker"
)

var OfficeLogin = ScriptManager.Register(&Script{
	Name: "office_login",
	Actions: inkwell.Seq(
		inkwell.Say(character.Paul, inkwell.Text{
			{Lang: "en", Text: "Under the keyboard. A sticky note, a login on it. “home box — don't lose this.”"},
			{Lang: "hu", Text: "A billentyűzet alatt. Egy cetli, rajta egy belépő. „otthoni gép — ezt el ne veszítsd.”"},
		}),
		inkwell.Say(speaker.Dex, inkwell.Text{
			{Lang: "en", Text: "The machine at his flat. This is the login the VPN was asking for."},
			{Lang: "hu", Text: "A lakásán lévő gép. Ez az a belépő, amit a VPN kért."},
		}),
		inkwell.SetFlag(inc.FlagOfficeLogin),
		inkwell.If(inkwell.Flag(inc.FlagVPNKnown),
			inkwell.Seq(
				inkwell.Say(character.Paul, inkwell.Text{
					{Lang: "en", Text: "Which means I can finish what his session started. The tunnel goes to an office nobody uses any more."},
					{Lang: "hu", Text: "Ami azt jelenti, hogy befejezhetem, amit a kapcsolata elkezdett. Az alagút egy irodába megy, amit már senki nem használ."},
				}),
				inkwell.SetFlag(inc.FlagAbandonedKnown),
			),
		),
	),
})
