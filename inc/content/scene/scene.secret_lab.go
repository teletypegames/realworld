package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/content/character"
	"realworld/inc/content/item"
	"realworld/inc/content/script"
	"realworld/inc/content/speaker"
)

type secretLab struct{}

var _ place = secretLab{}

var SecretLab = SceneManager.Register(&Scene{
	Name: "secret_lab",
	Title: inkwell.Text{
		{Lang: "en", Text: "SECRET LAB — INSIDE THE GOLDEN GATE"},
		{Lang: "hu", Text: "TITKOS LABOR — A GOLDEN GATE BELSEJÉBEN"},
	},
	Exits: []inkwell.Exit{
		{
			To: securityRoom{},
			Label: inkwell.Text{
				{Lang: "en", Text: "back to the observation room"},
				{Lang: "hu", Text: "vissza a megfigyelőszobába"},
			},
			Side: inkwell.ExitNear,
		},
	},
	Actors: []inkwell.SceneActor{
		{
			Character: character.Paul,
			At: inkwell.Point{
				X: 120,
				Y: 240,
			},
		},
		{
			Character: character.Norman,
			At: inkwell.Point{
				X: 400,
				Y: 205,
			},
		},
	},
	Triggers: secretLab{}.triggers(),
	Hotspots: secretLab{}.hotspots(),
})

func (secretLab) triggers() []inkwell.Trigger {
	plea := inkwell.Seq(
		inkwell.Say(speaker.TapeMystery, inkwell.Text{
			{Lang: "en", Text: "There. The machine that runs it. Put me in, Paul. It's the only way he wakes up."},
			{Lang: "hu", Text: "Ott. A gép, ami futtatja. Tegyél be, Paul. Csak így ébred fel."},
		}),
		inkwell.Say(speaker.TapeMystery, inkwell.Text{
			{Lang: "en", Text: "I carry what he knows. Once I'm inside, he can say it out loud, and the BVK is finished."},
			{Lang: "hu", Text: "Én hordozom, amit ő tud. Ha bent vagyok, ki tudja mondani, és a BVK-nak vége."},
		}),
		inkwell.Say(speaker.Dex, inkwell.Text{
			{Lang: "en", Text: "Paul. Think."},
			{Lang: "hu", Text: "Paul. Gondolkozz."},
		}),
	)
	return []inkwell.Trigger{
		{
			Name: "plea",
			When: inkwell.Given{
				Flag: inc.FlagTapeSpeaks,
			},
			Do:   plea,
			Once: true,
		},
	}
}

