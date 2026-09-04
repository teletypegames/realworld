package item

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/content/character"
	"realworld/inc/content/speaker"
)

var NoodleLetter = ItemManager.Register(&Item{
	Name: "noodle_letter",
	Description: Text{
		{Lang: "en", Text: "Noodle's letter"},
		{Lang: "hu", Text: "Noodle levele"},
	},
	OnUseSelf: inkwell.Seq(
		character.Paul.Say(Text{
			{Lang: "en", Text: "“Come out, Paul. Not a phone conversation.” That's all it says."},
			{Lang: "hu", Text: "„Gyere ki, Paul. Ez nem telefontéma.” Ennyi áll benne."},
		}),
		speaker.Dex.Say(Text{
			{Lang: "en", Text: "And now you're here. And he isn't."},
			{Lang: "hu", Text: "És most itt vagy. Ő meg nincs."},
		}),
	),
})
