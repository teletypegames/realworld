package ui

import (
	"testing"

	inkwell "git.teletypegames.org/engines/inkwell"

	"git.teletypegames.org/games/realworld/internal/world"
)

// The arrow keys walk the deck in registration order and wrap at both ends —
// until the beats grow doors, that walk is the only way through most screens,
// so a dead end at either end would strand the player on a greybox screen.
func TestScreenNavWrapsBothWays(t *testing.T) {
	g := inkwell.NewGame("t", ScreenW, ScreenH)
	for _, n := range []string{"a", "b", "c"} {
		g.SceneManager.Register(inkwell.Scene{Name: n, Background: "bg"})
	}
	w := world.New(g)
	nav := &screenNav{Name: "screen_nav", W: w}
	at := func(name string) { world.EnterScene(w, name).Start().Tick(&inkwell.Ctx{Game: g}) }

	for _, tc := range []struct {
		from string
		step int
		want string
	}{
		{"a", 1, "b"}, {"b", 1, "c"}, {"c", 1, "a"},
		{"c", -1, "b"}, {"b", -1, "a"}, {"a", -1, "c"},
	} {
		at(tc.from)
		if got := nav.neighbour(g, tc.step); got != tc.want {
			t.Errorf("from %q step %+d = %q, want %q", tc.from, tc.step, got, tc.want)
		}
	}
}
