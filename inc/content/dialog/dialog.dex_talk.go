package dialog

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/content/speaker"
)

var Dex = DialogManager.Register(&Dialog{
	Name:  "dex_talk",
	Start: "root",
	Nodes: []DialogNode{{
		Name: "root",
		Lines: []DialogLine{{
			Speaker: speaker.Dex,
			Text: Text{
				{Lang: "en", Text: "Talk."},
				{Lang: "hu", Text: "Beszélj."},
			},
		}},
		Choices: []DialogChoice{
			{
				Text: Text{
					{Lang: "en", Text: "What am I doing here, Dex?"},
					{Lang: "hu", Text: "Mit keresek én itt, Dex?"},
				},
				Actions: []Action{
					speaker.Dex.Say(Text{
						{Lang: "en", Text: "You got a letter from a man you hadn't seen in twenty years. And you came. That says more about you than it does about him."},
						{Lang: "hu", Text: "Levelet kaptál egy embertől, akit húsz éve nem láttál. És eljöttél. Ez rólad mond többet, mint róla."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: Text{
					{Lang: "en", Text: "What do you know about Noodle?"},
					{Lang: "hu", Text: "Mit tudsz Noodle-ról?"},
				},
				Show: Given{
					NotFlag: inc.FlagHasTape,
				},
				Actions: []Action{
					speaker.Dex.Say(Text{
						{Lang: "en", Text: "He traded cracked Armillas. That isn't charity, Paul, that's a business. The kind they take you in for."},
						{Lang: "hu", Text: "Feltört Armillákkal kereskedett. Az nem jótékonyság, Paul, az üzlet. Olyan, amiért bevisznek."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: Text{
					{Lang: "en", Text: "What do you make of this tape?"},
					{Lang: "hu", Text: "Mit gondolsz erről a kazettáról?"},
				},
				Show: Given{
					Flag: inc.FlagHasTape,
				},
				Actions: []Action{
					speaker.Dex.Say(Text{
						{Lang: "en", Text: "I make of it that you found a password-locked tape in a gap between two bins, on a tip from a government office. Count how many times that has ended well."},
						{Lang: "hu", Text: "Azt, hogy találtál egy jelszóval zárt kazettát két kuka közti résben, egy állami hivatal fülesére. Számold meg, hányszor végződött ez jól."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: Text{
					{Lang: "en", Text: "I miss you."},
					{Lang: "hu", Text: "Hiányzol."},
				},
				Once: true,
				Actions: []Action{
					speaker.Dex.Say(Text{
						{Lang: "en", Text: "I'm four kilobytes of a dead man. Don't do this to yourself."},
						{Lang: "hu", Text: "Egy halott ember négy kilobájtja vagyok. Ne tedd ezt magaddal."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: Text{
					{Lang: "en", Text: "Nothing. Forget it."},
					{Lang: "hu", Text: "Semmit. Hagyjuk."},
				},
				Actions: []Action{inkwell.EndDialog()},
			},
		},
	}},
})
