package dialog

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/content/speaker"
)

var CommentaryTape = DialogManager.Register(&Dialog{
	Name:  "commentary_tape",
	Start: "root",
	Nodes: []DialogNode{{
		Name: "root",
		Lines: []DialogLine{{
			Speaker: speaker.TapeCommentary,
			Text: Text{
				{Lang: "en", Text: "Still recording. Ask, Paul. I put the answers on here in the order I thought you'd want them."},
				{Lang: "hu", Text: "Még megy a felvétel. Kérdezz, Paul. Abban a sorrendben mondtam rá a válaszokat, ahogy gondoltam, hogy akarod."},
			},
		}},
		Choices: []DialogChoice{
			{
				Text: Text{
					{Lang: "en", Text: "Where did the other tape come from?"},
					{Lang: "hu", Text: "Honnan van a másik kazetta?"},
				},
				Actions: []Action{
					speaker.TapeCommentary.Say(Text{
						{Lang: "en", Text: "Norman. We were having beers, first time in a year. He put it on the table and said “have a look at it, you're good with these.” Like it was a broken watch."},
						{Lang: "hu", Text: "Normantól. Söröztünk, egy év után először. Letette az asztalra, és azt mondta: „nézd már meg, te értesz ezekhez.” Mintha egy rossz óra lenne."},
					}),
					speaker.TapeCommentary.Say(Text{
						{Lang: "en", Text: "It was his own Personal Tape, Paul. You don't hand that to a friend over a beer. Not unless something is already wrong."},
						{Lang: "hu", Text: "A saját Személyes Kazettája volt, Paul. Ilyet nem adsz oda egy barátodnak egy sör mellett. Csak ha már valami baj van."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: Text{
					{Lang: "en", Text: "What did you hear on it?"},
					{Lang: "hu", Text: "Mit hallottál rajta?"},
				},
				Actions: []Action{
					speaker.TapeCommentary.Say(Text{
						{Lang: "en", Text: "Not much. It's locked to the marrow. But what leaks out between the locks isn't Norman. It's too sure of itself."},
						{Lang: "hu", Text: "Nem sokat. Velejéig zárva. De ami a zárak között kiszivárog, az nem Norman. Túl biztos magában."},
					}),
					speaker.TapeCommentary.Say(Text{
						{Lang: "en", Text: "I didn't like it. I hid it behind the brick and I wrote you. If they've taken me in by the time you hear this, that's why."},
						{Lang: "hu", Text: "Nem tetszett. Eldugtam a tégla mögé, és írtam neked. Ha bevittek, mire ezt hallod, hát ezért."},
					}),
					speaker.Dex.Say(Text{
						{Lang: "en", Text: "He heard it before you did. And he still gave it to you."},
						{Lang: "hu", Text: "Ő előbb hallotta, mint te. És mégis neked adta."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: Text{
					{Lang: "en", Text: "The neural interface. Was that yours?"},
					{Lang: "hu", Text: "A neurális interfész. A tiéd volt?"},
				},
				Actions: []Action{
					speaker.TapeCommentary.Say(Text{
						{Lang: "en", Text: "The idea was mine. Run the ACP through a living brain, borrow the analogue horsepower. A bar napkin, ten years ago. Norman built it."},
						{Lang: "hu", Text: "Az ötlet az enyém volt. Futtasd az ACP-t egy élő agyon át, kölcsönözd az analóg lóerőt. Egy kocsmai szalvéta, tíz évvel ezelőtt. Norman megépítette."},
					}),
					speaker.TapeCommentary.Say(Text{
						{Lang: "en", Text: "I told him to push it further. Trance states, guided dreams. I was excited. I regret every word of it."},
						{Lang: "hu", Text: "Én mondtam neki, hogy vigye tovább. Transzállapotok, vezérelt álmok. Lelkes voltam. Minden szavát megbántam."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: Text{
					{Lang: "en", Text: "Why the worn keys?"},
					{Lang: "hu", Text: "Miért a kopott gombok?"},
				},
				Actions: []Action{
					speaker.TapeCommentary.Say(Text{
						{Lang: "en", Text: "Two and seven. I filed them a little. A code you have to guess is a code the police guess too, so I made it one you'd guess faster than they would."},
						{Lang: "hu", Text: "Kettes és hetes. Kicsit lereszeltem őket. Egy kód, amit találgatni kell, azt a rendőrség is találgatja, ezért olyat csináltam, amit te gyorsabban találsz ki, mint ők."},
					}),
					speaker.Dex.Say(Text{
						{Lang: "en", Text: "Sixteen combinations. He had a lot of faith in you."},
						{Lang: "hu", Text: "Tizenhat kombináció. Nagyon bízott benned."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: Text{
					{Lang: "en", Text: "That's enough for now."},
					{Lang: "hu", Text: "Ennyi elég most."},
				},
				Actions: []Action{
					speaker.TapeCommentary.Say(Text{
						{Lang: "en", Text: "Take care of yourself, Paul. And don't trust the other one."},
						{Lang: "hu", Text: "Vigyázz magadra, Paul. És a másikban ne bízz."},
					}),
					inkwell.EndDialog(),
				},
			},
		},
	}},
})
