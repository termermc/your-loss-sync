package config

// OutputFormatType is an output format type.
type OutputFormatType string

const (
	// OutputFormatTypeLossy is a lossy output format type.
	OutputFormatTypeLossy OutputFormatType = "lossy"

	// OutputFormatTypeLossless is a lossless output format type.
	OutputFormatTypeLossless OutputFormatType = "lossless"
)

// OutputFormat is an output format.
type OutputFormat struct {
	// The format type.
	Type OutputFormatType

	// The format name.
	Name string

	// The format's file extension.
	Extension string

	// The name of the FFmpeg encoder used to convert to the format.
	FfmpegEncoder string

	// Whether the container used by the format supports metadata.
	SupportsMetadata bool

	// Whether the container used by the format supports artwork.
	SupportsArtwork bool

	// The suggested bitrate for the format.
	// Does not apply to lossless formats.
	// Should be a multiple of 1000.
	SuggestedBitrate uint
}

// GetId returns the ID of the format.
// If the format is not in the supported formats map, the ID will be 0 and false will be returned.
func (f OutputFormat) GetId() (int, bool) {
	for id, format := range SupportedOutputFormats {
		if format == f {
			return id, true
		}
	}

	return 0, false
}

// SupportedOutputFormats is a mapping IDs to supported output formats.
// A map is used instead of a slice to allow for reordering or removal without changing IDs.
var SupportedOutputFormats = map[int]OutputFormat{
	0: {
		Type:             OutputFormatTypeLossy,
		Name:             "mp3",
		Extension:        "mp3",
		FfmpegEncoder:    "libmp3lame",
		SupportsMetadata: true,
		SupportsArtwork:  true,
		SuggestedBitrate: 320000,
	},
	1: {
		Type:             OutputFormatTypeLossless,
		Name:             "flac",
		Extension:        "flac",
		FfmpegEncoder:    "flac",
		SupportsMetadata: true,
		SupportsArtwork:  true,
		SuggestedBitrate: 0,
	},
	2: {
		Type:             OutputFormatTypeLossless,
		Name:             "wav",
		Extension:        "wav",
		FfmpegEncoder:    "pcm_s16le",
		SupportsMetadata: true,
		SupportsArtwork:  false,
		SuggestedBitrate: 0,
	},
	3: {
		Type:             OutputFormatTypeLossy,
		Name:             "opus",
		Extension:        "opus",
		FfmpegEncoder:    "libopus",
		SupportsMetadata: true,
		SupportsArtwork:  false,
		SuggestedBitrate: 120000,
	},
	4: {
		Type:             OutputFormatTypeLossy,
		Name:             "aac",
		Extension:        "m4a",
		FfmpegEncoder:    "aac",
		SupportsMetadata: true,
		SupportsArtwork:  true,
		SuggestedBitrate: 224000,
	},
	5: {
		Type:             OutputFormatTypeLossless,
		Name:             "alac",
		Extension:        "m4a",
		FfmpegEncoder:    "alac",
		SupportsMetadata: true,
		SupportsArtwork:  true,
		SuggestedBitrate: 0,
	},
	6: {
		Type:             OutputFormatTypeLossless,
		Name:             "aiff",
		Extension:        "aif",
		FfmpegEncoder:    "pcm_s16le",
		SupportsMetadata: true, // Limit, seems to only support title and comment
		SupportsArtwork:  false,
		SuggestedBitrate: 0,
	},
}

// OutputProfile is an output encoding profile.
type OutputProfile struct {
	// The profile name.
	// Names are expected to be unique and treated as primary keys.
	Name string

	// The output format.
	OutputFormat OutputFormat

	// The bitrate to use for the output format.
	// Only applies to lossy formats.
	// Should be a multiple of 1000.
	Bitrate uint
}

// DefaultOutputProfiles is a list of default output profiles.
// Names and descriptions are not literal; instead, they are filled in with translations.
var DefaultOutputProfiles = []*OutputProfile{
	{
		Name:         "{{profile.default.hq-mp3.name}}",
		OutputFormat: SupportedOutputFormats[0],
		Bitrate:      320000,
	},
	{
		Name:         "{{profile.default.flac.name}}",
		OutputFormat: SupportedOutputFormats[1],
		Bitrate:      0,
	},
	{
		Name:         "{{profile.default.wav.name}}",
		OutputFormat: SupportedOutputFormats[2],
		Bitrate:      0,
	},
	{
		Name:         "{{profile.default.hq-aac.name}}",
		OutputFormat: SupportedOutputFormats[4],
		Bitrate:      224000,
	},
	{
		Name:         "{{profile.default.alac.name}}",
		OutputFormat: SupportedOutputFormats[5],
		Bitrate:      0,
	},
}
