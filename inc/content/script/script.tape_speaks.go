package script

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/content/character"
	"realworld/inc/content/speaker"
)

var TapeSpeaks = ScriptManager.Register(&Script{
	Name: "tape_speaks",
	Actions: inkwell.Seq(
		inkwell.Wait(0.6),
		speaker.TapeMystery.Say(Text{
			{Lang: "en", Text: "The key. This is the key!"},
			{Lang: "hu", Text: "A kulcs. Ez a kulcs!"},
		}),
		speaker.Dex.Say(Text{
			{Lang: "en", Text: "Paul. The second slot just spoke. The second slot doesn't speak."},
			{Lang: "hu", Text: "Paul. A második foglalat most megszólalt. A második foglalat nem szólal meg."},
		}),
		character.Paul.Say(Text{
			{Lang: "en", Text: "Who are you?"},
			{Lang: "hu", Text: "Ki vagy?"},
		}),
		speaker.TapeMystery.Say(Text{
			{Lang: "en", Text: "Someone who was locked until you read that name. Norman Reds. Listen to me, because nobody else will tell you this."},
			{Lang: "hu", Text: "Valaki, aki zárva volt, amíg el nem olvastad azt a nevet. Norman Reds. Hallgass meg, mert ezt senki más nem fogja elmondani."},
		}),
		speaker.TapeMystery.Say(Text{
			{Lang: "en", Text: "Norman is a good man. He found out what the simulation is really for. Not therapy. Occupation. Of minds."},
			{Lang: "hu", Text: "Norman jó ember. Rájött, mire való igazából a szimuláció. Nem terápiára. Megszállásra. Elmékére."},
		}),
		speaker.TapeMystery.Say(Text{
			{Lang: "en", Text: "So the BVK put him where he couldn't talk. He is the first test subject. He is inside it right now."},
			{Lang: "hu", Text: "Ezért a BVK oda tette, ahol nem beszélhet. Ő az első kísérleti alany. Most is odabent van."},
		}),
		speaker.TapeMystery.Say(Text{
			{Lang: "en", Text: "Find him. Free him. He can stand up in public and say what he saw, and that will be the end of them."},
			{Lang: "hu", Text: "Találd meg. Szabadítsd ki. Kiállhat a nyilvánosság elé és elmondhatja, mit látott, és az a végük lesz."},
		}),
		speaker.Dex.Say(Text{
			{Lang: "en", Text: "A locked tape that opens on one name and comes out with a mission. I've heard cleaner sales pitches from the shopkeeper."},
			{Lang: "hu", Text: "Egy zárt kazetta, ami egy névre kinyílik, és küldetéssel jön elő. A boltostól tisztább rábeszélést hallottam."},
		}),
		character.Paul.Say(Text{
			{Lang: "en", Text: "Norman never called. Now I know why. That's enough for tonight."},
			{Lang: "hu", Text: "Norman soha nem hívott. Most már tudom, miért. Ma estére ennyi elég."},
		}),
		inkwell.SetFlag(inc.FlagTapeSpeaks),
		inkwell.SetVar(inc.VarDecay, inc.DecayTapeSpeaks),
	),
})
