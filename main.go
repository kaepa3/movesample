package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kaepa3/move/move"
)

func main() {
	game, err := move.NewGame()
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowSize(move.ScreenWidth, move.ScreenHeight)
	ebiten.SetWindowTitle("move")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
