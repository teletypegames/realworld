package script

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/content/character"
	"realworld/inc/content/speaker"
	"realworld/inc/system"
)

var BusTutorial = ScriptManager.Register(&Script{
	Name: "bus_tutorial",
	Actions: inkwell.Seq(
		system.System.SetMode(system.ModeCutscene),
		inkwell.Wait(0.5),
		inkwell.Say(speaker.Dex, inkwell.Text{
			{Lang: "en", Text: "Forty minutes to the city. Tell me again why a junk dealer flies across the country on a postcard."},
			{Lang: "hu", Text: "Negyven perc a városig. Mondd még egyszer, miért repül át egy ószeres az országon egy képeslap miatt."},
		}),
		inkwell.Say(character.Paul, inkwell.Text{
			{Lang: "en", Text: "Scavenger. Junk dealer is what they call it on the tax form. And Noodle isn't a postcard."},
			{Lang: "hu", Text: "Csóringer. Ószeres az, ami az adóbevalláson áll. És Noodle nem képeslap."},
		}),
		inkwell.Say(speaker.Dex, inkwell.Text{
			{Lang: "en", Text: "Twenty years since university. You two shared a taste for other people's antiques and nothing else."},
			{Lang: "hu", Text: "Húsz év az egyetem óta. Ti ketten a mások régiségei iránti ízlésen osztoztatok, és semmi máson."},
		}),
		inkwell.Say(character.Paul, inkwell.Text{
			{Lang: "en", Text: "That's more than I share with most people. Present company included."},
			{Lang: "hu", Text: "Ez több, mint amin a legtöbb emberrel osztozom. A jelenlévőket is beleértve."},
		}),
		inkwell.Say(speaker.Dex, inkwell.Text{
			{Lang: "en", Text: "Present company is four kilobytes of a dead man in your first slot. I'm not company. I'm a habit."},
			{Lang: "hu", Text: "A jelenlévő egy halott ember négy kilobájtja az első foglalatodban. Nem társaság vagyok. Szokás."},
		}),
		inkwell.Say(character.Paul, inkwell.Text{
			{Lang: "en", Text: "A habit that talks."},
			{Lang: "hu", Text: "Egy szokás, ami beszél."},
		}),
		inkwell.Say(speaker.Dex, inkwell.Text{
			{Lang: "en", Text: "So listen. Right-click on a thing and pick what you mean by it: look, use, talk, take. Left-click and you walk."},
			{Lang: "hu", Text: "Akkor figyelj. Jobb gomb egy dolgon, és válaszd ki, mit akarsz vele: nézd, húzd, szólj, vedd. Bal gomb, és mész."},
		}),
		inkwell.Say(speaker.Dex, inkwell.Text{
			{Lang: "en", Text: "What you pick up sits in the bar at the bottom. Pick one up and put it on something, and see what the world says."},
			{Lang: "hu", Text: "Amit felveszel, az alsó sávban ül. Vedd kézbe, tedd rá valamire, és nézd, mit szól a világ."},
		}),
		inkwell.Say(speaker.Dex, inkwell.Text{
			{Lang: "en", Text: "And the second slot on your wrist takes a tape. Anything you push in there gets a voice. Be careful whose."},
			{Lang: "hu", Text: "És a kettes foglalat a csuklódon kazettát fogad. Amit betolsz, hangot kap. Vigyázz, kiét."},
		}),
		inkwell.Say(character.Paul, inkwell.Text{
			{Lang: "en", Text: "I've worn this thing for six years, Dex."},
			{Lang: "hu", Text: "Hat éve hordom ezt a vackot, Dex."},
		}),
		inkwell.Say(speaker.Dex, inkwell.Text{
			{Lang: "en", Text: "And in six years the second slot has been empty. That's the part I'm worried about."},
			{Lang: "hu", Text: "És hat év alatt a kettes foglalat üres volt. Ez az, ami miatt aggódom."},
		}),
		inkwell.SetFlag(inc.FlagBusTutorialDone),
		system.System.SetMode(system.ModePlay),
	),
})
