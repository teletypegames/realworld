package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/content/character"
	"realworld/inc/content/dialog"
	"realworld/inc/content/speaker"
	"realworld/inc/system"
)

var TrinketShop = SceneManager.Register(&Scene{
	Name: "trinket_shop",
	Title: Text{
		{Lang: "en", Text: "CHINATOWN — TRINKET SHOP"},
		{Lang: "hu", Text: "KÍNAI NEGYED — CSECSEBECSEBOLT"},
	},
	Actors: []SceneActor{
		{
			Character: character.Paul,
			At: Point{
				X: 120,
				Y: 236,
			},
		},
		{
			Character: character.Shopkeeper,
			At: Point{
				X: 420,
				Y: 210,
			},
		},
	},
	Hotspots: []Hotspot{
		{
			Name: "shopkeeper",
			Label: Text{
				{Lang: "en", Text: "shopkeeper"},
				{Lang: "hu", Text: "boltos"},
			},
			Area: inkwell.Rect(396, 130, 48, 84),
			OnLook: character.Paul.Say(Text{
				{Lang: "en", Text: "Apron, round glasses, and a smile that's already decided what I'm buying."},
				{Lang: "hu", Text: "Kötény, kerek szemüveg, és egy mosoly, ami már eldöntötte, mit veszek."},
			}),
			OnTalk: system.System.Paused(inkwell.RunDialog(dialog.Shopkeeper)),
			OnUse: character.Shopkeeper.Say(Text{
				{Lang: "en", Text: "Please, sir, you touch, you buy."},
				{Lang: "hu", Text: "Kérem, uram, aki hozzányúl, az megveszi."},
			}),
			OnTake: speaker.Dex.Say(Text{
				{Lang: "en", Text: "He'd sell himself if you asked. Don't ask."},
				{Lang: "hu", Text: "Magát is eladná, ha kérnéd. Ne kérd."},
			}),
			OnUseWithAny: character.Shopkeeper.Say(Text{
				{Lang: "en", Text: "No refunds. No trades. Very good price though."},
				{Lang: "hu", Text: "Nincs visszaváltás. Nincs csere. De nagyon jó ár."},
			}),
		},
		{
			Name: "shelf",
			Label: Text{
				{Lang: "en", Text: "shelf of trinkets"},
				{Lang: "hu", Text: "csecsebecsés polc"},
			},
			Area: inkwell.Rect(40, 60, 200, 120),
			OnLook: character.Paul.Say(Text{
				{Lang: "en", Text: "Waving cats, jade that isn't, and a row of Personal Tapes with the seals scraped off."},
				{Lang: "hu", Text: "Integető macskák, jáde, ami nem az, és egy sor Személyes Kazetta, lekapart pecséttel."},
			}),
			OnUse: character.Shopkeeper.Say(Text{
				{Lang: "en", Text: "You like? You ask me, I give good price."},
				{Lang: "hu", Text: "Tetszik? Kérdezzen, jó árat adok."},
			}),
			OnTake: character.Shopkeeper.Say(Text{
				{Lang: "en", Text: "Sir. Sir. I am right here."},
				{Lang: "hu", Text: "Uram. Uram. Itt állok."},
			}),
		},
		{
			Name: "charms",
			Label: Text{
				{Lang: "en", Text: "bowl of charms"},
				{Lang: "hu", Text: "tál amulett"},
			},
			Area: inkwell.Rect(300, 176, 80, 36),
			OnLook: inkwell.Seq(
				character.Paul.Say(Text{
					{Lang: "en", Text: "Luck, health, wealth. One of them is a little sun on a blue disc. Sun under water."},
					{Lang: "hu", Text: "Szerencse, egészség, gazdagság. Egyikük egy kis nap egy kék korongon. Nap a víz alatt."},
				}),
				speaker.Dex.Say(Text{
					{Lang: "en", Text: "In a bowl of junk, in plain sight. Either it means nothing, or it's meant to be found by someone who knows."},
					{Lang: "hu", Text: "Egy tál kacat között, szem előtt. Vagy semmit nem jelent, vagy annak szánták, aki tudja."},
				}),
			),
			OnUse: character.Shopkeeper.Say(Text{
				{Lang: "en", Text: "Charms, two dollar. Three for five."},
				{Lang: "hu", Text: "Amulett, két dollár. Három öt dollár."},
			}),
			OnTake: character.Paul.Say(Text{
				{Lang: "en", Text: "I'll leave the bowl. I'm here for something that talks."},
				{Lang: "hu", Text: "A tálat itt hagyom. Valamiért jöttem, ami beszél."},
			}),
		},
		{
			Name: "curtain",
			Label: Text{
				{Lang: "en", Text: "bead curtain"},
				{Lang: "hu", Text: "gyöngyfüggöny"},
			},
			Area: inkwell.Rect(500, 60, 100, 130),
			OnLook: inkwell.If(inkwell.Flag(inc.FlagClubEntry),
				character.Paul.Say(Text{
					{Lang: "en", Text: "The way in. He's holding it open with his eyes."},
					{Lang: "hu", Text: "A bejárat. A szemével tartja félre."},
				}),
				character.Paul.Say(Text{
					{Lang: "en", Text: "Beads. A storeroom behind, if I had to guess. He doesn't want me guessing."},
					{Lang: "hu", Text: "Gyöngyök. Mögötte raktár, ha tippelnem kéne. Nem akarja, hogy tippeljek."},
				}),
			),
			OnUse: inkwell.If(inkwell.Flag(inc.FlagClubEntry),
				character.Shopkeeper.Say(Text{
					{Lang: "en", Text: "Yes, yes. Through there. Mind the step."},
					{Lang: "hu", Text: "Igen, igen. Arra tessék. Vigyázat, lépcső."},
				}),
				character.Shopkeeper.Say(Text{
					{Lang: "en", Text: "Staff only, sir. You want something, everything is out here."},
					{Lang: "hu", Text: "Csak személyzet, uram. Ami kell, az mind itt kint van."},
				}),
			),
			OnTake: character.Paul.Say(Text{
				{Lang: "en", Text: "I'm not stealing a curtain."},
				{Lang: "hu", Text: "Nem lopok el egy függönyt."},
			}),
		},
	},
})

var _ = TrinketShop.Leads(
	exitToSelector,
	Exit{
		To: SmallRestaurant,
		Label: Text{
			{Lang: "en", Text: "the eating place next door"},
			{Lang: "hu", Text: "a szomszédos falatozó"},
		},
		Side: inkwell.ExitRight,
	},
	Exit{
		To: SecretClub,
		Label: Text{
			{Lang: "en", Text: "the way in at the back"},
			{Lang: "hu", Text: "a hátsó bejárat"},
		},
		Side:  inkwell.ExitBack,
		Needs: inc.FlagClubEntry,
		Blocked: character.Paul.Say(Text{
			{Lang: "en", Text: "Just a shop, as far as the man behind the counter is concerned."},
			{Lang: "hu", Text: "Csak egy bolt, már ami a pult mögötti embert illeti."},
		}),
	},
)
