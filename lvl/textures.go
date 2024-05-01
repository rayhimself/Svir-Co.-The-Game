package lvl

type Object struct {
	OX       int
	OY       int
	IsObject bool
}

var MapObjects = map[string]Object{
	"g": Object{OX: 0, OY: 0, IsObject: false},
	"b": Object{OX: 64, OY: 0, IsObject: true},
}
