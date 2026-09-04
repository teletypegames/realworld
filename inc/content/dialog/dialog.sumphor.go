package dialog

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/content/character"
	"realworld/inc/content/item"
	"realworld/inc/content/speaker"
)

var Sumphor = DialogManager.Register(&Dialog{
	Name:  "sumphor",
	Start: "root",
	Nodes: []inkwell.DialogNode{{
		Name: "root",
		Lines: []inkwell.DialogLine{{
			Speaker: character.Sumphor,
			Text: inkwell.Text{
				{Lang: "en", Text: "You're not from up the hill. You walk like someone who owns his own shoes."},
				{Lang: "hu", Text: "Te nem a dombról vagy. Úgy jársz, mint akinek a saját cipője van."},
			},
		}},
		Choices: []inkwell.DialogChoice{
			{
				Text: inkwell.Text{
					{Lang: "en", Text: "I'm looking for Norman Reds."},
					{Lang: "hu", Text: "Norman Redset keresem."},
				},
				Show: inkwell.Given{
					NotFlag: inc.FlagSumphorTalked,
				},
				Actions: []inkwell.Action{
					inkwell.Say(character.Sumphor, inkwell.Text{
						{Lang: "en", Text: "Norman. Norman, Norman. Why? Who sent you? What do you know about The Project?"},
						{Lang: "hu", Text: "Norman. Norman, Norman. Miért? Ki küldött? Mit tudsz A Projektről?"},
					}),
					inkwell.Say(character.Paul, inkwell.Text{
						{Lang: "en", Text: "Less than you, from the sound of it."},
						{Lang: "hu", Text: "Kevesebbet, mint te, ha jól hallom."},
					}),
					inkwell.Say(character.Sumphor, inkwell.Text{
						{Lang: "en", Text: "We sat in the same room. Two desks. He built the thing and I sat in it. One of us is still sitting in it."},
						{Lang: "hu", Text: "Egy szobában ültünk. Két asztal. Ő megépítette, én beleültem. Egyikünk még mindig benne ül."},
					}),
					inkwell.Say(speaker.Dex, inkwell.Text{
						{Lang: "en", Text: "He isn't mad, Paul. He's precise about something that broke him. That's worse."},
						{Lang: "hu", Text: "Nem őrült, Paul. Pontosan beszél valamiről, ami összetörte. Az rosszabb."},
					}),
					inkwell.SetFlag(inc.FlagSumphorTalked),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: inkwell.Text{
					{Lang: "en", Text: "The tape I carry says it used to be yours."},
					{Lang: "hu", Text: "A kazetta, amit hordok, azt mondja, valaha a tiéd volt."},
				},
				Show: inkwell.Given{
					Flag: inc.FlagTapeSpeaks,
				},
				Once: true,
				Actions: []inkwell.Action{
					inkwell.Say(speaker.TapeMystery, inkwell.Text{
						{Lang: "en", Text: "This is the man. My developer. Look at what they did to him."},
						{Lang: "hu", Text: "Ez az az ember. A fejlesztőm. Nézd meg, mit tettek vele."},
					}),
					inkwell.Say(character.Sumphor, inkwell.Text{
						{Lang: "en", Text: "No. No. I never owned a tape. I never trusted one enough to. That's Norman's voice in there, or something wearing it."},
						{Lang: "hu", Text: "Nem. Nem. Nekem soha nem volt kazettám. Sosem bíztam egyikben sem annyira. Norman hangja van odabent, vagy valami, ami azt viseli."},
					}),
					inkwell.Say(speaker.TapeMystery, inkwell.Text{
						{Lang: "en", Text: "You hear that, Paul? That is the chair talking. It took his memory before it took his job."},
						{Lang: "hu", Text: "Hallod ezt, Paul? A szék beszél belőle. Előbb az emlékeit vitte el, aztán az állását."},
					}),
					inkwell.Say(character.Paul, inkwell.Text{
						{Lang: "en", Text: "One of you is lying and I can't tell which."},
						{Lang: "hu", Text: "Egyikőtök hazudik, és nem tudom, melyik."},
					}),
					inkwell.Say(character.Sumphor, inkwell.Text{
						{Lang: "en", Text: "Doesn't matter. Norman's alive. He's in the big one. Somebody has to get him out."},
						{Lang: "hu", Text: "Nem számít. Norman él. A nagyban van. Valakinek ki kell hoznia."},
					}),
					inkwell.Say(speaker.TapeMystery, inkwell.Text{
						{Lang: "en", Text: "On that, at least, we agree."},
						{Lang: "hu", Text: "Ebben legalább egyetértünk."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: inkwell.Text{
					{Lang: "en", Text: "Where is the abandoned office? The one Norman's VPN goes to."},
					{Lang: "hu", Text: "Hol van az elhagyatott iroda? Amelyikbe Norman VPN-je megy."},
				},
				Show: inkwell.Given{
					Flag:    inc.FlagAbandonedKnown,
					NotFlag: inc.FlagSumphorCard,
				},
				Actions: []inkwell.Action{
					inkwell.Say(character.Sumphor, inkwell.Text{
						{Lang: "en", Text: "You found the tunnel. Good. Then you're the one. Tower B, the lift, but not their floor. Ours."},
						{Lang: "hu", Text: "Megtaláltad az alagutat. Jó. Akkor te vagy az. B torony, a lift, de nem az ő emeletük. A miénk."},
					}),
					inkwell.Say(character.Sumphor, inkwell.Text{
						{Lang: "en", Text: "Take my card. Code is four-one-nine-one. They never revoked it. They never revoke anything, they just stop looking."},
						{Lang: "hu", Text: "Vidd a kártyámat. A kód négy-egy-kilenc-egy. Soha nem vonták vissza. Soha semmit nem vonnak vissza, csak nem néznek oda többet."},
					}),
					inkwell.Give(item.SumphorCard),
					inkwell.SetFlag(inc.FlagSumphorCard),
					inkwell.Say(speaker.Dex, inkwell.Text{
						{Lang: "en", Text: "He handed a stranger his keys. Either he trusts you, or he wants very badly for somebody to go in there instead of him."},
						{Lang: "hu", Text: "Egy idegennek adta a kulcsait. Vagy bízik benned, vagy nagyon akarja, hogy valaki más menjen be oda helyette."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: inkwell.Text{
					{Lang: "en", Text: "I found this tape on your old desk."},
					{Lang: "hu", Text: "Ezt a kazettát a régi asztalodon találtam."},
				},
				Show: inkwell.Given{
					Item: item.SumphorTape,
				},
				Once: true,
				Actions: []inkwell.Action{
					inkwell.Say(character.Sumphor, inkwell.Text{
						{Lang: "en", Text: "Mine? Was it mine? It could be Norman's. He had one just like it. Everybody has one just like it."},
						{Lang: "hu", Text: "Az enyém? Az enyém volt? Lehet, hogy Normané. Neki is pont ilyen volt. Mindenkinek pont ilyen van."},
					}),
					inkwell.Say(character.Paul, inkwell.Text{
						{Lang: "en", Text: "Do you want it back?"},
						{Lang: "hu", Text: "Visszakéred?"},
					}),
					inkwell.Say(character.Sumphor, inkwell.Text{
						{Lang: "en", Text: "Keep it. Whatever remembers that room can remember it for me."},
						{Lang: "hu", Text: "Tartsd meg. Ami emlékszik arra a szobára, emlékezzen rá helyettem."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: inkwell.Text{
					{Lang: "en", Text: "Take care of yourself."},
					{Lang: "hu", Text: "Vigyázz magadra."},
				},
				Actions: []inkwell.Action{
					inkwell.Say(character.Sumphor, inkwell.Text{
						{Lang: "en", Text: "I do. Every day. It's the only project I have left."},
						{Lang: "hu", Text: "Vigyázok. Minden nap. Ez az egyetlen projektem, ami maradt."},
					}),
					inkwell.EndDialog(),
				},
			},
		},
	}},
})
