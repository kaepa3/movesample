package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kaepa3/move/game"
)

func main() {
	g, err := game.NewGame()
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowSize(game.ScreenWidth, game.ScreenHeight)
	ebiten.SetWindowTitle("move")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
