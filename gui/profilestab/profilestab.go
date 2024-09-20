package profilestab

import (
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/termermc/your-loss-sync/config"
	ylwidget "github.com/termermc/your-loss-sync/gui/widget"
	"github.com/termermc/your-loss-sync/logic"
	"strconv"
)

type ProfilesTab struct {
	Widget fyne.CanvasObject

	setupForm func()
}

// New creates a new ProfilesTab
func New(s *logic.AppState, parent fyne.Window) ProfilesTab {
	var deleteProfById func(id int)

	list := ylwidget.NewTextList(
		func() int {
			// Length

			return len(s.Config.Profiles)
		},
		func(id widget.ListItemID) ylwidget.TextListItemData {
			// Update item

			return ylwidget.TextListItemData{
				Label:     s.Config.Profiles[id].Name,
				CanDelete: true,
				OnDelete: func() {
					inUse := false
					for _, sync := range s.Config.Syncs {
						if sync.Profile.Name == s.Config.Profiles[id].Name {
							inUse = true
							break
						}
					}

					if inUse {
						dialog.ShowError(errors.New(s.Locale.Tr("tab.profiles.error.in-use")), parent)
						return
					}

					deleteProfById(id)
				},
			}
		},
	)

	deleteProfById = func(id int) {
		prof := s.Config.Profiles[id]

		title := s.Locale.Tr("tab.profiles.delete-confirm.title")
		desc := s.Locale.Tr("tab.profiles.delete-confirm.description", prof.Name)
		dialog.ShowConfirm(title, desc, func(b bool) {
			if !b {
				return
			}

			s.Config.Profiles = append(s.Config.Profiles[:id], s.Config.Profiles[id+1:]...)

			if err := s.Save(); err != nil {
				dialog.ShowError(err, parent)
				return
			}

			list.Widget.UnselectAll()
			list.Widget.Refresh()
		}, parent)
	}

	// The target profile config that is being edited.
	// If nil, a new profile is being created.
	var targetProf *config.OutputProfile

	form := widget.NewForm()

	nameEntry := widget.NewEntry()
	isLosslessCheck := widget.NewCheck(s.Locale.Tr("tab.profiles.form.is-lossless"), func(_ bool) {})
	isLosslessCheck.Disable()
	supportsMetaCheck := widget.NewCheck(s.Locale.Tr("tab.profiles.form.supports-metadata"), func(_ bool) {})
	supportsMetaCheck.Disable()
	supportsArtworkCheck := widget.NewCheck(s.Locale.Tr("tab.profiles.form.supports-artwork"), func(_ bool) {})
	supportsArtworkCheck.Disable()
	bitrateEntry := widget.NewEntry()
	formatNames := make([]string, 0, len(config.SupportedOutputFormats))
	for _, format := range config.SupportedOutputFormats {
		formatNames = append(formatNames, format.Name)
	}
	var onFormatSelect func(name string)
	formatSelector := widget.NewSelect(formatNames, func(name string) {
		onFormatSelect(name)
	})
	onFormatSelect = func(name string) {
		format := config.GetOutputFormat(name)
		if format.IsLossless {
			bitrateEntry.Disable()
			bitrateEntry.SetText("")
		} else {
			bitrateEntry.SetText(strconv.Itoa(int(format.SuggestedBitrate)))
			bitrateEntry.Enable()
		}
		isLosslessCheck.SetChecked(format.IsLossless)
		supportsMetaCheck.SetChecked(format.SupportsMetadata)
		supportsArtworkCheck.SetChecked(format.SupportsArtwork)
	}

	var onSave func()

	saveBtn := widget.NewButton("", func() {
		onSave()
	})

	errMsg := widget.NewLabel("")

	setupForm := func() {
		errMsg.SetText("")

		if targetProf == nil {
			nameEntry.SetText("")
			format := config.SupportedOutputFormats[0]
			formatSelector.SetSelected(format.Name)
			onFormatSelect(format.Name)

			saveBtn.SetText(s.Locale.Tr("general.create"))
		} else {
			nameEntry.SetText(targetProf.Name)
			formatSelector.SetSelected(targetProf.OutputFormat.Name)
			onFormatSelect(targetProf.OutputFormat.Name)
			if !targetProf.OutputFormat.IsLossless {
				bitrateEntry.SetText(strconv.Itoa(int(targetProf.Bitrate)))
			}

			saveBtn.SetText(s.Locale.Tr("general.save"))
		}
	}

	setupForm()

	form.Append(s.Locale.Tr("tab.profiles.form.name"), nameEntry)
	form.Append(s.Locale.Tr("tab.profiles.form.format"), formatSelector)
	form.Append("", isLosslessCheck)
	form.Append("", supportsMetaCheck)
	form.Append("", supportsArtworkCheck)
	form.Append(s.Locale.Tr("tab.profiles.form.bitrate"), bitrateEntry)
	form.Append("", layout.NewSpacer())
	form.Append("", saveBtn)
	form.Append("", errMsg)

	createBtn := widget.NewButton(s.Locale.Tr("tab.profiles.create"), func() {
		list.Widget.UnselectAll()
		targetProf = nil
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
		targetProf = s.Config.GetProfile(s.Config.Profiles[id].Name)
		setupForm()
	}

	onSave = func() {
		errMsg.SetText("")

		if nameEntry.Text == "" {
			errMsg.SetText(s.Locale.Tr("tab.profiles.form.error.missing-name"))
			return
		}

		if targetProf == nil && s.Config.GetProfile(nameEntry.Text) != nil {
			errMsg.SetText(s.Locale.Tr("tab.profiles.form.error.name-exists"))
			return
		}

		// Formats are hardcoded, so we can assume it exists
		format := config.GetOutputFormat(formatSelector.Selected)

		bitrate := 0
		var err error
		if !format.IsLossless {
			bitrate, err = strconv.Atoi(bitrateEntry.Text)
			if err != nil || bitrate < 1 {
				errMsg.SetText(s.Locale.Tr("tab.profiles.form.error.invalid-bitrate"))
			}
		}

		// Config looks good, save it
		if targetProf == nil {
			newProf := &config.OutputProfile{
				Name:         nameEntry.Text,
				OutputFormat: *format,
				Bitrate:      uint(bitrate),
			}

			s.Config.Profiles = append(s.Config.Profiles, newProf)
			targetProf = newProf
		} else {
			targetProf.Name = nameEntry.Text
			targetProf.OutputFormat = *format
			targetProf.Bitrate = uint(bitrate)
		}

		err = s.Save()
		if err != nil {
			dialog.ShowError(err, parent)
			return
		}

		list.Widget.Refresh()
		list.Widget.Select(s.Config.GetProfileIndex(nameEntry.Text))
		parent.Content().Refresh()
	}

	return ProfilesTab{
		Widget:    cont,
		setupForm: setupForm,
	}
}

// ResetForm resets the form to its default state.
func (s *ProfilesTab) ResetForm() {
	s.setupForm()
}
