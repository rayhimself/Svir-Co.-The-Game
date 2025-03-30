package lvl
import (
	"image"
	"github.com/solarlune/resolv"
	"github.com/hajimehoshi/ebiten/v2"
)
type Cell struct {
	IsPlaced bool
	Name string
	EntityTextureId int
	TextureId string
}

type Object struct {
	OX       int
	OY       int
	isObject bool
}

type LvlMap struct {
	Height int
	Width int
	TextureSize int
	Model     *resolv.Object
	Grid [][]*Cell
}

var mapObjects = map[string]Object{
	"b_1_1":    Object{OX: 0, OY: 0, isObject: true},
	"b_1_2":    Object{OX: 0, OY: 16, isObject: true},
	"b_1_3":    Object{OX: 0, OY: 32, isObject: true},
	"b_2_1":    Object{OX: 16, OY: 0, isObject: true},
	"b_2_2":    Object{OX: 16, OY: 16, isObject: false},
	"b_2_3":    Object{OX: 16, OY: 32, isObject: true},
	"b_3_1":    Object{OX: 32, OY: 0, isObject: true},
	"b_3_2":    Object{OX: 32, OY: 16, isObject: true},
	"b_3_3":    Object{OX: 32, OY: 32, isObject: true},
}

func (l *LvlMap) CreateGrid() {
	mapHeight := l.Height/l.TextureSize-1
	mapWidth := l.Width/l.TextureSize-1
	l.Grid = make([][]*Cell, mapWidth)
    for i := range l.Grid {
        l.Grid[i] = make([]*Cell, mapHeight)
		for j := range l.Grid[i]{
			l.Grid[i][j] = &Cell{
				IsPlaced: false,
				Name: "",
				TextureId: "",
			}
		}
    }
	l.Grid[0][0].TextureId = "b_1_1"
	for i := 1; i < mapWidth; i++ {
		for j := 1; j < mapHeight; j++ {
			l.Grid[0][j].TextureId = "b_1_2"
			l.Grid[i][j].TextureId = "b_2_2"
			l.Grid[mapWidth-1][j].TextureId = "b_3_2"
		}
		l.Grid[i][mapHeight-1].TextureId = "b_2_3"
		l.Grid[i][0].TextureId = "b_2_1"
		l.Grid[mapWidth-1][mapHeight-1].TextureId = "b_3_3"
		l.Grid[mapWidth-1][0].TextureId = "b_3_1"
		l.Grid[0][mapHeight-1].TextureId = "b_1_3"
	}
	
}

func (l *LvlMap) Draw(bacgroundImage, entityImage, screen *ebiten.Image) {
	for i := range l.Grid {
		for j := range l.Grid[i] {
			op := &ebiten.DrawImageOptions{}
			sx, sy := mapObjects[l.Grid[i][j].TextureId].OX, mapObjects[l.Grid[i][j].TextureId].OY
			op.GeoM.Translate(float64(i*l.TextureSize), float64(j*l.TextureSize))
			screen.DrawImage(bacgroundImage.SubImage(image.Rect(sx, sy, sx+l.TextureSize, sy+l.TextureSize)).(*ebiten.Image), op)
			if l.Grid[i][j].IsPlaced {
				sx, sy = l.Grid[i][j].EntityTextureId*l.TextureSize, 0
				screen.DrawImage(entityImage.SubImage(image.Rect(sx, sy, sx+l.TextureSize, sy+l.TextureSize)).(*ebiten.Image), op)
			}
		}
	}
}

func (l *LvlMap) AddModel(space *resolv.Space) {
	space.Add(resolv.NewObject(0, 0, float64(l.Width), float64(l.TextureSize)))
	space.Add(resolv.NewObject(0, 0, float64(l.TextureSize), float64(l.Height)))
	space.Add(resolv.NewObject(0, float64(l.Height - l.TextureSize), float64(l.Width), float64(l.Height)))
	space.Add(resolv.NewObject(float64(l.Width - l.TextureSize), 0, float64(l.Width - l.TextureSize), float64(l.Height)))

}

