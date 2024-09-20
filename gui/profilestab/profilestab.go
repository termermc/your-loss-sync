package profilestab

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type ProfilesTab struct {
	Widget fyne.CanvasObject
}

// New creates a new ProfilesTab
func New() ProfilesTab {
	list := widget.NewList(
		func() int {
			return 0
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Profile 1")
		},
		func(id widget.ListItemID, o fyne.CanvasObject) {

		})
	cont := container.NewHBox(list, container.NewVBox())

	return ProfilesTab{
		Widget: cont,
	}
}
