package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/content/character"
	"realworld/inc/content/item"
	"realworld/inc/content/script"
)

var workplaceLobbyNeedCode = character.Paul.Say(Text{
	{Lang: "en", Text: "Dead panel. It wants a floor code and a card, and it wants both."},
	{Lang: "hu", Text: "Halott panel. Emeletkódot és kártyát kér, és mindkettőt akarja."},
})

var workplaceLobbyToActive = inkwell.Seq(
	character.Paul.Say(Text{
		{Lang: "en", Text: "Norman's code, Norman's card. Up we go."},
		{Lang: "hu", Text: "Norman kódja, Norman kártyája. Fel is megyünk."},
	}),
	inkwell.GoTo(ActiveOffice),
)

var workplaceLobbyToAbandoned = inkwell.If(inkwell.Flag(inc.FlagSumphorCard),
	inkwell.Seq(
		character.Paul.Say(Text{
			{Lang: "en", Text: "Sumphor's card, four-one-nine-one. The lift goes somewhere it isn't supposed to."},
			{Lang: "hu", Text: "Sumphor kártyája, négy-egy-kilenc-egy. A lift oda megy, ahová nem kéne."},
		}),
		inkwell.GoTo(AbandonedOffice),
	),
	workplaceLobbyNeedCode,
)

var WorkplaceLobby = SceneManager.Register(&Scene{
	Name: "workplace_lobby",
	Title: Text{
		{Lang: "en", Text: "TOWER B — LOBBY"},
		{Lang: "hu", Text: "B TORONY — ELŐTÉR"},
	},
	Layers: []SceneLayer{
		{
			Asset: &Asset{
				Name: "workplace_lobby",
				Path: "assets/bg/workplace_lobby.png",
			},
			Z: -1,
		},
	},
	OnEnter: inkwell.If(inkwell.Not(inkwell.Flag(inc.FlagNoodleReleased)),
		inkwell.RunScript(script.NoodleReleased),
	),
	Hotspots: []Hotspot{
		{
			Name: "lift",
			Label: Text{
				{Lang: "en", Text: "lift panel"},
				{Lang: "hu", Text: "liftpanel"},
			},
			Area: inkwell.Rect(280, 90, 70, 130),
			OnLook: character.Paul.Say(Text{
				{Lang: "en", Text: "A keypad and a card reader. No floor buttons at all. You go where you're allowed."},
				{Lang: "hu", Text: "Számbillentyűzet és kártyaolvasó. Emeletgomb sehol. Oda mész, ahová szabad."},
			}),
			OnUse: inkwell.If(inkwell.And(inkwell.Flag(inc.FlagLiftCode), inkwell.HasItem(item.NormanWallet)),
				workplaceLobbyToActive,
				workplaceLobbyNeedCode,
			),
			OnUseWith: map[*inkwell.Item]Action{
				item.SumphorCard: workplaceLobbyToAbandoned,
				item.NormanWallet: inkwell.If(inkwell.Flag(inc.FlagLiftCode),
					workplaceLobbyToActive,
					character.Paul.Say(Text{
						{Lang: "en", Text: "The card's fine. I still don't have his floor code."},
						{Lang: "hu", Text: "A kártya jó. Csak az emeletkódja nincs meg."},
					}),
				),
			},
			OnTake: character.Paul.Say(Text{
				{Lang: "en", Text: "It's a lift. It stays."},
				{Lang: "hu", Text: "Ez egy lift. Marad."},
			}),
		},
		{
			Name: "mailboxes",
			Label: Text{
				{Lang: "en", Text: "bank of mailboxes"},
				{Lang: "hu", Text: "postaláda-sor"},
			},
			Area: inkwell.Rect(40, 110, 120, 90),
			OnLook: character.Paul.Say(Text{
				{Lang: "en", Text: "Half of them taped over. One says REDS, N. in a hand that isn't his."},
				{Lang: "hu", Text: "A felük leragasztva. Az egyiken REDS, N., de nem az ő kézírásával."},
			}),
			OnUse: character.Paul.Say(Text{
				{Lang: "en", Text: "Empty. Someone cleared it out and wasn't gentle."},
				{Lang: "hu", Text: "Üres. Valaki kiürítette, és nem volt kíméletes."},
			}),
		},
		{
			Name: "plant",
			Label: Text{
				{Lang: "en", Text: "lobby plant"},
				{Lang: "hu", Text: "előtéri növény"},
			},
			Area: inkwell.Rect(500, 140, 60, 80),
			OnLook: character.Paul.Say(Text{
				{Lang: "en", Text: "Plastic. Dusted, though. Somebody still pretends this is a place people arrive at."},
				{Lang: "hu", Text: "Műanyag. Mégis letörölgetve. Valaki még úgy tesz, mintha ide emberek érkeznének."},
			}),
		},
	},
})

var _ = WorkplaceLobby.Leads(
	exitToSelector,
)
