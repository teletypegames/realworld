package item

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
)

func init() {
	Manager.Register(Item{
		Name: inc.ItemNoodleLetter,
		Description: inkwell.Text{
			{Lang: "en", Text: "Noodle's letter"},
			{Lang: "hu", Text: "Noodle levele"},
		},
		OnUseSelf: inkwell.Seq(
			inkwell.Say(inc.Paul, inkwell.Text{
				{Lang: "en", Text: "“Come out, Paul. Not a phone conversation.” That's all it says."},
				{Lang: "hu", Text: "„Gyere ki, Paul. Ez nem telefontéma.” Ennyi áll benne."},
			}),
			inkwell.Say(inc.Dex, inkwell.Text{
				{Lang: "en", Text: "And now you're here. And he isn't."},
				{Lang: "hu", Text: "És most itt vagy. Ő meg nincs."},
			}),
		),
	})
}
