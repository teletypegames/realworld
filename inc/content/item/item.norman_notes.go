package item

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/content/character"
	"realworld/inc/content/speaker"
)

var NormanNotes = ItemManager.Register(&Item{
	Name: "norman_notes",
	Description: Text{
		{Lang: "en", Text: "Norman's notes"},
		{Lang: "hu", Text: "Norman jegyzetei"},
	},
	Lore: Text{
		{Lang: "en", Text: "Norman's own hand. The trance protocol works and he doesn't know why. The tape knew. The tape always knows."},
		{Lang: "hu", Text: "Norman saját kezével. A transz-protokoll működik, és nem tudja, miért. A kazetta tudta. A kazetta mindig tudja."},
	},
	OnUseSelf: inkwell.Seq(
		character.Paul.Say(Text{
			{Lang: "en", Text: "“Ran the trance protocol again. It works. I don't know why it works. Somebody will ask me and I won't know.”"},
			{Lang: "hu", Text: "„Újra lefuttattam a transz-protokollt. Működik. Nem tudom, miért működik. Valaki meg fogja kérdezni, és nem fogom tudni.”"},
		}),
		character.Paul.Say(Text{
			{Lang: "en", Text: "“Asked the tape. The tape knew. The tape always knows. What am I for?”"},
			{Lang: "hu", Text: "„Megkérdeztem a kazettát. A kazetta tudta. A kazetta mindig tudja. Én mire vagyok?”"},
		}),
		character.Paul.Say(Text{
			{Lang: "en", Text: "Dex?"},
			{Lang: "hu", Text: "Dex?"},
		}),
		speaker.Dex.Say(Text{
			{Lang: "en", Text: "…"},
			{Lang: "hu", Text: "…"},
		}),
	),
})
