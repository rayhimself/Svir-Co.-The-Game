package main

import (
	"fmt"
	"math"
	"image"
	"image/png"
	_ "image/png"
	"log"
	"main/k8s"
	"main/lvl"
	"main/entity"
	"main/player"
	"os"
	"strings"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	input "github.com/quasilyte/ebitengine-input"
	"github.com/solarlune/resolv"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	restclient "k8s.io/client-go/rest"

)

var (
	mplusNormalFont font.Face
	test_kal = []*entity.Entity {
		{
			Name: "psheno",
			TextureId: 0,
			Model: resolv.NewObject(150, 30, bacgroundTextureSize, bacgroundTextureSize),
		},
		{
			Name: "plod",
			TextureId: 1,
			Model: resolv.NewObject(150, 80, bacgroundTextureSize, bacgroundTextureSize),
			},
	}
)

const (
	screenWidth          = 256+128
	screenHeight         = 256
	bacgroundTextureSize = 16
	frameWidth           = 48
	frameHeight          = 48
	frameCount           = 2
	moveSpd              = 4.0
	dpi                  = 72
)

var (
	playerImage    *ebiten.Image
	bacgroundImage *ebiten.Image
	entityImage *ebiten.Image
	logo           []image.Image
)

type Game struct {
	inputSystem input.System
	p           *player.Player
	e			[]*entity.Entity
	space       *resolv.Space
	lvl_map     []string
	curentLvl   int
	exitPosX    int
	exitPosY    int
	lvl_type    string
	count       int
	kubeConfig *restclient.Config
}

func (g *Game) Update() error {
	if g.p.Input.ActionIsJustPressed(player.ActionInteract) {
		for i := 0; i < len(g.e); i++ {
			if math.Abs(g.p.Model.Position.X-g.e[i].Model.Position.X) < 32 && math.Abs(g.p.Model.Position.Y-g.e[i].Model.Position.Y) < 32 {
				fmt.Println(k8s.GetDeploy("dev", g.kubeConfig))
				k8s.DeleteDeploy("dev", g.e[i].Name, g.kubeConfig)
			}
		}
	}
	if g.p.Input.ActionIsJustPressed(player.ActionPlant) {
		for i := 0; i < len(g.e); i++ {
			if math.Abs(g.p.Model.Position.X-g.e[i].Model.Position.X) < 32 && math.Abs(g.p.Model.Position.Y-g.e[i].Model.Position.Y) < 32 {
				fmt.Println(k8s.GetDeploy("dev", g.kubeConfig))
				k8s.CreateDeploy("dev", g.e[i].Name, g.kubeConfig)
			}
		}
	}
	g.p.Count++
	g.count++
	g.inputSystem.Update()
	g.p.Update(frameHeight, frameWidth)
	return nil
}

func startGame() *Game {
	g := &Game{}
	g.inputSystem.Init(input.SystemConfig{
		DevicesEnabled: input.AnyDevice,
	})
	g.kubeConfig = k8s.GerCubeConfig()
	tt, _ := opentype.Parse(fonts.MPlus1pRegular_ttf)

	mplusNormalFont, _ = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    12,
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})
	g.curentLvl = 1
	g.SetUpLevel(&lvl.Lvl1_data)
	g.SetUpPlayer(&lvl.Lvl1_data)

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
	for i := 0; i < len(g.e); i++ {
		g.e[i].Draw(entityImage, screen, bacgroundTextureSize, bacgroundTextureSize)
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) ClearLevel() *Game {
	g.space = nil
	g.p = nil
	return g
}

func (g *Game) SetUpLevel(lvl_data *lvl.Lvl_data) *Game {
	g.lvl_map = lvl_data.Lvl_map
	g.e = test_kal
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

func (g *Game) SetUpPlayer(lvl_data *lvl.Lvl_data) *Game {
	g.p = &player.Player{
		StartPosX: 70,
		StartPosY: 70,
		Input:     g.inputSystem.NewHandler(0, player.Keymap),
		FrameOX:   0,
		FrameOY:   0,
		Sprite:    playerImage,
		Model:     resolv.NewObject(70, 70, frameWidth/2, frameHeight/2),
	}
	g.space.Add(g.p.Model)
	return g
}

func main() {
	
	// Decode an image from the image file's byte slice.
	playerImage = Loader("player.png")
	bacgroundImage = Loader("TexturePack.png")
	logo_file, err := os.Open("_assets/Vano_face.png")
	entityImage = Loader("Entitys.png")
	if err != nil {
	}
	defer logo_file.Close()
	logo_img, _ := png.Decode(logo_file)
	logo = append(logo, logo_img)
	ebiten.SetWindowTitle("Svir-CO. - The Game")
	ebiten.SetWindowIcon(logo)
	if err := ebiten.RunGame(startGame()); err != nil {
		log.Fatal(err)
	}
}
