// models/objetos.go

package models

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func CreateObject() fyne.CanvasObject {
	obj := canvas.NewImageFromFile("assets/objeto.png")
	obj.FillMode = canvas.ImageFillContain
	obj.Resize(fyne.NewSize(50, 50))
	return obj
}

func CreateObjectContainer() *fyne.Container {
	objects := make([]fyne.CanvasObject, 5)

	for i := 0; i < 5; i++ {
		obj := CreateObject()
		objects[i] = obj
	}

	return container.NewVBox(objects...)
}
