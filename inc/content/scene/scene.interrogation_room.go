package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/content/character"
	"realworld/inc/content/dialog"
	"realworld/inc/content/speaker"
	"realworld/inc/system"
)

var InterrogationRoom = SceneManager.Register(&Scene{
	Name: "interrogation_room",
	Title: Text{
		{Lang: "en", Text: "SFPD — INTERVIEW ROOM"},
		{Lang: "hu", Text: "SFPD — KIHALLGATÓSZOBA"},
	},
	Actors: []SceneActor{
		{
			Character: character.Paul,
			At: Point{
				X: 140,
				Y: 240,
			},
		},
		{
			Character: character.Vasquez,
			At: Point{
				X: 300,
				Y: 236,
			},
		},
		{
			Character: character.Noodle,
			At: Point{
				X: 500,
				Y: 200,
			},
		},
	},
	Hotspots: []Hotspot{
		{
			Name: "vasquez",
			Label: Text{
				{Lang: "en", Text: "Detective Vasquez"},
				{Lang: "hu", Text: "Vasquez nyomozó"},
			},
			Area: inkwell.Rect(276, 156, 48, 84),
			OnLook: character.Paul.Say(Text{
				{Lang: "en", Text: "Red hair cut short on one side, gold earrings, brown blazer. Not in uniform. She was waiting for me."},
				{Lang: "hu", Text: "Vörös haj, egyik oldalon rövidre vágva, arany fülbevaló, barna blézer. Nem egyenruhás. Engem várt."},
			}),
			OnTalk: system.System.Paused(inkwell.RunDialog(dialog.Vasquez)),
			OnUse: character.Vasquez.Say(Text{
				{Lang: "en", Text: "Sit down, Paul."},
				{Lang: "hu", Text: "Üljön le, Paul."},
			}),
			OnUseWithAny: character.Vasquez.Say(Text{
				{Lang: "en", Text: "Whatever that is, I don't want it. Yet."},
				{Lang: "hu", Text: "Akármi is az, nem kérem. Még."},
			}),
		},
		{
			Name: "glass_wall",
			Label: Text{
				{Lang: "en", Text: "Noodle, behind the glass"},
				{Lang: "hu", Text: "Noodle, az üveg mögött"},
			},
			Area: inkwell.Rect(440, 100, 150, 120),
			OnLook: inkwell.Seq(
				character.Paul.Say(Text{
					{Lang: "en", Text: "Noodle. Fatter than I remember, monocle, a cigar behind his ear. Eating chips in a police station."},
					{Lang: "hu", Text: "Noodle. Kövérebb, mint ahogy emlékszem, monokli, szivar a füle mögött. Chipset eszik egy rendőrségen."},
				}),
				speaker.Dex.Say(Text{
					{Lang: "en", Text: "Twenty years and he still dresses like a man who lost a bet."},
					{Lang: "hu", Text: "Húsz év, és még mindig úgy öltözik, mint aki elvesztett egy fogadást."},
				}),
			),
			OnTalk: system.System.Paused(inkwell.RunDialog(dialog.NoodleGlass)),
			OnUse: character.Paul.Say(Text{
				{Lang: "en", Text: "Thick glass. A slit at the bottom for words to get through, barely."},
				{Lang: "hu", Text: "Vastag üveg. Egy rés alul, amin épphogy átjutnak a szavak."},
			}),
			OnTake: character.Paul.Say(Text{
				{Lang: "en", Text: "I'd take him if I could."},
				{Lang: "hu", Text: "Elvinném, ha tudnám."},
			}),
		},
		{
			Name: "one_way_mirror",
			Label: Text{
				{Lang: "en", Text: "one-way mirror"},
				{Lang: "hu", Text: "detektívtükör"},
			},
			Area: inkwell.Rect(40, 60, 150, 100),
			OnLook: character.Paul.Say(Text{
				{Lang: "en", Text: "Somebody is behind that. Somebody always is."},
				{Lang: "hu", Text: "Valaki áll mögötte. Mindig áll mögötte valaki."},
			}),
			OnUse: speaker.Dex.Say(Text{
				{Lang: "en", Text: "Don't wave. It's not that kind of mirror."},
				{Lang: "hu", Text: "Ne integess. Ez nem az a fajta tükör."},
			}),
		},
		{
			Name: "table",
			Label: Text{
				{Lang: "en", Text: "table with a recorder"},
				{Lang: "hu", Text: "asztal a magnóval"},
			},
			Area: inkwell.Rect(220, 190, 130, 40),
			OnLook: character.Paul.Say(Text{
				{Lang: "en", Text: "Steel table, a tape recorder, two cups. Neither of them for me."},
				{Lang: "hu", Text: "Acélasztal, egy magnó, két pohár. Egyik sem nekem."},
			}),
			OnUse: character.Vasquez.Say(Text{
				{Lang: "en", Text: "It's running. It's always running."},
				{Lang: "hu", Text: "Megy. Mindig megy."},
			}),
			OnTake: character.Paul.Say(Text{
				{Lang: "en", Text: "Stealing the recorder from an interview room. Even I have limits."},
				{Lang: "hu", Text: "Ellopni a magnót egy kihallgatószobából. Még nekem is vannak határaim."},
			}),
		},
	},
})

var _ = InterrogationRoom.Leads(
	Exit{
		To: PoliceStation,
		Label: Text{
			{Lang: "en", Text: "back to the front desk"},
			{Lang: "hu", Text: "vissza a recepcióhoz"},
		},
		Side: inkwell.ExitNear,
	},
)
