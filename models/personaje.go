package models

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type Personaje struct {
	posX, posY float32
	status     bool
	Perso      *canvas.Image
	width      float32
}

func NewPersonaje(posX, posY float32, img *canvas.Image) *Personaje {
	return &Personaje{
		posX:   posX,
		posY:   posY,
		status: true,
		Perso:  img,
	}
}

func (p *Personaje) GetPosX() float32 {
	return p.posX
}

func (p *Personaje) GetPosY() float32 {
	return p.posY
}

func (p *Personaje) GetPosition() fyne.Position {
	return fyne.NewPos(p.posX, p.posY)
}

func (p *Personaje) Run(window fyne.Window) {
	p.status = true

	//for p.status {
	window.Canvas().SetOnTypedKey(func(key *fyne.KeyEvent) {
		switch key.Name {
		case fyne.KeyRight:
			fmt.Println("right")
			p.right()
		case fyne.KeyUp:
			fmt.Println("up")
			p.up()
		case fyne.KeyDown:
			fmt.Println("down")
			p.down()
		case fyne.KeyLeft:
			fmt.Println("left")
			p.left()
		default:
			fmt.Println("moved not valid")
		}
	})

	//}
}

func (p *Personaje) SetStatus(status bool) {
	p.status = status
}

func (p *Personaje) right() {
	fmt.Println(p.posX, p.posY, "antes if")
	if p.status && p.posX <= 1170 {
		fmt.Println(p.posX, p.posY, "if")
		p.posX += 10

	} else {
		p.posX = 1170
	}
	p.Perso.Move(fyne.NewPos(p.posX, p.posY))
	p.Perso.Refresh()
}

func (p *Personaje) left() {
	fmt.Println(p.posX, p.posY, "antes if")
	if p.status && p.posX >= 0 {
		fmt.Println(p.posX, p.posY, "if")
		p.posX -= 10

	} else {
		p.posX = 0
	}
	p.Perso.Move(fyne.NewPos(p.posX, p.posY))
	p.Perso.Refresh()
}

func (p *Personaje) down() {
	fmt.Println(p.posX, p.posY, "antes if")
	if p.status && p.posY <= 660 {
		fmt.Println(p.posX, p.posY, "if")
		p.posY += 10

	} else {
		p.posY = 660
	}
	p.Perso.Move(fyne.NewPos(p.posX, p.posY))
	p.Perso.Refresh()
}

func (p *Personaje) up() {
	fmt.Println(p.posX, p.posY, "antes if")
	if p.status && p.posY >= 0 {
		fmt.Println(p.posX, p.posY, "if")
		p.posY -= 10

	} else {
		p.posY = 0
	}
	p.Perso.Move(fyne.NewPos(p.posX, p.posY))
	p.Perso.Refresh()
}
