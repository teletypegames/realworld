package dialog

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
)

func init() {
	Manager.Register(Dialog{
		Name:  inc.DlgMysteryTape,
		Start: "root",
		Nodes: []inkwell.DialogueNode{{
			Name: "root",
			Lines: []inkwell.DialogueLine{{
				Speaker: inc.Dex,
				Text:    "Nothing. Warm, spinning, and silent.",
			}},
			Choices: []inkwell.DialogueChoice{
				{
					Text: "Are you sure it works?",
					Actions: []inkwell.Action{
						inkwell.Say(inc.Dex, "It works. It just isn't talking to you. That is not the same as silence."),
						inkwell.GotoNode("root"),
					},
				},
				{
					Text:    "Fine. It'll speak when it speaks.",
					Actions: []inkwell.Action{inkwell.EndDialogue()},
				},
			},
		}},
	})
}
