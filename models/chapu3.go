package models

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"time"
)

type Chapu3 struct {
	posX, posY float32
	status     bool
	Cha3       *canvas.Image
	width      float64
	height     float32
	puntos     int
	endGameCh  chan struct{}
}

func NewChapu3(posx float32, posy float32, img *canvas.Image) *Chapu3 {
	return &Chapu3{
		posX:   posx,
		posY:   posy,
		status: true,
		Cha3:   img,
		puntos: 0,
	}
}

func (c3 *Chapu3) GetPosition() fyne.Position {
	return fyne.NewPos(c3.posX, c3.posY)
}

func (c3 *Chapu3) GetDimensions() fyne.Position {
	return c3.GetPosition()
}

func (c3 *Chapu3) Run() {
	var incX float32 = 50
	var incY float32 = 50
	c3.status = true
	for c3.status {
		if c3.posX < 5 || c3.posX > 1150 {
			incX *= -1
		}
		if c3.posY < 5 || c3.posY > 540 {
			incY *= -1
		}
		select {
		case <-c3.endGameCh:
			c3.status = false // Finaliza el juego cuando se recibe una señal
		default:
		}
		c3.posX += incX
		c3.posY += incY
		c3.Cha3.Move(fyne.NewPos(float32(c3.posX), float32(c3.posY)))
		time.Sleep(100 * time.Millisecond)
	}
}

func (c3 *Chapu3) SetStatus(status bool) {
	c3.status = status
}

func (c3 Chapu3) SetStat(b bool) {
}
