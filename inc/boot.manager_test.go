package inc

import (
	"image/png"
	"os"
	"path/filepath"
	"strings"
	"testing"

	inkwell "git.teletypegames.org/engines/inkwell"
)

func newTestGame(t *testing.T) *inkwell.Game {
	t.Helper()
	g := New(Opts{})
	if err := g.Validate(); err != nil {
		t.Fatalf("Validate: %v", err)
	}
	return g
}

// The engine's Validate only checks scene→asset/character references. A typo in
// a dialogue, script or item name would surface at runtime as a silent no-op
// (RunDialogue and RunScript fail quietly), so check them here.
func TestReferencesResolve(t *testing.T) {
	g := newTestGame(t)

	for _, n := range []string{DlgDex, DlgMysteryTape} {
		if !g.DialogueManager.Has(n) {
			t.Errorf("missing dialogue: %q", n)
		}
	}
	for _, n := range []string{ScriptTapeInsert, ScriptFinale} {
		if !g.ScriptManager.Has(n) {
			t.Errorf("missing script: %q", n)
		}
	}
	for _, n := range []string{ItemNoodleLetter, ItemMysteryTape, ItemBlackArmilla} {
		if !g.ItemManager.Has(n) {
			t.Errorf("missing item: %q", n)
		}
	}
	for n := range Tapes {
		if !g.CharacterManager.Has(n) {
			t.Errorf("tape character not registered: %q", n)
		}
	}
	// Every loadable tape needs an item, a character and a dialogue.
	for item, char := range TapeItems {
		if !g.ItemManager.Has(item) {
			t.Errorf("tape item not registered: %q", item)
		}
		if !g.CharacterManager.Has(char) {
			t.Errorf("tape character not registered: %q", char)
		}
		if !g.DialogueManager.Has(TapeDialogue(item)) {
			t.Errorf("tape %q has no dialogue: %q", item, TapeDialogue(item))
		}
	}
}

// The two-channel rule: every tape needs its own voice colour, distinct from
// the world's. Without that, "if something is amber, a tape said it" is not
// readable.
func TestTapeVoicesAreDistinct(t *testing.T) {
	g := newTestGame(t)
	paul, _ := g.CharacterManager.Get(Paul)
	seen := map[[4]uint32]string{}
	for n := range Tapes {
		c, _ := g.CharacterManager.Get(n)
		if c.SpeechColor == nil {
			t.Errorf("%s: no SpeechColor", n)
			continue
		}
		if c.SpeechColor == paul.SpeechColor {
			t.Errorf("%s speaks in the same colour as Paul", n)
		}
		r, gr, b, a := c.SpeechColor.RGBA()
		key := [4]uint32{r, gr, b, a}
		if other, dup := seen[key]; dup {
			t.Errorf("%s and %s share a voice colour", n, other)
		}
		seen[key] = n
	}
}

// HUD geometry: the two channels must not overlap, and both must stay inside
// the HUD strip.
func TestHUDGeometry(t *testing.T) {
	g := newTestGame(t)

	ch, ok := g.UIManager.Get("channel")
	if !ok {
		t.Fatal("no 'channel' widget")
	}
	tc := ch.(*TapeChannel)
	if tc.Bounds.X < DividerX {
		t.Errorf("tape channel crosses into the world side: X=%v < %v", tc.Bounds.X, DividerX)
	}
	if tc.Bounds.X+tc.Bounds.W > ScreenW {
		t.Errorf("tape channel runs off screen: %v", tc.Bounds.X+tc.Bounds.W)
	}
	if tc.Bounds.Y < HUDTop {
		t.Errorf("tape channel reaches into the scene: Y=%v < %v", tc.Bounds.Y, HUDTop)
	}

	sl, ok := g.UIManager.Get("tapes")
	if !ok {
		t.Fatal("no 'tapes' widget")
	}
	if ts := sl.(*TapeSlots); ts.Bounds.X+ts.Bounds.W > DividerX {
		t.Errorf("tape slots cross the divider: %v > %v", ts.Bounds.X+ts.Bounds.W, DividerX)
	}

	// The dialogue box deliberately covers only the left region, so a tape can
	// comment alongside the conversation.
	db, ok := g.UIManager.Get("dialog")
	if !ok {
		t.Fatal("no 'dialog' widget")
	}
	if box := db.(*inkwell.DialogBox); box.Bounds.X+box.Bounds.W > DividerX {
		t.Errorf("dialogue box would cover the tape channel: %v > %v",
			box.Bounds.X+box.Bounds.W, DividerX)
	}
}

