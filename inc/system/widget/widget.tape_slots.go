package widget

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"realworld/inc"
	"realworld/inc/content/tape"
	"realworld/inc/system"
	"realworld/inc/system/armilla"
)

func init() {
	WidgetManager.Register(&tapeSlots{
		Name:   "tapes",
		Bounds: inkwell.Rect(192, slotsY, 194, slotsHeight),
	})
}

type tapeSlots struct {
	Name   string
	Bounds inkwell.Rectangle
}

func (slots *tapeSlots) GetName() string                { return slots.Name }
func (slots *tapeSlots) VisibleWhen() inkwell.Condition { return whenPlaying }
func (slots *tapeSlots) Layer() inkwell.Layer           { return inkwell.LayerHUD }

func (slots *tapeSlots) BlocksClickAt(point inkwell.Point) bool {
	return slots.Bounds.Contains(point)
}

func (slots *tapeSlots) slotRectangles() (inkwell.Rectangle, inkwell.Rectangle) {
	bounds := slots.Bounds
	width := (bounds.W - Padding) / 2
	return inkwell.Rect(bounds.X, bounds.Y, width, bounds.H),
		inkwell.Rect(bounds.X+width+Padding, bounds.Y, width, bounds.H)
}

func (slots *tapeSlots) Tick(context *inkwell.UICtx) {
	game := context.Game
	mousePoint := game.Input.Point()
	firstSlot, secondSlot := slots.slotRectangles()

	switch {
	case firstSlot.Contains(mousePoint):
		game.SetHoverLabel(game.CharacterLabel(inc.Dex).String())
	case secondSlot.Contains(mousePoint):
		if loaded := armilla.Slot2(); loaded != "" {
			game.SetHoverLabel(game.ItemLabel(loaded).String())
		} else {
			game.SetHoverLabel(slotEmptyLabel.String())
		}
	}

	if !game.Input.LeftClicked() {
		return
	}
	if !firstSlot.Contains(mousePoint) && !secondSlot.Contains(mousePoint) {
		return
	}
	game.Input.ConsumeLeft()

	clickedSecond := secondSlot.Contains(mousePoint)
	selected := game.Inventory.Selected()

	if clickedSecond && selected != "" {
		game.Inventory.Select("")
		if tape.IsItem(selected) {
			armilla.SetPending(selected)
			system.System.Do(inkwell.RunScript(inc.ScriptTapeInsert))
			return
		}
		system.System.Do(inkwell.Say(inc.Dex, inkwell.Text{
			{Lang: "en", Text: "That isn't a tape, Paul. That's an object. There is a difference."},
			{Lang: "hu", Text: "Az nem kazetta, Paul. Az egy tárgy. Van különbség."},
		}))
		return
	}
	if !clickedSecond && selected != "" {
		game.Inventory.Select("")
		system.System.Do(inkwell.Say(inc.Dex, inkwell.Text{
			{Lang: "en", Text: "The first slot is mine. I'm not moving."},
			{Lang: "hu", Text: "Az első foglalat az enyém. Nem megyek sehová."},
		}))
		return
	}

	switch game.SelectedVerb() {
	case "talk":
		if !clickedSecond {
			system.System.Do(system.System.Paused(inkwell.RunDialog(inc.DialogDex)))
			return
		}
		if loaded := armilla.Slot2(); loaded != "" {
			system.System.Do(system.System.Paused(inkwell.RunDialog(tape.Dialog(loaded))))
			return
		}
		system.System.Do(inkwell.Say(inc.Dex, inkwell.Text{
			{Lang: "en", Text: "Empty. Nobody to talk to."},
			{Lang: "hu", Text: "Üres. Nincs kivel beszélni."},
		}))
	default:
		if !clickedSecond {
			system.System.Do(inkwell.Say(inc.Dex, inkwell.Text{
				{Lang: "en", Text: "Me. Four kilobytes of a dead man. Be grateful."},
				{Lang: "hu", Text: "Én. Egy halott ember négy kilobájtja. Légy hálás."},
			}))
			return
		}
		if loaded := armilla.Slot2(); loaded != "" {
			system.System.Do(inkwell.Say(inc.Dex, inkwell.Text{
				{Lang: "en", Text: "That is the second slot. Be careful what you let in."},
				{Lang: "hu", Text: "Az a második foglalat. Vigyázz, mit engedsz be."},
			}))
			return
		}
		system.System.Do(inkwell.Say(inc.Dex, inkwell.Text{
			{Lang: "en", Text: "The second slot is empty. That's the rarer state."},
			{Lang: "hu", Text: "A második foglalat üres. Ez a ritkább állapot."},
		}))
	}
}

func (slots *tapeSlots) Draw(destination *ebiten.Image, context *inkwell.UICtx) {
	game := context.Game
	theme := game.Theme()
	firstSlot, secondSlot := slots.slotRectangles()

	drawSlot := func(rectangle inkwell.Rectangle, label, body string, filled bool) {
		background, border, textColor := theme.InventorySlotBG, theme.ChatLogSystem, theme.ChatLogSystem
		if filled {
			background, border, textColor = theme.PanelBG, theme.CharacterPanelBorder, theme.ChatLogResponse
		}
		vector.DrawFilledRect(destination, float32(rectangle.X), float32(rectangle.Y),
			float32(rectangle.W), float32(rectangle.H), background, false)
		vector.StrokeRect(destination, float32(rectangle.X), float32(rectangle.Y),
			float32(rectangle.W), float32(rectangle.H), 1, border, false)

		textX, innerWidth := int(rectangle.X)+Padding, int(rectangle.W)-Padding*2
		game.DrawText(destination, label, textX, int(rectangle.Y)+Padding/2, theme.ChatLogSystem)
		game.DrawText(destination, inkwell.ClipText(body, innerWidth), textX,
			int(rectangle.Y)+Padding/2+LineHeight, textColor)
	}

	drawSlot(firstSlot, slot1Label.String(), game.CharacterLabel(inc.Dex).String(), true)
	if loaded := armilla.Slot2(); loaded != "" {
		drawSlot(secondSlot, slot2Label.String(), game.ItemLabel(loaded).String(), true)
	} else {
		drawSlot(secondSlot, slot2Label.String(), slotEmpty.String(), false)
	}
}

var (
	slot1Label = inkwell.Text{
		{Lang: "en", Text: "SLOT 1"},
		{Lang: "hu", Text: "1. FOGLALAT"},
	}
	slot2Label = inkwell.Text{
		{Lang: "en", Text: "SLOT 2"},
		{Lang: "hu", Text: "2. FOGLALAT"},
	}
	slotEmpty = inkwell.Text{
		{Lang: "en", Text: "empty"},
		{Lang: "hu", Text: "üres"},
	}
	slotEmptyLabel = inkwell.Text{
		{Lang: "en", Text: "empty tape slot"},
		{Lang: "hu", Text: "üres kazettafoglalat"},
	}
)
