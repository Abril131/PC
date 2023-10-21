package scenes

import (
	"Juego/models"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"time"
)

type GameScene struct {
	window     fyne.Window
	posX, posY float32
	app        fyne.App
}

var c *models.Chapu
var p *models.Personaje
var s *models.Sprite
var c2 *models.Chapu2
var c3 *models.Chapu3

func NewGamesScene(window fyne.Window, app fyne.App) *GameScene {
	return &GameScene{window: window, app: app}
}

func (me *GameScene) Show() {
	backgroundImage := canvas.NewImageFromURI(storage.NewFileURI("./assets/fondo.png"))
	backgroundImage.Resize(fyne.NewSize(1240, 720))
	backgroundImage.Move(fyne.NewPos(0, 0))
	personaje := canvas.NewImageFromURI(storage.NewFileURI("./assets/personaje.png"))
	personaje.Resize(fyne.NewSize(50, 50))
	personaje.Move(fyne.NewPos(100, 100))
	chapu := canvas.NewImageFromURI(storage.NewFileURI("./assets/chapu.png"))
	chapu.Resize(fyne.NewSize(50, 50))
	chapu.Move(fyne.NewPos(600, 100))
	sprite := canvas.NewImageFromURI(storage.NewFileURI("./assets/Sprite.png"))
	sprite.Resize(fyne.NewSize(50, 50))
	sprite.Move(fyne.NewPos(600, 100))
	chapu2 := canvas.NewImageFromURI(storage.NewFileURI("./assets/chapu.png"))
	chapu2.Resize(fyne.NewSize(50, 50))
	chapu2.Move(fyne.NewPos(600, 100))
	chapu3 := canvas.NewImageFromURI(storage.NewFileURI("./assets/chapu.png"))
	chapu3.Resize(fyne.NewSize(50, 50))
	chapu3.Move(fyne.NewPos(600, 100))

	//Creamos el modelo
	p = models.NewPersonaje(100, 100, personaje)
	c = models.NewChapu(50, 100, chapu)
	c2 = models.NewChapu2(50, 200, chapu2)
	c3 = models.NewChapu3(50, 300, chapu3)
	s = models.NewSprite(100, 100, sprite)

	botonIniciar := widget.NewButton("Start Game", func() {
		me.StartGame(me.window)
	})
	botonIniciar.Resize(fyne.NewSize(150, 30))
	botonIniciar.Move(fyne.NewPos(100, 10))

	botonDetener := widget.NewButton("Stop Game", me.StopGame)
	botonDetener.Resize(fyne.NewSize(150, 30))
	botonDetener.Move(fyne.NewPos(300, 10))

	me.window.SetContent(container.NewWithoutLayout(backgroundImage, personaje, chapu, chapu2, chapu3, sprite, botonIniciar, botonDetener))
}

func (me *GameScene) StartGame(window fyne.Window) {
	p.Run(window)
	go c.Run()
	go c2.Run()
	go c3.Run()
	go s.Run()

	collisionCh := make(chan bool)
	collisionCh2 := make(chan bool)

	time.Sleep(1 * time.Second)

	// Goroutine para detectar la colisión
	go func() {
		for {
			if me.CheckCollision(p, s) {
				collisionCh <- true
				return
			}
			time.Sleep(300 * time.Millisecond)
		}
	}()

	go func() {
		for {
			if me.EndCollision(p, c) {
				collisionCh2 <- true
				return
			}
			time.Sleep(300 * time.Millisecond)
		}
	}()

	go func() {
		for {
			if me.End2Collision(p, c2) {
				collisionCh2 <- true
				return
			}
			time.Sleep(300 * time.Millisecond)
		}
	}()

	go func() {
		for {
			if me.End3Collision(p, c3) {
				collisionCh2 <- true
				return
			}
			time.Sleep(300 * time.Millisecond)
		}
	}()

	go func() {
		for {
			select {
			case <-collisionCh:
				me.ShowEnd()
				return
			default:
				time.Sleep(100 * time.Millisecond) // Espera sin bloquear
			}
		}
	}()

	go func() {
		for {
			select {
			case <-collisionCh2:
				me.Ends()
				return
			default:
				time.Sleep(100 * time.Millisecond) // Espera sin bloquear
			}
		}
	}()
}

func (me *GameScene) CheckCollision(p *models.Personaje, s *models.Sprite) bool {
	pos1 := p.GetPosition()
	pos2 := s.GetPosition()

	width1 := float32(p.Perso.Size().Width)
	height1 := float32(p.Perso.Size().Height)
	width2 := float32(s.Spri.Size().Width)
	height2 := float32(s.Spri.Size().Height)

	if pos1.X < pos2.X+width2 &&
		pos1.X+width1 > pos2.X &&
		pos1.Y < pos2.Y+height2 &&
		pos1.Y+height1 > pos2.Y {
		// Ha ocurrido una colisión
		return true
	}
	return false
}

func (me *GameScene) EndCollision(p *models.Personaje, c *models.Chapu) bool {
	pos1 := p.GetPosition()
	pos2 := c.GetPosition()

	width1 := float32(p.Perso.Size().Width)
	height1 := float32(p.Perso.Size().Height)
	width2 := float32(c.Cha.Size().Width)
	height2 := float32(c.Cha.Size().Height)

	if pos1.X < pos2.X+width2 &&
		pos1.X+width1 > pos2.X &&
		pos1.Y < pos2.Y+height2 &&
		pos1.Y+height1 > pos2.Y {
		// Ha ocurrido una colisión
		return true
	}
	return false
}

func (me *GameScene) End2Collision(p *models.Personaje, c2 *models.Chapu2) bool {
	pos1 := p.GetPosition()
	pos2 := c2.GetPosition()

	width1 := float32(p.Perso.Size().Width)
	height1 := float32(p.Perso.Size().Height)
	width2 := float32(c2.Cha2.Size().Width)
	height2 := float32(c2.Cha2.Size().Height)

	if pos1.X < pos2.X+width2 &&
		pos1.X+width1 > pos2.X &&
		pos1.Y < pos2.Y+height2 &&
		pos1.Y+height1 > pos2.Y {
		// Ha ocurrido una colisión
		return true
	}
	return false
}
func (me *GameScene) End3Collision(p *models.Personaje, c3 *models.Chapu3) bool {
	pos1 := p.GetPosition()
	pos2 := c3.GetPosition()

	width1 := float32(p.Perso.Size().Width)
	height1 := float32(p.Perso.Size().Height)
	width2 := float32(c3.Cha3.Size().Width)
	height2 := float32(c3.Cha3.Size().Height)

	if pos1.X < pos2.X+width2 &&
		pos1.X+width1 > pos2.X &&
		pos1.Y < pos2.Y+height2 &&
		pos1.Y+height1 > pos2.Y {
		// Ha ocurrido una colisión
		return true
	}
	return false
}

func (me *GameScene) StopGame() {
	c2.SetStatus(false)
	c3.SetStatus(false)
	p.SetStatus(false)
	c.SetStatus(false)
	s.SetStatus(false)
}

func (me *GameScene) ShowEnd() {
	scene := NewWinScenes(me.window, me.app)
	scene.Win()

}

func (me *GameScene) Ends() {
	scene := NewEndScenes(me.window, me.app)
	scene.End()

}
