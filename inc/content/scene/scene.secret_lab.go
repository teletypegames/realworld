package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

type secretLab struct{}

var _ = SceneManager.Register(Scene{
	Name: inc.SceneSecretLab,
	Title: inkwell.Text{
		{Lang: "en", Text: "SECRET LAB — INSIDE THE GOLDEN GATE"},
		{Lang: "hu", Text: "TITKOS LABOR — A GOLDEN GATE BELSEJÉBEN"},
	},
	Exits: []inkwell.Exit{
		{
			To: inc.SceneSecurityRoom,
			Label: inkwell.Text{
				{Lang: "en", Text: "back to the observation room"},
				{Lang: "hu", Text: "vissza a megfigyelőszobába"},
			},
			Side: inkwell.ExitNear,
		},
	},
	Actors: []inkwell.SceneActor{
		{
			CharacterName: inc.Paul,
			At: inkwell.Point{
				X: 120,
				Y: 240,
			},
		},
		{
			CharacterName: inc.Norman,
			At: inkwell.Point{
				X: 400,
				Y: 205,
			},
		},
	},
	Triggers: []inkwell.Trigger{
		{
			Name: "plea",
			When: inkwell.Given{
				Flag: inc.FlagTapeSpeaks,
			},
			Do:   secretLab{}.plea(),
			Once: true,
		},
	},
	Hotspots: secretLab{}.hotspots(),
})

func (secretLab) plea() inkwell.Action {
	return inkwell.Seq(
		inkwell.Say(inc.TapeMystery, inkwell.Text{
			{Lang: "en", Text: "There. The machine that runs it. Put me in, Paul. It's the only way he wakes up."},
			{Lang: "hu", Text: "Ott. A gép, ami futtatja. Tegyél be, Paul. Csak így ébred fel."},
		}),
		inkwell.Say(inc.TapeMystery, inkwell.Text{
			{Lang: "en", Text: "I carry what he knows. Once I'm inside, he can say it out loud, and the BVK is finished."},
			{Lang: "hu", Text: "Én hordozom, amit ő tud. Ha bent vagyok, ki tudja mondani, és a BVK-nak vége."},
		}),
		inkwell.Say(inc.Dex, inkwell.Text{
			{Lang: "en", Text: "Paul. Think."},
			{Lang: "hu", Text: "Paul. Gondolkozz."},
		}),
	)
}

func (place secretLab) hotspots() []inkwell.Hotspot {
	return []inkwell.Hotspot{
		place.norman(),
		place.monitor(),
		place.simulator(),
	}
}

func (secretLab) norman() inkwell.Hotspot {
	return inkwell.Hotspot{
		Name: "norman",
		Label: inkwell.Text{
			{Lang: "en", Text: "Norman, in the chair"},
			{Lang: "hu", Text: "Norman, a székben"},
		},
		Area: inkwell.Rect(376, 125, 48, 84),
		OnLook: inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "Asleep. Every so often a hand twitches, like he's typing."},
			{Lang: "hu", Text: "Alszik. Néha megrándul a keze, mintha gépelne."},
		}),
		OnTalk: inkwell.Seq(
			inkwell.Say(inc.Paul, inkwell.Text{
				{Lang: "en", Text: "Norman. Norman, it's Paul."},
				{Lang: "hu", Text: "Norman. Norman, Paul vagyok."},
			}),
			inkwell.Say(inc.Dex, inkwell.Text{
				{Lang: "en", Text: "He's somewhere else. Somewhere that has a desk and a coffee machine, if that monitor is honest."},
				{Lang: "hu", Text: "Ő máshol van. Valahol, ahol van íróasztal és kávégép, ha az a monitor nem hazudik."},
			}),
		),
		OnTake: inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "Pull the cables and I don't know what comes out with them."},
			{Lang: "hu", Text: "Ha kihúzom a kábeleket, nem tudom, mi jön ki velük."},
		}),
		OnUse: inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "The helmet is locked on. This is not a thing you undo by hand."},
			{Lang: "hu", Text: "A sisak rázárva. Ezt nem kézzel szedi le az ember."},
		}),
	}
}

func (secretLab) monitor() inkwell.Hotspot {
	return inkwell.Hotspot{
		Name: "monitor",
		Label: inkwell.Text{
			{Lang: "en", Text: "monitor behind the chair"},
			{Lang: "hu", Text: "monitor a szék mögött"},
		},
		Area: inkwell.Rect(440, 40, 150, 90),
		OnLook: inkwell.Seq(
			inkwell.Say(inc.Paul, inkwell.Text{
				{Lang: "en", Text: "A small office. A man at a desk, a coffee machine, a plant nobody waters. It's rendered. It's a game."},
				{Lang: "hu", Text: "Egy kis iroda. Egy ember az asztalnál, kávégép, egy növény, amit senki nem öntöz. Renderelt. Ez egy játék."},
			}),
			inkwell.Say(inc.Dex, inkwell.Text{
				{Lang: "en", Text: "That's where he is. And the man at the desk keeps looking over his shoulder."},
				{Lang: "hu", Text: "Ott van ő. És az ember az asztalnál folyton hátranéz a válla fölött."},
			}),
		),
		OnUse: inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "No keyboard on this side. It's a window, not a door."},
			{Lang: "hu", Text: "Ezen az oldalon nincs billentyűzet. Ez ablak, nem ajtó."},
		}),
	}
}

func (secretLab) simulator() inkwell.Hotspot {
	nothingItWants := inkwell.If(inkwell.HasItem(inc.ItemMysteryTape),
		inkwell.Say(inc.Dex, inkwell.Text{
			{Lang: "en", Text: "You're holding the tape. If you're going to do this, you do it deliberately: slot it, or push it in. Not by accident."},
			{Lang: "hu", Text: "A kezedben a kazetta. Ha megteszed, szándékosan tedd: rakd a foglalatba, vagy told be. Ne véletlenül."},
		}),
		inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "A tape slot on the front panel. I have nothing it wants."},
			{Lang: "hu", Text: "Kazettafoglalat az előlapon. Nincs nálam semmi, amit akarna."},
		}),
	)
	return inkwell.Hotspot{
		Name: "simulator",
		Label: inkwell.Text{
			{Lang: "en", Text: "the simulator's main computer"},
			{Lang: "hu", Text: "a szimulátor központi gépe"},
		},
		Area: inkwell.Rect(200, 100, 130, 130),
		OnLook: inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "Rack-sized. Warm. One empty tape slot on the front panel, lit up like it's waiting."},
			{Lang: "hu", Text: "Rack-méretű. Meleg. Egy üres kazettafoglalat az előlapon, kivilágítva, mintha várna."},
		}),
		OnUse: inkwell.If(inkwell.VarEq(inc.VarSlot2, inc.ItemMysteryTape),
			inkwell.Seq(
				inkwell.SetVar(inc.VarSlot2, ""),
				inkwell.RunScript(inc.ScriptFinale),
			),
			nothingItWants,
		),
		OnUseWith: map[string]inkwell.Action{
			inc.ItemMysteryTape: inkwell.Seq(
				inkwell.TakeAway(inc.ItemMysteryTape),
				inkwell.RunScript(inc.ScriptFinale),
			),
		},
		OnUseWithAny: inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "The slot is tape-shaped. That isn't."},
			{Lang: "hu", Text: "A foglalat kazetta alakú. Az nem."},
		}),
	}
}
