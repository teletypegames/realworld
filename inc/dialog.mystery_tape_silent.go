package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func dialogMysteryTapeSilent() inkwell.Dialogue {
	return inkwell.Dialogue{
		Name:  DlgMysteryTape,
		Start: "root",
		Nodes: []inkwell.DialogueNode{{
			Name: "root",
			Lines: []inkwell.DialogueLine{{
				Speaker: Dex,
				Text:    "Nothing. Warm, spinning, and silent.",
			}},
			Choices: []inkwell.DialogueChoice{
				{
					Text: "Are you sure it works?",
					Actions: []inkwell.Action{
						inkwell.Say(Dex, "It works. It just isn't talking to you. That is not the same as silence."),
						inkwell.GotoNode("root"),
					},
				},
				{
					Text:    "Fine. It'll speak when it speaks.",
					Actions: []inkwell.Action{inkwell.EndDialogue()},
				},
			},
		}},
	}
}
