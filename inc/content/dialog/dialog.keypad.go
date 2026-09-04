package dialog

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/content/character"
	"realworld/inc/content/speaker"
)

var Keypad = DialogManager.Register(&Dialog{
	Name:  "keypad",
	Start: "root",
	Nodes: keypadNodes(),
})

var keypadPairs = []string{"22", "27", "72", "77"}

const keypadCode = "2772"

func keypadNodes() []inkwell.DialogNode {
	label := func(pair string) inkwell.Text {
		spaced := string(pair[0]) + " " + string(pair[1])
		return inkwell.Text{
			{Lang: "en", Text: spaced},
			{Lang: "hu", Text: spaced},
		}
	}
	stepAway := inkwell.DialogChoice{
		Text: inkwell.Text{
			{Lang: "en", Text: "Step away."},
			{Lang: "hu", Text: "Hátralépni."},
		},
		Actions: []inkwell.Action{inkwell.EndDialog()},
	}
	attempt := func(code string) []inkwell.Action {
		if code == keypadCode {
			return []inkwell.Action{
				inkwell.Say(character.Paul, inkwell.Text{
					{Lang: "en", Text: "Click."},
					{Lang: "hu", Text: "Katt."},
				}),
				inkwell.Say(speaker.Dex, inkwell.Text{
					{Lang: "en", Text: "Well. Somebody knew that."},
					{Lang: "hu", Text: "Nos. Ezt valaki tudta."},
				}),
				inkwell.SetFlag(inc.FlagSecretRoomOpen),
				inkwell.EndDialog(),
			}
		}
		return []inkwell.Action{
			inkwell.Say(character.Paul, inkwell.Text{
				{Lang: "en", Text: "Nothing."},
				{Lang: "hu", Text: "Semmi."},
			}),
			inkwell.Say(speaker.Dex, inkwell.Text{
				{Lang: "en", Text: "Fifteen to go. Or one, if you'd ask the man who set it."},
				{Lang: "hu", Text: "Még tizenöt. Vagy egy, ha megkérdeznéd azt, aki beállította."},
			}),
			inkwell.GotoNode("root"),
		}
	}
	first := []inkwell.DialogChoice{}
	for _, pair := range keypadPairs {
		first = append(first, inkwell.DialogChoice{
			Text:    label(pair),
			Actions: []inkwell.Action{inkwell.GotoNode("first_" + pair)},
		})
	}
	nodes := []inkwell.DialogNode{{
		Name: "root",
		Lines: []inkwell.DialogLine{{
			Speaker: speaker.Dex,
			Text: inkwell.Text{
				{Lang: "en", Text: "Four digits. Two and seven are worn. First two?"},
				{Lang: "hu", Text: "Négy számjegy. A kettes és a hetes kopott. Az első kettő?"},
			},
		}},
		Choices: append(first, stepAway),
	}}
	for _, firstPair := range keypadPairs {
		second := []inkwell.DialogChoice{}
		for _, secondPair := range keypadPairs {
			second = append(second, inkwell.DialogChoice{
				Text:    label(secondPair),
				Actions: attempt(firstPair + secondPair),
			})
		}
		nodes = append(nodes, inkwell.DialogNode{
			Name: "first_" + firstPair,
			Lines: []inkwell.DialogLine{{
				Speaker: speaker.Dex,
				Text: inkwell.Text{
					{Lang: "en", Text: "Last two."},
					{Lang: "hu", Text: "Az utolsó kettő."},
				},
			}},
			Choices: append(second, stepAway),
		})
	}
	return nodes
}
