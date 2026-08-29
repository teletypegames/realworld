package dialog

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/internal/names"
)

// dexTalk is the standing conversation with Dex. Its branches are gated on
// story flags, so the same entry point stays useful as the game progresses.
func dexTalk() inkwell.Dialogue {
	return inkwell.Dialogue{
		Name:  names.DlgDex,
		Start: "root",
		Nodes: []inkwell.DialogueNode{{
			Name:  "root",
			Lines: []inkwell.DialogueLine{{Speaker: names.Dex, Text: "Talk."}},
			Choices: []inkwell.DialogueChoice{
				{
					Text: "What am I doing here, Dex?",
					Actions: []inkwell.Action{
						inkwell.Say(names.Dex, "You got a letter from a man you hadn't seen in twenty years. And you came. That says more about you than it does about him."),
						inkwell.GotoNode("root"),
					},
				},
				{
					Text: "What do you know about Noodle?",
					Show: inkwell.Not(inkwell.Flag(names.FlagHasTape)),
					Actions: []inkwell.Action{
						inkwell.Say(names.Dex, "He traded cracked Armillas. That isn't charity, Paul, that's a business. The kind they take you in for."),
						inkwell.GotoNode("root"),
					},
				},
				{
					Text: "What do you make of this tape?",
					Show: inkwell.Flag(names.FlagHasTape),
					Actions: []inkwell.Action{
						inkwell.Say(names.Dex, "I make of it that you found a password-locked tape in a gap between two bins, on a tip from a government office. Count how many times that has ended well."),
						inkwell.GotoNode("root"),
					},
				},
				{
					Text: "I miss you.",
					Once: true,
					Actions: []inkwell.Action{
						inkwell.Say(names.Dex, "I'm four kilobytes of a dead man. Don't do this to yourself."),
						inkwell.GotoNode("root"),
					},
				},
				{
					Text:    "Nothing. Forget it.",
					Actions: []inkwell.Action{inkwell.EndDialogue()},
				},
			},
		}},
	}
}
