package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/termermc/your-loss-sync/gui/inprogresstab"
	"github.com/termermc/your-loss-sync/gui/profilestab"
	"github.com/termermc/your-loss-sync/gui/settingstab"
	"github.com/termermc/your-loss-sync/gui/syncstab"
	"strconv"
	"time"
)

func (g Gui) initMainWindow() {
	w := g.App.NewWindow("Your Loss! Sync")
	g.Win = w

	hasQuit := false
	w.Resize(fyne.NewSize(800, 500))
	w.SetOnClosed(func() {
		hasQuit = true
		g.App.Quit()
	})

	syncsTab := syncstab.New(g.State, w)
	profilesTab := profilestab.New(g.State, w)
	settingsTab := settingstab.New()
	inProgressTab := inprogresstab.New(g.State, w)

	inProgressTabItem := container.NewTabItem("", inProgressTab.Widget)

	tabs := container.NewAppTabs(
		container.NewTabItem(g.State.Locale.Tr("shell.tab.syncs"), syncsTab.Widget),
		container.NewTabItem(g.State.Locale.Tr("shell.tab.profiles"), profilesTab.Widget),
		container.NewTabItem(g.State.Locale.Tr("shell.tab.settings"), settingsTab.Widget),
		inProgressTabItem,
	)
	tabs.SetTabLocation(container.TabLocationTop)
	tabs.OnSelected = func(_ *container.TabItem) {
		syncsTab.ResetForm()
		profilesTab.ResetForm()
		inProgressTab.ResetForm()
	}

	w.SetContent(tabs)

	// Update in-progress tab periodically
	go func() {
		for !hasQuit {
			inProgressTabItem.Text = g.State.Locale.Tr(
				"shell.tab.in-progress",
				strconv.FormatInt(g.State.Progress.Completed.Load(), 10),
				strconv.FormatInt(g.State.Progress.Total.Load(), 10),
			)
			tabs.Refresh()
			time.Sleep(time.Second)
		}
	}()

	w.Show()
}
