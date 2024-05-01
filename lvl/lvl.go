package lvl

import (
	"main/npc"
)

type Lvl_data struct {
	Npces   []*npc.Npc_data
	Lvl_map []string
}

var lvl1_map = [...]string{
	"b b b b b b b b b b b b b b b b",
	"b g g g g g g g g g g g g g g b",
	"b g g g g g g g g g g g g g g b",
	"b g g g g g g g g g g g g g g b",
	"b g g g g g g g g g g g g g g b",
	"b g g g g g g g g g g g g g g b",
	"b g g g g g g g g g g g g g g b",
	"b g g g g g g g g g g g g g g b",
	"b g g g g g g g g g g g g g g b",
	"b g g g g g g g g g g g g g g b",
	"b g g g g g g g g g g g g g g b",
	"b g g g g g g g g g g g g g g b",
	"b g g g g g g g g g g g g g g b",
	"b g g g g g g g g g g g g g g b",
	"b g g g g g g g g g g g g g g b",
	"b g g g g g g g g g g g g g g b",
	"b b b b b b b b b b b b b b b b",
}

var Lvl1_data = Lvl_data{
	Lvl_map: lvl1_map[:],
	Npces: []*npc.Npc_data{
		&npc.Npc_data{
			StartPosX:    200,
			StartPosY:    130,
			Sprite_asset: "Mama.png",
		},
		&npc.Npc_data{
			StartPosX:    420,
			StartPosY:    150,
			Sprite_asset: "Vitalik.png",
		},
	},
}
