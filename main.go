package main

import (
	"fmt"
	"image"
	"image/png"
	_ "image/png"
	"log"
	"main/k8s"
	"main/lvl"
	"main/npc"
	"main/player"
	"math"
	"os"
	"strings"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	input "github.com/quasilyte/ebitengine-input"
	"github.com/solarlune/resolv"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var (
	mplusNormalFont font.Face
)

const (
	screenWidth          = 512
	screenHeight         = 256
	bacgroundTextureSize = 16
	frameWidth           = 16
	frameHeight          = 32
	frameCount           = 6
	moveSpd              = 4.0
	dpi                  = 72
)

var (
	playerImage    *ebiten.Image
	bacgroundImage *ebiten.Image
	logo           []image.Image
)

type Game struct {
	inputSystem input.System
	p           *player.Player
	n           []*npc.Npc
	space       *resolv.Space
	lvl_map     []string
	curentLvl   int
	exitPosX    int
	exitPosY    int
	lvl_type    string
	count       int
}

func (g *Game) Update() error {
	lvlNumber := g.curentLvl
	if g.p.Input.ActionIsJustPressed(player.ActionInteract) {
		for i := 0; i < len(g.n); i++ {
			if math.Abs(g.p.Model.Position.X-g.n[i].Model.Position.X) < 32 && math.Abs(g.p.Model.Position.Y-g.n[i].Model.Position.Y) < 32 {

			}
		}
		if math.Abs(g.p.Model.Position.X-float64(g.exitPosX)) < 32 && math.Abs(g.p.Model.Position.Y-float64(g.exitPosY)) < 32 {
			lvlNumber++
		}
	}
	g.p.Count++
	g.count++
	g.inputSystem.Update()
	g.p.Update()
	for i := 0; i < len(g.n); i++ {
		g.n[i].Update()
		g.n[i].Count++
	}
	return nil
}

func startGame() *Game {
	g := &Game{}
	g.inputSystem.Init(input.SystemConfig{
		DevicesEnabled: input.AnyDevice,
	})
	tt, _ := opentype.Parse(fonts.MPlus1pRegular_ttf)

	mplusNormalFont, _ = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    12,
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})
	g.curentLvl = 1
	g.SetUpLevel(&lvl.Lvl1_data)
	g.SetUpPlayer(&lvl.Lvl1_data)
	g.SetUpNpces(&lvl.Lvl1_data)
	return g
}

func Loader(path string) *ebiten.Image {
	asset, _, err := ebitenutil.NewImageFromFile("_assets/" + path)
	if err != nil {
		log.Fatal(err)
	}
	ebitenImage := ebiten.NewImageFromImage(asset)
	return ebitenImage
}

func (g *Game) Draw(screen *ebiten.Image) {
	for i := 0; i < len(g.lvl_map); i++ {
		line := strings.Split(g.lvl_map[i], " ")
		for j := 0; j < len(line); j++ {
			op := &ebiten.DrawImageOptions{}
			sx, sy := lvl.MapObjects[line[j]].OX, lvl.MapObjects[line[j]].OY
			op.GeoM.Translate(float64(i*bacgroundTextureSize), float64(j*bacgroundTextureSize))
			screen.DrawImage(bacgroundImage.SubImage(image.Rect(sx, sy, sx+bacgroundTextureSize, sy+bacgroundTextureSize)).(*ebiten.Image), op)
		}
	}
	g.p.Draw(screen, frameHeight, frameWidth, frameCount)
	for i := 0; i < len(g.n); i++ {
		g.n[i].Draw(screen, frameHeight, frameWidth, frameCount)
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) ClearLevel() *Game {
	g.space = nil
	g.n = nil
	g.p = nil
	return g
}

func (g *Game) SetUpLevel(lvl_data *lvl.Lvl_data) *Game {
	g.lvl_map = lvl_data.Lvl_map
	g.space = resolv.NewSpace(screenWidth, screenHeight, bacgroundTextureSize, bacgroundTextureSize)
	for i := 0; i < len(g.lvl_map); i++ {
		line := strings.Split(g.lvl_map[i], " ")
		for j := 0; j < len(line); j++ {
			if lvl.MapObjects[line[j]].IsObject {
				g.space.Add(resolv.NewObject(float64(i*bacgroundTextureSize), float64(j*bacgroundTextureSize), bacgroundTextureSize, bacgroundTextureSize))
			}
		}
	}
	return g
}

func (g *Game) SetUpNpces(lvl_data *lvl.Lvl_data) *Game {
	for i := 0; i < len(lvl_data.Npces); i++ {
		g.n = append(g.n, &npc.Npc{
			StartPosX: lvl_data.Npces[i].StartPosX,
			StartPosY: lvl_data.Npces[i].StartPosY,
			FrameOX:   0,
			FrameOY:   0,
			Sprite:    Loader(lvl_data.Npces[i].Sprite_asset),
			Model:     resolv.NewObject(float64(lvl_data.Npces[i].StartPosX), float64(lvl_data.Npces[i].StartPosY+frameHeight/2), frameWidth, frameHeight/2),
		})
	}
	for i := 0; i < len(g.n); i++ {
		g.space.Add(g.n[i].Model)
	}
	return g
}

func (g *Game) SetUpPlayer(lvl_data *lvl.Lvl_data) *Game {
	g.p = &player.Player{
		StartPosX: 70,
		StartPosY: 70,
		Input:     g.inputSystem.NewHandler(0, player.Keymap),
		FrameOX:   0,
		FrameOY:   0,
		Sprite:    playerImage,
		Model:     resolv.NewObject(70, 70+frameHeight/2, frameWidth, frameHeight/2),
	}
	g.space.Add(g.p.Model)
	return g
}

func main() {
	// Decode an image from the image file's byte slice.
	playerImage = Loader("Vano.png")
	bacgroundImage = Loader("TexturePack.png")
	logo_file, err := os.Open("_assets/Vano_face.png")
	if err != nil {
	}
	defer logo_file.Close()
	logo_img, _ := png.Decode(logo_file)
	logo = append(logo, logo_img)
	ebiten.SetWindowTitle("Simulyator Stoyaniya V Uglu")
	ebiten.SetWindowIcon(logo)
	fmt.Print(k8s.Get_Pods("kube_system"))
	if err := ebiten.RunGame(startGame()); err != nil {
		log.Fatal(err)
	}
}
