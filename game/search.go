package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type GameSubMode int

const (
	Move GameSubMode = iota
	Ask
)

type SearchGame struct {
	SubMode GameSubMode
	Lock    bool
	Tick    int
}

func (s *SearchGame) Update() {
	switch s.SubMode {
	case Move:
		if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
			if mainChar.PosX+1 <= ScreenWidth {
				mainChar.PosX += 1
			}
		} else if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
			if mainChar.PosX-1 > 0 {
				mainChar.PosX -= 1
			}
		}
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			if robotChar.IsNear(mainChar) {
				s.SubMode = Ask
			}
		}
	case Ask:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			if !s.Lock {
				s.SubMode = Move
			}
		}
	}
}

func (s *SearchGame) Draw(screen *ebiten.Image) {
	mainChar.Draw(screen)
	robotChar.Draw(screen)
	switch s.SubMode {
	case Ask:
		op := &text.DrawOptions{}
		op.GeoM.Translate(0, groundPos+charSize)
		text.Draw(screen, robotChar.Text, fontFace, op)
	}
	pop := &text.DrawOptions{}
	serifu := fmt.Sprintf("%d:%d %d:%d", mainChar.PosX, mainChar.PosY, robotChar.PosX, robotChar.PosY)
	text.Draw(screen, serifu, fontFace, pop)
}
