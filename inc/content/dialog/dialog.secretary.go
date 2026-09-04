package dialog

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/content/character"
	"realworld/inc/content/speaker"
)

var Secretary = DialogManager.Register(&Dialog{
	Name:  "secretary",
	Start: "root",
	Nodes: []inkwell.DialogNode{{
		Name: "root",
		Lines: []inkwell.DialogLine{{
			Speaker: character.Secretary,
			Text: inkwell.Text{
				{Lang: "en", Text: "You're late. By two weeks. Sit anywhere, there's plenty of anywhere."},
				{Lang: "hu", Text: "Elkésett. Két héttel. Üljön bárhová, bárhol van bőven."},
			},
		}},
		Choices: []inkwell.DialogChoice{
			{
				Text: inkwell.Text{
					{Lang: "en", Text: "What was the talk about?"},
					{Lang: "hu", Text: "Miről szólt az előadás?"},
				},
				Actions: []inkwell.Action{
					inkwell.Say(character.Secretary, inkwell.Text{
						{Lang: "en", Text: "The Underwater Sun held it. A troupe, if you like. They bring what the BVK doesn't print, and they read it out loud."},
						{Lang: "hu", Text: "A Vízalatti Nap tartotta. Egy társulat, ha úgy tetszik. Elhozzák, amit a BVK nem nyomtat ki, és hangosan felolvassák."},
					}),
					inkwell.Say(character.Secretary, inkwell.Text{
						{Lang: "en", Text: "This time it was about the ACPs. Or rather: about how the ACPs were never the point."},
						{Lang: "hu", Text: "Most az ACP-kről volt szó. Vagyis: arról, hogy az ACP-k sosem voltak a lényeg."},
					}),
					inkwell.GotoNode("point"),
				},
			},
			{
				Text: inkwell.Text{
					{Lang: "en", Text: "I'm looking for a key. Forged here, the saying goes."},
					{Lang: "hu", Text: "Egy kulcsot keresek. Itt kovácsolták, úgy mondják."},
				},
				Actions: []inkwell.Action{
					inkwell.Say(character.Secretary, inkwell.Text{
						{Lang: "en", Text: "A key? We don't forge anything. We fold chairs. Whoever told you that was speaking in pictures."},
						{Lang: "hu", Text: "Kulcsot? Mi nem kovácsolunk semmit. Székeket hajtogatunk. Aki ezt mondta, képekben beszélt."},
					}),
					inkwell.Say(speaker.Dex, inkwell.Text{
						{Lang: "en", Text: "Dead end. Or the key isn't a thing you can hold."},
						{Lang: "hu", Text: "Zsákutca. Vagy a kulcs nem olyasmi, amit kézben lehet tartani."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: inkwell.Text{
					{Lang: "en", Text: "Ever seen a Personal Tape like this one? Unmarked, locked."},
					{Lang: "hu", Text: "Látott már ilyen Személyes Kazettát? Jelöletlen, zárt."},
				},
				Actions: []inkwell.Action{
					inkwell.Say(character.Secretary, inkwell.Text{
						{Lang: "en", Text: "Half the people who come here carry an unmarked tape. The other half carry two. I don't ask. That's the rule."},
						{Lang: "hu", Text: "Aki idejön, annak a fele jelöletlen kazettát hord. A másik fele kettőt. Nem kérdezek. Ez a szabály."},
					}),
					inkwell.Say(character.Secretary, inkwell.Text{
						{Lang: "en", Text: "Yours means nothing to me. Sorry."},
						{Lang: "hu", Text: "A magáé nekem semmit nem mond. Sajnálom."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: inkwell.Text{
					{Lang: "en", Text: "What's the terminal in the corner?"},
					{Lang: "hu", Text: "Mi az a terminál a sarokban?"},
				},
				Actions: []inkwell.Action{
					inkwell.Say(character.Secretary, inkwell.Text{
						{Lang: "en", Text: "Not the BVK's. It looks like LECTOR because it has to. Underneath, it's ECHO. Peer to peer, no registry, no log."},
						{Lang: "hu", Text: "Nem a BVK-é. LECTOR-nak néz ki, mert annak kell kinéznie. Alatta ECHO fut. Peer to peer, nyilvántartás nélkül, log nélkül."},
					}),
					inkwell.Say(character.Secretary, inkwell.Text{
						{Lang: "en", Text: "Read what you like. Nobody here will know what you read, and nobody out there either."},
						{Lang: "hu", Text: "Olvasson, amit akar. Itt senki nem fogja tudni, mit olvasott, és odakint sem."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: inkwell.Text{
					{Lang: "en", Text: "I'll leave you to your chairs."},
					{Lang: "hu", Text: "Hagyom a székeivel."},
				},
				Actions: []inkwell.Action{inkwell.EndDialog()},
			},
		},
	}, {
		Name: "point",
		Lines: []inkwell.DialogLine{
			{
				Speaker: character.Secretary,
				Text: inkwell.Text{
					{Lang: "en", Text: "The BVK doesn't only want the machines quiet. It wants the people quiet. Not silenced. Reprogrammed."},
					{Lang: "hu", Text: "A BVK nem csak a gépeket akarja csendben tudni. Az embereket is. Nem elhallgattatva. Átprogramozva."},
				},
			},
			{
				Speaker: character.Secretary,
				Text: inkwell.Text{
					{Lang: "en", Text: "Their next research goes straight into the head. Keep a person in a dream, a trance, and steer what they think in there. Rewrite it, if need be."},
					{Lang: "hu", Text: "A következő kutatásuk egyenesen a fejbe nyúl. Tartsd az embert álomban, transzban, és irányítsd, mit gondol odabent. Írd át, ha kell."},
				},
			},
			{
				Speaker: character.Secretary,
				Text: inkwell.Text{
					{Lang: "en", Text: "Start with a few politicians. The ones that matter. Then you don't need to hold the ACPs down at all. You hold the hands that hold them."},
					{Lang: "hu", Text: "Kezdd pár politikussal. Akik számítanak. Utána az ACP-ket már nem is kell lefogni. Azokat a kezeket fogod, amelyek őket fogják."},
				},
			},
			{
				Speaker: speaker.Dex,
				Text: inkwell.Text{
					{Lang: "en", Text: "“Brain above the machine.” The brochure wasn't lying. It just didn't say whose hand was on the brain."},
					{Lang: "hu", Text: "„Agy a gép fölött.” A brosúra nem hazudott. Csak azt nem mondta, kinek a keze van az agyon."},
				},
			},
		},
		Choices: []inkwell.DialogChoice{
			{
				Text: inkwell.Text{
					{Lang: "en", Text: "Where's the proof?"},
					{Lang: "hu", Text: "Hol a bizonyíték?"},
				},
				Actions: []inkwell.Action{
					inkwell.Say(character.Secretary, inkwell.Text{
						{Lang: "en", Text: "In the corner. Read it yourself. I only take the minutes."},
						{Lang: "hu", Text: "A sarokban. Olvassa el maga. Én csak a jegyzőkönyvet vezetem."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: inkwell.Text{
					{Lang: "en", Text: "Sounds like a story people tell in basements."},
					{Lang: "hu", Text: "Olyan sztorinak hangzik, amit pincékben mesélnek."},
				},
				Actions: []inkwell.Action{
					inkwell.Say(character.Secretary, inkwell.Text{
						{Lang: "en", Text: "Every true thing was told in a basement first. Ask the terminal. It doesn't have an opinion, only records."},
						{Lang: "hu", Text: "Minden igaz dolgot először pincében mondtak el. Kérdezze a terminált. Annak nincs véleménye, csak feljegyzései."},
					}),
					inkwell.GotoNode("root"),
				},
			},
		},
	}},
})
