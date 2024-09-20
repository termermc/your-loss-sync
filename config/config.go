package config

import (
	"github.com/termermc/your-loss-sync/lang"
	"os"
	"path/filepath"
)

// DirName is the name of the configuration directory.
const DirName = "your-loss"

// FileName is the name of the configuration file.
const FileName = "config.json"

// GetDirPath returns the path to the application's configuration directory.
// If the user's config directory can't be determined, an error will be returned.
func GetDirPath() (string, error) {
	cfgDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(cfgDir, DirName), nil
}

// GetFilePath returns the path to the application's configuration file.
func GetFilePath() (string, error) {
	cfgDir, err := GetDirPath()
	if err != nil {
		return "", err
	}

	return filepath.Join(cfgDir, FileName), nil
}

// SyncConfig is the configuration for a sync.
type SyncConfig struct {
	// The sync's name.
	Name string

	// The source directory to sync from.
	SourceDir string

	// The destination directory to sync to.
	DestDir string

	// The output profile to use.
	Profile *OutputProfile

	// Whether to escape filenames.
	// Default: true
	EscapeFilenames bool

	// Whether to reencode files with the same format.
	// Default: false
	ReencodeSameFormat bool
}

// Config is the application configuration.
type Config struct {
	// The language code to use.
	LangCode string

	// All output profiles.
	Profiles []*OutputProfile

	// All sync configurations.
	Syncs []*SyncConfig
}

// CreateDefault creates a default configuration.
func CreateDefault(locale lang.Locale) *Config {
	res := Config{
		LangCode: lang.DefaultLangCode,
		Profiles: DefaultOutputProfiles,
		Syncs:    []*SyncConfig{},
	}

	for i := range res.Profiles {
		profile := res.Profiles[i]
		profile.Name = locale.TrTemplate(profile.Name)
	}

	return &res
}

// GetProfile returns the output profile with the specified name.
// If no profile with the specified name exists, nil will be returned.
func (c *Config) GetProfile(name string) *OutputProfile {
	for _, profile := range c.Profiles {
		if profile.Name == name {
			return profile
		}
	}

	return nil
}

// GetSync returns the sync with the specified name.
// If no sync with the specified name exists, nil will be returned.
func (c *Config) GetSync(name string) *SyncConfig {
	for _, sync := range c.Syncs {
		if sync.Name == name {
			return sync
		}
	}
	return nil
}

// GetSyncIndex returns the index of the sync with the specified name.
// If no sync with the specified name exists, -1 will be returned.
func (c *Config) GetSyncIndex(name string) int {
	for i, sync := range c.Syncs {
		if sync.Name == name {
			return i
		}
	}

	return -1
}
