package settingstab

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type SettingsTab struct {
	Widget fyne.CanvasObject
}

// New creates a new SettingsTab
func New() SettingsTab {
	list := widget.NewList(
		func() int {
			return 0
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Sync 1")
		},
		func(id widget.ListItemID, o fyne.CanvasObject) {

		})
	cont := container.NewHBox(list, container.NewVBox())

	return SettingsTab{
		Widget: cont,
	}
}
