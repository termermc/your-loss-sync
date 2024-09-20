package setup

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/termermc/your-loss-sync/lang"
)

// Result is the result of the setup process.
type Result struct {
	Completed bool
	LangCode  string
}

// ShowSetup shows the setup window and returns the result.
func ShowSetup(app fyne.App) Result {
	resChan := make(chan Result)
	var locale lang.Locale

	w := app.NewWindow("")
	w.Resize(fyne.NewSize(300, 200))
	w.SetFixedSize(true)

	langLabel := widget.NewLabel("")
	confirmButton := widget.NewButton("", func() {
		resChan <- Result{
			Completed: true,
			LangCode:  locale.LangCode,
		}

		w.Close()
	})

	setLocale := func(langCode string) {
		locale = lang.NewLocale(langCode)

		w.SetTitle(locale.Tr("setup.title"))
		langLabel.SetText(locale.Tr("setup.select-language"))
		confirmButton.SetText(locale.Tr("general.confirm"))
	}
	setLocale(lang.DefaultLangCode)

	langSelect := widget.NewSelect(lang.GetLangNames(), func(langName string) {
		setLocale(lang.GetLangCodeFromName(langName))
	})
	langSelect.SetSelected(lang.Languages[lang.DefaultLangCode])

	w.SetContent(
		container.NewPadded(
			container.NewVScroll(
				container.NewVBox(
					container.NewCenter(
						langLabel,
					),
					langSelect,
					confirmButton,
				),
			),
		),
	)

	w.SetOnClosed(func() {
		resChan <- Result{
			Completed: false,
		}
	})

	w.Show()
	return <-resChan
}
