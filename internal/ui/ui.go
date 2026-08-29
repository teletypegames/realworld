package ui

// HUD layout.
//
//	     0                                     394│396                    639
//	   ┌────────────────────────────────────────────────────────────────────┐
//	 0 │ ALLEY — paused                              SRP: 1 tape            │ TopBar (20)
//	20 ├────────────────────────────────────────────────────────────────────┤
//	   │                                                                    │
//	   │            scene — hotspots, characters, SpeechBubble              │ Scene (242)
//	   │                                                                    │
//	262├─────────────────────────────────────────┬──────────────────────────┤
//	   │ Use hook on: floor grate                │ DEX                      │
//	   │ ┌───┬───┬───┬───┐   ┌───────┐┌────────┐ │ "The hook is a hand      │ HUD (118)
//	   │ ├───┼───┼───┼───┤   │ 1 DEX ││ 2  —   │ │  shorter than the gap."  │
//	379└─────────────────────────────────────────┴──────────────────────────┘
//	     InventoryBar          TapeSlots           TapeChannel
//
// On the resolution: the screen is exactly one painted background. inkwell
// stretches a scene background over the WHOLE screen (core.engine.go), not over
// the region left free by the HUD, so any screen size other than the art's own
// squashes every painting — and the deck is drawn at 640×380
// (games/realworld_screen_assets). Hence 640×380: the art lands pixel for
// pixel, and the HUD is an opaque strip laid over its foot.
//
// That still satisfies what the wiki's "320×200 VGA look" brief was after — a
// low, wide VGA frame at an exact 2× of a 320-wide canvas — and it still keeps
// the engine's fixed 6×16 debug font readable: a 16px glyph is 4% of 380, the
// proportion the design was drawn for. (At 320×190 it would be 8%, rows would
// overlap and the top bar could not hold a line.)
//
// Everything vertical is derived from LineH and from ScreenH below, never from
// the wireframe deck's ratios, so the layout cannot drift out of step with the
// font — or with the art — again.

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"git.teletypegames.org/games/realworld/internal/names"
	"git.teletypegames.org/games/realworld/internal/world"
)

// Screen and layout, in internal (unscaled) pixels.
const (
	// The screen is the size of one painted background, because the engine
	// scales a background to the full screen: at any other size the art is
	// squashed. See the note above.
	ScreenW = 640
	ScreenH = 380

	// LineH is one text row: the glyph cell plus leading. Every text-bearing
	// box is sized as a multiple of this, so rows can never overlap.
	LineH = glyphH + 2 // 18
	Pad   = 6          // inner padding for every panel

	TopBarH  = LineH + 2 // 20 — one row plus a hairline of breathing room
	DividerX = 394       // world | tape channel (61.6% of the width)

	// The HUD strip is as short as its own contents allow, and is measured
	// from the bottom of the screen — every pixel it takes is a pixel of the
	// painting it covers. It has to hold the status row and, under it, the two
	// rows of inventory slots (the tallest thing in the strip), padded above
	// and below.
	HUDH   = Pad + LineH + 4 + invSlot*2 + invGap + Pad // 118
	HUDTop = ScreenH - HUDH                             // 262 — the scene | HUD divider

	// Derived HUD rows.
	statusY = HUDTop + Pad         // 268
	slotsY  = HUDTop + Pad + LineH // 286
	slotsH  = LineH*2 + Pad*2      // 48 — two rows plus padding
	invY    = slotsY + 4           // 290
	invSlot = 40
	invGap  = 4
)

