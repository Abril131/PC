// models/sprites.go
package models

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"time"
)

type Chapu struct {
	posX, posY float32
	status     bool
	Cha        *canvas.Image
	width      float64
	height     float32
	puntos     int
	endGameCh  chan struct{}
}

func NewChapu(posx float32, posy float32, img *canvas.Image) *Chapu {
	return &Chapu{
		posX:   posx,
		posY:   posy,
		status: true,
		Cha:    img,
		puntos: 0,
	}
}

func (c *Chapu) GetPosition() fyne.Position {
	return fyne.NewPos(c.posX, c.posY)
}

func (c *Chapu) Run() {
	var incX float32 = 50
	var incY float32 = 50
	c.status = true
	for c.status {
		if c.posX < 10 || c.posX > 1150 {
			incX *= -1
		}
		if c.posY < 10 || c.posY > 540 {
			incY *= -1
		}
		select {
		case <-c.endGameCh:
			c.status = false // Finaliza el juego cuando se recibe una señal
		default:
		}
		c.posX += incX
		c.posY += incY
		c.Cha.Move(fyne.NewPos(float32(c.posX), float32(c.posY)))
		time.Sleep(150 * time.Millisecond)
	}
}

func (c *Chapu) SetStatus(status bool) {
	c.status = status
}

func (c Chapu) SetStat(b bool) {
}
