package dialog

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

type keypad struct{}

var _ = DialogManager.Register(Dialog{
	Name:  inc.DialogKeypad,
	Start: "root",
	Nodes: keypad{}.nodes(),
})

var keypadPairs = []string{"22", "27", "72", "77"}

const keypadCode = "2772"

func (pad keypad) nodes() []inkwell.DialogNode {
	nodes := []inkwell.DialogNode{pad.root()}
	for _, first := range keypadPairs {
		nodes = append(nodes, pad.second(first))
	}
	return nodes
}

func (pad keypad) root() inkwell.DialogNode {
	choices := []inkwell.DialogChoice{}
	for _, first := range keypadPairs {
		choices = append(choices, inkwell.DialogChoice{
			Text:    pad.label(first),
			Actions: []inkwell.Action{inkwell.GotoNode("first_" + first)},
		})
	}
	choices = append(choices, pad.stepAway())
	return inkwell.DialogNode{
		Name: "root",
		Lines: []inkwell.DialogLine{{
			Speaker: inc.Dex,
			Text: inkwell.Text{
				{Lang: "en", Text: "Four digits. Two and seven are worn. First two?"},
				{Lang: "hu", Text: "Négy számjegy. A kettes és a hetes kopott. Az első kettő?"},
			},
		}},
		Choices: choices,
	}
}

func (pad keypad) second(first string) inkwell.DialogNode {
	choices := []inkwell.DialogChoice{}
	for _, second := range keypadPairs {
		choices = append(choices, inkwell.DialogChoice{
			Text:    pad.label(second),
			Actions: pad.attempt(first + second),
		})
	}
	choices = append(choices, pad.stepAway())
	return inkwell.DialogNode{
		Name: "first_" + first,
		Lines: []inkwell.DialogLine{{
			Speaker: inc.Dex,
			Text: inkwell.Text{
				{Lang: "en", Text: "Last two."},
				{Lang: "hu", Text: "Az utolsó kettő."},
			},
		}},
		Choices: choices,
	}
}

func (keypad) attempt(code string) []inkwell.Action {
	if code == keypadCode {
		return []inkwell.Action{
			inkwell.Say(inc.Paul, inkwell.Text{
				{Lang: "en", Text: "Click."},
				{Lang: "hu", Text: "Katt."},
			}),
			inkwell.Say(inc.Dex, inkwell.Text{
				{Lang: "en", Text: "Well. Somebody knew that."},
				{Lang: "hu", Text: "Nos. Ezt valaki tudta."},
			}),
			inkwell.SetFlag(inc.FlagSecretRoomOpen),
			inkwell.EndDialog(),
		}
	}
	return []inkwell.Action{
		inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "Nothing."},
			{Lang: "hu", Text: "Semmi."},
		}),
		inkwell.Say(inc.Dex, inkwell.Text{
			{Lang: "en", Text: "Fifteen to go. Or one, if you'd ask the man who set it."},
			{Lang: "hu", Text: "Még tizenöt. Vagy egy, ha megkérdeznéd azt, aki beállította."},
		}),
		inkwell.GotoNode("root"),
	}
}

func (keypad) label(pair string) inkwell.Text {
	spaced := string(pair[0]) + " " + string(pair[1])
	return inkwell.Text{
		{Lang: "en", Text: spaced},
		{Lang: "hu", Text: spaced},
	}
}

func (keypad) stepAway() inkwell.DialogChoice {
	return inkwell.DialogChoice{
		Text: inkwell.Text{
			{Lang: "en", Text: "Step away."},
			{Lang: "hu", Text: "Hátralépni."},
		},
		Actions: []inkwell.Action{inkwell.EndDialog()},
	}
}
