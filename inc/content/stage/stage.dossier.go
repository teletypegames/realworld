package stage

import (
	"sort"

	inkwell "git.teletypegames.org/engines/inkwell"
	"github.com/hajimehoshi/ebiten/v2"

	"realworld/inc/system/widget"
)

var dossierPeople = Text{
	{Lang: "en", Text: "People"},
	{Lang: "hu", Text: "Személyek"},
}

var dossierThings = Text{
	{Lang: "en", Text: "Things"},
	{Lang: "hu", Text: "Tárgyak"},
}

var Dossier = StageManager.Register(&Stage{
	Name: "dossier",
	Title: Text{
		{Lang: "en", Text: "DOSSIER"},
		{Lang: "hu", Text: "DOSSZIÉ"},
	},
})

var _ = widget.WidgetManager.Register(&inkwell.Menu{
	Name:   "dossier_menu",
	When:   inkwell.OnStage(Dossier),
	Sticky: true,
	Bounds: window(480, 300),
	Root: MenuItem{
		Label: Text{
			{Lang: "en", Text: "DOSSIER"},
			{Lang: "hu", Text: "DOSSZIÉ"},
		},
		Items: []MenuItem{
			{
				Label:  dossierPeople,
				Expand: dossierPeopleExpand,
			},
			{
				Label:  dossierThings,
				Expand: dossierThingsExpand,
			},
			{
				Label: back,
				Do:    inkwell.GoBack(),
			},
		},
	},
})

var _ = widget.WidgetManager.Register(&inkwell.KeyBinding{
	Name: "dossier_escape",
	When: inkwell.OnStage(Dossier),
	Key:  ebiten.KeyEscape,
	Do:   inkwell.GoBack(),
})

var _ = widget.WidgetManager.Register(&inkwell.KeyBinding{
	Name: "dossier_close",
	When: inkwell.OnStage(Dossier),
	Key:  ebiten.KeyTab,
	Do:   inkwell.GoBack(),
})

type dossierEntry struct {
	label Text
	lore  Text
	order int
}

func dossierRows(entries []dossierEntry) []MenuItem {
	sort.Slice(entries, func(left, right int) bool {
		return entries[left].order < entries[right].order
	})
	var rows []MenuItem
	for _, entry := range entries {
		rows = append(rows, MenuItem{
			Label: entry.label,
			Body:  entry.lore,
		})
	}
	return rows
}

func dossierPeopleExpand(context *inkwell.Ctx) []MenuItem {
	game := context.Game.(*inkwell.Game)
	state := game.State
	var entries []dossierEntry
	for _, person := range game.CharacterManager.All() {
		if state.Known(person) > 0 && len(person.Lore) > 0 {
			entries = append(entries, dossierEntry{
				label: game.SpeakerLabel(person.Name),
				lore:  person.Lore,
				order: state.Known(person),
			})
		}
	}
	for _, voice := range game.SpeakerManager.All() {
		if state.Known(voice) > 0 && len(voice.Lore) > 0 {
			entries = append(entries, dossierEntry{
				label: voice.DisplayName(),
				lore:  voice.Lore,
				order: state.Known(voice),
			})
		}
	}
	return dossierRows(entries)
}

func dossierThingsExpand(context *inkwell.Ctx) []MenuItem {
	game := context.Game.(*inkwell.Game)
	state := game.State
	var entries []dossierEntry
	for _, thing := range game.ItemManager.All() {
		if state.Known(thing) > 0 && len(thing.Lore) > 0 {
			entries = append(entries, dossierEntry{
				label: thing.Description,
				lore:  thing.Lore,
				order: state.Known(thing),
			})
		}
	}
	return dossierRows(entries)
}
