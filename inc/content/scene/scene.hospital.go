package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/content/character"
	"realworld/inc/content/dialog"
	"realworld/inc/content/speaker"
	"realworld/inc/system"
)

var Hospital = SceneManager.Register(&Scene{
	Name: "hospital",
	Title: inkwell.Text{
		{Lang: "en", Text: "HOSPITAL — PSYCHIATRIC WING"},
		{Lang: "hu", Text: "KÓRHÁZ — PSZICHIÁTRIAI SZÁRNY"},
	},
	Actors: []inkwell.SceneActor{
		{
			Character: character.Paul,
			At: inkwell.Point{
				X: 120,
				Y: 240,
			},
		},
		{
			Character: character.DoctorMaren,
			At: inkwell.Point{
				X: 420,
				Y: 236,
			},
		},
	},
	Hotspots: []inkwell.Hotspot{
		{
			Name: "doctor_maren",
			Label: inkwell.Text{
				{Lang: "en", Text: "Dr. Maren"},
				{Lang: "hu", Text: "Dr. Maren"},
			},
			Area: inkwell.Rect(396, 156, 48, 84),
			OnLook: inkwell.Say(character.Paul, inkwell.Text{
				{Lang: "en", Text: "White coat, no badge. She reads a chart like it owes her money."},
				{Lang: "hu", Text: "Fehér köpeny, kitűző nélkül. Úgy olvas egy kartont, mintha az tartozna neki."},
			}),
			OnTalk: system.System.Paused(inkwell.RunDialog(dialog.DoctorMaren)),
			OnUse: inkwell.Say(character.Paul, inkwell.Text{
				{Lang: "en", Text: "I'll talk to her. She doesn't look like she's for using."},
				{Lang: "hu", Text: "Beszélek vele. Nem úgy néz ki, mint akit használni lehet."},
			}),
			OnTake: inkwell.Say(speaker.Dex, inkwell.Text{
				{Lang: "en", Text: "No."},
				{Lang: "hu", Text: "Nem."},
			}),
		},
		{
			Name: "empty_bed",
			Label: inkwell.Text{
				{Lang: "en", Text: "empty bed"},
				{Lang: "hu", Text: "üres ágy"},
			},
			Area: inkwell.Rect(180, 150, 160, 90),
			OnLook: inkwell.Seq(
				inkwell.Say(character.Paul, inkwell.Text{
					{Lang: "en", Text: "A chart still clipped to the foot. REDS, N. Admitted, observed, discharged to \"family\". Diagnosis: blank."},
					{Lang: "hu", Text: "Egy karton még az ágy végére csíptetve. REDS, N. Felvéve, megfigyelve, elbocsátva a „családhoz”. Diagnózis: üres."},
				}),
				inkwell.Say(speaker.Dex, inkwell.Text{
					{Lang: "en", Text: "Norman doesn't have family. Not the kind that signs forms."},
					{Lang: "hu", Text: "Normannak nincs családja. Olyan nincs, amelyik papírt ír alá."},
				}),
			),
			OnUse: inkwell.Say(character.Paul, inkwell.Text{
				{Lang: "en", Text: "I'm not lying down in a psychiatric ward. Not voluntarily."},
				{Lang: "hu", Text: "Nem fekszem le egy pszichiátrián. Önként nem."},
			}),
			OnTake: inkwell.Say(character.Paul, inkwell.Text{
				{Lang: "en", Text: "The chart is the hospital's. Taking it makes me a story on that chart."},
				{Lang: "hu", Text: "A karton a kórházé. Ha elviszem, én leszek a következő bejegyzés rajta."},
			}),
		},
		{
			Name: "window",
			Label: inkwell.Text{
				{Lang: "en", Text: "window"},
				{Lang: "hu", Text: "ablak"},
			},
			Area: inkwell.Rect(40, 40, 110, 120),
			OnLook: inkwell.Say(character.Paul, inkwell.Text{
				{Lang: "en", Text: "Frosted glass, wire mesh inside it. The light gets in. Nothing else does."},
				{Lang: "hu", Text: "Tejüveg, benne drótháló. A fény bejut. Más semmi."},
			}),
		},
	},
})

var _ = Hospital.Leads(
	exitToSelector,
)
