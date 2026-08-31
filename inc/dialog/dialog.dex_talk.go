package dialog

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
)

func init() {
	Manager.Register(Dialog{
		Name:  inc.DlgDex,
		Start: "root",
		Nodes: []inkwell.DialogueNode{{
			Name: "root",
			Lines: []inkwell.DialogueLine{{
				Speaker: inc.Dex,
				Text: inkwell.Text{
					{Lang: "en", Text: "Talk."},
					{Lang: "hu", Text: "Beszélj."},
				},
			}},
			Choices: []inkwell.DialogueChoice{
				{
					Text: inkwell.Text{
						{Lang: "en", Text: "What am I doing here, Dex?"},
						{Lang: "hu", Text: "Mit keresek én itt, Dex?"},
					},
					Actions: []inkwell.Action{
						inkwell.Say(inc.Dex, inkwell.Text{
							{Lang: "en", Text: "You got a letter from a man you hadn't seen in twenty years. And you came. That says more about you than it does about him."},
							{Lang: "hu", Text: "Levelet kaptál egy embertől, akit húsz éve nem láttál. És eljöttél. Ez rólad mond többet, mint róla."},
						}),
						inkwell.GotoNode("root"),
					},
				},
				{
					Text: inkwell.Text{
						{Lang: "en", Text: "What do you know about Noodle?"},
						{Lang: "hu", Text: "Mit tudsz Noodle-ról?"},
					},
					Show: inkwell.Not(inkwell.Flag(inc.FlagHasTape)),
					Actions: []inkwell.Action{
						inkwell.Say(inc.Dex, inkwell.Text{
							{Lang: "en", Text: "He traded cracked Armillas. That isn't charity, Paul, that's a business. The kind they take you in for."},
							{Lang: "hu", Text: "Feltört Armillákkal kereskedett. Az nem jótékonyság, Paul, az üzlet. Olyan, amiért bevisznek."},
						}),
						inkwell.GotoNode("root"),
					},
				},
				{
					Text: inkwell.Text{
						{Lang: "en", Text: "What do you make of this tape?"},
						{Lang: "hu", Text: "Mit gondolsz erről a kazettáról?"},
					},
					Show: inkwell.Flag(inc.FlagHasTape),
					Actions: []inkwell.Action{
						inkwell.Say(inc.Dex, inkwell.Text{
							{Lang: "en", Text: "I make of it that you found a password-locked tape in a gap between two bins, on a tip from a government office. Count how many times that has ended well."},
							{Lang: "hu", Text: "Azt, hogy találtál egy jelszóval zárt kazettát két kuka közti résben, egy állami hivatal fülesére. Számold meg, hányszor végződött ez jól."},
						}),
						inkwell.GotoNode("root"),
					},
				},
				{
					Text: inkwell.Text{
						{Lang: "en", Text: "I miss you."},
						{Lang: "hu", Text: "Hiányzol."},
					},
					Once: true,
					Actions: []inkwell.Action{
						inkwell.Say(inc.Dex, inkwell.Text{
							{Lang: "en", Text: "I'm four kilobytes of a dead man. Don't do this to yourself."},
							{Lang: "hu", Text: "Egy halott ember négy kilobájtja vagyok. Ne tedd ezt magaddal."},
						}),
						inkwell.GotoNode("root"),
					},
				},
				{
					Text: inkwell.Text{
						{Lang: "en", Text: "Nothing. Forget it."},
						{Lang: "hu", Text: "Semmit. Hagyjuk."},
					},
					Actions: []inkwell.Action{inkwell.EndDialogue()},
				},
			},
		}},
	})
}
