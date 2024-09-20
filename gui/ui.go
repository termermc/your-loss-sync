package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/termermc/your-loss-sync/gui/profilestab"
	"github.com/termermc/your-loss-sync/gui/settingstab"
	"github.com/termermc/your-loss-sync/gui/syncstab"
)

func (g Gui) initMainWindow() {
	w := g.App.NewWindow("Your Loss! Sync")
	g.Win = w

	w.Resize(fyne.NewSize(800, 500))
	w.SetOnClosed(func() {
		g.App.Quit()
	})

	syncsTab := syncstab.New(g.State, w)
	profilesTab := profilestab.New()
	settingsTab := settingstab.New()

	tabs := container.NewAppTabs(
		container.NewTabItem(g.State.Locale.Tr("shell.tab.syncs"), syncsTab.Widget),
		container.NewTabItem(g.State.Locale.Tr("shell.tab.profiles"), profilesTab.Widget),
		container.NewTabItem(g.State.Locale.Tr("shell.tab.settings"), settingsTab.Widget),
	)
	tabs.SetTabLocation(container.TabLocationTop)

	w.SetContent(tabs)

	w.Show()
}
