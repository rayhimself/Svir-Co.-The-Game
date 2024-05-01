package player

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	input "github.com/quasilyte/ebitengine-input"
	"github.com/solarlune/resolv"
)

type Player struct {
	StartPosX int
	StartPosY int
	Input     *input.Handler
	Sprite    *ebiten.Image
	Model     *resolv.Object
	Count     int
	FrameOX   int
	FrameOY   int
	IsLocked  bool
}

func (p *Player) Draw(screen *ebiten.Image, frameHeight, frameWidth, frameCount int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(p.Model.Position.X), p.Model.Position.Y-float64(frameHeight/2))
	i := (p.Count / 12) % frameCount
	sx, sy := p.FrameOX+i*frameWidth, p.FrameOY
	screen.DrawImage(p.Sprite.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)
}

func (p *Player) Update() {
	p.FrameOY = 0
	p.FrameOX = 0
	dx, dy := 0.0, 0.0
	moveSpd := 4.0

	if p.Input.ActionIsPressed(ActionMoveLeft) {
		dx = -moveSpd
		p.FrameOY = 32
		p.FrameOX = 32 * 6
	}
	if p.Input.ActionIsPressed(ActionMoveRight) {
		dx += moveSpd
		p.FrameOY = 32
		p.FrameOX = 0
	}
	if p.Input.ActionIsPressed(ActionMoveDown) {
		dy += moveSpd
		p.FrameOY = 32
		p.FrameOX = 32 * 9
	}
	if p.Input.ActionIsPressed(ActionMoveTop) {
		dy = -moveSpd
		p.FrameOY = 32
		p.FrameOX = 32 * 3
	}
	if collision := p.Model.Check(dx, 0); collision != nil {
		dx = 0
	}
	if collision := p.Model.Check(0, dy); collision != nil {
		dy = 0
	}
	if p.IsLocked {
		dx, dy = 0, 0
		p.FrameOY = 0
		p.FrameOX = 0
	}
	p.Model.Position.X += dx
	p.Model.Position.Y += dy
	p.Model.Update()
}
