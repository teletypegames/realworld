package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/content/character"
	"realworld/inc/content/item"
	"realworld/inc/content/script"
	"realworld/inc/content/speaker"
)

var abandonedOfficeTaken = inkwell.Say(character.Paul, inkwell.Text{
	{Lang: "en", Text: "Nothing else on the desk worth carrying."},
	{Lang: "hu", Text: "Semmi más az asztalon, amit érdemes elvinni."},
})

var abandonedOfficePickUp = inkwell.Seq(
	inkwell.Give(item.SumphorTape),
	inkwell.Say(character.Paul, inkwell.Text{
		{Lang: "en", Text: "A Personal Tape, label peeled off. Sumphor's desk, so presumably Sumphor's."},
		{Lang: "hu", Text: "Egy Személyes Kazetta, a címke lehúzva. Sumphor asztala, tehát feltehetően Sumphoré."},
	}),
	inkwell.If(inkwell.Flag(inc.FlagTapeSpeaks),
		inkwell.Say(speaker.TapeMystery, inkwell.Text{
			{Lang: "en", Text: "That one is Norman's. It got mixed up with Sumphor's things. Bring it, it belongs with me."},
			{Lang: "hu", Text: "Az Normané. Összekeveredett Sumphor holmijával. Hozd, hozzám tartozik."},
		}),
		inkwell.Say(speaker.Dex, inkwell.Text{
			{Lang: "en", Text: "Or it's exactly what it looks like: a broken man's tape on a broken man's desk."},
			{Lang: "hu", Text: "Vagy pontosan az, aminek látszik: egy összetört ember kazettája egy összetört ember asztalán."},
		}),
	),
)

var abandonedOfficeOpen = inkwell.Say(character.Paul, inkwell.Text{
	{Lang: "en", Text: "Dead and swung aside. The room behind it is open."},
	{Lang: "hu", Text: "Kihalt, és félretolva. A mögötte lévő szoba nyitva."},
})

