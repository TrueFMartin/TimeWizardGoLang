package main

import (
	//sprites "github.com/TrueFMartin/TimeWizardGolang/sprites"
	"github.com/dradtke/ecs-go"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "image/png"
	"log"
)

const (
	screenW, screenH = 720, 408
)

var (
	background   *ebiten.Image
	bgOp         *ebiten.DrawImageOptions
	world        *ecs.World
	err          error
	player       *ecs.Object
	skeletonBase *ecs.Object
	pipeBase     *ecs.Object
	globalScreen *ebiten.Image
	skel         skeleton
)

func init() {
	world = ecs.NewWorld()
	background, _, err = ebitenutil.NewImageFromFile(
		"resources/background/background.png")
	if err != nil {
		log.Fatal(err)
	}
	bgOp = &ebiten.DrawImageOptions{}
	bgOp.GeoM.Scale(1.5, 1.5)
	//tempImage, err, _ := ebitenutil.NewImageFromFile("resources/wizard/wizard (1)")
	//if err != nil {
	//	log.Fatal(err)
	//}
	player = ecs.NewObject(
		Image{image: getImage("resources/wizard/wizard (1).png"), numImages: 8, url: "resources/wizard/wizard", options: setScale(1, 1)},
		Position{x: 100, px: 100, y: 400, py: 400},
		Size{w: 67, h: 69},
		Velocity{})
	//tempImage, err, _ = ebitenutil.NewImageFromFile("resources/skeleton/skeleton alive")
	//if err != nil {
	//	log.Fatal(err)
	//}
	skeletonBase = ecs.NewObject(
		Image{image: getImage("resources/skeleton/skeleton alive (1).png"), numImages: 12, url: "resources/skeleton/skeleton alive", options: setScale(1.5, 1.5)},
		Position{x: 500, px: 500, y: 400, py: 400},
		Size{w: 60, h: 55.5},
		Velocity{})
	skel = skeleton{
		i: Image{image: getImage("resources/skeleton/skeleton alive (1).png"), numImages: 12, url: "resources/skeleton/skeleton alive", options: setScale(1.5, 1.5)},

		p: Position{x: 500, px: 500, y: 400, py: 400},

		s: Size{w: 60, h: 55.5},

		v: Velocity{},
	}
}

type skeleton struct {
	i Image
	p Position
	s Size
	v Velocity
}

func setScale(x, y float64) *ebiten.DrawImageOptions {
	scale := ebiten.DrawImageOptions{}
	scale.GeoM.Scale(x, y)
	return &scale
}
func getImage(url string) *ebiten.Image {
	var i, _, err = ebitenutil.NewImageFromFile(url)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

type Game struct{}

// Components
type (
	hasLoadedImg bool
	Image        struct {
		image     *ebiten.Image
		images    []*ebiten.Image
		options   *ebiten.DrawImageOptions
		num       int
		numImages int
		url       string
	}
	Position struct {
		x, px float64
		y, py float64
	}
	Size struct {
		h float64
		w float64
	}
	Velocity struct {
		x, y           float64
		speedX, speedY float64
	}
)

func movePlayer(v Velocity) Velocity {
	//FIXME add in movement v.prevY, player.prevX = player.y, player.x
	v.speedX = 3
	v.speedY = 3
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		v.y -= v.speedY
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		v.y += v.speedY
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		v.x -= v.speedX
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		v.x += v.speedX
	}
	return v
}
func loadImage(img hasLoadedImg) {

}
func drawSprite(i Image, p Position) Image {
	i.options.GeoM.Translate(p.x-p.px, p.y-p.py)
	globalScreen.DrawImage(i.image, i.options)
	return i
}
func updateSprite(p Position, v Velocity) (Position, Velocity) {
	p.px = p.x
	p.py = p.y
	p.x += v.x
	p.y += v.y
	v.x = 0
	v.y = 0
	return p, v
}
func (g *Game) Update() error {
	//update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(background, bgOp)
	screen.DrawImage(skel.i.image, skel.i.options)
	//screen.DrawImage(playerImg, player.imgOp)
	//screen.DrawImage(skeletonImg, skeleton.imgOp)
	globalScreen = screen
	world.Run()

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return screenW, screenH
}

func main() {

	ebiten.SetWindowSize(screenW, screenH)
	ebiten.SetWindowTitle("Time Wizard")
	world.AddSystem(ecs.System{Func: drawSprite})
	world.AddSystem(ecs.System{Func: movePlayer})
	world.AddSystem(ecs.System{Func: updateSprite})

	world.AddObject(skeletonBase)
	world.AddObject(player)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
