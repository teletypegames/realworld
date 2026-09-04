package dialog

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/content/speaker"
)

var SupportTape = DialogManager.Register(&Dialog{
	Name:  "support_tape",
	Start: "root",
	Nodes: []DialogNode{{
		Name: "root",
		Lines: []DialogLine{{
			Speaker: speaker.TapeSupport,
			Text: Text{
				{Lang: "en", Text: "Thank you for holding. My name is Raj. May I have your ticket number?"},
				{Lang: "hu", Text: "Köszönöm, hogy tartotta a vonalat. A nevem Raj. Megkaphatnám a hibajegy számát?"},
			},
		}},
		Choices: []DialogChoice{
			{
				Text: Text{
					{Lang: "en", Text: "I don't have a ticket number."},
					{Lang: "hu", Text: "Nincs hibajegyszámom."},
				},
				Actions: []Action{
					speaker.TapeSupport.Say(Text{
						{Lang: "en", Text: "No problem, sir. I have opened a ticket for you. Your ticket number is one. Please quote it on future calls."},
						{Lang: "hu", Text: "Semmi gond, uram. Nyitottam önnek egy hibajegyet. A hibajegy száma: egy. Kérem, hivatkozzon rá a jövőben."},
					}),
					speaker.Dex.Say(Text{
						{Lang: "en", Text: "It's the same ticket every time. I've counted."},
						{Lang: "hu", Text: "Mindig ugyanaz a hibajegy. Számoltam."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: Text{
					{Lang: "en", Text: "There's a mirror that flickers."},
					{Lang: "hu", Text: "Van egy tükör, ami vibrál."},
				},
				Show: Given{
					NotFlag: inc.FlagMirrorPassword,
				},
				Actions: []Action{
					speaker.TapeSupport.Say(Text{
						{Lang: "en", Text: "I understand, sir, and I apologise for the inconvenience. Have you tried turning it off and on again?"},
						{Lang: "hu", Text: "Értem, uram, és elnézését kérem a kellemetlenségért. Próbálta már ki- és bekapcsolni?"},
					}),
					speaker.Dex.Say(Text{
						{Lang: "en", Text: "Of course it said that. It says that to the weather."},
						{Lang: "hu", Text: "Persze hogy ezt mondta. Az időjárásnak is ezt mondja."},
					}),
					speaker.TapeSupport.Say(Text{
						{Lang: "en", Text: "Please make sure I am connected when you do, sir. The procedure is only supported with an agent on the line."},
						{Lang: "hu", Text: "Kérem, gondoskodjon róla, hogy csatlakoztatva legyek, amikor megteszi, uram. Az eljárást csak ügyintézővel a vonalon támogatjuk."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: Text{
					{Lang: "en", Text: "There's a fridge I need to get behind."},
					{Lang: "hu", Text: "Van egy hűtő, ami mögé be kell jutnom."},
				},
				Show: Given{
					Flag:    inc.FlagAbandonedKnown,
					NotFlag: inc.FlagFridgeOpen,
				},
				Actions: []Action{
					speaker.TapeSupport.Say(Text{
						{Lang: "en", Text: "A refrigeration unit, sir. Is it plugged in? Is the light on? Have you tried turning it off and on again?"},
						{Lang: "hu", Text: "Egy hűtőegység, uram. Be van dugva? Ég a lámpa? Próbálta már ki- és bekapcsolni?"},
					}),
					speaker.Dex.Say(Text{
						{Lang: "en", Text: "It worked on the mirror. I'm not saying it'll work on a fridge. I'm saying I've stopped ruling it out."},
						{Lang: "hu", Text: "A tükörnél bevált. Nem mondom, hogy egy hűtőnél is. Azt mondom, hogy már nem zárom ki."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: Text{
					{Lang: "en", Text: "There's a security system I need kept busy."},
					{Lang: "hu", Text: "Van egy biztonsági rendszer, amit le kéne kötni."},
				},
				Show: Given{
					Flag: inc.FlagLabLocated,
				},
				Actions: []Action{
					speaker.TapeSupport.Say(Text{
						{Lang: "en", Text: "Sir, I can keep any system on the line for hours. Verification, hold music, a survey at the end. It is my one talent."},
						{Lang: "hu", Text: "Uram, bármilyen rendszert órákig a vonalon tudok tartani. Azonosítás, várakozási zene, kérdőív a végén. Ez az egyetlen tehetségem."},
					}),
					speaker.Dex.Say(Text{
						{Lang: "en", Text: "For once I believe every word."},
						{Lang: "hu", Text: "Egyszer az életben minden szavát elhiszem."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: Text{
					{Lang: "en", Text: "What do you actually know?"},
					{Lang: "hu", Text: "Mit tudsz te egyáltalán?"},
				},
				Actions: []Action{
					speaker.TapeSupport.Say(Text{
						{Lang: "en", Text: "Everything, sir, at a basic level. An Armilla is worn on the wrist. A tape goes in the slot. Water is wet. Would you like me to read the warranty?"},
						{Lang: "hu", Text: "Mindent, uram, alapszinten. Az Armillát a csuklón hordják. A kazetta a foglalatba megy. A víz nedves. Felolvassam a garanciát?"},
					}),
					speaker.Dex.Say(Text{
						{Lang: "en", Text: "I'm four kilobytes and I know more than that. Paul, take it out."},
						{Lang: "hu", Text: "Négy kilobájt vagyok, és többet tudok ennél. Paul, vedd ki."},
					}),
					speaker.TapeSupport.Say(Text{
						{Lang: "en", Text: "The other agent is not authorised on this ticket, sir."},
						{Lang: "hu", Text: "A másik ügyintéző nem jogosult ezen a hibajegyen, uram."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: Text{
					{Lang: "en", Text: "Goodbye, Raj."},
					{Lang: "hu", Text: "Viszlát, Raj."},
				},
				Actions: []Action{
					speaker.TapeSupport.Say(Text{
						{Lang: "en", Text: "Thank you for calling. Please stay on the line for a short survey."},
						{Lang: "hu", Text: "Köszönjük a hívását. Kérem, maradjon a vonalban egy rövid kérdőívhez."},
					}),
					inkwell.EndDialog(),
				},
			},
		},
	}},
})
