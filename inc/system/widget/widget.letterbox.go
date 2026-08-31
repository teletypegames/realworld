package widget

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"realworld/inc"
	"realworld/inc/system/armilla"
	"realworld/inc/util"
)

func init() {
	WidgetManager.Register(&letterbox{
		Name: "letterbox",
		Bar:  36,
	})
}

type letterbox struct {
	Name string
	Bar  float64
}

var tapeWaitingHint = inkwell.Text{
	{Lang: "en", Text: "SPACE — %s has something to say"},
	{Lang: "hu", Text: "SZÓKÖZ — %s mondani akar valamit"},
}

func (box *letterbox) GetName() string                { return box.Name }
func (box *letterbox) VisibleWhen() inkwell.Condition { return whenCutscene }
func (box *letterbox) Layer() inkwell.Layer           { return inkwell.LayerScene }

func (box *letterbox) Tick(context *inkwell.UICtx) {
	if !armilla.Waiting() {
		return
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		armilla.TakeOffer()
	}
}

func (box *letterbox) Draw(destination *ebiten.Image, context *inkwell.UICtx) {
	game := context.Game
	theme := game.Theme()
	bar := box.Bar
	if bar <= 0 {
		bar = 18
	}
	width, height := float32(game.Width), float32(game.Height)
	black := util.RGB(0x000000)
	vector.DrawFilledRect(destination, 0, 0, width, float32(bar), black, false)
	vector.DrawFilledRect(destination, 0, height-float32(bar), width, float32(bar), black, false)

	if armilla.Waiting() {
		hint := inkwell.Format(tapeWaitingHint, game.CharacterLabel(inc.Dex)).String()
		game.DrawText(destination, hint, game.Width-inkwell.TextWidth(hint)-4, game.Height-int(bar)+4, theme.ChatLogResponse)
	}
}
