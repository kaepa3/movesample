package game

import (
	"bytes"
	_ "embed"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/kaepa3/move/game/charactor"
)

const (
	ScreenWidth  = 600
	ScreenHeight = 420
	boardSize    = 4
	fontSize     = 24
	dpi          = 72
	groundPos    = 300
	charSize     = 20
)

var (
	fontFace  *text.GoTextFace
	mainChar  *charactor.Charactor
	robotChar *charactor.Charactor
)

type GameMode int

const (
	Move GameMode = iota
	Ask
)

type Game struct {
	Mode GameMode
	Lock bool
}

func NewGame() (*Game, error) {
	g := Game{}
	return &g, nil
}

func init() {
	rand.NewSource(time.Now().UnixNano())
	//font
	src, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}
	fontFace = &text.GoTextFace{Source: src, Size: fontSize}

	mainChar, err = charactor.CharactorCreate("./game/1123.png", 0, groundPos, charSize, "")
	if err != nil {
		log.Fatal(err)
	}
	robotChar, err = charactor.CharactorCreate("./game/robot.png", 0, groundPos, charSize+20, "nemui")
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Update() error {
	switch g.Mode {
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
				g.Mode = Ask
			}
		}
	case Ask:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			if !g.Lock {
				g.Mode = Move
			}
		}
	}
	return nil
}

var once sync.Once

func (g *Game) Draw(screen *ebiten.Image) {
	mainChar.Draw(screen)
	robotChar.Draw(screen)
	switch g.Mode {
	case Ask:
		op := &text.DrawOptions{}
		op.GeoM.Translate(0, groundPos+charSize)
		text.Draw(screen, robotChar.Text, fontFace, op)
	}
	pop := &text.DrawOptions{}
	serifu := fmt.Sprintf("%d:%d %d:%d", mainChar.PosX, mainChar.PosY, robotChar.PosX, robotChar.PosY)
	text.Draw(screen, serifu, fontFace, pop)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
