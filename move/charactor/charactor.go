package charactor

import (
	"bytes"
	"image"
	"io"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	NearThreshold = 5
)

type Charactor struct {
	PosX int
	PosY int
	Img  *ebiten.Image
	Size int
	Text string
}

func CharactorCreate(path string, x, y, size int, text string) (*Charactor, error) {
	fp, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	ary, err := io.ReadAll(fp)
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(bytes.NewReader(ary))
	if err != nil {
		return nil, err
	}
	eImg := ebiten.NewImageFromImage(img)
	ch := Charactor{x, y - size, eImg, size, text}
	return &ch, nil
}

func (c *Charactor) Draw(screen *ebiten.Image) {
	w, _ := c.Img.Bounds().Dx(), c.Img.Bounds().Dy()
	scale := float64(c.Size) / float64(w)
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(float64(c.PosX), float64(c.PosY))
	op.Filter = ebiten.FilterLinear

	screen.DrawImage(c.Img, op)
}

func (c *Charactor) IsNear(o *Charactor) bool {
	myLeft := c.PosX - NearThreshold
	myRight := c.PosX + c.Size + NearThreshold

	if myLeft < o.PosX && myRight > o.PosX {
		return true
	}
	return false
}
