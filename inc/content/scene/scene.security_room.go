package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/content/character"
	"realworld/inc/content/dialog"
	"realworld/inc/content/item"
	"realworld/inc/content/script"
	"realworld/inc/content/speaker"
	"realworld/inc/system"
)

var securityRoomRefused = inkwell.Seq(
	character.SecurityACP.Say(Text{
		{Lang: "en", Text: "ACCESS DENIED. THE SIMULATION MAY NOT BE INTERRUPTED."},
		{Lang: "hu", Text: "HOZZÁFÉRÉS MEGTAGADVA. A SZIMULÁCIÓ NEM SZAKÍTHATÓ MEG."},
	}),
	character.SecurityACP.Say(Text{
		{Lang: "en", Text: "NO TAPE MAY BE CARRIED IN. STANDING ORDER OF THE OPERATOR."},
		{Lang: "hu", Text: "KAZETTA NEM VIHETŐ BE. AZ OPERÁTOR ÁLLANDÓ UTASÍTÁSA."},
	}),
)

var SecurityRoom = SceneManager.Register(&Scene{
	Name: "security_room",
	Title: Text{
		{Lang: "en", Text: "OBSERVATION ROOM — INSIDE THE STRUCTURE"},
		{Lang: "hu", Text: "MEGFIGYELŐSZOBA — A SZERKEZET BELSEJÉBEN"},
	},
	Layers: []SceneLayer{
		{
			Asset: &Asset{
				Name: "security_room",
				Path: "assets/bg/security_room.png",
			},
			Z: -1,
		},
	},
	Triggers: []Trigger{
		{
			Name: "found",
			When: inkwell.Not(inkwell.Flag(inc.FlagNormanFound)),
			Do:   inkwell.RunScript(script.NormanFound),
			Once: true,
		},
	},
	Hotspots: []Hotspot{
		{
			Name: "glass_wall",
			Label: Text{
				{Lang: "en", Text: "one-way glass"},
				{Lang: "hu", Text: "tükörüveg"},
			},
			Area: inkwell.Rect(40, 40, 300, 150),
			OnLook: character.Paul.Say(Text{
				{Lang: "en", Text: "Norman. In a chair, cables running out of a helmet. A monitor behind him, playing something."},
				{Lang: "hu", Text: "Norman. Egy székben, kábelek futnak ki egy sisakból. Mögötte egy monitor, valami megy rajta."},
			}),
			OnUse: character.Paul.Say(Text{
				{Lang: "en", Text: "It's a wall that happens to be glass. It doesn't open."},
				{Lang: "hu", Text: "Egy fal, ami történetesen üveg. Nem nyílik."},
			}),
			OnTalk: speaker.Dex.Say(Text{
				{Lang: "en", Text: "He can't hear you. He can't hear anything on this side."},
				{Lang: "hu", Text: "Nem hall téged. Semmit nem hall ezen az oldalon."},
			}),
		},
		{
			Name: "monitors",
			Label: Text{
				{Lang: "en", Text: "bank of monitors"},
				{Lang: "hu", Text: "monitorfal"},
			},
			Area: inkwell.Rect(360, 40, 250, 110),
			OnLook: character.Paul.Say(Text{
				{Lang: "en", Text: "The same sleeping man from six angles. One camera is pointed at the monitor behind him."},
				{Lang: "hu", Text: "Ugyanaz az alvó ember hat szögből. Egy kamera a mögötte lévő monitort nézi."},
			}),
			OnUse: speaker.Dex.Say(Text{
				{Lang: "en", Text: "Look, don't touch. Somebody is very interested in every twitch he makes."},
				{Lang: "hu", Text: "Nézd, ne fogd. Valakit nagyon érdekel minden rándulása."},
			}),
		},
		{
			Name: "acp_terminal",
			Label: Text{
				{Lang: "en", Text: "security terminal"},
				{Lang: "hu", Text: "biztonsági terminál"},
			},
			Area: inkwell.Rect(420, 160, 150, 90),
			OnLook: character.Paul.Say(Text{
				{Lang: "en", Text: "Green phosphor. One line, repeating: ACCESS DENIED. There's an ACP behind it, and it's listening."},
				{Lang: "hu", Text: "Zöld foszfor. Egy sor, ismétlődik: HOZZÁFÉRÉS MEGTAGADVA. Egy ACP van mögötte, és figyel."},
			}),
			OnTalk: system.System.Paused(inkwell.RunDialog(dialog.SecurityACP)),
			OnUse: inkwell.If(inkwell.Flag(inc.FlagACPDistracted),
				character.Paul.Say(Text{
					{Lang: "en", Text: "Door release is unlocked. It's still on hold with Raj."},
					{Lang: "hu", Text: "Az ajtónyitó szabad. Még mindig várakoztatja Raj."},
				}),
				inkwell.If(inkwell.VarEq(inc.VarSlot2, item.SupportTape.Name),
					inkwell.RunScript(script.ACPDistracted),
					securityRoomRefused,
				),
			),
			OnUseWithAny: character.SecurityACP.Say(Text{
				{Lang: "en", Text: "FOREIGN OBJECT. ACCESS DENIED."},
				{Lang: "hu", Text: "IDEGEN TÁRGY. HOZZÁFÉRÉS MEGTAGADVA."},
			}),
		},
	},
})

var _ = SecurityRoom.Leads(
	Exit{
		To: BridgeFoot,
		Label: Text{
			{Lang: "en", Text: "back down"},
			{Lang: "hu", Text: "vissza le"},
		},
		Side: inkwell.ExitNear,
	},
	Exit{
		To: SecretLab,
		Label: Text{
			{Lang: "en", Text: "the chamber door"},
			{Lang: "hu", Text: "a kamra ajtaja"},
		},
		Side:  inkwell.ExitBack,
		Needs: inc.FlagACPDistracted,
		Blocked: character.SecurityACP.Say(Text{
			{Lang: "en", Text: "ACCESS DENIED."},
			{Lang: "hu", Text: "HOZZÁFÉRÉS MEGTAGADVA."},
		}),
	},
)
