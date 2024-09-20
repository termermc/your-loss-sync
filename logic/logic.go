package logic

import (
	"github.com/termermc/your-loss-sync/config"
	"github.com/termermc/your-loss-sync/config/json"
	"github.com/termermc/your-loss-sync/lang"
	"os"
)

// AppState is the state of the application.
type AppState struct {
	Config     *config.Config
	ConfigDir  string
	ConfigFile string
	Locale     lang.Locale
}

// Save saves the application state to disk, including configuration.
func (s *AppState) Save() error {
	oldCfgPath := s.ConfigFile + ".bak"
	err := os.Rename(s.ConfigFile, oldCfgPath)
	if err != nil {
		return err
	}

	newCfgPath := s.ConfigFile
	cfgFile, err := os.Create(newCfgPath)
	if err != nil {
		return err
	}
	defer func() {
		_ = cfgFile.Close()
	}()

	err = json.SerializeToJson(s.Config, cfgFile)
	if err != nil {
		return err
	}

	return nil
}

// Init initializes the application and returns the state.
func Init() (*AppState, error) {
	cfgDir, err := config.GetDirPath()
	if err != nil {
		return nil, err
	}
	cfgPath, err := config.GetFilePath()
	if err != nil {
		return nil, err
	}
	cfgFile, err := os.Open(cfgPath)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = cfgFile.Close()
	}()

	cfg, err := json.DeserializeFromJson(cfgFile)
	if err != nil {
		return nil, err
	}

	return &AppState{
		Config:     cfg,
		ConfigDir:  cfgDir,
		ConfigFile: cfgPath,
		Locale:     lang.NewLocale(cfg.LangCode),
	}, nil
}
