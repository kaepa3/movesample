package game

import (
	"bytes"
	_ "embed"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
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
	Title GameMode = iota
	Search
)

type Game struct {
	Mode       GameMode
	TitleGame  TitleGame
	SearchGame SearchGame
}

func NewGame() (*Game, error) {
	g := Game{Title, TitleGame{1}, SearchGame{}}
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
	case Title:
		g.Mode = g.TitleGame.Update()
	case Search:
		g.SearchGame.Update()
	}
	return nil
}

var once sync.Once

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.Mode {
	case Title:
		g.TitleGame.Draw(screen)
	case Search:
		g.SearchGame.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
