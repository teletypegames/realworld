package stage

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/system/widget"
)

var MockInventory = mockChrome(StageManager.Register(&Stage{
	Name: "mock_inventory",
	Title: Text{
		{Lang: "en", Text: "MOCK — INVENTORY"},
		{Lang: "hu", Text: "MAKETT — LELTÁR"},
	},
}))

var _ = widget.WidgetManager.Register(&widget.ArtMock{
	Name: "mock_inventory_art",
	When: inkwell.OnStage(MockInventory),
	Dim:  true,
	Caption: Text{
		{Lang: "en", Text: "ALLEY — paused"},
		{Lang: "hu", Text: "SIKÁTOR — szünet"},
	},
})

var _ = widget.WidgetManager.Register(&widget.InventoryPane{
	Name: "mock_inventory_items",
	When: inkwell.OnStage(MockInventory),
	Header: Text{
		{Lang: "en", Text: "CARRYING — 7 THINGS"},
		{Lang: "hu", Text: "NÁLAD VAN — 7 DOLOG"},
	},
	Hint: Text{
		{Lang: "en", Text: "pick one, then pick what to do with it"},
		{Lang: "hu", Text: "válassz egyet, aztán hogy mit kezdesz vele"},
	},
	Items: []widget.Choice{
		{
			Label: Text{
				{Lang: "en", Text: "unmarked tape"},
				{Lang: "hu", Text: "címkézetlen kazetta"},
			},
		},
		{
			Label: Text{
				{Lang: "en", Text: "trinket"},
				{Lang: "hu", Text: "csecsebecse"},
			},
		},
		{
			Label: Text{
				{Lang: "en", Text: "Noodle's letter"},
				{Lang: "hu", Text: "Noodle levele"},
			},
		},
		{
			Label: Text{
				{Lang: "en", Text: "black Armilla"},
				{Lang: "hu", Text: "fekete Armilla"},
			},
		},
		{
			Label: Text{
				{Lang: "en", Text: "… 3 more"},
				{Lang: "hu", Text: "… még 3"},
			},
			Dim: true,
		},
	},
	ArtLabel: Text{
		{Lang: "en", Text: "[ ART ]"},
		{Lang: "hu", Text: "[ KÉP ]"},
	},
	Description: Text{
		{Lang: "en", Text: "Password-locked. The seal has been lifted and pressed back, badly."},
		{Lang: "hu", Text: "Jelszóval zárva. A pecsétet felfeszítették és visszanyomták, ügyetlenül."},
	},
	Verbs: []widget.Choice{
		{
			Label: Text{
				{Lang: "en", Text: "Look closer"},
				{Lang: "hu", Text: "Nézd meg közelebbről"},
			},
		},
		{
			Label: Text{
				{Lang: "en", Text: "Put it away"},
				{Lang: "hu", Text: "Tedd el"},
			},
		},
		{
			Label: Text{
				{Lang: "en", Text: "Use with…"},
				{Lang: "hu", Text: "Használd ezzel…"},
			},
		},
	},
})

var _ = widget.WidgetManager.Register(&widget.Console{
	Name:   "mock_inventory_console",
	When:   inkwell.OnStage(MockInventory),
	Tokens: 3,
	Lines: []widget.ConsoleLine{
		{
			Body: Text{
				{Lang: "en", Text: "The series prefix matches Noodle's Armilla. Make of that what you like."},
				{Lang: "hu", Text: "A sorozatjel egyezik Noodle Armillájával. Kezdj vele, amit akarsz."},
			},
			Cost: Text{
				{Lang: "en", Text: "— free — you showed me"},
				{Lang: "hu", Text: "— ingyen — te mutattad meg"},
			},
		},
	},
})
