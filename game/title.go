package game

import (
	"bytes"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var (
	titleFace *text.GoTextFace
	basicFace *text.GoTextFace
)

const (
	titleFontSize = 48
	interval      = 100
)

var (
	counter int
)

func init() {
	//font
	src, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}
	titleFace = &text.GoTextFace{Source: src, Size: titleFontSize}
}

type TitleGame struct {
	Tick int
}

func (t *TitleGame) Update() GameMode {
	ret := Title
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		ret = Search
	}
	t.Tick++
	return ret
}

func (t *TitleGame) Draw(screen *ebiten.Image) {
	drawTitle(screen, "Kanetori")
	drawCommand(screen, float64(t.Tick), "-press space key-")
}

func drawTitle(s *ebiten.Image, title string) {
	w, h := text.Measure(title, titleFace, 0)
	posx := (ScreenWidth - w) / 2
	posy := ScreenHeight/2 - (h * 2)
	op := &text.DrawOptions{}
	op.GeoM.Translate(float64(posx), float64(posy))
	text.Draw(s, title, titleFace, op)
}
func drawCommand(s *ebiten.Image, tick float64, cmd string) {
	w, h := text.Measure(cmd, fontFace, 0)
	posx := (ScreenWidth - w) / 2
	posy := ScreenHeight/2 + (h * 2)
	op := &text.DrawOptions{}
	op.GeoM.Translate(float64(posx), float64(posy))
	alpha := math.Sin(float64(tick*2*math.Pi/180)) + 1
	log.Println(alpha)
	op.ColorScale.ScaleAlpha(float32(alpha))
	text.Draw(s, cmd, fontFace, op)
}
