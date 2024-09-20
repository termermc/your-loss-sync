package json

type V1OutputProfile struct {
	Name           string `json:"name"`
	OutputFormatId int    `json:"outputFormatId"`
	Bitrate        uint   `json:"bitrate"`
}

// V1Sync is the JSON format version 1 representation of a sync.
type V1Sync struct {
	Name               string `json:"name"`
	SourceDir          string `json:"sourceDir"`
	DestDir            string `json:"destDir"`
	ProfileName        string `json:"profileName"`
	EscapeFilenames    bool   `json:"escapeFilenames"`
	ReencodeSameFormat bool   `json:"reencodeSameFormat"`
}

// V1 is the JSON format version 1 representation of the application configuration.
type V1 struct {
	Version  int               `json:"version"` // Should be Version1
	LangCode string            `json:"langCode"`
	Syncs    []V1Sync          `json:"syncs"`
	Profiles []V1OutputProfile `json:"profiles"`
}
