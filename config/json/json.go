package json

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/termermc/your-loss-sync/config"
	"io"
)

const (
	// Version1 is config JSON version 1.
	Version1 = 1
)

// ConfigBase is the struct that all JSON configs must implement.
// It is used to check the version needed for deserialization.
type ConfigBase struct {
	// The config JSON version.
	Version int `json:"version"`
}

// ErrUnsupportedVersion is returned when the JSON config version is not supported.
var ErrUnsupportedVersion = errors.New("{{config.error.unsupported-version}}")

// ErrUnknownProfile is returned when a config contains a reference to an unknown profile.
var ErrUnknownProfile = errors.New("{{config.error.unknown-profile}}")

// ErrUnknownFormat is returned when a config contains a reference to an unknown format.
var ErrUnknownFormat = errors.New("{{config.error.unknown-format}}")

// DeserializeFromJson deserializes a JSON configuration from the given reader.
// If the config version is not supported, ErrUnsupportedVersion is returned.
// If the config contains an unknown profile, ErrUnknownProfile is returned.
// If the config contains an unknown format, ErrUnknownFormat is returned.
// The returned config will be valid if no error is returned.
func DeserializeFromJson(reader io.Reader) (*config.Config, error) {
	// Buffer entire reader to memory so that we can read it again if needed.
	buffer := make([]byte, 1024)
	_, err := reader.Read(buffer)
	if err != nil {
		return nil, err
	}

	bufReader := bytes.NewReader(buffer)

	var container ConfigBase
	err = json.NewDecoder(bufReader).Decode(&container)
	if err != nil {
		// The buffer wasn't a valid config at all
		return nil, err
	}

	bufReader.Reset(buffer)

	switch container.Version {
	case Version1:
		var v1 V1
		err = json.NewDecoder(bufReader).Decode(&v1)
		if err != nil {
			return nil, err
		}

		resProfiles := make([]*config.OutputProfile, len(v1.Profiles))
		for i, v1Profile := range v1.Profiles {
			format, ok := config.SupportedOutputFormats[v1Profile.OutputFormatId]
			if !ok {
				return nil, ErrUnknownFormat
			}

			resProfiles[i] = &config.OutputProfile{
				Name:         v1Profile.Name,
				OutputFormat: format,
				Bitrate:      v1Profile.Bitrate,
			}
		}

		resSyncs := make([]*config.SyncConfig, len(v1.Syncs))
		for i, v1Sync := range v1.Syncs {
			var profile *config.OutputProfile
			for _, p := range resProfiles {
				if p.Name == v1Sync.ProfileName {
					profile = p
					break
				}
			}
			if profile == nil {
				return nil, ErrUnknownProfile
			}

			resSyncs[i] = &config.SyncConfig{
				Name:               v1Sync.Name,
				SourceDir:          v1Sync.SourceDir,
				DestDir:            v1Sync.DestDir,
				Profile:            profile,
				EscapeFilenames:    v1Sync.EscapeFilenames,
				ReencodeSameFormat: v1Sync.ReencodeSameFormat,
			}
		}

		return &config.Config{
			LangCode: v1.LangCode,
			Profiles: resProfiles,
			Syncs:    resSyncs,
		}, nil

	default:
		return nil, ErrUnsupportedVersion

	}
}

// SerializeToJson serializes a config to the given writer.
func SerializeToJson(config *config.Config, writer io.Writer) error {
	res := V1{
		Version:  Version1,
		LangCode: config.LangCode,
		Syncs:    make([]V1Sync, len(config.Syncs)),
		Profiles: make([]V1OutputProfile, len(config.Profiles)),
	}

	for i, sync := range config.Syncs {
		res.Syncs[i] = V1Sync{
			Name:               sync.Name,
			SourceDir:          sync.SourceDir,
			DestDir:            sync.DestDir,
			ProfileName:        sync.Profile.Name,
			EscapeFilenames:    sync.EscapeFilenames,
			ReencodeSameFormat: sync.ReencodeSameFormat,
		}
	}

	for i, profile := range config.Profiles {
		formatId, ok := profile.OutputFormat.GetId()
		if !ok {
			return ErrUnknownFormat
		}

		res.Profiles[i] = V1OutputProfile{
			Name:           profile.Name,
			OutputFormatId: formatId,
			Bitrate:        profile.Bitrate,
		}
	}

	return json.NewEncoder(writer).Encode(res)
}
