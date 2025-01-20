package entity

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
)
type Entity struct {
	Name string
	StartPosX int
	StartPosY int
	TextureId int
	Model     *resolv.Object
}

func (e *Entity) Draw(sprite, screen *ebiten.Image, frameHeight, frameWidth int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(e.Model.Position.X), e.Model.Position.Y)
	sx, sy := e.TextureId*frameWidth, 0
	screen.DrawImage(sprite.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)
}
