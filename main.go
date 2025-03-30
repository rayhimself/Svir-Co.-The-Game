package main

import (
	"image"
	"image/png"
	_ "image/png"
	"fmt"
	"main/k8s"
	"main/lvl"
	"main/player"
	"os"
	"log"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	input "github.com/quasilyte/ebitengine-input"
	"github.com/solarlune/resolv"
	restclient "k8s.io/client-go/rest"

)

const (
	screenWidth          = 256+128
	screenHeight         = 256
	bacgroundTextureSize = 16
	frameWidth           = 16
	frameHeight          = 32
	frameCount           = 2
	moveSpd              = 2.0
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
	space       *resolv.Space
	lvl_map     *lvl.LvlMap
	count       int
	kubeConfig *restclient.Config
}

func (g *Game) Update() error {
	if g.p.Input.ActionIsJustPressed(player.ActionInteract) {
		switch g.p.Item {
			case 1: g.lvl_map.Break(g.p.TargetX, g.p.TargetY, "dev", g.kubeConfig)
			case 2: g.lvl_map.Plant(g.p.TargetX, g.p.TargetY, "apple", "dev", g.kubeConfig)
			case 3: g.lvl_map.Plant(g.p.TargetX, g.p.TargetY, "wheat", "dev", g.kubeConfig)
		}
	}
	// сюда интеракт с логикой что делать при каком предмете 
	g.p.Count++
	g.count++
	g.inputSystem.Update()
	g.p.Update(frameHeight, frameWidth, bacgroundTextureSize)
	return nil
}

func startGame() *Game {
	g := &Game{}
	g.inputSystem.Init(input.SystemConfig{
		DevicesEnabled: input.AnyDevice,
	})
	g.kubeConfig = k8s.GerCubeConfig()

	g.SetUpLevel()
	g.SetUpPlayer()

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
	g.lvl_map.Draw(bacgroundImage, entityImage, screen)
	g.p.Draw(screen, frameHeight, frameWidth, frameCount, bacgroundTextureSize)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Now holding: %s", player.Items[g.p.Item]))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) SetUpLevel() *Game {
	g.lvl_map = &lvl.LvlMap{
		Height: screenHeight,
		Width: screenWidth,
		TextureSize: bacgroundTextureSize,
	}
	g.space = resolv.NewSpace(screenWidth, screenHeight, bacgroundTextureSize, bacgroundTextureSize)
	g.lvl_map.AddModel(g.space)
	g.lvl_map.CreateGrid()
	g.lvl_map.FillLvl("dev", g.kubeConfig)
	return g
}

func (g *Game) SetUpPlayer() *Game {
	g.p = &player.Player{
		StartPosX: 70,
		StartPosY: 70,
		Input:     g.inputSystem.NewHandler(0, player.Keymap),
		FrameOX:   0,
		FrameOY:   0,
		Sprite:    playerImage,
		Model:     resolv.NewObject(70, 70+frameHeight/2, frameWidth, frameHeight/2),
		Direction: "Down",
		Item:      1,
	}
	g.space.Add(g.p.Model)
	return g
}

func main() {
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
