package dialog

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

var _ = DialogManager.Register(Dialog{
	Name:  inc.DialogHotDogVendor,
	Start: "root",
	Nodes: []inkwell.DialogNode{{
		Name: "root",
		Lines: []inkwell.DialogLine{{
			Speaker: inc.HotDogVendor,
			Text: inkwell.Text{
				{Lang: "en", Text: "What'll it be? Onions are extra, opinions are free."},
				{Lang: "hu", Text: "Mi legyen? A hagyma felár, a vélemény ingyen van."},
			},
		}},
		Choices: []inkwell.DialogChoice{
			{
				Text: inkwell.Text{
					{Lang: "en", Text: "Do you know the man who owns this card?"},
					{Lang: "hu", Text: "Ismeri az embert, akié ez a kártya?"},
				},
				Show: inkwell.Given{
					Item:    inc.ItemNormanWallet,
					NotFlag: inc.FlagWorkplaceKnown,
				},
				Actions: []inkwell.Action{
					inkwell.Say(inc.HotDogVendor, inkwell.Text{
						{Lang: "en", Text: "Eight stamps? That's Norman. Every day at six, one with everything, never says a word to anybody but his wrist."},
						{Lang: "hu", Text: "Nyolc pecsét? Az Norman. Minden nap hatkor, egy mindennel, senkihez egy szót se, csak a csuklójához."},
					}),
					inkwell.Say(inc.HotDogVendor, inkwell.Text{
						{Lang: "en", Text: "Haven't seen him in weeks. He works up the hill, the residential complex with the lift that never works. Tower B."},
						{Lang: "hu", Text: "Hetek óta nem láttam. Fent dolgozik a dombon, a lakókomplexumban, amelyikben sosem működik a lift. B torony."},
					}),
					inkwell.Say(inc.Paul, inkwell.Text{
						{Lang: "en", Text: "Tower B. Thanks."},
						{Lang: "hu", Text: "B torony. Köszönöm."},
					}),
					inkwell.SetFlag(inc.FlagWorkplaceKnown),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: inkwell.Text{
					{Lang: "en", Text: "I'd like to order a delivery. For the office up the hill."},
					{Lang: "hu", Text: "Kiszállítást szeretnék rendelni. A dombon lévő irodába."},
				},
				Show: inkwell.Given{
					Flag:    inc.FlagWorkplaceKnown,
					NotItem: inc.ItemHotDogOrder,
				},
				Actions: []inkwell.Action{
					inkwell.Say(inc.HotDogVendor, inkwell.Text{
						{Lang: "en", Text: "Tower B lunch run? Sure. Got their staff cards? It's half price on the cards, and I don't carry the tray up without."},
						{Lang: "hu", Text: "B torony, ebédfutár? Mehet. Vannak a dolgozói kártyáik? Kártyával féláron van, és anélkül nem cipelem fel a tálcát."},
					}),
					inkwell.GotoNode("order"),
				},
			},
			{
				Text: inkwell.Text{
					{Lang: "en", Text: "Who's the man on the kerb?"},
					{Lang: "hu", Text: "Ki az a férfi a járdaszélen?"},
				},
				Once: true,
				Actions: []inkwell.Action{
					inkwell.Say(inc.HotDogVendor, inkwell.Text{
						{Lang: "en", Text: "Him? Used to work up the hill too. They let him go, and something let go of him. Harmless. Buys a coffee, talks to the air."},
						{Lang: "hu", Text: "Ő? Régen ő is fent dolgozott a dombon. Elengedték, és őt is elengedte valami. Ártalmatlan. Vesz egy kávét, beszél a levegőhöz."},
					}),
					inkwell.Say(inc.Dex, inkwell.Text{
						{Lang: "en", Text: "So do you, technically."},
						{Lang: "hu", Text: "Te is, technikailag."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: inkwell.Text{
					{Lang: "en", Text: "Nothing today. Thanks."},
					{Lang: "hu", Text: "Ma semmit. Köszönöm."},
				},
				Actions: []inkwell.Action{inkwell.EndDialog()},
			},
		},
	}, {
		Name: "order",
		Choices: []inkwell.DialogChoice{
			{
				Text: inkwell.Text{
					{Lang: "en", Text: "Here are their cards."},
					{Lang: "hu", Text: "Itt vannak a kártyáik."},
				},
				Show: inkwell.Given{
					Item: inc.ItemOfficeCards,
				},
				Actions: []inkwell.Action{
					inkwell.Say(inc.HotDogVendor, inkwell.Text{
						{Lang: "en", Text: "Six with everything, on the cards. Tray's yours. Tell them the mustard is the good kind."},
						{Lang: "hu", Text: "Hat mindennel, kártyára. A tálca a tiéd. Mondd meg nekik, hogy a mustár a jobbik fajta."},
					}),
					inkwell.Give(inc.ItemHotDogOrder),
					inkwell.Say(inc.Dex, inkwell.Text{
						{Lang: "en", Text: "You are now a delivery boy. This is the most honest job you've had all week."},
						{Lang: "hu", Text: "Mostantól kifutófiú vagy. Ez a legtisztességesebb munkád a héten."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: inkwell.Text{
					{Lang: "en", Text: "I don't have their cards."},
					{Lang: "hu", Text: "Nincsenek nálam a kártyáik."},
				},
				Show: inkwell.Given{
					NotItem: inc.ItemOfficeCards,
				},
				Actions: []inkwell.Action{
					inkwell.Say(inc.HotDogVendor, inkwell.Text{
						{Lang: "en", Text: "Then come back when you do. They leave them lying about on their desks, it's not hard."},
						{Lang: "hu", Text: "Akkor gyere vissza, ha lesznek. Ott hagyják az asztalukon szanaszét, nem nagy ügy."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: inkwell.Text{
					{Lang: "en", Text: "Let me think about it."},
					{Lang: "hu", Text: "Hadd gondoljam át."},
				},
				Actions: []inkwell.Action{inkwell.GotoNode("root")},
			},
		},
	}},
})
