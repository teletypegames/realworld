package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/content/character"
	"realworld/inc/content/item"
	"realworld/inc/content/script"
	"realworld/inc/content/speaker"
)

var noodleSecretRoomTake = inkwell.Seq(
	inkwell.Give(item.CommentaryTape),
	character.Paul.Say(Text{
		{Lang: "en", Text: "A tape. Label says PAUL, in that handwriting."},
		{Lang: "hu", Text: "Egy kazetta. A címkén PAUL, azzal a kézírással."},
	}),
	speaker.Dex.Say(Text{
		{Lang: "en", Text: "He labelled it for you and put it where only you'd look. That's the most sentimental thing Noodle ever did."},
		{Lang: "hu", Text: "Neked címkézte, és oda tette, ahol csak te keresnéd. Ez a legérzelmesebb dolog, amit Noodle valaha tett."},
	}),
)

var noodleSecretRoomGone = character.Paul.Say(Text{
	{Lang: "en", Text: "Empty shelf. An evidence tag where something used to sit."},
	{Lang: "hu", Text: "Üres polc. Egy bűnjelcédula ott, ahol valami állt."},
})

var NoodleSecretRoom = SceneManager.Register(&Scene{
	Name: "noodle_secret_room",
	Title: Text{
		{Lang: "en", Text: "NOODLE'S BACK ROOM"},
		{Lang: "hu", Text: "NOODLE HÁTSÓ SZOBÁJA"},
	},
	Triggers: []Trigger{
		{
			Name: "raided",
			When: Given{
				Flag:    inc.FlagNoodleCode,
				NotFlag: inc.FlagRoomRaided,
			},
			Do:   inkwell.RunScript(script.SecretRoomEmpty),
			Once: true,
		},
	},
	Hotspots: []Hotspot{
		{
			Name: "shelf",
			Label: Text{
				{Lang: "en", Text: "shelf over the desk"},
				{Lang: "hu", Text: "polc az asztal fölött"},
			},
			Area: inkwell.Rect(260, 70, 140, 50),
			OnLook: inkwell.If(inkwell.Flag(inc.FlagNoodleCode),
				noodleSecretRoomGone,
				inkwell.If(inkwell.HasItem(item.CommentaryTape),
					character.Paul.Say(Text{
						{Lang: "en", Text: "Dust, and the clean rectangle where the tape was."},
						{Lang: "hu", Text: "Por, és egy tiszta téglalap, ahol a kazetta volt."},
					}),
					character.Paul.Say(Text{
						{Lang: "en", Text: "One tape, set apart from the rest. Somebody wanted it seen."},
						{Lang: "hu", Text: "Egy kazetta, külön a többitől. Valaki azt akarta, hogy észrevegyék."},
					}),
				),
			),
			OnUse: inkwell.If(inkwell.Flag(inc.FlagNoodleCode),
				noodleSecretRoomGone,
				inkwell.If(inkwell.HasItem(item.CommentaryTape),
					character.Paul.Say(Text{
						{Lang: "en", Text: "Nothing else on it worth having."},
						{Lang: "hu", Text: "Más nincs rajta, amit érdemes lenne."},
					}),
					noodleSecretRoomTake,
				),
			),
			OnTake: inkwell.If(inkwell.Flag(inc.FlagNoodleCode),
				noodleSecretRoomGone,
				inkwell.If(inkwell.HasItem(item.CommentaryTape),
					character.Paul.Say(Text{
						{Lang: "en", Text: "I've got it."},
						{Lang: "hu", Text: "Megvan."},
					}),
					noodleSecretRoomTake,
				),
			),
		},
		{
			Name: "workbench",
			Label: Text{
				{Lang: "en", Text: "workbench of cracked Armillas"},
				{Lang: "hu", Text: "munkaasztal feltört Armillákkal"},
			},
			Area: inkwell.Rect(220, 130, 220, 70),
			OnLook: inkwell.If(inkwell.Flag(inc.FlagNoodleCode),
				character.Paul.Say(Text{
					{Lang: "en", Text: "Stripped. Even the soldering iron's gone."},
					{Lang: "hu", Text: "Kipucolva. Még a forrasztópáka is eltűnt."},
				}),
				inkwell.Seq(
					character.Paul.Say(Text{
						{Lang: "en", Text: "Six Armillas open on the bench, every one of them mid-surgery."},
						{Lang: "hu", Text: "Hat felnyitott Armilla az asztalon, mind műtét közben."},
					}),
					speaker.Dex.Say(Text{
						{Lang: "en", Text: "This is the business. The letter was about something else."},
						{Lang: "hu", Text: "Ez az üzlet. A levél másról szólt."},
					}),
				),
			),
			OnUse: character.Paul.Say(Text{
				{Lang: "en", Text: "Not my bench, not my tools."},
				{Lang: "hu", Text: "Nem az én asztalom, nem az én szerszámaim."},
			}),
			OnTake: speaker.Dex.Say(Text{
				{Lang: "en", Text: "Carrying evidence out of a sealed house. Think it through."},
				{Lang: "hu", Text: "Bűnjelet kihordani egy lepecsételt házból. Gondold végig."},
			}),
		},
		{
			Name: "cot",
			Label: Text{
				{Lang: "en", Text: "camp cot"},
				{Lang: "hu", Text: "tábori ágy"},
			},
			Area: inkwell.Rect(480, 150, 130, 70),
			OnLook: character.Paul.Say(Text{
				{Lang: "en", Text: "He slept back here. Next to the work. Some things don't change."},
				{Lang: "hu", Text: "Itt hátul aludt. A munka mellett. Bizonyos dolgok nem változnak."},
			}),
			OnUse: character.Paul.Say(Text{
				{Lang: "en", Text: "I'm not tired. I'm something else."},
				{Lang: "hu", Text: "Nem vagyok fáradt. Valami más vagyok."},
			}),
			OnTake: character.Paul.Say(Text{
				{Lang: "en", Text: "No."},
				{Lang: "hu", Text: "Nem."},
			}),
		},
	},
})

var _ = NoodleSecretRoom.Leads(
	Exit{
		To: Alley,
		Label: Text{
			{Lang: "en", Text: "back out to the alley"},
			{Lang: "hu", Text: "vissza a sikátorba"},
		},
		Side: inkwell.ExitNear,
	},
)
