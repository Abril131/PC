package models

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"time"
)

type Chapu2 struct {
	posX, posY float32
	status     bool
	Cha2       *canvas.Image
	width      float64
	height     float32
	puntos     int
	endGameCh  chan struct{}
}

func NewChapu2(posx float32, posy float32, img *canvas.Image) *Chapu2 {
	return &Chapu2{
		posX:   posx,
		posY:   posy,
		status: true,
		Cha2:   img,
		puntos: 0,
	}
}

func (c2 *Chapu2) GetPosition() fyne.Position {
	return fyne.NewPos(c2.posX, c2.posY)
}

func (c2 *Chapu2) Run() {
	var incX float32 = 50
	var incY float32 = 50
	c2.status = true
	for c2.status {
		if c2.posX < 20 || c2.posX > 1150 {
			incX *= -1
		}
		if c2.posY < 20 || c2.posY > 540 {
			incY *= -1
		}
		select {
		case <-c2.endGameCh:
			c2.status = false // Finaliza el juego cuando se recibe una señal
		default:
		}
		c2.posX += incX
		c2.posY += incY
		c2.Cha2.Move(fyne.NewPos(float32(c2.posX), float32(c2.posY)))
		time.Sleep(200 * time.Millisecond)
	}
}

func (c2 *Chapu2) SetStatus(status bool) {
	c2.status = status
}

func (c2 Chapu2) SetStat(b bool) {
}
