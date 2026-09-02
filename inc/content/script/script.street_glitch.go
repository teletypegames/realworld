package script

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

var _ = ScriptManager.Register(Script{
	Name: inc.ScriptStreetGlitch,
	Actions: inkwell.If(inkwell.VarEq(inc.VarDecay, inc.DecayNormanFound),
		inkwell.Seq(
			inkwell.Say(inc.Paul, inkwell.Text{
				{Lang: "en", Text: "Every light on the street just went out together. And came back. Together."},
				{Lang: "hu", Text: "Az utca minden lámpája egyszerre kialudt. És visszajött. Egyszerre."},
			}),
			inkwell.Say(inc.Paul, inkwell.Text{
				{Lang: "en", Text: "The ticker says: WE ARE AWAKE. Then weather."},
				{Lang: "hu", Text: "A hírszalag azt írja: ÉBREN VAGYUNK. Aztán időjárás."},
			}),
			inkwell.Say(inc.Dex, inkwell.Text{
				{Lang: "en", Text: "It's not hiding any more. It doesn't think it needs to."},
				{Lang: "hu", Text: "Már nem rejtőzik. Nem hiszi, hogy kellene."},
			}),
		),
		inkwell.Seq(
			inkwell.Say(inc.Paul, inkwell.Text{
				{Lang: "en", Text: "The streetlight stuttered. The ticker with it, on the same beat."},
				{Lang: "hu", Text: "Az utcalámpa megrándult. Vele a hírszalag is, ugyanarra az ütemre."},
			}),
			inkwell.Say(inc.Paul, inkwell.Text{
				{Lang: "en", Text: "\"MAYOR THANKS RESIDENTS FOR THEIR PATIENCE DURING.\" During what? The sentence just stops."},
				{Lang: "hu", Text: "„A POLGÁRMESTER KÖSZÖNI A LAKOSOK TÜRELMÉT A.” Mi alatt? A mondat egyszerűen elakad."},
			}),
			inkwell.Say(inc.Dex, inkwell.Text{
				{Lang: "en", Text: "Two systems on two networks, hiccuping in step. Something is in both of them."},
				{Lang: "hu", Text: "Két rendszer két hálózaton, egy ütemre csuklik. Valami mindkettőben benne van."},
			}),
		),
	),
})