func (secretLab) hotspots() []inkwell.Hotspot {
	nothingItWants := inkwell.If(inkwell.HasItem(item.MysteryTape),
		inkwell.Say(speaker.Dex, inkwell.Text{
			{Lang: "en", Text: "You're holding the tape. If you're going to do this, you do it deliberately: slot it, or push it in. Not by accident."},
			{Lang: "hu", Text: "A kezedben a kazetta. Ha megteszed, szándékosan tedd: rakd a foglalatba, vagy told be. Ne véletlenül."},
		}),
		inkwell.Say(character.Paul, inkwell.Text{
			{Lang: "en", Text: "A tape slot on the front panel. I have nothing it wants."},
			{Lang: "hu", Text: "Kazettafoglalat az előlapon. Nincs nálam semmi, amit akarna."},
		}),
	)
	return []inkwell.Hotspot{
		{
			Name: "norman",
			Label: inkwell.Text{
				{Lang: "en", Text: "Norman, in the chair"},
				{Lang: "hu", Text: "Norman, a székben"},
			},
			Area: inkwell.Rect(376, 125, 48, 84),
			OnLook: inkwell.Say(character.Paul, inkwell.Text{
				{Lang: "en", Text: "Asleep. Every so often a hand twitches, like he's typing."},
				{Lang: "hu", Text: "Alszik. Néha megrándul a keze, mintha gépelne."},
			}),
			OnTalk: inkwell.Seq(
				inkwell.Say(character.Paul, inkwell.Text{
					{Lang: "en", Text: "Norman. Norman, it's Paul."},
					{Lang: "hu", Text: "Norman. Norman, Paul vagyok."},
				}),
				inkwell.Say(speaker.Dex, inkwell.Text{
					{Lang: "en", Text: "He's somewhere else. Somewhere that has a desk and a coffee machine, if that monitor is honest."},
					{Lang: "hu", Text: "Ő máshol van. Valahol, ahol van íróasztal és kávégép, ha az a monitor nem hazudik."},
				}),
			),
			OnTake: inkwell.Say(character.Paul, inkwell.Text{
				{Lang: "en", Text: "Pull the cables and I don't know what comes out with them."},
				{Lang: "hu", Text: "Ha kihúzom a kábeleket, nem tudom, mi jön ki velük."},
			}),
			OnUse: inkwell.Say(character.Paul, inkwell.Text{
				{Lang: "en", Text: "The helmet is locked on. This is not a thing you undo by hand."},
				{Lang: "hu", Text: "A sisak rázárva. Ezt nem kézzel szedi le az ember."},
			}),
		},
		{
			Name: "monitor",
			Label: inkwell.Text{
				{Lang: "en", Text: "monitor behind the chair"},
				{Lang: "hu", Text: "monitor a szék mögött"},
			},
			Area: inkwell.Rect(440, 40, 150, 90),
			OnLook: inkwell.Seq(
				inkwell.Say(character.Paul, inkwell.Text{
					{Lang: "en", Text: "A small office. A man at a desk, a coffee machine, a plant nobody waters. It's rendered. It's a game."},
					{Lang: "hu", Text: "Egy kis iroda. Egy ember az asztalnál, kávégép, egy növény, amit senki nem öntöz. Renderelt. Ez egy játék."},
				}),
				inkwell.Say(speaker.Dex, inkwell.Text{
					{Lang: "en", Text: "That's where he is. And the man at the desk keeps looking over his shoulder."},
					{Lang: "hu", Text: "Ott van ő. És az ember az asztalnál folyton hátranéz a válla fölött."},
				}),
			),
			OnUse: inkwell.Say(character.Paul, inkwell.Text{
				{Lang: "en", Text: "No keyboard on this side. It's a window, not a door."},
				{Lang: "hu", Text: "Ezen az oldalon nincs billentyűzet. Ez ablak, nem ajtó."},
			}),
		},
		{
			Name: "simulator",
			Label: inkwell.Text{
				{Lang: "en", Text: "the simulator's main computer"},
				{Lang: "hu", Text: "a szimulátor központi gépe"},
			},
			Area: inkwell.Rect(200, 100, 130, 130),
			OnLook: inkwell.Say(character.Paul, inkwell.Text{
				{Lang: "en", Text: "Rack-sized. Warm. One empty tape slot on the front panel, lit up like it's waiting."},
				{Lang: "hu", Text: "Rack-méretű. Meleg. Egy üres kazettafoglalat az előlapon, kivilágítva, mintha várna."},
			}),
			OnUse: inkwell.If(inkwell.VarEq(inc.VarSlot2, item.MysteryTape),
				inkwell.Seq(
					inkwell.SetVar(inc.VarSlot2, ""),
					inkwell.RunScript(script.Finale),
				),
				nothingItWants,
			),
			OnUseWith: map[*inkwell.Item]inkwell.Action{
				item.MysteryTape: inkwell.Seq(
					inkwell.TakeAway(item.MysteryTape),
					inkwell.RunScript(script.Finale),
				),
			},
			OnUseWithAny: inkwell.Say(character.Paul, inkwell.Text{
				{Lang: "en", Text: "The slot is tape-shaped. That isn't."},
				{Lang: "hu", Text: "A foglalat kazetta alakú. Az nem."},
			}),
		},
	}
}

func (secretLab) Scene() *inkwell.Scene { return SecretLab }
