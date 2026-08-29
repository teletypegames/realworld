package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	ScreenW = 640
	ScreenH = 380

	LineH = glyphH + 2
	Pad   = 6

	TopBarH  = LineH + 2
	DividerX = 394

	HUDH   = Pad + LineH + 4 + invSlot*2 + invGap + Pad
	HUDTop = ScreenH - HUDH

	statusY = HUDTop + Pad
	slotsY  = HUDTop + Pad + LineH
	slotsH  = LineH*2 + Pad*2
	invY    = slotsY + 4
	invSlot = 40
	invGap  = 4
)

func registerUI() {
	g := World.G

	g.UIManager.Register(&useWithGuard{
		Name: "usewith_guard",
	})
	g.UIManager.Register(&actionPump{
		Name: "pump",
	})
	g.UIManager.Register(&screenNav{
		Name: "screen_nav",
	})
	g.UIManager.Register(&windowSizer{
		Name:  "window",
		Scale: 2,
	})
	g.UIManager.Register(&inkwell.HotspotDebug{
		Name: "hotspot_debug",
	})

	top := &inkwell.TopBar{
		Name:    "topbar",
		Height:  TopBarH,
		TimeVar: VarArmillaStrip,
	}
	World.SetTopBar(top)
	g.UIManager.Register(gated("topbar_gate", top))

	g.UIManager.Register(&letterbox{
		Name: "letterbox",
		Bar:  36,
	})
	g.UIManager.Register(&hudFrame{
		Name: "hudframe",
	})

	g.UIManager.Register(gated("status_gate", &inkwell.StatusLine{
		Name:        "status",
		Y:           statusY,
		Align:       inkwell.AlignLeft,
		ScreenWidth: DividerX,
	}))
	g.UIManager.Register(gated("inventory_gate", &inkwell.InventoryBar{
		Name: "inventory",
		Origin: inkwell.Point{
			X: Pad + 2,
			Y: invY,
		},
		Slots:    8,
		Cols:     4,
		SlotSize: invSlot,
		Gap:      invGap,
	}))
	g.UIManager.Register(&tapeSlots{
		Name:   "tapes",
		Bounds: inkwell.Rect(192, slotsY, 194, slotsH),
	})
	g.UIManager.Register(&tapeChannel{
		Name:   "channel",
		Bounds: inkwell.Rect(DividerX+2, HUDTop+2, ScreenW-DividerX-4, ScreenH-HUDTop-4),
	})

	g.UIManager.Register(&inkwell.SpeechBubble{
		Name:      "speech",
		MaxWidth:  360,
		Padding:   Pad,
		OffsetY:   8,
		FallbackY: TopBarH + Pad,
	})

	g.UIManager.Register(&inkwell.DialogBox{
		Name:       "dialog",
		Bounds:     inkwell.Rect(0, ScreenH-LineH*9, DividerX, LineH*9),
		LineHeight: LineH,
		Padding:    Pad,
	})

	g.UIManager.Register(&inkwell.RadialVerbs{
		Name:    "verbs",
		Trigger: inkwell.MouseButtonRight,
		Radius:  58,
		Labels: map[string]string{
			"look": "Look", "use": "Use", "talk": "Talk", "take": "Take",
		},
	})
	g.UIManager.Register(&inkwell.EndCard{
		Name: "endcard",
	})
	g.UIManager.Register(&inkwell.Cursor{
		Name: "cursor",
	})
}

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

type gate struct {
	Name  string
	Inner inkwell.Widget
}

func gated(name string, inner inkwell.Widget) inkwell.Widget {
	return &gate{
		Name:  name,
		Inner: inner,
	}
}

func (n *gate) GetName() string { return n.Name }

func (n *gate) Tick(ctx *inkwell.UICtx) {
	if World.HUDVisible() {
		n.Inner.Tick(ctx)
	}
}

func (n *gate) Draw(dst *ebiten.Image, ctx *inkwell.UICtx) {
	if World.HUDVisible() {
		n.Inner.Draw(dst, ctx)
	}
}

func (n *gate) BlocksClickAt(p inkwell.Point) bool {
	if !World.HUDVisible() {
		return false
	}
	b, ok := n.Inner.(interface{ BlocksClickAt(inkwell.Point) bool })
	return ok && b.BlocksClickAt(p)
}

type hudFrame struct {
	Name string
}

func (h *hudFrame) GetName() string         { return h.Name }
func (h *hudFrame) Tick(ctx *inkwell.UICtx) {}

func (h *hudFrame) Draw(dst *ebiten.Image, ctx *inkwell.UICtx) {
	if !World.HUDVisible() {
		return
	}
	th := ctx.Game.Theme()
	vector.DrawFilledRect(dst, 0, HUDTop+1, ScreenW, ScreenH-HUDTop-1, th.PanelBG, false)

	vector.StrokeLine(dst, 0, HUDTop, ScreenW, HUDTop, 1, th.CharacterPanelBorder, false)

	vector.StrokeLine(dst, DividerX, HUDTop+1, DividerX, ScreenH, 1, th.CharacterPanelBorder, false)
}

func (h *hudFrame) BlocksClickAt(p inkwell.Point) bool {
	return World.HUDVisible() && p.Y >= HUDTop
}
