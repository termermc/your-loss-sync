package widget

import (
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	nativedialog "github.com/sqweek/dialog"
	"github.com/termermc/your-loss-sync/lang"
	"os"
)

type FilePicker struct {
	// The current path.
	// Empty is none.
	Path binding.String

	// Whether the file picker is disabled.
	IsDisabled binding.Bool

	// Whether the file picker is a directory picker.
	// If false, the file picker is a normal file picker.
	IsDirectoryPicker binding.Bool

	// The underlying widget.
	Widget fyne.CanvasObject

	parent fyne.Window
	locale lang.Locale
}

// NewFilePicker creates a new FilePicker.
func NewFilePicker(parent fyne.Window, locale lang.Locale) *FilePicker {
	fp := FilePicker{
		Path:              binding.NewString(),
		IsDisabled:        binding.NewBool(),
		IsDirectoryPicker: binding.NewBool(),
		locale:            locale,
	}

	showPicker := func() {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			dialog.ShowError(err, parent)
			return
		}

		if val, _ := fp.IsDirectoryPicker.Get(); val {
			browser := nativedialog.Directory()

			if res, _ := fp.Path.Get(); res == "" {
				browser.SetStartDir(homeDir)
			} else {
				browser.SetStartDir(res)
			}

			path, err := browser.Browse()
			if err != nil {
				if errors.Is(err, nativedialog.ErrCancelled) {
					return
				}

				dialog.ShowError(err, parent)
				return
			}

			if err := fp.Path.Set(path); err != nil {
				dialog.ShowError(err, parent)
				return
			}
		} else {
			browser := nativedialog.File()

			if res, _ := fp.Path.Get(); res == "" {
				browser.SetStartDir(homeDir)
			} else {
				browser.SetStartDir(res)
			}

			path, err := browser.Load()
			if err != nil {
				if errors.Is(err, nativedialog.ErrCancelled) {
					return
				}

				dialog.ShowError(err, parent)
				return
			}

			if err := fp.Path.Set(path); err != nil {
				dialog.ShowError(err, parent)
				return
			}
		}
	}

	btn := widget.NewButton("", showPicker)
	pathEntry := widget.NewEntry()
	pathEntry.Disable()

	entryCont := container.NewHScroll(pathEntry)
	entryCont.SetMinSize(fyne.NewSize(150, 30))
	fp.Widget = container.NewHBox(
		btn,
		entryCont,
	)

	fp.Path.AddListener(NewDataListener(func() {
		val, err := fp.Path.Get()
		if err != nil {
			dialog.ShowError(err, parent)
			return
		}

		pathEntry.SetText(val)
	}))
	fp.IsDirectoryPicker.AddListener(NewDataListener(func() {
		if val, _ := fp.IsDirectoryPicker.Get(); val {
			btn.SetText(locale.Tr("widget.file-picker.select-folder"))
		} else {
			btn.SetText(locale.Tr("widget.file-picker.select-file"))
		}
	}))
	fp.IsDisabled.AddListener(NewDataListener(func() {
		if val, _ := fp.IsDisabled.Get(); val {
			btn.Disable()
		} else {
			btn.Enable()
		}
	}))

	_ = fp.Path.Set("")
	_ = fp.IsDirectoryPicker.Set(false)
	_ = fp.IsDisabled.Set(false)

	return &fp
}

// Clear clears the path.
func (f *FilePicker) Clear() {
	_ = f.Path.Set("")
}

// HasPath returns whether the path is not empty.
func (f *FilePicker) HasPath() bool {
	val, _ := f.Path.Get()
	return val != ""
}
