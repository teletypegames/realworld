package dialog

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/internal/names"
)

// mysteryTapeSilent covers the stretch before the club trigger (beat 6), while
// the tape is still mute. Talking to it is really Dex talking about it — the
// tape itself says nothing.
func mysteryTapeSilent() inkwell.Dialogue {
	return inkwell.Dialogue{
		Name:  names.DlgMysteryTape,
		Start: "root",
		Nodes: []inkwell.DialogueNode{{
			Name:  "root",
			Lines: []inkwell.DialogueLine{{Speaker: names.Dex, Text: "Nothing. Warm, spinning, and silent."}},
			Choices: []inkwell.DialogueChoice{
				{
					Text: "Are you sure it works?",
					Actions: []inkwell.Action{
						inkwell.Say(names.Dex, "It works. It just isn't talking to you. That is not the same as silence."),
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
