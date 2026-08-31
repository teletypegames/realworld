package widget

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"realworld/inc"
	"realworld/inc/content/tape"
	"realworld/inc/system/world"
)

func init() {
	Manager.Register(&tapeSlots{
		Name:   "tapes",
		Bounds: inkwell.Rect(192, slotsY, 194, slotsH),
	})
}

type tapeSlots struct {
	Name   string
	Bounds inkwell.Rectangle
}

func (t *tapeSlots) GetName() string                { return t.Name }
func (t *tapeSlots) VisibleWhen() inkwell.Condition { return whenPlaying }
func (t *tapeSlots) Layer() inkwell.Layer           { return inkwell.LayerHUD }

func (t *tapeSlots) BlocksClickAt(p inkwell.Point) bool { return t.Bounds.Contains(p) }

func (t *tapeSlots) slotRects() (inkwell.Rectangle, inkwell.Rectangle) {
	b := t.Bounds
	w := (b.W - Pad) / 2
	return inkwell.Rect(b.X, b.Y, w, b.H), inkwell.Rect(b.X+w+Pad, b.Y, w, b.H)
}

func (t *tapeSlots) Tick(ctx *inkwell.UICtx) {
	g := ctx.Game
	mp := g.Input.Point()
	r1, r2 := t.slotRects()

	switch {
	case r1.Contains(mp):
		g.SetHoverLabel(g.CharacterLabel(inc.Dex).String())
	case r2.Contains(mp):
		if s := world.Slot2(); s != "" {
			g.SetHoverLabel(g.ItemLabel(s).String())
		} else {
			g.SetHoverLabel(slotEmptyLabel.String())
		}
	}

	if !g.Input.LeftClicked() {
		return
	}
	if !r1.Contains(mp) && !r2.Contains(mp) {
		return
	}
	g.Input.ConsumeLeft()

	slot2 := r2.Contains(mp)
	sel := g.Inventory.Selected()

	if slot2 && sel != "" {
		g.Inventory.Select("")
		if tape.IsItem(sel) {
			world.SetPending(sel)
			world.Do(inkwell.RunScript(inc.ScriptTapeInsert))
			return
		}
		world.Do(inkwell.Say(inc.Dex, inkwell.Text{
			{Lang: "en", Text: "That isn't a tape, Paul. That's an object. There is a difference."},
			{Lang: "hu", Text: "Az nem kazetta, Paul. Az egy tárgy. Van különbség."},
		}))
		return
	}
	if !slot2 && sel != "" {
		g.Inventory.Select("")
		world.Do(inkwell.Say(inc.Dex, inkwell.Text{
			{Lang: "en", Text: "The first slot is mine. I'm not moving."},
			{Lang: "hu", Text: "Az első foglalat az enyém. Nem megyek sehová."},
		}))
		return
	}

	switch g.SelectedVerb() {
	case "talk":
		if !slot2 {
			world.Do(world.Paused(inkwell.RunDialogue(inc.DlgDex)))
			return
		}
		if s := world.Slot2(); s != "" {
			world.Do(world.Paused(inkwell.RunDialogue(tape.Dialogue(s))))
			return
		}
		world.Do(inkwell.Say(inc.Dex, inkwell.Text{
			{Lang: "en", Text: "Empty. Nobody to talk to."},
			{Lang: "hu", Text: "Üres. Nincs kivel beszélni."},
		}))
	default:
		if !slot2 {
			world.Do(inkwell.Say(inc.Dex, inkwell.Text{
				{Lang: "en", Text: "Me. Four kilobytes of a dead man. Be grateful."},
				{Lang: "hu", Text: "Én. Egy halott ember négy kilobájtja. Légy hálás."},
			}))
			return
		}
		if s := world.Slot2(); s != "" {
			world.Do(inkwell.Say(inc.Dex, inkwell.Text{
				{Lang: "en", Text: "That is the second slot. Be careful what you let in."},
				{Lang: "hu", Text: "Az a második foglalat. Vigyázz, mit engedsz be."},
			}))
			return
		}
		world.Do(inkwell.Say(inc.Dex, inkwell.Text{
			{Lang: "en", Text: "The second slot is empty. That's the rarer state."},
			{Lang: "hu", Text: "A második foglalat üres. Ez a ritkább állapot."},
		}))
	}
}

func (t *tapeSlots) Draw(dst *ebiten.Image, ctx *inkwell.UICtx) {
	g := ctx.Game
	th := g.Theme()
	r1, r2 := t.slotRects()

	drawSlot := func(r inkwell.Rectangle, label, body string, filled bool) {
		bg, border, col := th.InventorySlotBG, th.ChatLogSystem, th.ChatLogSystem
		if filled {
			bg, border, col = th.PanelBG, th.CharacterPanelBorder, th.ChatLogResponse
		}
		vector.DrawFilledRect(dst, float32(r.X), float32(r.Y), float32(r.W), float32(r.H), bg, false)
		vector.StrokeRect(dst, float32(r.X), float32(r.Y), float32(r.W), float32(r.H), 1, border, false)

		x, inner := int(r.X)+Pad, int(r.W)-Pad*2
		g.DrawText(dst, label, x, int(r.Y)+Pad/2, th.ChatLogSystem)
		g.DrawText(dst, inkwell.ClipText(body, inner), x, int(r.Y)+Pad/2+LineH, col)
	}

	drawSlot(r1, slot1.String(), g.CharacterLabel(inc.Dex).String(), true)
	if s := world.Slot2(); s != "" {
		drawSlot(r2, slot2Label.String(), g.ItemLabel(s).String(), true)
	} else {
		drawSlot(r2, slot2Label.String(), slotEmpty.String(), false)
	}
}

var (
	slot1 = inkwell.Text{
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
