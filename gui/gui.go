package gui

import (
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"github.com/termermc/your-loss-sync/config"
	"github.com/termermc/your-loss-sync/config/json"
	"github.com/termermc/your-loss-sync/gui/setup"
	"github.com/termermc/your-loss-sync/lang"
	"github.com/termermc/your-loss-sync/logic"
	"os"
	"time"
)

var defaultDialogSize = fyne.NewSize(400, 200)

// FallbackErrorDialog creates a simple dialog box without needing a parent window.
func FallbackErrorDialog(err error, app fyne.App) {
	closeChan := make(chan struct{})

	w := app.NewWindow("Error")
	w.Resize(defaultDialogSize)
	w.SetFixedSize(true)

	infDialog := dialog.NewError(err, w)
	infDialog.Resize(defaultDialogSize)

	infDialog.SetOnClosed(func() {
		w.Close()
		closeChan <- struct{}{}
	})

	go func() {
		time.Sleep(250 * time.Millisecond)
		infDialog.Show()
	}()

	w.Show()
	<-closeChan
}

// Gui is the application GUI.
type Gui struct {
	State *logic.AppState
	App   fyne.App
	Win   fyne.Window
}

func New() Gui {
	return Gui{
		App: app.New(),
	}
}

func (g Gui) ShowError(err error) {

}

func (g Gui) Run() {
	go func() {
		checkErr := func(err error) {
			if err == nil {
				return
			}

			FallbackErrorDialog(err, g.App)
			g.App.Quit()
		}

		cfgPath, err := config.GetFilePath()
		checkErr(err)

		// Run the setup if the config file doesn't exist.
		_, err = os.Stat(cfgPath)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				setupRes := setup.ShowSetup(g.App)

				if !setupRes.Completed {
					g.App.Quit()
					return
				}

				cfg := config.CreateDefault(lang.NewLocale(setupRes.LangCode))
				cfg.LangCode = setupRes.LangCode

				// Create config dir
				cfgDir, err := config.GetDirPath()
				checkErr(err)
				err = os.MkdirAll(cfgDir, os.ModePerm)
				checkErr(err)

				// Create config file
				newCfgFile, err := os.Create(cfgPath)
				checkErr(err)
				defer func() {
					_ = newCfgFile.Close()
				}()

				err = json.SerializeToJson(cfg, newCfgFile)
				checkErr(err)
			} else {
				checkErr(err)
			}
		}

		state, err := logic.Init()
		checkErr(err)

		g.State = state
		g.initMainWindow()
	}()

	g.App.Run()
}
