package dialog

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/content/character"
	"realworld/inc/content/speaker"
)

var DoctorMaren = DialogManager.Register(&Dialog{
	Name:  "doctor_maren",
	Start: "root",
	Nodes: []DialogNode{{
		Name: "root",
		Lines: []DialogLine{{
			Speaker: character.DoctorMaren,
			Text: Text{
				{Lang: "en", Text: "Visiting hours are over. Unless you're here about a patient I no longer have."},
				{Lang: "hu", Text: "A látogatási idő lejárt. Hacsak nem egy olyan beteg miatt jött, aki már nincs nálam."},
			},
		}},
		Choices: []DialogChoice{
			{
				Text: Text{
					{Lang: "en", Text: "Norman Reds. You examined him."},
					{Lang: "hu", Text: "Norman Reds. Ön vizsgálta."},
				},
				Once: true,
				Actions: []Action{
					character.DoctorMaren.Say(Text{
						{Lang: "en", Text: "For eleven hours, before they collected him. No neurological cause. No toxin. Bloodwork clean."},
						{Lang: "hu", Text: "Tizenegy órán át, mielőtt elvitték. Semmi neurológiai ok. Semmi toxin. Tiszta vérkép."},
					}),
					character.DoctorMaren.Say(Text{
						{Lang: "en", Text: "The EEG looked like deep REM. Except REM ends. His didn't."},
						{Lang: "hu", Text: "Az EEG mély REM-nek látszott. Csak a REM véget ér. Az övé nem."},
					}),
					inkwell.SetFlag(inc.FlagMarenTalked),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: Text{
					{Lang: "en", Text: "What did the official description say?"},
					{Lang: "hu", Text: "Mit mondott a hivatalos leírás?"},
				},
				Show: Given{
					Flag: inc.FlagMarenTalked,
				},
				Actions: []Action{
					character.DoctorMaren.Say(Text{
						{Lang: "en", Text: "\"Exhaustion, workplace related. Recovering with family.\" I was handed it typed. Nobody types a diagnosis before the doctor writes one."},
						{Lang: "hu", Text: "„Kimerültség, munkahelyi eredetű. Családja körében lábadozik.” Gépelve kaptam. Diagnózist senki nem gépel le, mielőtt az orvos megírja."},
					}),
					character.DoctorMaren.Say(Text{
						{Lang: "en", Text: "It doesn't match what I saw. That's as far as I go with a stranger."},
						{Lang: "hu", Text: "Nem egyezik azzal, amit láttam. Egy idegennek ennyit mondok, nem többet."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: Text{
					{Lang: "en", Text: "I know where he is. Inside the Golden Gate, wired into a machine."},
					{Lang: "hu", Text: "Tudom, hol van. A Golden Gate belsejében, egy gépbe kötve."},
				},
				Show: Given{
					Flag: inc.FlagNormanFound,
				},
				Once: true,
				Actions: []Action{
					character.DoctorMaren.Say(Text{
						{Lang: "en", Text: "Then listen to me, because I am only saying this once. A brain held in that state does not rest. It burns."},
						{Lang: "hu", Text: "Akkor hallgasson meg, mert csak egyszer mondom. Egy agy, amit abban az állapotban tartanak, nem pihen. Ég."},
					}),
					character.DoctorMaren.Say(Text{
						{Lang: "en", Text: "Get him out of whatever is holding it. Not gently. Soon."},
						{Lang: "hu", Text: "Hozza ki abból, ami bent tartja. Ne finoman. Hamar."},
					}),
					speaker.Dex.Say(Text{
						{Lang: "en", Text: "That's the first doctor I've heard prescribe hurry."},
						{Lang: "hu", Text: "Ez az első orvos, akitől sietséget hallok felírni."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: Text{
					{Lang: "en", Text: "Who collected him?"},
					{Lang: "hu", Text: "Ki vitte el?"},
				},
				Show: Given{
					Flag: inc.FlagMarenTalked,
				},
				Actions: []Action{
					character.DoctorMaren.Say(Text{
						{Lang: "en", Text: "Two people with the right paperwork and no names I could read. The paperwork had a seal on it I've only seen on licences."},
						{Lang: "hu", Text: "Két ember a megfelelő papírokkal és olvashatatlan nevekkel. A papírokon olyan pecsét volt, amit eddig csak licenceken láttam."},
					}),
					character.Paul.Say(Text{
						{Lang: "en", Text: "BVK."},
						{Lang: "hu", Text: "BVK."},
					}),
					character.DoctorMaren.Say(Text{
						{Lang: "en", Text: "I didn't say that."},
						{Lang: "hu", Text: "Ezt nem én mondtam."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: Text{
					{Lang: "en", Text: "Thank you, doctor."},
					{Lang: "hu", Text: "Köszönöm, doktornő."},
				},
				Actions: []Action{
					character.DoctorMaren.Say(Text{
						{Lang: "en", Text: "You were never here. That's how this ward works."},
						{Lang: "hu", Text: "Ön itt sem volt. Ez az osztály így működik."},
					}),
					inkwell.EndDialog(),
				},
			},
		},
	}},
})