// The guard must be registered FIRST so it ticks LAST among the widgets —
// otherwise it would swallow HUD clicks.
func TestWidgetOrder(t *testing.T) {
	g := newTestGame(t)
	n := g.UIManager.Names()
	if len(n) == 0 || n[0] != "usewith_guard" {
		t.Errorf("usewith_guard is not registered first: %v", n)
	}
	if n[len(n)-1] != "cursor" {
		t.Errorf("cursor is not registered last: %v", n)
	}
}

// Every screen needs a background, the drawn ones need the file to be where the
// asset says it is, and the file has to be exactly screen-sized.
//
// The path check: a renamed or mistyped path shows up in the running game only
// as an inkwell placeholder, which is exactly what a still-greyboxed screen
// looks like — invisible as a bug.
//
// The size check: the engine scales a background over the WHOLE screen, so a
// painting that is not ScreenW×ScreenH is silently squashed to fit. That is why
// the screen is 640×380 in the first place; this keeps the two in step.
func TestSceneBackgroundsResolve(t *testing.T) {
	g := newTestGame(t)

	// Not drawn yet: these render as placeholders on purpose. The alley is
	// missing from the concept-art deck; the selector is the wiki's menu
	// screen, which the deck does not cover at all.
	undrawn := map[string]bool{BgAlley: true, BgSelector: true}

	for _, n := range g.SceneManager.Names() {
		s := g.SceneManager.MustGet(n)
		a, ok := g.AssetManager.Get(s.Background)
		if !ok {
			t.Errorf("scene %q: unknown background %q", n, s.Background)
			continue
		}
		if undrawn[a.Name] {
			continue
		}
		f, err := os.Open(filepath.Join("..", a.Path))
		if err != nil {
			t.Errorf("scene %q: background file missing: %v", n, err)
			continue
		}
		cfg, err := png.DecodeConfig(f)
		f.Close()
		if err != nil {
			t.Errorf("scene %q: unreadable background: %v", n, err)
			continue
		}
		if cfg.Width != ScreenW || cfg.Height != ScreenH {
			t.Errorf("scene %q: background is %d×%d, want %d×%d — it will be stretched",
				n, cfg.Width, cfg.Height, ScreenW, ScreenH)
		}
	}
}

// exits reads the connection graph back out of the registered screens. The
// exit hotspots are named "exit:<target>", which is what keeps the graph
// machine-readable — see screen.exit.go.
func exits(g *inkwell.Game, screen string) []string {
	var out []string
	for _, h := range g.SceneManager.MustGet(screen).Hotspots {
		if to, ok := strings.CutPrefix(h.Name, "exit:"); ok && to != "back" {
			out = append(out, to)
		}
	}
	return out
}

// An exit that names a screen nobody registered is a dead end the player walks
// into: the engine logs "changeScene: unknown" and nothing happens on screen.
func TestExitsResolve(t *testing.T) {
	g := newTestGame(t)
	for _, n := range g.SceneManager.Names() {
		for _, to := range exits(g, n) {
			if !g.SceneManager.Has(to) {
				t.Errorf("screen %q: exit to unknown screen %q", n, to)
			}
		}
	}
}

// Every screen has to be walkable to from the start of the game, through the
// exits alone — the arrow keys are a browsing tool, not the map. A screen that
// only the arrow keys can reach is content no player would ever see.
//
// The walk starts at Paul's shop, where the letter arrives, and mostly runs
// through the selector: that is what the wiki's location graph is centred on.
func TestEveryScreenIsReachable(t *testing.T) {
	g := newTestGame(t)

	seen := map[string]bool{ScreenPaulShop: true}
	queue := []string{ScreenPaulShop}
	for len(queue) > 0 {
		at := queue[0]
		queue = queue[1:]
		for _, to := range exits(g, at) {
			if !seen[to] && g.SceneManager.Has(to) {
				seen[to] = true
				queue = append(queue, to)
			}
		}
	}

	for _, n := range g.SceneManager.Names() {
		if !seen[n] {
			t.Errorf("screen %q cannot be reached from %q by its exits", n, ScreenPaulShop)
		}
	}
}

// A screen with no way out at all is a trap: the player walks in and the only
// way on is the debug walk.
func TestEveryScreenHasAWayOut(t *testing.T) {
	g := newTestGame(t)
	for _, n := range g.SceneManager.Names() {
		var ways int
		for _, h := range g.SceneManager.MustGet(n).Hotspots {
			if strings.HasPrefix(h.Name, "exit:") {
				ways++
			}
		}
		if ways == 0 {
			t.Errorf("screen %q has no exit", n)
		}
	}
}
