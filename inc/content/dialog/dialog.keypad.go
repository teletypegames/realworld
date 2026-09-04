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

func keypadNodes() []DialogNode {
	label := func(pair string) Text {
		spaced := string(pair[0]) + " " + string(pair[1])
		return Text{
			{Lang: "en", Text: spaced},
			{Lang: "hu", Text: spaced},
		}
	}
	stepAway := DialogChoice{
		Text: Text{
			{Lang: "en", Text: "Step away."},
			{Lang: "hu", Text: "Hátralépni."},
		},
		Actions: []Action{inkwell.EndDialog()},
	}
	attempt := func(code string) []Action {
		if code == keypadCode {
			return []Action{
				character.Paul.Say(Text{
					{Lang: "en", Text: "Click."},
					{Lang: "hu", Text: "Katt."},
				}),
				speaker.Dex.Say(Text{
					{Lang: "en", Text: "Well. Somebody knew that."},
					{Lang: "hu", Text: "Nos. Ezt valaki tudta."},
				}),
				inkwell.SetFlag(inc.FlagSecretRoomOpen),
				inkwell.EndDialog(),
			}
		}
		return []Action{
			character.Paul.Say(Text{
				{Lang: "en", Text: "Nothing."},
				{Lang: "hu", Text: "Semmi."},
			}),
			speaker.Dex.Say(Text{
				{Lang: "en", Text: "Fifteen to go. Or one, if you'd ask the man who set it."},
				{Lang: "hu", Text: "Még tizenöt. Vagy egy, ha megkérdeznéd azt, aki beállította."},
			}),
			inkwell.GotoNode("root"),
		}
	}
	first := []DialogChoice{}
	for _, pair := range keypadPairs {
		first = append(first, DialogChoice{
			Text:    label(pair),
			Actions: []Action{inkwell.GotoNode("first_" + pair)},
		})
	}
	nodes := []DialogNode{{
		Name: "root",
		Lines: []DialogLine{{
			Speaker: speaker.Dex,
			Text: Text{
				{Lang: "en", Text: "Four digits. Two and seven are worn. First two?"},
				{Lang: "hu", Text: "Négy számjegy. A kettes és a hetes kopott. Az első kettő?"},
			},
		}},
		Choices: append(first, stepAway),
	}}
	for _, firstPair := range keypadPairs {
		second := []DialogChoice{}
		for _, secondPair := range keypadPairs {
			second = append(second, DialogChoice{
				Text:    label(secondPair),
				Actions: attempt(firstPair + secondPair),
			})
		}
		nodes = append(nodes, DialogNode{
			Name: "first_" + firstPair,
			Lines: []DialogLine{{
				Speaker: speaker.Dex,
				Text: Text{
					{Lang: "en", Text: "Last two."},
					{Lang: "hu", Text: "Az utolsó kettő."},
				},
			}},
			Choices: append(second, stepAway),
		})
	}
	return nodes
}
