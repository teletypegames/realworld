package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/content/character"
	"realworld/inc/content/item"
	"realworld/inc/content/script"
	"realworld/inc/content/speaker"
)

type activeOffice struct{}

var _ place = activeOffice{}

var ActiveOffice = SceneManager.Register(&Scene{
	Name: "active_office",
	Title: inkwell.Text{
		{Lang: "en", Text: "TOWER B — ACTIVE FLOOR"},
		{Lang: "hu", Text: "B TORONY — AKTÍV EMELET"},
	},
	Exits: []inkwell.Exit{
		{
			To: workplaceLobby{},
			Label: inkwell.Text{
				{Lang: "en", Text: "the lift down"},
				{Lang: "hu", Text: "a lift lefelé"},
			},
			Side: inkwell.ExitNear,
		},
	},
	Hotspots: activeOffice{}.hotspots(),
})

func (activeOffice) hotspots() []inkwell.Hotspot {
	admitted := inkwell.Say(character.Paul, inkwell.Text{
		{Lang: "en", Text: "Front desk. Nobody's manning it. Nobody's manning anything."},
		{Lang: "hu", Text: "Recepció. Senki nem ül ott. Senki nem ül sehol."},
	})
	blocked := inkwell.Say(character.Paul, inkwell.Text{
		{Lang: "en", Text: "A kid glances up. “Deliveries only past this point.” Then back to his screen."},
		{Lang: "hu", Text: "Egy kölyök felnéz. „Innentől csak kiszállítás.” Aztán vissza a képernyőre."},
	})
	taken := inkwell.Say(character.Paul, inkwell.Text{
		{Lang: "en", Text: "Already lifted their cards. They still haven't noticed."},
		{Lang: "hu", Text: "Már elemeltem a kártyáikat. Még mindig nem vették észre."},
	})
	normanDeskBlocked := inkwell.Say(character.Paul, inkwell.Text{
		{Lang: "en", Text: "Not while the kid's watching. Deliveries only, remember."},
		{Lang: "hu", Text: "Nem, amíg a kölyök figyel. Csak kiszállítás, emlékszel."},
	})
	normanDeskAdmitted := inkwell.If(inkwell.Flag(inc.FlagOfficeLogin),
		inkwell.Say(character.Paul, inkwell.Text{
			{Lang: "en", Text: "I have the login already. It's for the machine at his flat."},
			{Lang: "hu", Text: "A belépő már megvan. A lakásán lévő géphez való."},
		}),
		inkwell.RunScript(script.OfficeLogin),
	)
	return []inkwell.Hotspot{
		{
			Name: "desk",
			Label: inkwell.Text{
				{Lang: "en", Text: "reception desk"},
				{Lang: "hu", Text: "recepciós pult"},
			},
			Area:   inkwell.Rect(60, 120, 110, 90),
			OnLook: inkwell.If(inkwell.Flag(inc.FlagOfficeAdmitted), admitted, blocked),
			OnTalk: inkwell.If(inkwell.Flag(inc.FlagOfficeAdmitted),
				inkwell.Say(character.Paul, inkwell.Text{
					{Lang: "en", Text: "We're past the talking. I'm in."},
					{Lang: "hu", Text: "Túl vagyunk a beszéden. Bent vagyok."},
				}),
				blocked,
			),
			OnUse: inkwell.If(inkwell.Flag(inc.FlagOfficeAdmitted), admitted, blocked),
			OnUseWith: map[*inkwell.Item]inkwell.Action{
				item.HotDogOrder: inkwell.Seq(
					inkwell.Say(character.Paul, inkwell.Text{
						{Lang: "en", Text: "Hot dogs for the floor. Six with everything."},
						{Lang: "hu", Text: "Hotdog az emeletnek. Hat mindennel."},
					}),
					inkwell.Say(speaker.Dex, inkwell.Text{
						{Lang: "en", Text: "He waved you in without looking up. That's the whole security model."},
						{Lang: "hu", Text: "Fel se nézett, úgy intett be. Ennyi az egész biztonsági modell."},
					}),
					inkwell.TakeAway(item.HotDogOrder),
					inkwell.SetFlag(inc.FlagOfficeAdmitted),
				),
			},
		},
		{
			Name: "cubicle",
			Label: inkwell.Text{
				{Lang: "en", Text: "empty cubicle"},
				{Lang: "hu", Text: "üres bokszasztal"},
			},
			Area: inkwell.Rect(200, 120, 100, 90),
			OnLook: inkwell.If(inkwell.HasItem(item.OfficeCards),
				taken,
				inkwell.Say(character.Paul, inkwell.Text{
					{Lang: "en", Text: "Lunch break. Six staff cards on the divider, on a lanyard, like nobody ever leaves."},
					{Lang: "hu", Text: "Ebédszünet. Hat dolgozói kártya a válaszfalon, egy szalagon, mintha soha senki nem menne el."},
				}),
			),
			OnTake: inkwell.If(inkwell.HasItem(item.OfficeCards),
				taken,
				inkwell.Seq(
					inkwell.Give(item.OfficeCards),
					inkwell.Say(character.Paul, inkwell.Text{
						{Lang: "en", Text: "Borrowing these. For the hot dogs."},
						{Lang: "hu", Text: "Kölcsönveszem. A hotdoghoz."},
					}),
				),
			),
			OnUse: inkwell.If(inkwell.HasItem(item.OfficeCards),
				taken,
				inkwell.Say(character.Paul, inkwell.Text{
					{Lang: "en", Text: "The cards, then. Quick."},
					{Lang: "hu", Text: "Akkor a kártyák. Gyorsan."},
				}),
			),
		},
		{
			Name: "norman_desk",
			Label: inkwell.Text{
				{Lang: "en", Text: "Norman's desk"},
				{Lang: "hu", Text: "Norman asztala"},
			},
			Area: inkwell.Rect(360, 120, 100, 90),
			OnLook: inkwell.If(inkwell.Flag(inc.FlagOfficeAdmitted),
				inkwell.Say(character.Paul, inkwell.Text{
					{Lang: "en", Text: "Cleared, mostly. A monitor, a dead plant, and a coffee ring nobody wiped."},
					{Lang: "hu", Text: "Nagyrészt leürítve. Egy monitor, egy halott növény, és egy kávéfolt, amit senki nem törölt le."},
				}),
				normanDeskBlocked,
			),
			OnUse: inkwell.If(inkwell.Flag(inc.FlagOfficeAdmitted), normanDeskAdmitted, normanDeskBlocked),
		},
		{
			Name: "whiteboard",
			Label: inkwell.Text{
				{Lang: "en", Text: "whiteboard"},
				{Lang: "hu", Text: "tábla"},
			},
			Area: inkwell.Rect(500, 50, 110, 70),
			OnLook: inkwell.Say(character.Paul, inkwell.Text{
				{Lang: "en", Text: "“Q3 DELIVERABLES.” “SYNERGY.” “ASK NORMAN.” The last one is half rubbed out."},
				{Lang: "hu", Text: "„Q3 EREDMÉNYEK.” „SZINERGIA.” „KÉRDEZD NORMANT.” Az utolsó félig letörölve."},
			}),
		},
	}
}

func (activeOffice) Scene() *inkwell.Scene { return ActiveOffice }
