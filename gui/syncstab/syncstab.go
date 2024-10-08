package syncstab

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/termermc/your-loss-sync/config"
	ylwidget "github.com/termermc/your-loss-sync/gui/widget"
	"github.com/termermc/your-loss-sync/logic"
	"os"
)

type SyncsTab struct {
	Widget    fyne.CanvasObject
	setupForm func()
}

// New creates a new SyncsTab
func New(s *logic.AppState, parent fyne.Window) SyncsTab {
	var deleteSyncById func(id int)

	list := ylwidget.NewTextList(
		func() int {
			// Length

			return len(s.Config.Syncs)
		},
		func(id widget.ListItemID) ylwidget.TextListItemData {
			// Update item

			return ylwidget.TextListItemData{
				Label:     s.Config.Syncs[id].Name,
				CanDelete: true,
				OnDelete: func() {
					deleteSyncById(id)
				},
			}
		},
	)

	deleteSyncById = func(id int) {
		sync := s.Config.Syncs[id]

		title := s.Locale.Tr("tab.syncs.delete-confirm.title")
		desc := s.Locale.Tr("tab.syncs.delete-confirm.description", sync.Name)
		dialog.ShowConfirm(title, desc, func(b bool) {
			if !b {
				return
			}

			s.Config.Syncs = append(s.Config.Syncs[:id], s.Config.Syncs[id+1:]...)

			if err := s.Save(); err != nil {
				dialog.ShowError(err, parent)
				return
			}

			list.Widget.UnselectAll()
			list.Widget.Refresh()
		}, parent)
	}

	// The target sync config that is being edited.
	// If nil, a new sync is being created.
	var targetSync *config.SyncConfig

	form := widget.NewForm()

	nameEntry := widget.NewEntry()
	srcDirPicker := ylwidget.NewFilePicker(parent, s.Locale)
	_ = srcDirPicker.IsDirectoryPicker.Set(true)
	destDirPicker := ylwidget.NewFilePicker(parent, s.Locale)
	_ = destDirPicker.IsDirectoryPicker.Set(true)
	profileSelector := widget.NewSelect([]string{}, func(_ string) {})
	escapeFilenamesCheck := widget.NewCheck(s.Locale.Tr("tab.syncs.form.escape-filenames"), func(_ bool) {})
	reencodeSameFormatCheck := widget.NewCheck(s.Locale.Tr("tab.syncs.form.reencode-same-format"), func(_ bool) {})

	var onSave func()

	saveBtn := widget.NewButton("", func() {
		onSave()
	})

	errMsg := widget.NewLabel("")

	setupForm := func() {
		errMsg.SetText("")

		profileNames := make([]string, 0, len(s.Config.Profiles))
		for _, profile := range s.Config.Profiles {
			profileNames = append(profileNames, profile.Name)
		}
		profileSelector.SetOptions(profileNames)

		if targetSync == nil {
			nameEntry.SetText("")
			srcDirPicker.Clear()
			destDirPicker.Clear()
			if len(s.Config.Profiles) > 0 {
				profileSelector.SetSelected(s.Config.Profiles[0].Name)
			} else {
				profileSelector.SetSelected("")
			}
			escapeFilenamesCheck.SetChecked(true)
			reencodeSameFormatCheck.SetChecked(false)

			saveBtn.SetText(s.Locale.Tr("general.create"))
		} else {
			nameEntry.SetText(targetSync.Name)
			_ = srcDirPicker.Path.Set(targetSync.SourceDir)
			_ = destDirPicker.Path.Set(targetSync.DestDir)
			profileSelector.SetSelected(targetSync.Profile.Name)
			escapeFilenamesCheck.SetChecked(targetSync.EscapeFilenames)
			reencodeSameFormatCheck.SetChecked(targetSync.ReencodeSameFormat)

			saveBtn.SetText(s.Locale.Tr("general.save"))
		}
	}

	setupForm()

	form.Append(s.Locale.Tr("tab.syncs.form.name"), nameEntry)
	form.Append(s.Locale.Tr("tab.syncs.form.source-dir"), srcDirPicker.Widget)
	form.Append(s.Locale.Tr("tab.syncs.form.dest-dir"), destDirPicker.Widget)
	form.Append(s.Locale.Tr("tab.syncs.form.profile"), profileSelector)
	form.Append("", escapeFilenamesCheck)
	form.Append("", reencodeSameFormatCheck)
	form.Append("", layout.NewSpacer())
	form.Append("", saveBtn)
	form.Append("", errMsg)

	createBtn := widget.NewButton(s.Locale.Tr("tab.syncs.create"), func() {
		list.Widget.UnselectAll()
		targetSync = nil
		setupForm()
	})

	listScroll := container.NewScroll(list.Widget)
	listScroll.SetMinSize(fyne.NewSize(200, 500))

	formScroll := container.NewScroll(
		container.New(
			layout.NewFormLayout(),
			form,
		),
	)
	formScroll.SetMinSize(fyne.NewSize(600, 500))

	cont := container.NewHBox(
		container.NewVBox(
			listScroll,
			createBtn,
		),
		formScroll,
	)

	list.Widget.OnSelected = func(id widget.ListItemID) {
		targetSync = s.Config.GetSync(s.Config.Syncs[id].Name)
		setupForm()
	}

	onSave = func() {
		errMsg.SetText("")

		if nameEntry.Text == "" {
			errMsg.SetText(s.Locale.Tr("tab.syncs.form.error.missing-name"))
			return
		}

		if !srcDirPicker.HasPath() {
			errMsg.SetText(s.Locale.Tr("tab.syncs.form.error.missing-source-dir"))
			return
		}

		if !destDirPicker.HasPath() {
			errMsg.SetText(s.Locale.Tr("tab.syncs.form.error.missing-dest-dir"))
			return
		}

		srcDirPath, err := srcDirPicker.Path.Get()
		if err != nil {
			dialog.ShowError(err, parent)
			return
		}
		destDirPath, err := destDirPicker.Path.Get()
		if err != nil {
			dialog.ShowError(err, parent)
			return
		}
		if srcDirPath == destDirPath {
			errMsg.SetText(s.Locale.Tr("tab.syncs.form.error.source-dest-dirs-same"))
			return
		}

		if targetSync == nil {
			if s.Config.GetSync(nameEntry.Text) != nil {
				errMsg.SetText(s.Locale.Tr("tab.syncs.form.error.name-exists"))
				return
			}
		} else {
			if nameEntry.Text != targetSync.Name && s.Config.GetSync(nameEntry.Text) != nil {
				errMsg.SetText(s.Locale.Tr("tab.syncs.form.error.name-exists"))
				return
			}
		}

		// Check if paths exist and are indeed directories
		srcStat, err := os.Stat(srcDirPath)
		if err != nil {
			if os.IsNotExist(err) {
				errMsg.SetText(s.Locale.Tr("tab.syncs.form.error.source-dir-nonexistent"))
				return
			}

			dialog.ShowError(err, parent)
			return
		}
		if !srcStat.IsDir() {
			errMsg.SetText(s.Locale.Tr("tab.syncs.form.error.source-dir-not-dir"))
			return
		}

		destStat, err := os.Stat(destDirPath)
		if err != nil {
			if os.IsNotExist(err) {
				errMsg.SetText(s.Locale.Tr("tab.syncs.form.error.dest-dir-nonexistent"))
				return
			}

			dialog.ShowError(err, parent)
			return
		}
		if !destStat.IsDir() {
			errMsg.SetText(s.Locale.Tr("tab.syncs.form.error.dest-dir-not-dir"))
			return
		}

		// Config looks good, save it
		if targetSync == nil {
			newSync := &config.SyncConfig{
				Name:               nameEntry.Text,
				SourceDir:          srcDirPath,
				DestDir:            destDirPath,
				Profile:            s.Config.GetProfile(profileSelector.Selected),
				EscapeFilenames:    escapeFilenamesCheck.Checked,
				ReencodeSameFormat: reencodeSameFormatCheck.Checked,
			}

			s.Config.Syncs = append(s.Config.Syncs, newSync)
			targetSync = newSync
		} else {
			targetSync.Name = nameEntry.Text
			targetSync.SourceDir = srcDirPath
			targetSync.DestDir = destDirPath
			targetSync.Profile = s.Config.GetProfile(profileSelector.Selected)
			targetSync.EscapeFilenames = escapeFilenamesCheck.Checked
			targetSync.ReencodeSameFormat = reencodeSameFormatCheck.Checked
		}

		err = s.Save()
		if err != nil {
			dialog.ShowError(err, parent)
			return
		}

		list.Widget.Refresh()
		list.Widget.Select(s.Config.GetSyncIndex(nameEntry.Text))
	}

	return SyncsTab{
		Widget:    cont,
		setupForm: setupForm,
	}
}

// ResetForm resets the form to its default state.
func (s *SyncsTab) ResetForm() {
	s.setupForm()
}
