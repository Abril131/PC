// models/persona.go

package models

import (
	"fyne.io/fyne/theme"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func CreateCharacter() fyne.CanvasObject {
	character := canvas.NewImageFromFile("assets/personaje.png")
	character.FillMode = canvas.ImageFillContain
	character.Resize(fyne.NewSize(50, 50))
	return character
}

func CreateCharacterContainer() *fyne.Container {
	character := CreateCharacter()
	blankSpace := canvas.NewRectangle(theme.BackgroundColor())
	blankSpace.SetMinSize(fyne.NewSize(50, 50))

	characterContainer := container.NewVBox(blankSpace, character)

	return characterContainer
}
