package dialog

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

var _ = DialogManager.Register(Dialog{
	Name:  inc.DialogVasquez,
	Start: "root",
	Nodes: []inkwell.DialogNode{{
		Name: "root",
		Lines: []inkwell.DialogLine{{
			Speaker: inc.Vasquez,
			Text: inkwell.Text{
				{Lang: "en", Text: "Paul. Antiques. Noodle's friend from way back. And you turn up the week we go to seize a tape from his house. Sit."},
				{Lang: "hu", Text: "Paul. Régiségek. Noodle régi barátja. És pont azon a héten bukkan fel, amikor egy kazettát megyünk lefoglalni a házából. Üljön le."},
			},
		}},
		Choices: []inkwell.DialogChoice{
			{
				Text: inkwell.Text{
					{Lang: "en", Text: "I'm a collector, not a dealer. I've never carried anything dodgy in my life."},
					{Lang: "hu", Text: "Gyűjtő vagyok, nem kereskedő. Soha az életben nem volt nálam semmi kétes."},
				},
				Once: true,
				Actions: []inkwell.Action{
					inkwell.Say(inc.Vasquez, inkwell.Text{
						{Lang: "en", Text: "Everybody in this building is a collector. Your friend collects cracked Armillas."},
						{Lang: "hu", Text: "Ebben az épületben mindenki gyűjtő. A barátja feltört Armillákat gyűjt."},
					}),
					inkwell.Say(inc.Vasquez, inkwell.Text{
						{Lang: "en", Text: "Here is what matters. Noodle does not know what he found. And it is very important that it does not end up in the wrong hands."},
						{Lang: "hu", Text: "Ami számít: Noodle nem tudja, mit talált. És nagyon fontos, hogy ne kerüljön rossz kezekbe."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: inkwell.Text{
					{Lang: "en", Text: "What did he find?"},
					{Lang: "hu", Text: "Mit talált?"},
				},
				Actions: []inkwell.Action{
					inkwell.Say(inc.Vasquez, inkwell.Text{
						{Lang: "en", Text: "If I knew that, you would not be sitting here."},
						{Lang: "hu", Text: "Ha azt tudnám, most nem itt ülne."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: inkwell.Text{
					{Lang: "en", Text: "So what do you want from me?"},
					{Lang: "hu", Text: "Akkor mit akar tőlem?"},
				},
				Once: true,
				Actions: []inkwell.Action{
					inkwell.Say(inc.Vasquez, inkwell.Text{
						{Lang: "en", Text: "Nothing. I don't care about your petty little deals, either of you."},
						{Lang: "hu", Text: "Semmit. Nem érdekelnek a pitiáner üzleteik, egyiküké sem."},
					}),
					inkwell.Say(inc.Vasquez, inkwell.Text{
						{Lang: "en", Text: "But if you two find something strange. Something dodgy. You tell me. You may be walking into something you very much do not want to be in."},
						{Lang: "hu", Text: "De ha maguk ketten találnak valami furcsát. Valami kéteset. Jelzik nekem. Belekeveredhetnek valamibe, amiben nagyon nem akarnak benne lenni."},
					}),
					inkwell.Say(inc.Vasquez, inkwell.Text{
						{Lang: "en", Text: "Or I find you anyway, sooner or later. Either way. Go on, talk to him."},
						{Lang: "hu", Text: "Vagy előbb-utóbb úgyis megtalálom magukat. Így is, úgy is. Menjen, beszéljen vele."},
					}),
					inkwell.SetFlag(inc.FlagVasquezDeal),
					inkwell.Say(inc.Dex, inkwell.Text{
						{Lang: "en", Text: "That wasn't permission, Paul. That was a trade. Remember which one you agreed to."},
						{Lang: "hu", Text: "Ez nem engedély volt, Paul. Ez egy csere. Jegyezd meg, melyikbe egyeztél bele."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: inkwell.Text{
					{Lang: "en", Text: "Why is a detective in a blazer handling a fencing case?"},
					{Lang: "hu", Text: "Miért egy blézeres nyomozó visz egy orgazdaügyet?"},
				},
				Once: true,
				Actions: []inkwell.Action{
					inkwell.Say(inc.Vasquez, inkwell.Text{
						{Lang: "en", Text: "Because it isn't one. Good question, though. Ask fewer of them."},
						{Lang: "hu", Text: "Mert nem az. Jó kérdés, viszont. Kevesebbet tegyen fel belőlük."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: inkwell.Text{
					{Lang: "en", Text: "No, I haven't been to the alley."},
					{Lang: "hu", Text: "Nem, nem voltam a sikátorban."},
				},
				Show: inkwell.Given{
					Flag: inc.FlagHasTape,
				},
				Once: true,
				Actions: []inkwell.Action{
					inkwell.Say(inc.Vasquez, inkwell.Text{
						{Lang: "en", Text: "I didn't ask about an alley."},
						{Lang: "hu", Text: "Én nem kérdeztem semmilyen sikátorról."},
					}),
					inkwell.Say(inc.Paul, inkwell.Text{
						{Lang: "en", Text: "…No. You didn't."},
						{Lang: "hu", Text: "…Nem. Nem kérdezett."},
					}),
					inkwell.Say(inc.Vasquez, inkwell.Text{
						{Lang: "en", Text: "Remember our arrangement, Paul."},
						{Lang: "hu", Text: "Ne feledje a megállapodásunkat, Paul."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: inkwell.Text{
					{Lang: "en", Text: "It's a weather readout."},
					{Lang: "hu", Text: "Az egy időjárás-kijelző."},
				},
				Show: inkwell.Given{
					Flag: inc.FlagTapeSpeaks,
				},
				Once: true,
				Actions: []inkwell.Action{
					inkwell.Say(inc.Vasquez, inkwell.Text{
						{Lang: "en", Text: "Your Armilla. The strip. It's been scrolling since you sat down."},
						{Lang: "hu", Text: "Az Armillája. A csík. Azóta fut, hogy leült."},
					}),
					inkwell.Say(inc.Paul, inkwell.Text{
						{Lang: "en", Text: "Fog later. Always fog later."},
						{Lang: "hu", Text: "Később köd. Mindig később köd."},
					}),
					inkwell.Say(inc.Vasquez, inkwell.Text{
						{Lang: "en", Text: "Two slots, Paul. Licensed ones have one."},
						{Lang: "hu", Text: "Két foglalat, Paul. A licenszelteknek egy van."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: inkwell.Text{
					{Lang: "en", Text: "I'll go talk to him."},
					{Lang: "hu", Text: "Megyek, beszélek vele."},
				},
				Actions: []inkwell.Action{inkwell.EndDialog()},
			},
		},
	}},
})
