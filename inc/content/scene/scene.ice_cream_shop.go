package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/content/character"
	"realworld/inc/content/dialog"
	"realworld/inc/content/item"
	"realworld/inc/content/script"
	"realworld/inc/content/speaker"
	"realworld/inc/system"
)

type iceCreamShop struct{}

var _ place = iceCreamShop{}

var IceCreamShop = SceneManager.Register(&Scene{
	Name: "ice_cream_shop",
	Title: inkwell.Text{
		{Lang: "en", Text: "ICE CREAM SHOP — WHERE THE BAR WAS"},
		{Lang: "hu", Text: "FAGYIZÓ — A RÉGI BÁR HELYÉN"},
	},
	Exits: []inkwell.Exit{
		exitToSelector,
	},
	Actors: []inkwell.SceneActor{
		{
			Character: character.Paul,
			At: inkwell.Point{
				X: 120,
				Y: 236,
			},
		},
		{
			Character: character.IceCreamVendor,
			At: inkwell.Point{
				X: 300,
				Y: 206,
			},
		},
		{
			Character: character.CandyCrusher,
			At: inkwell.Point{
				X: 520,
				Y: 244,
			},
		},
	},
	Hotspots: iceCreamShop{}.hotspots(),
})

func (iceCreamShop) hotspots() []inkwell.Hotspot {
	gone := inkwell.Say(character.Paul, inkwell.Text{
		{Lang: "en", Text: "Empty booth. A BVK compliance notice where his sundae was."},
		{Lang: "hu", Text: "Üres boksz. Egy BVK megfelelőségi értesítő a fagyija helyén."},
	})
	return []inkwell.Hotspot{
		{
			Name: "vendor",
			Label: inkwell.Text{
				{Lang: "en", Text: "ice cream vendor"},
				{Lang: "hu", Text: "fagyiárus"},
			},
			Area: inkwell.Rect(276, 126, 48, 84),
			OnLook: inkwell.Say(character.Paul, inkwell.Text{
				{Lang: "en", Text: "Paper hat. Big smile. Standing exactly where Manny used to pull pints."},
				{Lang: "hu", Text: "Papírsapka. Nagy mosoly. Pont ott áll, ahol Manny csapolta a sört."},
			}),
			OnTalk: system.System.Paused(inkwell.RunDialog(dialog.IceCreamVendor)),
			OnUse: inkwell.Say(character.Paul, inkwell.Text{
				{Lang: "en", Text: "I'll talk to him. I'm not going to use him."},
				{Lang: "hu", Text: "Beszélek vele. Használni nem fogom."},
			}),
			OnTake: inkwell.Say(speaker.Dex, inkwell.Text{
				{Lang: "en", Text: "He comes with the shop, Paul."},
				{Lang: "hu", Text: "Ő a bolthoz tartozik, Paul."},
			}),
		},
		{
			Name: "candy_crusher",
			Label: inkwell.Text{
				{Lang: "en", Text: "man with the handheld"},
				{Lang: "hu", Text: "kézikonzolos fickó"},
			},
			Area: inkwell.Rect(496, 164, 48, 84),
			OnLook: inkwell.If(inkwell.Flag(inc.FlagBVKLogged),
				gone,
				inkwell.Say(character.Paul, inkwell.Text{
					{Lang: "en", Text: "Thumbs going like pistons. He's parked in front of the only terminal in the place."},
					{Lang: "hu", Text: "A hüvelykujjai úgy járnak, mint a dugattyúk. A hely egyetlen termináljának elé parkolt."},
				}),
			),
			OnTalk: inkwell.If(inkwell.Flag(inc.FlagBVKLogged),
				gone,
				system.System.Paused(inkwell.RunDialog(dialog.CandyCrusher)),
			),
			OnUse: inkwell.If(inkwell.Flag(inc.FlagBVKLogged),
				gone,
				inkwell.Say(character.Paul, inkwell.Text{
					{Lang: "en", Text: "Shoving him won't get me the terminal. It'll get me the manager."},
					{Lang: "hu", Text: "Ha félretolom, nem a terminált kapom meg. Az üzletvezetőt."},
				}),
			),
			OnTake: inkwell.Say(speaker.Dex, inkwell.Text{
				{Lang: "en", Text: "You can't take a man. You can take his seat, if he lets you."},
				{Lang: "hu", Text: "Egy embert nem vehetsz el. A helyét igen, ha engedi."},
			}),
			OnUseWith: map[*inkwell.Item]inkwell.Action{
				item.IceCream: inkwell.If(inkwell.Flag(inc.FlagBVKLogged),
					gone,
					inkwell.If(inkwell.Flag(inc.FlagTerminalAllowed),
						inkwell.Say(character.Paul, inkwell.Text{
							{Lang: "en", Text: "He's already moved over. I'll keep the ice cream."},
							{Lang: "hu", Text: "Már félrehúzódott. A fagyit megtartom."},
						}),
						inkwell.Seq(
							inkwell.Say(character.Paul, inkwell.Text{
								{Lang: "en", Text: "Pistachio. On the house, if I can have the terminal for two minutes."},
								{Lang: "hu", Text: "Pisztácia. A ház ajándéka, ha megkaphatom a terminált két percre."},
							}),
							inkwell.Say(character.CandyCrusher, inkwell.Text{
								{Lang: "en", Text: "Oh. Well. I mean. Two minutes. Don't close my game."},
								{Lang: "hu", Text: "Ó. Hát. Mármint. Két perc. Ne zárd be a játékomat."},
							}),
							inkwell.TakeAway(item.IceCream),
							inkwell.SetFlag(inc.FlagTerminalAllowed),
							inkwell.Say(speaker.Dex, inkwell.Text{
								{Lang: "en", Text: "Bought with a scoop of pistachio. That's the going rate for a BVK terminal now."},
								{Lang: "hu", Text: "Egy gombóc pisztáciáért. Ennyi most egy BVK-terminál ára."},
							}),
						),
					),
				),
			},
		},
		{
			Name: "terminal",
			Label: inkwell.Text{
				{Lang: "en", Text: "BVK terminal"},
				{Lang: "hu", Text: "BVK-terminál"},
			},
			Area: inkwell.Rect(556, 96, 64, 72),
			OnLook: inkwell.Say(character.Paul, inkwell.Text{
				{Lang: "en", Text: "Registered BVK terminal. The seal is newer than the paint around it."},
				{Lang: "hu", Text: "Bejegyzett BVK-terminál. A pecsét újabb, mint a festék körülötte."},
			}),
			OnUse: inkwell.If(inkwell.Flag(inc.FlagTerminalSearched),
				inkwell.Say(character.Paul, inkwell.Text{
					{Lang: "en", Text: "One logged search is one too many. I'm not typing that in again."},
					{Lang: "hu", Text: "Egy logolt keresés is több a kelleténél. Nem gépelem be újra."},
				}),
				inkwell.If(inkwell.Flag(inc.FlagTerminalAllowed),
					inkwell.RunScript(script.TerminalSearch),
					inkwell.Say(character.CandyCrusher, inkwell.Text{
						{Lang: "en", Text: "Hey. Hey. I'm on this. Level ninety-one."},
						{Lang: "hu", Text: "Hé. Hé. Ezen én vagyok. Kilencvenegyedik szint."},
					}),
				),
			),
			OnTake: inkwell.Say(speaker.Dex, inkwell.Text{
				{Lang: "en", Text: "Bolted down, sealed, and logged. Three reasons, pick one."},
				{Lang: "hu", Text: "Lecsavarozva, lepecsételve, logolva. Három ok, válassz egyet."},
			}),
		},
		{
			Name: "counter",
			Label: inkwell.Text{
				{Lang: "en", Text: "counter"},
				{Lang: "hu", Text: "pult"},
			},
			Area: inkwell.Rect(220, 150, 170, 60),
			OnLook: inkwell.Seq(
				inkwell.Say(character.Paul, inkwell.Text{
					{Lang: "en", Text: "Thirty-one flavours on a board where the taps used to be."},
					{Lang: "hu", Text: "Harmincegy íz egy táblán, ott, ahol a csapok voltak."},
				}),
				inkwell.Say(speaker.Dex, inkwell.Text{
					{Lang: "en", Text: "The taps were here. I remember the taps. I'm the one who doesn't forget."},
					{Lang: "hu", Text: "A csapok itt voltak. Emlékszem a csapokra. Én vagyok az, aki nem felejt."},
				}),
			),
			OnUse: inkwell.Say(character.Paul, inkwell.Text{
				{Lang: "en", Text: "Sticky. It was sticky before, too. Different reason."},
				{Lang: "hu", Text: "Ragad. Régen is ragadt. Más okból."},
			}),
			OnTake: inkwell.Say(character.Paul, inkwell.Text{
				{Lang: "en", Text: "I'm not taking a counter."},
				{Lang: "hu", Text: "Nem viszek el egy pultot."},
			}),
		},
	}
}

func (iceCreamShop) Scene() *inkwell.Scene { return IceCreamShop }
