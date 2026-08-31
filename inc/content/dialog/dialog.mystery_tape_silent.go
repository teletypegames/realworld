package dialog

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

func init() {
	DialogManager.Register(Dialog{
		Name:  inc.DlgMysteryTape,
		Start: "root",
		Nodes: []inkwell.DialogueNode{{
			Name: "root",
			Lines: []inkwell.DialogueLine{{
				Speaker: inc.Dex,
				Text: inkwell.Text{
					{Lang: "en", Text: "Nothing. Warm, spinning, and silent."},
					{Lang: "hu", Text: "Semmi. Meleg, forog, és néma."},
				},
			}},
			Choices: []inkwell.DialogueChoice{
				{
					Text: inkwell.Text{
						{Lang: "en", Text: "Are you sure it works?"},
						{Lang: "hu", Text: "Biztos, hogy működik?"},
					},
					Actions: []inkwell.Action{
						inkwell.Say(inc.Dex, inkwell.Text{
							{Lang: "en", Text: "It works. It just isn't talking to you. That is not the same as silence."},
							{Lang: "hu", Text: "Működik. Csak nem beszél hozzád. Ez nem ugyanaz, mint a csend."},
						}),
						inkwell.GotoNode("root"),
					},
				},
				{
					Text: inkwell.Text{
						{Lang: "en", Text: "Fine. It'll speak when it speaks."},
						{Lang: "hu", Text: "Jó. Majd megszólal, ha megszólal."},
					},
					Actions: []inkwell.Action{inkwell.EndDialogue()},
				},
			},
		}},
	})
}
