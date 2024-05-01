package npc

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
)

type Npc struct {
	StartPosX int
	StartPosY int
	Model     *resolv.Object
	Sprite    *ebiten.Image
	State     int
	Count     int
	FrameOX   int
	FrameOY   int
	Name      string
}
type Npc_data struct {
	StartPosX    int
	StartPosY    int
	Sprite_asset string
}

func (n *Npc) Draw(screen *ebiten.Image, frameHeight, frameWidth, frameCount int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(n.Model.Position.X), n.Model.Position.Y-float64(frameHeight/2))
	i := (n.Count / 12) % frameCount
	sx, sy := n.FrameOX+i*frameWidth, n.FrameOY
	screen.DrawImage(n.Sprite.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)
}

func (n *Npc) Update() {
}
