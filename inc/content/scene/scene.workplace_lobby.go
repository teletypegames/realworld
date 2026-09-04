package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/content/character"
	"realworld/inc/content/item"
	"realworld/inc/content/script"
)

type workplaceLobby struct{}

var _ place = workplaceLobby{}

var WorkplaceLobby = SceneManager.Register(&Scene{
	Name: "workplace_lobby",
	Title: inkwell.Text{
		{Lang: "en", Text: "TOWER B — LOBBY"},
		{Lang: "hu", Text: "B TORONY — ELŐTÉR"},
	},
	Exits: []inkwell.Exit{
		exitToSelector,
	},
	OnEnter: inkwell.If(inkwell.Not(inkwell.Flag(inc.FlagNoodleReleased)),
		inkwell.RunScript(script.NoodleReleased),
	),
	Hotspots: workplaceLobby{}.hotspots(),
})

func (workplaceLobby) hotspots() []inkwell.Hotspot {
	needCode := inkwell.Say(character.Paul, inkwell.Text{
		{Lang: "en", Text: "Dead panel. It wants a floor code and a card, and it wants both."},
		{Lang: "hu", Text: "Halott panel. Emeletkódot és kártyát kér, és mindkettőt akarja."},
	})
	toActive := inkwell.Seq(
		inkwell.Say(character.Paul, inkwell.Text{
			{Lang: "en", Text: "Norman's code, Norman's card. Up we go."},
			{Lang: "hu", Text: "Norman kódja, Norman kártyája. Fel is megyünk."},
		}),
		inkwell.GoTo(activeOffice{}),
	)
	toAbandoned := inkwell.If(inkwell.Flag(inc.FlagSumphorCard),
		inkwell.Seq(
			inkwell.Say(character.Paul, inkwell.Text{
				{Lang: "en", Text: "Sumphor's card, four-one-nine-one. The lift goes somewhere it isn't supposed to."},
				{Lang: "hu", Text: "Sumphor kártyája, négy-egy-kilenc-egy. A lift oda megy, ahová nem kéne."},
			}),
			inkwell.GoTo(abandonedOffice{}),
		),
		needCode,
	)
	return []inkwell.Hotspot{
		{
			Name: "lift",
			Label: inkwell.Text{
				{Lang: "en", Text: "lift panel"},
				{Lang: "hu", Text: "liftpanel"},
			},
			Area: inkwell.Rect(280, 90, 70, 130),
			OnLook: inkwell.Say(character.Paul, inkwell.Text{
				{Lang: "en", Text: "A keypad and a card reader. No floor buttons at all. You go where you're allowed."},
				{Lang: "hu", Text: "Számbillentyűzet és kártyaolvasó. Emeletgomb sehol. Oda mész, ahová szabad."},
			}),
			OnUse: inkwell.If(inkwell.And(inkwell.Flag(inc.FlagLiftCode), inkwell.HasItem(item.NormanWallet)),
				toActive,
				needCode,
			),
			OnUseWith: map[*inkwell.Item]inkwell.Action{
				item.SumphorCard: toAbandoned,
				item.NormanWallet: inkwell.If(inkwell.Flag(inc.FlagLiftCode),
					toActive,
					inkwell.Say(character.Paul, inkwell.Text{
						{Lang: "en", Text: "The card's fine. I still don't have his floor code."},
						{Lang: "hu", Text: "A kártya jó. Csak az emeletkódja nincs meg."},
					}),
				),
			},
			OnTake: inkwell.Say(character.Paul, inkwell.Text{
				{Lang: "en", Text: "It's a lift. It stays."},
				{Lang: "hu", Text: "Ez egy lift. Marad."},
			}),
		},
		{
			Name: "mailboxes",
			Label: inkwell.Text{
				{Lang: "en", Text: "bank of mailboxes"},
				{Lang: "hu", Text: "postaláda-sor"},
			},
			Area: inkwell.Rect(40, 110, 120, 90),
			OnLook: inkwell.Say(character.Paul, inkwell.Text{
				{Lang: "en", Text: "Half of them taped over. One says REDS, N. in a hand that isn't his."},
				{Lang: "hu", Text: "A felük leragasztva. Az egyiken REDS, N., de nem az ő kézírásával."},
			}),
			OnUse: inkwell.Say(character.Paul, inkwell.Text{
				{Lang: "en", Text: "Empty. Someone cleared it out and wasn't gentle."},
				{Lang: "hu", Text: "Üres. Valaki kiürítette, és nem volt kíméletes."},
			}),
		},
		{
			Name: "plant",
			Label: inkwell.Text{
				{Lang: "en", Text: "lobby plant"},
				{Lang: "hu", Text: "előtéri növény"},
			},
			Area: inkwell.Rect(500, 140, 60, 80),
			OnLook: inkwell.Say(character.Paul, inkwell.Text{
				{Lang: "en", Text: "Plastic. Dusted, though. Somebody still pretends this is a place people arrive at."},
				{Lang: "hu", Text: "Műanyag. Mégis letörölgetve. Valaki még úgy tesz, mintha ide emberek érkeznének."},
			}),
		},
	}
}

func (workplaceLobby) Scene() *inkwell.Scene { return WorkplaceLobby }
