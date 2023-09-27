package scenes

import (
	"Juego/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func CreateScene() fyne.Window {
	a := app.New()
	w := a.NewWindow("Escenario")
	characterContainer := models.CreateCharacterContainer()
	objectContainer := models.CreateObjectContainer()

	exitButton := widget.NewButton("Salir", func() {
		a.Quit()
	})
	exitButton.Resize(fyne.NewSize(50, 30))

	buttonContainer := container.NewVBox(exitButton)

	blankSpace := canvas.NewRectangle(theme.BackgroundColor())
	blankSpace.SetMinSize(fyne.NewSize(50, 30))

	mainContainer := container.NewVBox(objectContainer, blankSpace, characterContainer)

	buttonContainer := container.NewVBox(exitButton)

	scene := container.NewBorder(nil, nil, buttonContainer, nil, mainContainer)

	w.SetContent(scene)

	w.Resize(fyne.NewSize(1100, 700))
	return w

}
