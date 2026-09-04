package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/content/character"
	"realworld/inc/content/script"
	"realworld/inc/content/speaker"
)

type street struct{}

var _ place = street{}

var Street = SceneManager.Register(&Scene{
	Name: "street",
	Title: inkwell.Text{
		{Lang: "en", Text: "STREET"},
		{Lang: "hu", Text: "UTCA"},
	},
	Exits: []inkwell.Exit{
		exitToSelector,
	},
	Triggers: []inkwell.Trigger{
		{
			Name: "glitch",
			When: inkwell.Not(inkwell.VarEq(inc.VarDecay, nil)),
			Do:   inkwell.RunScript(script.StreetGlitch),
		},
	},
	Hotspots: street{}.hotspots(),
})

func (street) hotspots() []inkwell.Hotspot {
	return []inkwell.Hotspot{
		{
			Name: "news_ticker",
			Label: inkwell.Text{
				{Lang: "en", Text: "news ticker"},
				{Lang: "hu", Text: "hírszalag"},
			},
			Area: inkwell.Rect(180, 30, 300, 40),
			OnLook: inkwell.Reply(character.Paul, func(state *inkwell.State) inkwell.Text {
				switch state.Var(inc.VarDecay) {
				case inc.DecayNormanFound:
					return inkwell.Text{
						{Lang: "en", Text: "\"BVK ANNOUNCES BVK ANNOUNCES BVK ANNOUNCES.\" The letters are drifting apart from each other."},
						{Lang: "hu", Text: "„A BVK BEJELENTI A BVK BEJELENTI A BVK BEJELENTI.” A betűk lassan elúsznak egymástól."},
					}
				case inc.DecayTapeSpeaks:
					return inkwell.Text{
						{Lang: "en", Text: "\"TRANSIT DELAYS CITYWIDE. CAUSE: ROUTINE.\" That's not a cause. That's a shrug."},
						{Lang: "hu", Text: "„KÖZLEKEDÉSI FENNAKADÁSOK A VÁROSBAN. OK: RUTIN.” Ez nem ok. Ez egy vállrándítás."},
					}
				}
				return inkwell.Text{
					{Lang: "en", Text: "Weather, sports, a licence amnesty from the BVK. An ordinary Tuesday."},
					{Lang: "hu", Text: "Időjárás, sport, licencamnesztia a BVK-tól. Egy közönséges kedd."},
				}
			}),
			OnUse: inkwell.Say(character.Paul, inkwell.Text{
				{Lang: "en", Text: "It's four storeys up. I read it, I don't operate it."},
				{Lang: "hu", Text: "Négy emelet magasan van. Olvasom, nem kezelem."},
			}),
		},
		{
			Name: "streetlight",
			Label: inkwell.Text{
				{Lang: "en", Text: "flickering streetlight"},
				{Lang: "hu", Text: "pislákoló utcalámpa"},
			},
			Area: inkwell.Rect(520, 30, 40, 200),
			OnLook: inkwell.Reply(character.Paul, func(state *inkwell.State) inkwell.Text {
				if state.Var(inc.VarDecay) == nil {
					return inkwell.Text{
						{Lang: "en", Text: "On in daylight. The city has never fixed a lamp in my lifetime."},
						{Lang: "hu", Text: "Nappal is ég. A város az én életemben még nem javított meg egy lámpát sem."},
					}
				}
				return inkwell.Text{
					{Lang: "en", Text: "It flickers in a pattern. Long, short, short. It did it again. The same pattern."},
					{Lang: "hu", Text: "Mintázatban pislog. Hosszú, rövid, rövid. Megint. Ugyanaz a mintázat."},
				}
			}),
			OnUse: inkwell.Say(speaker.Dex, inkwell.Text{
				{Lang: "en", Text: "Don't climb it. Whatever is talking through the grid, it isn't talking to you."},
				{Lang: "hu", Text: "Ne mássz fel rá. Ami a hálózaton keresztül beszél, az nem hozzád beszél."},
			}),
		},
		{
			Name: "bus_stop",
			Label: inkwell.Text{
				{Lang: "en", Text: "bus stop"},
				{Lang: "hu", Text: "buszmegálló"},
			},
			Area: inkwell.Rect(40, 110, 120, 130),
			OnLook: inkwell.Say(character.Paul, inkwell.Text{
				{Lang: "en", Text: "A timetable and a BBS terminal with the screen kicked in. Somebody didn't like what it said."},
				{Lang: "hu", Text: "Egy menetrend és egy BBS-terminál berúgott képernyővel. Valakinek nem tetszett, amit írt."},
			}),
			OnUse: inkwell.Say(character.Paul, inkwell.Text{
				{Lang: "en", Text: "I'm not waiting for a bus. I've had my bus."},
				{Lang: "hu", Text: "Nem várok buszra. Buszból megvolt a mai."},
			}),
		},
	}
}

func (street) Scene() *inkwell.Scene { return Street }
