package dialog

import (
	"strings"
	"unicode"
	"unicode/utf8"

	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/content/speaker"
	"realworld/inc/system"
)

var dexTalkHintTarget = Text{
	{Lang: "en", Text: "Look at it from where I'm sitting: %s, %s. That's where this goes next. I didn't say it was a good idea."},
	{Lang: "hu", Text: "Innen nézve: %s, %s. Ott megy tovább a dolog. Azt nem mondtam, hogy jó ötlet."},
}

var dexTalkHintItem = Text{
	{Lang: "en", Text: "You're not carrying %s for the weight. Try it on this: %s, %s."},
	{Lang: "hu", Text: "Nem a súlya miatt hordod magadnál: %s. Próbáld ki ezen: %s, %s."},
}

var dexTalkHintNothing = Text{
	{Lang: "en", Text: "Nothing, Paul. For once in my life I've got nothing."},
	{Lang: "hu", Text: "Semmi, Paul. Életemben először nincs semmim."},
}

func dexTalkHint(*inkwell.State) Text {
	game := system.System.Game()
	situation := system.System.Situation()
	var target inkwell.Move
	for _, move := range situation.Next {
		if move.Dialog == nil {
			target = move
		}
	}
	if !situation.Found || target.Scene == "" {
		return dexTalkHintNothing
	}
	place := dexTalkPlace(game, target.Scene)
	label := target.Label
	if destination, isExit := strings.CutPrefix(target.Hotspot, "exit:"); isExit {
		label = dexTalkPlace(game, destination)
	}
	if target.With != nil {
		return inkwell.Format(dexTalkHintItem, game.ItemLabel(target.Item.Name), game.ItemLabel(target.With.Name), place)
	}
	if target.Item != nil {
		return inkwell.Format(dexTalkHintItem, game.ItemLabel(target.Item.Name), label, place)
	}
	return inkwell.Format(dexTalkHintTarget, label, place)
}

func dexTalkPlace(game *inkwell.Game, name string) Text {
	scene, found := game.SceneManager.Get(name)
	if !found || scene.Title.Empty() {
		return Text{
			{Lang: "en", Text: name},
		}
	}
	return scene.Title.Map(func(title string) string {
		place, _, _ := strings.Cut(title, " — ")
		words := strings.Fields(strings.ToLower(place))
		for index, word := range words {
			first, size := utf8.DecodeRuneInString(word)
			words[index] = string(unicode.ToUpper(first)) + word[size:]
		}
		return strings.Join(words, " ")
	})
}

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
					{Lang: "en", Text: "Where do I even go from here?"},
					{Lang: "hu", Text: "Innen merre tovább?"},
				},
				Actions: []Action{
					speaker.Dex.Reply(dexTalkHint),
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
