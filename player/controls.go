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
)

var Keymap = input.Keymap{
	ActionMoveLeft:  {input.KeyGamepadLeft, input.KeyLeft, input.KeyA},
	ActionMoveRight: {input.KeyGamepadRight, input.KeyRight, input.KeyD},
	ActionMoveTop:   {input.KeyGamepadLeft, input.KeyUp, input.KeyW},
	ActionMoveDown:  {input.KeyGamepadRight, input.KeyDown, input.KeyS},
	ActionInteract:  {input.KeyGamepadA, input.KeyE},
	ActionUnbound:   {},
}
