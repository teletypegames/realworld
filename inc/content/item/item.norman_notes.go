package item

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/content/character"
	"realworld/inc/content/speaker"
)

var NormanNotes = ItemManager.Register(&Item{
	Name: "norman_notes",
	Description: inkwell.Text{
		{Lang: "en", Text: "Norman's notes"},
		{Lang: "hu", Text: "Norman jegyzetei"},
	},
	OnUseSelf: inkwell.Seq(
		inkwell.Say(character.Paul, inkwell.Text{
			{Lang: "en", Text: "“Ran the trance protocol again. It works. I don't know why it works. Somebody will ask me and I won't know.”"},
			{Lang: "hu", Text: "„Újra lefuttattam a transz-protokollt. Működik. Nem tudom, miért működik. Valaki meg fogja kérdezni, és nem fogom tudni.”"},
		}),
		inkwell.Say(character.Paul, inkwell.Text{
			{Lang: "en", Text: "“Asked the tape. The tape knew. The tape always knows. What am I for?”"},
			{Lang: "hu", Text: "„Megkérdeztem a kazettát. A kazetta tudta. A kazetta mindig tudja. Én mire vagyok?”"},
		}),
		inkwell.Say(character.Paul, inkwell.Text{
			{Lang: "en", Text: "Dex?"},
			{Lang: "hu", Text: "Dex?"},
		}),
		inkwell.Say(speaker.Dex, inkwell.Text{
			{Lang: "en", Text: "…"},
			{Lang: "hu", Text: "…"},
		}),
	),
})
