package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/content/character"
	"realworld/inc/content/item"
	"realworld/inc/content/script"
	"realworld/inc/content/speaker"
)

type noodleSecretRoom struct{}

var NoodleSecretRoom = SceneManager.Register(&Scene{
	Name: "noodle_secret_room",
	Title: inkwell.Text{
		{Lang: "en", Text: "NOODLE'S BACK ROOM"},
		{Lang: "hu", Text: "NOODLE HÁTSÓ SZOBÁJA"},
	},
	Exits: []inkwell.Exit{
		{
			To: alley{},
			Label: inkwell.Text{
				{Lang: "en", Text: "back out to the alley"},
				{Lang: "hu", Text: "vissza a sikátorba"},
			},
			Side: inkwell.ExitNear,
		},
	},
	Triggers: []inkwell.Trigger{
		{
			Name: "raided",
			When: inkwell.Given{
				Flag:    inc.FlagNoodleCode,
				NotFlag: inc.FlagRoomRaided,
			},
			Do:   inkwell.RunScript(script.SecretRoomEmpty),
			Once: true,
		},
	},
	Hotspots: noodleSecretRoom{}.hotspots(),
})

func (room noodleSecretRoom) hotspots() []inkwell.Hotspot {
	return []inkwell.Hotspot{
		room.shelf(),
		room.workbench(),
		room.cot(),
	}
}

func (noodleSecretRoom) shelf() inkwell.Hotspot {
	take := inkwell.Seq(
		inkwell.Give(item.CommentaryTape),
		inkwell.Say(character.Paul, inkwell.Text{
			{Lang: "en", Text: "A tape. Label says PAUL, in that handwriting."},
			{Lang: "hu", Text: "Egy kazetta. A címkén PAUL, azzal a kézírással."},
		}),
		inkwell.Say(speaker.Dex, inkwell.Text{
			{Lang: "en", Text: "He labelled it for you and put it where only you'd look. That's the most sentimental thing Noodle ever did."},
			{Lang: "hu", Text: "Neked címkézte, és oda tette, ahol csak te keresnéd. Ez a legérzelmesebb dolog, amit Noodle valaha tett."},
		}),
	)
	gone := inkwell.Say(character.Paul, inkwell.Text{
		{Lang: "en", Text: "Empty shelf. An evidence tag where something used to sit."},
		{Lang: "hu", Text: "Üres polc. Egy bűnjelcédula ott, ahol valami állt."},
	})
	return inkwell.Hotspot{
		Name: "shelf",
		Label: inkwell.Text{
			{Lang: "en", Text: "shelf over the desk"},
			{Lang: "hu", Text: "polc az asztal fölött"},
		},
		Area: inkwell.Rect(260, 70, 140, 50),
		OnLook: inkwell.If(inkwell.Flag(inc.FlagNoodleCode),
			gone,
			inkwell.If(inkwell.HasItem(item.CommentaryTape),
				inkwell.Say(character.Paul, inkwell.Text{
					{Lang: "en", Text: "Dust, and the clean rectangle where the tape was."},
					{Lang: "hu", Text: "Por, és egy tiszta téglalap, ahol a kazetta volt."},
				}),
				inkwell.Say(character.Paul, inkwell.Text{
					{Lang: "en", Text: "One tape, set apart from the rest. Somebody wanted it seen."},
					{Lang: "hu", Text: "Egy kazetta, külön a többitől. Valaki azt akarta, hogy észrevegyék."},
				}),
			),
		),
		OnUse: inkwell.If(inkwell.Flag(inc.FlagNoodleCode),
			gone,
			inkwell.If(inkwell.HasItem(item.CommentaryTape),
				inkwell.Say(character.Paul, inkwell.Text{
					{Lang: "en", Text: "Nothing else on it worth having."},
					{Lang: "hu", Text: "Más nincs rajta, amit érdemes lenne."},
				}),
				take,
			),
		),
		OnTake: inkwell.If(inkwell.Flag(inc.FlagNoodleCode),
			gone,
			inkwell.If(inkwell.HasItem(item.CommentaryTape),
				inkwell.Say(character.Paul, inkwell.Text{
					{Lang: "en", Text: "I've got it."},
					{Lang: "hu", Text: "Megvan."},
				}),
				take,
			),
		),
	}
}

func (noodleSecretRoom) workbench() inkwell.Hotspot {
	return inkwell.Hotspot{
		Name: "workbench",
		Label: inkwell.Text{
			{Lang: "en", Text: "workbench of cracked Armillas"},
			{Lang: "hu", Text: "munkaasztal feltört Armillákkal"},
		},
		Area: inkwell.Rect(220, 130, 220, 70),
		OnLook: inkwell.If(inkwell.Flag(inc.FlagNoodleCode),
			inkwell.Say(character.Paul, inkwell.Text{
				{Lang: "en", Text: "Stripped. Even the soldering iron's gone."},
				{Lang: "hu", Text: "Kipucolva. Még a forrasztópáka is eltűnt."},
			}),
			inkwell.Seq(
				inkwell.Say(character.Paul, inkwell.Text{
					{Lang: "en", Text: "Six Armillas open on the bench, every one of them mid-surgery."},
					{Lang: "hu", Text: "Hat felnyitott Armilla az asztalon, mind műtét közben."},
				}),
				inkwell.Say(speaker.Dex, inkwell.Text{
					{Lang: "en", Text: "This is the business. The letter was about something else."},
					{Lang: "hu", Text: "Ez az üzlet. A levél másról szólt."},
				}),
			),
		),
		OnUse: inkwell.Say(character.Paul, inkwell.Text{
			{Lang: "en", Text: "Not my bench, not my tools."},
			{Lang: "hu", Text: "Nem az én asztalom, nem az én szerszámaim."},
		}),
		OnTake: inkwell.Say(speaker.Dex, inkwell.Text{
			{Lang: "en", Text: "Carrying evidence out of a sealed house. Think it through."},
			{Lang: "hu", Text: "Bűnjelet kihordani egy lepecsételt házból. Gondold végig."},
		}),
	}
}

func (noodleSecretRoom) cot() inkwell.Hotspot {
	return inkwell.Hotspot{
		Name: "cot",
		Label: inkwell.Text{
			{Lang: "en", Text: "camp cot"},
			{Lang: "hu", Text: "tábori ágy"},
		},
		Area: inkwell.Rect(480, 150, 130, 70),
		OnLook: inkwell.Say(character.Paul, inkwell.Text{
			{Lang: "en", Text: "He slept back here. Next to the work. Some things don't change."},
			{Lang: "hu", Text: "Itt hátul aludt. A munka mellett. Bizonyos dolgok nem változnak."},
		}),
		OnUse: inkwell.Say(character.Paul, inkwell.Text{
			{Lang: "en", Text: "I'm not tired. I'm something else."},
			{Lang: "hu", Text: "Nem vagyok fáradt. Valami más vagyok."},
		}),
		OnTake: inkwell.Say(character.Paul, inkwell.Text{
			{Lang: "en", Text: "No."},
			{Lang: "hu", Text: "Nem."},
		}),
	}
}

func (noodleSecretRoom) Scene() *inkwell.Scene { return NoodleSecretRoom }
