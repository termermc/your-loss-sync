package inprogresstab

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/termermc/your-loss-sync/logic"
	"math"
	"strconv"
	"sync"
	"time"
)

type InProgressTab struct {
	Widget    fyne.CanvasObject
	setupForm func()
}

// New creates a new InProgressTab
func New(s *logic.AppState, parent fyne.Window) InProgressTab {
	// Atomic logging
	logOut := make(chan string, 1000)
	logMultiline := widget.NewMultiLineEntry()
	logMultilineLock := sync.Mutex{}
	logMultiline.Disable()
	multilineScroll := container.NewScroll(logMultiline)
	multilineScroll.SetMinSize(fyne.NewSize(600, 400))

	go func() {
		for msg := range logOut {
			logMultilineLock.Lock()
			logMultiline.Append(msg + "\n")
			logMultilineLock.Unlock()
		}
	}()

	statusLabel := widget.NewLabel("")
	progressBar := widget.NewProgressBar()

	syncsSelector := widget.NewSelect([]string{}, func(_ string) {})
	selectorScroll := container.NewScroll(syncsSelector)
	selectorScroll.SetMinSize(fyne.NewSize(300, 40))
	actionBtn := widget.NewButton("", func() {
		if s.Progress.Sync.Load() == nil {
			syncConf := s.Config.GetSync(syncsSelector.Selected)

			logMultilineLock.Lock()
			logMultiline.SetText("")
			logMultilineLock.Unlock()
			go logic.StartSync(s, syncConf, logOut)
		} else {
			s.Progress.Sync.Store(nil)
		}
	})

	// Periodically update status
	go func() {
		for {
			time.Sleep(100 * time.Millisecond)

			statusLabel.SetText(s.Locale.Tr(
				"tab.progress.status-label",
				strconv.Itoa(int(s.Progress.Completed.Load())),
				strconv.Itoa(int(s.Progress.Total.Load())),
				strconv.Itoa(int(s.Progress.Failed.Load()))),
			)

			percent := float64(s.Progress.Completed.Load()+s.Progress.Failed.Load()) / float64(s.Progress.Total.Load())
			if math.IsNaN(percent) {
				percent = 0
			}

			progressBar.SetValue(percent)

			if s.Progress.Sync.Load() == nil {
				syncsSelector.Enable()
				actionBtn.SetText(s.Locale.Tr("tab.progress.start"))
			} else {
				syncsSelector.Disable()
				actionBtn.SetText(s.Locale.Tr("tab.progress.cancel"))
			}
		}
	}()

	cont := container.NewVBox(
		container.NewHBox(
			selectorScroll,
			actionBtn,
		),
		statusLabel,
		progressBar,
		multilineScroll,
	)

	setupForm := func() {
		syncNames := make([]string, 0, len(s.Config.Syncs))
		for _, syncConf := range s.Config.Syncs {
			syncNames = append(syncNames, syncConf.Name)
		}
		syncsSelector.SetOptions(syncNames)

		if s.Progress.Sync.Load() == nil {
			if len(s.Config.Syncs) > 0 {
				syncsSelector.Enable()
				syncsSelector.SetSelected(s.Config.Syncs[0].Name)
			} else {
				syncsSelector.Disable()
			}

			actionBtn.SetText(s.Locale.Tr("tab.progress.start"))
		} else {
			syncsSelector.SetSelected(s.Progress.Sync.Load().Name)

			actionBtn.SetText(s.Locale.Tr("tab.progress.cancel"))
		}
	}

	setupForm()

	return InProgressTab{
		Widget:    cont,
		setupForm: setupForm,
	}
}

// ResetForm resets the form to its default state.
func (s *InProgressTab) ResetForm() {
	s.setupForm()
}