var AbandonedOffice = SceneManager.Register(&Scene{
	Name: "abandoned_office",
	Title: inkwell.Text{
		{Lang: "en", Text: "ABANDONED OFFICE"},
		{Lang: "hu", Text: "ELHAGYATOTT IRODA"},
	},
	Triggers: []inkwell.Trigger{
		{
			Name: "sumphor_enters",
			Once: true,
			When: inkwell.Given{
				Flag:    inc.FlagLabLocated,
				NotFlag: inc.FlagSumphorSeen,
			},
			Do: inkwell.RunScript(script.SumphorEnters),
		},
	},
	Hotspots: []inkwell.Hotspot{
		{
			Name: "computer",
			Label: inkwell.Text{
				{Lang: "en", Text: "Norman's second machine"},
				{Lang: "hu", Text: "Norman második gépe"},
			},
			Area: inkwell.Rect(80, 120, 100, 80),
			OnLook: inkwell.Say(character.Paul, inkwell.Text{
				{Lang: "en", Text: "Still on a UPS, humming to itself in the dark. He never turned it off."},
				{Lang: "hu", Text: "Még mindig szünetmentesen, magának dúdol a sötétben. Sosem kapcsolta ki."},
			}),
			OnUse: inkwell.Seq(
				inkwell.Say(character.Paul, inkwell.Text{
					{Lang: "en", Text: "Dev logs. Two names on every commit: Reds and one they scrubbed to initials. The trance protocol, version after version."},
					{Lang: "hu", Text: "Fejlesztői naplók. Minden bejegyzésen két név: Reds és egy, amit kezdőbetűkre töröltek. A transz-protokoll, verzióról verzióra."},
				}),
				inkwell.Say(speaker.Dex, inkwell.Text{
					{Lang: "en", Text: "The scrubbed one is Sumphor. They didn't fire him. They deleted him."},
					{Lang: "hu", Text: "A törölt Sumphor. Nem kirúgták. Kitörölték."},
				}),
			),
		},
		{
			Name: "sumphor_desk",
			Label: inkwell.Text{
				{Lang: "en", Text: "Sumphor's desk"},
				{Lang: "hu", Text: "Sumphor asztala"},
			},
			Area: inkwell.Rect(220, 130, 100, 70),
			OnLook: inkwell.If(inkwell.HasItem(item.SumphorTape),
				abandonedOfficeTaken,
				inkwell.Say(character.Paul, inkwell.Text{
					{Lang: "en", Text: "Untouched since the day. A cold coffee, a chewed pen, and a tape sitting square in the middle."},
					{Lang: "hu", Text: "Aznap óta érintetlen. Egy hideg kávé, egy megrágott toll, és egy kazetta pontosan középen."},
				}),
			),
			OnTake: inkwell.If(inkwell.HasItem(item.SumphorTape), abandonedOfficeTaken, abandonedOfficePickUp),
			OnUse:  inkwell.If(inkwell.HasItem(item.SumphorTape), abandonedOfficeTaken, abandonedOfficePickUp),
		},
		{
			Name: "fridge",
			Label: inkwell.Text{
				{Lang: "en", Text: "humming fridge"},
				{Lang: "hu", Text: "zümmögő hűtő"},
			},
			Area: inkwell.Rect(420, 110, 80, 110),
			OnLook: inkwell.If(inkwell.Flag(inc.FlagFridgeOpen),
				abandonedOfficeOpen,
				inkwell.Say(character.Paul, inkwell.Text{
					{Lang: "en", Text: "A full-height fridge against the back wall. Humming. There's a draught coming from behind it."},
					{Lang: "hu", Text: "Egy padlótól plafonig érő hűtő a hátsó falnál. Zümmög. Huzat jön mögüle."},
				}),
			),
			OnUse: inkwell.If(inkwell.Flag(inc.FlagFridgeOpen),
				abandonedOfficeOpen,
				inkwell.If(inkwell.VarEq(inc.VarSlot2, item.SupportTape),
					inkwell.Seq(
						inkwell.Say(speaker.TapeSupport, inkwell.Text{
							{Lang: "en", Text: "Have you tried turning it off and on again?"},
							{Lang: "hu", Text: "Próbálta már ki- és bekapcsolni?"},
						}),
						inkwell.Say(character.Paul, inkwell.Text{
							{Lang: "en", Text: "…fine. Pulling the plug."},
							{Lang: "hu", Text: "…jó. Kihúzom."},
						}),
						inkwell.Say(speaker.Dex, inkwell.Text{
							{Lang: "en", Text: "The hum stops. The whole thing rolls on a hidden castor. There's your door."},
							{Lang: "hu", Text: "A zümmögés eláll. Az egész elgurul egy rejtett görgőn. Ott az ajtód."},
						}),
						inkwell.SetFlag(inc.FlagFridgeOpen),
					),
					inkwell.Seq(
						inkwell.Say(character.Paul, inkwell.Text{
							{Lang: "en", Text: "Won't budge. It's plugged in and it does not want to move."},
							{Lang: "hu", Text: "Meg se mozdul. Be van dugva, és nem akar elmozdulni."},
						}),
						inkwell.Say(speaker.Dex, inkwell.Text{
							{Lang: "en", Text: "Listen to the compressor. It has a rhythm. Kill the power and it'll roll."},
							{Lang: "hu", Text: "Hallgasd a kompresszort. Van ritmusa. Vedd le róla az áramot, és elgurul."},
						}),
					),
				),
			),
		},
	},
})

var _ = AbandonedOffice.Leads(
	inkwell.Exit{
		To: WorkplaceLobby,
		Label: inkwell.Text{
			{Lang: "en", Text: "the lift down"},
			{Lang: "hu", Text: "a lift lefelé"},
		},
		Side: inkwell.ExitNear,
	},
	inkwell.Exit{
		To: PrototypeRoom,
		Label: inkwell.Text{
			{Lang: "en", Text: "behind the fridge"},
			{Lang: "hu", Text: "a hűtő mögött"},
		},
		Side:  inkwell.ExitBack,
		Needs: inc.FlagFridgeOpen,
		Blocked: inkwell.Say(character.Paul, inkwell.Text{
			{Lang: "en", Text: "There's a room back there, but the fridge is bolted to the wall over the gap."},
			{Lang: "hu", Text: "Hátul van egy szoba, de a hűtő rá van csavarozva a falra a rés fölött."},
		}),
	},
)
