package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// TextListItemData is data required to render a TextList item.
type TextListItemData struct {
	// The item's label text.
	Label string

	// Whether the item can be deleted.
	CanDelete bool

	// The function to call when attempting to delete the item.
	OnDelete func()
}

// TextList is a list of text items.
type TextList struct {
	// The underlying widget.
	Widget *widget.List
}

// NewTextList creates a new TextList.
func NewTextList(lenFunc func() int, updateFunc func(id widget.ListItemID) TextListItemData) *TextList {
	var refresh func()

	list := widget.NewList(
		lenFunc,
		func() fyne.CanvasObject {
			// Create item

			scroll := container.NewHScroll(container.NewHBox())
			scroll.SetMinSize(fyne.NewSize(140, 40))
			return scroll
		},
		func(id widget.ListItemID, o fyne.CanvasObject) {
			// Update item

			updateRes := updateFunc(id)

			label := widget.NewLabel("")
			btn := widget.NewButton("✕", func() {
				updateRes.OnDelete()
				refresh()
			})

			scroll := o.(*container.Scroll)
			hbox := scroll.Content.(*fyne.Container)

			labelScroll := container.NewHScroll(label)
			labelScroll.SetMinSize(fyne.NewSize(165, 40))
			hbox.Objects = []fyne.CanvasObject{
				labelScroll,
				btn,
			}

			label.SetText(updateRes.Label)
			if !updateRes.CanDelete {
				btn.Hide()
			}

			hbox.Refresh()
		},
	)

	refresh = func() {
		list.Refresh()
	}

	return &TextList{
		Widget: list,
	}
}
