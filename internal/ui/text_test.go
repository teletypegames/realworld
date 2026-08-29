package ui

import "testing"

// The engine's textWidth counts bytes, which mismeasures any non-ASCII string.
// Ours counts runes — pin that down, because wrapping depends on it.
func TestTextWidthCountsRunes(t *testing.T) {
	if got, want := textW("naïve"), 5*glyphW; got != want {
		t.Errorf("textW(naïve) = %d, want %d", got, want)
	}
	if got, want := textW("ab"), 2*glyphW; got != want {
		t.Errorf("textW(ab) = %d, want %d", got, want)
	}
}

func TestWrapRespectsWidth(t *testing.T) {
	const maxPx = 60
	lines := wrap("An Armilla without a tape is a stupid watch, remember that.", maxPx)
	if len(lines) < 2 {
		t.Fatalf("did not wrap: %v", lines)
	}
	for _, l := range lines {
		if textW(l) > maxPx && len(wrap(l, maxPx)) > 1 {
			t.Errorf("line too long (%d px): %q", textW(l), l)
		}
	}
}

func TestClipFits(t *testing.T) {
	if got := clip("black-market Armilla", 8*glyphW); textW(got) > 8*glyphW {
		t.Errorf("clip too wide: %q", got)
	}
	if got := clip("dex", 8*glyphW); got != "dex" {
		t.Errorf("clip shortened a fitting string: %q", got)
	}
}

// Regression guard for the bug that made the first build unreadable: the
// layout was derived from the wireframe deck's ratios, where text is 2.8% of
// the screen height, while the engine's debug font is a fixed 6×16 cell. At
// 320×200 that is 8% — so rows overlapped and the top bar overflowed into the
// scene. Every text row must be at least as tall as the glyph cell.
func TestLayoutFitsTheFont(t *testing.T) {
	if LineH < glyphH {
		t.Errorf("LineH (%d) is shorter than the glyph cell (%d): rows will overlap", LineH, glyphH)
	}
	if TopBarH < glyphH {
		t.Errorf("TopBarH (%d) cannot hold a %dpx glyph", TopBarH, glyphH)
	}
	// The tape channel must fit a header, a rule and at least four body rows —
	// below that the dialogue is unreadable.
	const channelH = ScreenH - HUDTop - 4
	body := (channelH - Pad*2 - LineH - 6) / LineH
	if body < 4 {
		t.Errorf("tape channel only fits %d body rows, want >= 4", body)
	}
	// The two tape slots must fit their two rows inside their border.
	if slotsH < LineH*2+Pad {
		t.Errorf("slotsH (%d) cannot hold two %dpx rows", slotsH, LineH)
	}
}

// The HUD must not reach outside the screen, and the scene must keep a usable
// aspect ratio for a point & click background.
func TestSceneProportions(t *testing.T) {
	if invY+invSlot*2+invGap > ScreenH {
		t.Errorf("inventory runs off the bottom: %d > %d", invY+invSlot*2+invGap, ScreenH)
	}
	sceneH := HUDTop - TopBarH
	if ratio := float64(ScreenW) / float64(sceneH); ratio > 3.0 {
		t.Errorf("scene is too letterboxed for a background: %.2f:1", ratio)
	}
}