// Register assembles the whole HUD in the conventional Z-order: the debug
// overlay at the back, the cursor in front.
func Register(w *world.World) {
	g := w.G

	// The guard goes FIRST so it ticks LAST among the widgets — widgets tick
	// in reverse registration order, so every HUD widget gets the click before
	// it does, and it only ever catches clicks that land on the scene.
	g.UIManager.Register(&UseWithGuard{Name: "usewith_guard", W: w})
	g.UIManager.Register(&world.ActionPump{Name: "pump", W: w})
	g.UIManager.Register(&screenNav{Name: "screen_nav", W: w})
	g.UIManager.Register(&windowSizer{Name: "window", Scale: 2})
	g.UIManager.Register(&inkwell.HotspotDebug{Name: "hotspot_debug"})

	top := &inkwell.TopBar{
		Name:   "topbar",
		Height: TopBarH,
		// The right section prints State.Var(TimeVar). We deliberately reuse
		// it for the Armilla's single-line phosphor strip: in canon the device
		// display is one line of text, not a phone screen.
		TimeVar: names.VarArmillaStrip,
	}
	w.SetTopBar(top)
	g.UIManager.Register(gated(w, "topbar_gate", top))

	g.UIManager.Register(&Letterbox{Name: "letterbox", W: w, Bar: 36})
	g.UIManager.Register(&hudFrame{Name: "hudframe", W: w})

	g.UIManager.Register(gated(w, "status_gate", &inkwell.StatusLine{
		Name: "status", Y: statusY, Align: inkwell.AlignLeft, ScreenWidth: DividerX,
	}))
	g.UIManager.Register(gated(w, "inventory_gate", &inkwell.InventoryBar{
		Name: "inventory", Origin: inkwell.Point{X: Pad + 2, Y: invY},
		Slots: 8, Cols: 4, SlotSize: invSlot, Gap: invGap,
	}))
	g.UIManager.Register(&TapeSlots{
		Name: "tapes", W: w,
		Bounds: inkwell.Rect(192, slotsY, 194, slotsH),
	})
	g.UIManager.Register(&TapeChannel{
		Name: "channel", W: w,
		Bounds: inkwell.Rect(DividerX+2, HUDTop+2, ScreenW-DividerX-4, ScreenH-HUDTop-4),
	})

	g.UIManager.Register(&inkwell.SpeechBubble{
		Name: "speech", MaxWidth: 360, Padding: Pad, OffsetY: 8, FallbackY: TopBarH + Pad,
	})
	// The dialogue box deliberately covers only the left region, so a tape can
	// comment ALONGSIDE the conversation. It consumes every click while active,
	// so the tape channel stays visible but inert — which is what we want.
	g.UIManager.Register(&inkwell.DialogBox{
		Name: "dialog", Bounds: inkwell.Rect(0, ScreenH-LineH*9, DividerX, LineH*9),
		LineHeight: LineH, Padding: Pad,
	})
	// A verb coin, not a verb bar: a permanent VerbBar would eat the left
	// region, leaving room for neither the inventory nor the tape slots.
	g.UIManager.Register(&inkwell.RadialVerbs{
		Name: "verbs", Trigger: inkwell.MouseButtonRight, Radius: 58,
		Labels: map[string]string{
			"look": "Look", "use": "Use", "talk": "Talk", "take": "Take",
		},
	})
	g.UIManager.Register(&inkwell.EndCard{Name: "endcard"})
	g.UIManager.Register(&inkwell.Cursor{Name: "cursor"})
}

// ----- windowSizer ------------------------------------------------------

// windowSizer resizes the window once, on the first frame.
//
// inkwell.Run hardcodes a 4× window (core.dsl.go), which at 640×380 would be
// 2560×1520 — larger than most laptop screens. Run sets the size before
// entering the loop, so the only place to override it is the first tick.
type windowSizer struct {
	Name  string
	Scale int
	done  bool
}

func (s *windowSizer) GetName() string                            { return s.Name }
func (s *windowSizer) Draw(dst *ebiten.Image, ctx *inkwell.UICtx) {}

func (s *windowSizer) Tick(ctx *inkwell.UICtx) {
	if s.done {
		return
	}
	s.done = true
	scale := s.Scale
	if scale <= 0 {
		scale = 2
	}
	ebiten.SetWindowSize(ctx.Game.Width*scale, ctx.Game.Height*scale)
}

// ----- gate -------------------------------------------------------------

// gate switches built-in widgets off in cutscene and menu mode. inkwell's
// widgets have no Visible field and the Manager cannot unregister, so we wrap.
// BlocksClickAt is forwarded, otherwise the verb coin would open over the HUD.
type gate struct {
	Name  string
	W     *world.World
	Inner inkwell.Widget
}

func gated(w *world.World, name string, inner inkwell.Widget) inkwell.Widget {
	return &gate{Name: name, W: w, Inner: inner}
}

func (n *gate) GetName() string { return n.Name }

func (n *gate) Tick(ctx *inkwell.UICtx) {
	if n.W.HUDVisible() {
		n.Inner.Tick(ctx)
	}
}

func (n *gate) Draw(dst *ebiten.Image, ctx *inkwell.UICtx) {
	if n.W.HUDVisible() {
		n.Inner.Draw(dst, ctx)
	}
}

func (n *gate) BlocksClickAt(p inkwell.Point) bool {
	if !n.W.HUDVisible() {
		return false
	}
	b, ok := n.Inner.(interface{ BlocksClickAt(inkwell.Point) bool })
	return ok && b.BlocksClickAt(p)
}

// ----- hudFrame ---------------------------------------------------------

// hudFrame paints the HUD strip's backdrop and the two dividers. The rule: the
// two channels always have a visible, continuous separator between them.
type hudFrame struct {
	Name string
	W    *world.World
}

func (h *hudFrame) GetName() string         { return h.Name }
func (h *hudFrame) Tick(ctx *inkwell.UICtx) {}

func (h *hudFrame) Draw(dst *ebiten.Image, ctx *inkwell.UICtx) {
	if !h.W.HUDVisible() {
		return
	}
	th := ctx.Game.Theme()
	vector.DrawFilledRect(dst, 0, HUDTop+1, ScreenW, ScreenH-HUDTop-1, th.PanelBG, false)
	// horizontal: scene | HUD
	vector.StrokeLine(dst, 0, HUDTop, ScreenW, HUDTop, 1, th.CharacterPanelBorder, false)
	// vertical: world | tape
	vector.StrokeLine(dst, DividerX, HUDTop+1, DividerX, ScreenH, 1, th.CharacterPanelBorder, false)
}

func (h *hudFrame) BlocksClickAt(p inkwell.Point) bool {
	return h.W.HUDVisible() && p.Y >= HUDTop
}
