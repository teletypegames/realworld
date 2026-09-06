package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/content/character"
	"realworld/inc/content/script"
	"realworld/inc/content/speaker"
)

var prototypeRoomRefuse = inkwell.Seq(
	character.Paul.Say(Text{
		{Lang: "en", Text: "A chair, a helmet, a fist of cables. I am not sitting in that."},
		{Lang: "hu", Text: "Egy szék, egy sisak, egy marék kábel. Nem ülök bele."},
	}),
	speaker.Dex.Say(Text{
		{Lang: "en", Text: "Good. That is the chair that emptied Sumphor out. You look at it. You do not sit in it."},
		{Lang: "hu", Text: "Helyes. Ez az a szék, ami kiürítette Sumphort. Ránézel. Nem ülsz bele."},
	}),
)

var PrototypeRoom = SceneManager.Register(&Scene{
	Name: "prototype_room",
	Title: Text{
		{Lang: "en", Text: "PROTOTYPE ROOM"},
		{Lang: "hu", Text: "PROTOTÍPUS SZOBA"},
	},
	Layers: []SceneLayer{
		{
			Asset: &Asset{
				Name: "prototype_room",
				Path: "assets/bg/prototype_room.png",
			},
			Z: -1,
		},
	},
	Hotspots: []Hotspot{
		{
			Name: "chair",
			Label: Text{
				{Lang: "en", Text: "chair with the helmet"},
				{Lang: "hu", Text: "szék a sisakkal"},
			},
			Area:   inkwell.Rect(240, 120, 90, 110),
			OnLook: prototypeRoomRefuse,
			OnUse:  prototypeRoomRefuse,
			OnTake: character.Paul.Say(Text{
				{Lang: "en", Text: "I wouldn't carry that out if it were the last chair in the city."},
				{Lang: "hu", Text: "Ezt nem vinném ki, ha ez lenne az utolsó szék a városban."},
			}),
		},
		{
			Name: "terminal",
			Label: Text{
				{Lang: "en", Text: "offline terminal"},
				{Lang: "hu", Text: "offline terminál"},
			},
			Area: inkwell.Rect(60, 130, 90, 70),
			OnLook: character.Paul.Say(Text{
				{Lang: "en", Text: "No network light. Just a log, spooling to a disk nobody reads."},
				{Lang: "hu", Text: "Nincs hálózati fény. Csak egy napló, ami egy lemezre ír, amit senki nem olvas."},
			}),
			OnUse: inkwell.If(inkwell.Flag(inc.FlagLabLocated),
				character.Paul.Say(Text{
					{Lang: "en", Text: "I've read it. The next prototype went into the Golden Gate. That's where he is."},
					{Lang: "hu", Text: "Elolvastam. A következő prototípus a Golden Gate-be került. Ott van."},
				}),
				inkwell.RunScript(script.PrototypeLogs),
			),
		},
		{
			Name: "camera",
			Label: Text{
				{Lang: "en", Text: "live camera feed"},
				{Lang: "hu", Text: "élő kamerakép"},
			},
			Area: inkwell.Rect(460, 60, 120, 90),
			OnLook: inkwell.Seq(
				character.Paul.Say(Text{
					{Lang: "en", Text: "A monitor, one live feed. A man asleep in a chair like this one, only bigger. Wired to the ceiling."},
					{Lang: "hu", Text: "Egy monitor, egy élő kép. Egy alvó ember egy ehhez hasonló székben, csak nagyobban. A plafonhoz drótozva."},
				}),
				character.Paul.Say(Text{
					{Lang: "en", Text: "That's Norman. That's him, right now, somewhere."},
					{Lang: "hu", Text: "Ez Norman. Ez ő, most, valahol."},
				}),
				speaker.Dex.Say(Text{
					{Lang: "en", Text: "Then the feed comes from wherever they put him. Find the wire and you find the room."},
					{Lang: "hu", Text: "Akkor a kép onnan jön, ahová bedugták. Találd meg a drótot, és megtalálod a szobát."},
				}),
			),
		},
	},
})

var _ = PrototypeRoom.Leads(
	Exit{
		To: AbandonedOffice,
		Label: Text{
			{Lang: "en", Text: "back out past the fridge"},
			{Lang: "hu", Text: "vissza ki a hűtő mellett"},
		},
		Side: inkwell.ExitNear,
	},
)
