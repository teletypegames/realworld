package dialog

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

var _ = DialogManager.Register(Dialog{
	Name:  inc.DialogSumphorTape,
	Start: "root",
	Nodes: []inkwell.DialogNode{{
		Name: "root",
		Lines: []inkwell.DialogLine{{
			Speaker: inc.TapeSumphor,
			Text: inkwell.Text{
				{Lang: "en", Text: "Desk four. Coffee at ten. The fridge on the back wall hums in E. Ask me anything about the room."},
				{Lang: "hu", Text: "Négyes asztal. Kávé tízkor. A hátsó fal hűtője E-ben zümmög. Kérdezz bármit a szobáról."},
			},
		}},
		Choices: []inkwell.DialogChoice{
			{
				Text: inkwell.Text{
					{Lang: "en", Text: "Whose tape are you?"},
					{Lang: "hu", Text: "Kinek a kazettája vagy?"},
				},
				Actions: []inkwell.Action{
					inkwell.Say(inc.TapeSumphor, inkwell.Text{
						{Lang: "en", Text: "I don't remember. I remember the room perfectly and the owner not at all. Is that strange? It feels strange."},
						{Lang: "hu", Text: "Nem emlékszem. A szobára tökéletesen emlékszem, a gazdámra egyáltalán nem. Ez fura? Furának érződik."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: inkwell.Text{
					{Lang: "en", Text: "Tell me about Norman."},
					{Lang: "hu", Text: "Mesélj Normanról."},
				},
				Actions: []inkwell.Action{
					inkwell.Say(inc.TapeSumphor, inkwell.Text{
						{Lang: "en", Text: "He sat over there. He talked to his wrist a lot. More than a man talks to a colleague. It answered him more than I would have."},
						{Lang: "hu", Text: "Ott ült. Sokat beszélt a csuklójához. Többet, mint amennyit egy ember a kollégájához. Az többet válaszolt neki, mint én tettem volna."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: inkwell.Text{
					{Lang: "en", Text: "What about the chair with the helmet?"},
					{Lang: "hu", Text: "És a szék a sisakkal?"},
				},
				Actions: []inkwell.Action{
					inkwell.Say(inc.TapeSumphor, inkwell.Text{
						{Lang: "en", Text: "…I don't discuss the chair."},
						{Lang: "hu", Text: "…a székről nem beszélek."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: inkwell.Text{
					{Lang: "en", Text: "That's enough for now."},
					{Lang: "hu", Text: "Egyelőre ennyi elég."},
				},
				Actions: []inkwell.Action{inkwell.EndDialog()},
			},
		},
	}},
})
