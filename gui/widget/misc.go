package widget

import "fyne.io/fyne/v2/data/binding"

type simpleDataListener struct {
	f func()
}

func (s *simpleDataListener) DataChanged() {
	s.f()
}

// NewDataListener returns a new DataListener that calls the given function when
// the data changes.
func NewDataListener(f func()) binding.DataListener {
	return &simpleDataListener{f: f}
}
