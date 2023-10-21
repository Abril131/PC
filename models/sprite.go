package models

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"time"
)

type Sprite struct {
	posX, posY float32
	status     bool
	Spri       *canvas.Image
	width      float64
	height     float32
	puntos     int
	endGameCh  chan struct{}
}

func NewSprite(posx float32, posy float32, img *canvas.Image) *Sprite {
	return &Sprite{
		posX:   posx,
		posY:   posy,
		status: true,
		Spri:   img,
		puntos: 0,
	}

}

func (s *Sprite) GetPosition() fyne.Position {
	return fyne.NewPos(s.posX, s.posY)
}

func (s *Sprite) Run() {
	var incX float32 = 50
	var incY float32 = 50
	s.status = true
	for s.status {
		if s.posX < 50 || s.posX > 1150 {
			incX *= -1
		}
		if s.posY < 50 || s.posY > 540 {
			incY *= -1
		}
		select {
		case <-s.endGameCh:
			s.status = false // Finaliza el juego cuando se recibe una señal
		default:
		}
		s.posX += incX
		s.posY += incY
		s.Spri.Move(fyne.NewPos(float32(s.posX), float32(s.posY)))
		time.Sleep(300 * time.Millisecond)
	}
}

func (s *Sprite) SetStatus(status bool) {
	s.status = status
}
func (s *Sprite) SetStat(b bool) {

}

func (s *Sprite) AumentarPuntos() {
	s.puntos += 10 // Aumenta en 10 puntos cada vez
}

func (s *Sprite) GetPuntos() int {
	return s.puntos
}
