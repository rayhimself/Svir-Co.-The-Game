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
	FirstItem
	SecondItem
	ThirdItem
	FourthItem
	ActionChangeItemUp
	ActionChangeItemDown
)

var Keymap = input.Keymap{
	ActionMoveLeft:  {input.KeyGamepadLeft, input.KeyLeft, input.KeyA},
	ActionMoveRight: {input.KeyGamepadRight, input.KeyRight, input.KeyD},
	ActionMoveTop:   {input.KeyGamepadUp, input.KeyUp, input.KeyW},
	ActionMoveDown:  {input.KeyGamepadDown, input.KeyDown, input.KeyS},
	ActionInteract:  {input.KeyGamepadA, input.KeyE, input.KeyMouseLeft},
	ActionChangeItemUp: {input.KeyWheelVertical},
	ActionChangeItemDown: {input.KeyWheelDown},
	FirstItem:  {input.KeyGamepadA, input.Key1},
	SecondItem: {input.KeyGamepadB, input.Key2},
	ThirdItem:  {input.KeyGamepadX, input.Key3},
	FourthItem: {input.KeyGamepadY, input.Key4},
	ActionUnbound:   {},
}
