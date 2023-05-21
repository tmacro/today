package config

import (
	"os"

	"github.com/tmacro/today/pkg/utils"
	"gopkg.in/yaml.v3"
)

type TodayConfig struct {
	DateFormat string        `yaml:"date_format"`
	Notes      NotesConfig   `yaml:"notes"`
	Scratch    ScratchConfig `yaml:"scratch"`
}

type NotesConfig struct {
	Extension string `yaml:"extension"`
	Directory string `yaml:"directory"`
	Editor    string `yaml:"editor"`
	Viewer    string `yaml:"viewer"`
	Search    string `yaml:"search"`
}

type ScratchConfig struct {
	Directory string `yaml:"directory"`
	Viewer    string `yaml:"viewer"`
	Search    string `yaml:"search"`
	Find      string `yaml:"find"`
}

const DefaultConfigPath = "~/.config/today/config.yaml"

var DefaultTodayConfig = TodayConfig{
	DateFormat: "2006.01.02",
	Notes: NotesConfig{
		Extension: ".md",
		Directory: "~/notes",
		Editor:    "nano -Y markdown -f /usr/share/nano/markdown.nanorc",
		Viewer:    "nano -Y markdown -f /usr/share/nano/markdown.nanorc --view",
		Search:    "grep -i -n -H -T -r --color=auto",
	},
	Scratch: ScratchConfig{
		Directory: "~/scratch",
		Viewer:    "ls -altr --color=auto {{ .Directory }}",
		Search:    "grep -i -n -H -T -r --color=auto {{ .Expression }} {{ .Directory }}",
		Find:      "find {{ .Directory }} -type f -iname \"*{{ .Expression }}*\"",
	},
}

// Merge a user given config with the default config
func MergeWithDefault(userConfig *TodayConfig) {
	if userConfig.DateFormat == "" {
		userConfig.DateFormat = DefaultTodayConfig.DateFormat
	}

	if userConfig.Notes.Extension == "" {
		userConfig.Notes.Extension = DefaultTodayConfig.Notes.Extension
	}

	if userConfig.Notes.Directory == "" {
		userConfig.Notes.Directory = DefaultTodayConfig.Notes.Directory
	}

	if userConfig.Notes.Editor == "" {
		userConfig.Notes.Editor = DefaultTodayConfig.Notes.Editor
	}

	if userConfig.Notes.Viewer == "" {
		userConfig.Notes.Viewer = DefaultTodayConfig.Notes.Viewer
	}

	if userConfig.Notes.Search == "" {
		userConfig.Notes.Search = DefaultTodayConfig.Notes.Search
	}

	if userConfig.Scratch.Directory == "" {
		userConfig.Scratch.Directory = DefaultTodayConfig.Scratch.Directory
	}

	if userConfig.Scratch.Viewer == "" {
		userConfig.Scratch.Viewer = DefaultTodayConfig.Scratch.Viewer
	}

	if userConfig.Scratch.Search == "" {
		userConfig.Scratch.Search = DefaultTodayConfig.Scratch.Search
	}

	if userConfig.Scratch.Find == "" {
		userConfig.Scratch.Find = DefaultTodayConfig.Scratch.Find
	}
}

// Read the config file from a given path
func ReadConfig(opts ...string) (*TodayConfig, error) {
	var config TodayConfig
	path := DefaultConfigPath

	if len(opts) > 0 {
		path = opts[0]
	}

	resolved, err := utils.ResolvePath(path)
	if err != nil {
		return nil, err
	}

	// Check if the file exists
	_, err = os.Stat(resolved)
	if err != nil {
		if os.IsNotExist(err) {
			// If the file doesn't exist, return the default config
			return &DefaultTodayConfig, nil
		}
		return nil, err
	}

	data, err := os.ReadFile(resolved)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	MergeWithDefault(&config)
	return &config, nil
}
