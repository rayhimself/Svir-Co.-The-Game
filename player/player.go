package player

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	input "github.com/quasilyte/ebitengine-input"
	"github.com/solarlune/resolv"
)

var mapDirection = map[string]int{ 
	"Down": 0,
	"Top": 1,
	"Left": 2,
	"Right": 3,
}

type Player struct {
	StartPosX int
	StartPosY int
	Input     *input.Handler
	Sprite    *ebiten.Image
	Model     *resolv.Object
	Count     int
	FrameOX   int
	FrameOY   int
	Direction string
	TargetX int
	TargetY int
	Item int
}

func (p *Player) Draw(screen *ebiten.Image, frameHeight, frameWidth, frameCount, textureSize int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(p.Model.Position.X), p.Model.Position.Y-float64(frameHeight/2))
	i := (p.Count / 12) % frameCount
	sx, sy := p.FrameOX+i*frameWidth, p.FrameOY
	screen.DrawImage(p.Sprite.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)
	tp := &ebiten.DrawImageOptions{}
	tp.GeoM.Translate(float64(p.TargetX*textureSize), float64(p.TargetY*textureSize))
	sx, sy = 128, 0
	screen.DrawImage(p.Sprite.SubImage(image.Rect(sx, sy, sx+textureSize, sy+textureSize)).(*ebiten.Image), tp)
}

func (p *Player) Update(frameHeight, frameWidth, textureSize int) {
	p.FrameOY = mapDirection[p.Direction] * frameHeight
	p.FrameOX = 0

	var (
		deltaX, deltaY float64
		moveSpeed      = 2.0
	)

	switch p.Direction {
	case "Down":
		p.TargetX = (int(p.Model.Position.X) + textureSize/2) / textureSize
		p.TargetY = (int(p.Model.Position.Y) + textureSize/2) / textureSize + 1
	case "Top":
		p.TargetX = (int(p.Model.Position.X) + textureSize/2) / textureSize
		p.TargetY = (int(p.Model.Position.Y) + textureSize/2) / textureSize - 1
	case "Left":
		p.TargetX = (int(p.Model.Position.X) + textureSize/2) / textureSize - 1
		p.TargetY = (int(p.Model.Position.Y) + textureSize/2) / textureSize
	case "Right":
		p.TargetX = (int(p.Model.Position.X) + textureSize/2) / textureSize + 1
		p.TargetY = (int(p.Model.Position.Y) + textureSize/2) / textureSize
	}

	updatePosition := func(direction string, dx, dy float64, frameOffset int) {
		p.Direction = direction
		deltaX, deltaY = dx, dy
		p.FrameOY = mapDirection[direction] * frameHeight
		p.FrameOX = frameWidth * frameOffset
	}

	if p.Input.ActionIsPressed(ActionMoveTop) {
		updatePosition("Top", 0, -moveSpeed, 2)
	}
	if p.Input.ActionIsPressed(ActionMoveDown) {
		updatePosition("Down", 0, moveSpeed, 2)
	}
	if p.Input.ActionIsPressed(ActionMoveLeft) {
		updatePosition("Left", -moveSpeed, 0, 2)
	}
	if p.Input.ActionIsPressed(ActionMoveRight) {
		updatePosition("Right", moveSpeed, 0, 2)
	}
	if p.Input.ActionIsPressed(ActionInteract) {
		p.FrameOY = mapDirection[p.Direction] * frameHeight
		if p.Item == 1 {
			p.FrameOX = frameWidth * 4
		} else {
			p.FrameOX = frameWidth * 6
		}
	}
	if p.Input.ActionIsPressed(FirstItem) {
		p.Item = 1
	}
	if p.Input.ActionIsPressed(SecondItem) {
		p.Item = 2
	}
	if p.Input.ActionIsPressed(ThirdItem) {
		p.Item = 3
	}
	// if p.Input.ActionIsPressed(ActionChangeItemUp) {
	// 	p.ChangeItem("Up")
	// }
	// if p.Input.ActionIsPressed(ActionChangeItemDown) {
	// 	p.ChangeItem("Down")
	// }
	if collision := p.Model.Check(deltaX, 0); collision != nil {
		deltaX = 0
	}
	if collision := p.Model.Check(0, deltaY); collision != nil {
		deltaY = 0
	}

	p.Model.Position.X += deltaX
	p.Model.Position.Y += deltaY
	p.Model.Update()
}
