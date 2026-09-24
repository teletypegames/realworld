package widget

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"realworld/inc/util"
)

var artHatch = util.RGBA(0xFFFFFF, 0x0A)

var artDim = util.RGBA(0x000000, 0x8C)

type ArtBox struct {
	Area  inkwell.Rectangle
	Label Text
	Solid bool
}

type ArtMock struct {
	Name     string
	When     inkwell.Condition
	Caption  Text
	Dim      bool
	Boxes    []ArtBox
	Bubble   Text
	BubbleAt inkwell.Rectangle
	Tag      Text
	TagAt    inkwell.Point
	Subtitle Text
}

func (art *ArtMock) GetName() string                { return art.Name }
func (art *ArtMock) VisibleWhen() inkwell.Condition { return art.When }
func (art *ArtMock) Layer() inkwell.Layer           { return inkwell.LayerScene }
func (art *ArtMock) Tick(context *inkwell.UICtx)    {}

func (art *ArtMock) Draw(destination *ebiten.Image, context *inkwell.UICtx) {
	game := context.Game.(*inkwell.Game)
	theme := game.Theme()

	top := float32(uiTop(game))
	for x := -top; x < ScreenWidth; x += 20 {
		vector.StrokeLine(destination, x, top, x+top, 0, 8, artHatch, false)
	}
	if art.Dim {
		fill(destination, inkwell.Rect(0, 0, ScreenWidth, float64(top)), artDim)
	}

	for _, box := range art.Boxes {
		if box.Solid {
			fill(destination, box.Area, theme.VerbButtonBG)
			stroke(destination, box.Area, theme.StatusText)
		} else {
			strokeDashed(destination, box.Area, theme.TopBarText)
		}
		if len(box.Label) > 0 {
			centered(destination, game, box.Label.String(), box.Area,
				int(box.Area.Y+box.Area.H/2)-inkwell.GlyphH/2, theme.TopBarText)
		}
	}

	if len(art.Bubble) > 0 {
		lines := inkwell.WrapText(art.Bubble.String(), int(art.BubbleAt.W)-Padding*2)
		frame := inkwell.Rect(art.BubbleAt.X, art.BubbleAt.Y, art.BubbleAt.W,
			float64(len(lines)*LineHeight+Padding*2))
		fill(destination, frame, theme.SpeechBubbleBG)
		stroke(destination, frame, theme.CharacterPanelBorder)
		for index, line := range lines {
			game.DrawText(destination, line, int(frame.X)+Padding,
				int(frame.Y)+Padding+index*LineHeight, theme.SpeechDefaultText)
		}
	}

	if len(art.Caption) > 0 {
		plate(destination, game, art.Caption.String(), 8, 8, theme.TopBarText)
	}

	if len(art.Tag) > 0 {
		game.DrawText(destination, "→", int(art.TagAt.X), int(art.TagAt.Y), theme.StatusText)
		plate(destination, game, art.Tag.String(), int(art.TagAt.X)+10, int(art.TagAt.Y), theme.StatusText)
	}

	if len(art.Subtitle) > 0 {
		platedCenter(destination, game, art.Subtitle.String(),
			int(top)-inkwell.GlyphH-14, theme.StatusText)
	}
}
