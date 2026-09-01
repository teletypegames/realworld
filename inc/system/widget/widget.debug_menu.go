package widget

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"github.com/hajimehoshi/ebiten/v2"
)

var _ = WidgetManager.Register(&inkwell.Menu{
	Name:       "debug_menu",
	When:       whenDebug,
	Root:       inkwell.DebugMenuRoot,
	ToggleKey:  ebiten.KeyD,
	LineHeight: LineHeight,
	Padding:    Padding,
})
