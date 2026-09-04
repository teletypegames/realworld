package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/content/character"
	"realworld/inc/content/item"
	"realworld/inc/content/script"
	"realworld/inc/content/speaker"
)

type restaurantToilet struct{}

var _ place = restaurantToilet{}

var RestaurantToilet = SceneManager.Register(&Scene{
	Name: "restaurant_toilet",
	Title: inkwell.Text{
		{Lang: "en", Text: "TOILET — THE LIT MIRROR"},
		{Lang: "hu", Text: "MOSDÓ — A KIVILÁGÍTOTT TÜKÖR"},
	},
	Exits: []inkwell.Exit{
		{
			To: smallRestaurant{},
			Label: inkwell.Text{
				{Lang: "en", Text: "back into the restaurant"},
				{Lang: "hu", Text: "vissza az étterembe"},
			},
			Side: inkwell.ExitNear,
		},
	},
	Hotspots: restaurantToilet{}.hotspots(),
})

func (restaurantToilet) hotspots() []inkwell.Hotspot {
	toggle := inkwell.If(inkwell.Flag(inc.FlagMirrorPassword),
		inkwell.Say(character.Paul, inkwell.Text{
			{Lang: "en", Text: "It's dead, and the sun is still there behind it. I've got what it was hiding."},
			{Lang: "hu", Text: "Kihalt, és a nap még mindig ott van mögötte. Megvan, amit rejtett."},
		}),
		inkwell.If(inkwell.VarEq(inc.VarSlot2, item.SupportTape),
			inkwell.RunScript(script.MirrorPassword),
			inkwell.Seq(
				inkwell.Say(character.Paul, inkwell.Text{
					{Lang: "en", Text: "Off. On. The mirror comes back, brighter than before. Nothing behind it."},
					{Lang: "hu", Text: "Ki. Be. A tükör visszajön, fényesebben, mint előtte. Semmi mögötte."},
				}),
				inkwell.Say(speaker.Dex, inkwell.Text{
					{Lang: "en", Text: "There's a trick to it, and I don't know the trick. That's a rarer sentence than you think."},
					{Lang: "hu", Text: "Van ennek valami trükkje, és én nem ismerem. Ez ritkább mondat, mint gondolnád."},
				}),
			),
		),
	)
	return []inkwell.Hotspot{
		{
			Name: "mirror",
			Label: inkwell.Text{
				{Lang: "en", Text: "lit mirror"},
				{Lang: "hu", Text: "kivilágított tükör"},
			},
			Area: inkwell.Rect(240, 40, 160, 130),
			OnLook: inkwell.If(inkwell.Flag(inc.FlagMirrorPassword),
				inkwell.Say(character.Paul, inkwell.Text{
					{Lang: "en", Text: "A sun under water, glowing in the dark glass. And the word beneath it: Lector."},
					{Lang: "hu", Text: "Nap a víz alatt, a sötét üvegben világít. Alatta a szó: Lector."},
				}),
				inkwell.Seq(
					inkwell.Say(character.Paul, inkwell.Text{
						{Lang: "en", Text: "The light behind the glass flickers. Every few seconds there's a shape in it, then there isn't."},
						{Lang: "hu", Text: "Az üveg mögött vibrál a fény. Pár másodpercenként kirajzolódik benne valami, aztán eltűnik."},
					}),
					inkwell.Say(speaker.Dex, inkwell.Text{
						{Lang: "en", Text: "Something's painted on the back of the mirror. You only see it when the light dies."},
						{Lang: "hu", Text: "Valami a tükör hátára van festve. Csak akkor látszik, amikor a fény kialszik."},
					}),
				),
			),
			OnUse: toggle,
			OnTake: inkwell.Say(character.Paul, inkwell.Text{
				{Lang: "en", Text: "Bolted to the tiles. And I'd only see myself in it anyway."},
				{Lang: "hu", Text: "A csempéhez csavarozva. És úgyis csak magamat látnám benne."},
			}),
			OnUseWith: map[*inkwell.Item]inkwell.Action{
				item.SupportTape: inkwell.Say(speaker.Dex, inkwell.Text{
					{Lang: "en", Text: "Holding it up to the glass does nothing. A tape talks from a slot, Paul."},
					{Lang: "hu", Text: "Az üveghez tartva semmit nem csinál. Egy kazetta a foglalatból beszél, Paul."},
				}),
			},
		},
		{
			Name: "light_switch",
			Label: inkwell.Text{
				{Lang: "en", Text: "light switch"},
				{Lang: "hu", Text: "villanykapcsoló"},
			},
			Area: inkwell.Rect(440, 110, 30, 40),
			OnLook: inkwell.Say(character.Paul, inkwell.Text{
				{Lang: "en", Text: "One switch for the mirror light. Worn smooth. People fiddle with it."},
				{Lang: "hu", Text: "Egy kapcsoló a tükörfénynek. Simára kopott. Babrálnak vele."},
			}),
			OnUse: toggle,
			OnTake: inkwell.Say(character.Paul, inkwell.Text{
				{Lang: "en", Text: "It's in the wall."},
				{Lang: "hu", Text: "A falban van."},
			}),
		},
		{
			Name: "sink",
			Label: inkwell.Text{
				{Lang: "en", Text: "sink"},
				{Lang: "hu", Text: "mosdókagyló"},
			},
			Area: inkwell.Rect(250, 176, 140, 60),
			OnLook: inkwell.Say(character.Paul, inkwell.Text{
				{Lang: "en", Text: "Cracked porcelain. The tap drips in time with the mirror."},
				{Lang: "hu", Text: "Repedt porcelán. A csap a tükörrel egy ütemre csöpög."},
			}),
			OnUse: inkwell.Say(character.Paul, inkwell.Text{
				{Lang: "en", Text: "Cold water. It helps less than it used to."},
				{Lang: "hu", Text: "Hideg víz. Kevésbé segít, mint régen."},
			}),
			OnTake: inkwell.Say(speaker.Dex, inkwell.Text{
				{Lang: "en", Text: "Plumbing, Paul. Leave the plumbing."},
				{Lang: "hu", Text: "Vízvezeték, Paul. Hagyd a vízvezetéket."},
			}),
		},
	}
}

func (restaurantToilet) Scene() *inkwell.Scene { return RestaurantToilet }
