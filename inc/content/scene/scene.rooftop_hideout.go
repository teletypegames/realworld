package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/content/character"
	"realworld/inc/content/item"
	"realworld/inc/content/speaker"
)

type rooftopHideout struct{}

var _ place = rooftopHideout{}

var RooftopHideout = SceneManager.Register(&Scene{
	Name: "rooftop_hideout",
	Title: inkwell.Text{
		{Lang: "en", Text: "NORMAN'S ROOFTOP HIDEOUT"},
		{Lang: "hu", Text: "NORMAN BÚVÓHELYE A TETŐN"},
	},
	Exits: []inkwell.Exit{
		{
			To: normanApartment{},
			Label: inkwell.Text{
				{Lang: "en", Text: "back down into the flat"},
				{Lang: "hu", Text: "vissza le a lakásba"},
			},
			Side: inkwell.ExitNear,
		},
		{
			To: bbsTerminal{},
			Label: inkwell.Text{
				{Lang: "en", Text: "the old terminal"},
				{Lang: "hu", Text: "a régi terminál"},
			},
			Side: inkwell.ExitRight,
		},
	},
	Hotspots: rooftopHideout{}.hotspots(),
})

func (rooftopHideout) hotspots() []inkwell.Hotspot {
	taken := inkwell.Say(character.Paul, inkwell.Text{
		{Lang: "en", Text: "Empty crate. The notes are with me."},
		{Lang: "hu", Text: "Üres láda. A jegyzetek nálam vannak."},
	})
	return []inkwell.Hotspot{
		{
			Name: "crate",
			Label: inkwell.Text{
				{Lang: "en", Text: "crate of papers"},
				{Lang: "hu", Text: "papírokkal teli láda"},
			},
			Area: inkwell.Rect(120, 170, 80, 60),
			OnLook: inkwell.If(inkwell.HasItem(item.NormanNotes),
				taken,
				inkwell.Say(character.Paul, inkwell.Text{
					{Lang: "en", Text: "A fruit crate under a tarp. Notebooks, loose sheets, a pencil worn to the metal."},
					{Lang: "hu", Text: "Gyümölcsösláda ponyva alatt. Füzetek, kitépett lapok, egy fémig kopott ceruza."},
				}),
			),
			OnTake: inkwell.If(inkwell.HasItem(item.NormanNotes),
				taken,
				inkwell.Seq(
					inkwell.Give(item.NormanNotes),
					inkwell.Say(character.Paul, inkwell.Text{
						{Lang: "en", Text: "Norman's handwriting. All of it."},
						{Lang: "hu", Text: "Norman kézírása. Az egész."},
					}),
					inkwell.Say(speaker.Dex, inkwell.Text{
						{Lang: "en", Text: "This is where he came to be nobody's engineer. Read it later, Paul. Not up here."},
						{Lang: "hu", Text: "Ide jött fel, hogy ne legyen senki mérnöke. Később olvasd el, Paul. Ne itt fent."},
					}),
				),
			),
			OnUse: inkwell.If(inkwell.HasItem(item.NormanNotes),
				taken,
				inkwell.Say(character.Paul, inkwell.Text{
					{Lang: "en", Text: "I'll take the lot. It'll get rained on otherwise."},
					{Lang: "hu", Text: "Elviszem az egészet. Különben eláztatja az eső."},
				}),
			),
		},
		{
			Name: "chair",
			Label: inkwell.Text{
				{Lang: "en", Text: "folding chair"},
				{Lang: "hu", Text: "kempingszék"},
			},
			Area: inkwell.Rect(260, 160, 60, 70),
			OnLook: inkwell.Say(character.Paul, inkwell.Text{
				{Lang: "en", Text: "One chair, facing the bridge. Nobody was ever meant to sit next to him up here."},
				{Lang: "hu", Text: "Egy szék, a híd felé fordítva. Ide fent soha nem kellett volna senkinek mellé ülnie."},
			}),
			OnUse: inkwell.Say(speaker.Dex, inkwell.Text{
				{Lang: "en", Text: "Sit if you want. He won't mind. He can't."},
				{Lang: "hu", Text: "Ülj le, ha akarsz. Nem bánja. Nem tudja."},
			}),
		},
		{
			Name: "view",
			Label: inkwell.Text{
				{Lang: "en", Text: "the bay"},
				{Lang: "hu", Text: "az öböl"},
			},
			Area: inkwell.Rect(360, 30, 260, 110),
			OnLook: inkwell.If(inkwell.Flag(inc.FlagLabLocated),
				inkwell.Say(character.Paul, inkwell.Text{
					{Lang: "en", Text: "The Golden Gate, right there. He sat up here every night looking straight at the place they'd put him."},
					{Lang: "hu", Text: "A Golden Gate, pont ott. Minden este itt ült, és egyenesen arra a helyre nézett, ahová aztán bedugták."},
				}),
				inkwell.Say(character.Paul, inkwell.Text{
					{Lang: "en", Text: "Fog, the bridge, and a city that keeps saying nothing happened here."},
					{Lang: "hu", Text: "Köd, a híd, és egy város, ami folyton azt mondja, itt nem történt semmi."},
				}),
			),
		},
	}
}

func (rooftopHideout) Scene() *inkwell.Scene { return RooftopHideout }
