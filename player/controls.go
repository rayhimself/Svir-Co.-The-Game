package player

import input "github.com/quasilyte/ebitengine-input"

const (
	ActionUnknown input.Action = iota
	ActionMoveLeft
	ActionMoveRight
	ActionMoveTop
	ActionMoveDown
	ActionUnbound
	ActionInteract
	ActionPlant
)

var Keymap = input.Keymap{
	ActionMoveLeft:  {input.KeyGamepadLeft, input.KeyLeft, input.KeyA},
	ActionMoveRight: {input.KeyGamepadRight, input.KeyRight, input.KeyD},
	ActionMoveTop:   {input.KeyGamepadUp, input.KeyUp, input.KeyW},
	ActionMoveDown:  {input.KeyGamepadDown, input.KeyDown, input.KeyS},
	ActionInteract:  {input.KeyGamepadA, input.KeyE},
	ActionPlant:  {input.KeyGamepadB, input.KeyF},
	ActionUnbound:   {},
}
