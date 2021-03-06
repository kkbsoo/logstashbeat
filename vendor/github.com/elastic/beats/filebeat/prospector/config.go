package prospector

import (
	"fmt"
	"regexp"
	"time"

	cfg "github.com/elastic/beats/filebeat/config"
)

var (
	defaultConfig = prospectorConfig{
		IgnoreOlder:   0,
		ScanFrequency: 10 * time.Second,
		InputType:     cfg.DefaultInputType,
		CleanOlder:    0,
		CleanRemoved:  false,
	}
)

type prospectorConfig struct {
	ExcludeFiles  []*regexp.Regexp `config:"exclude_files"`
	IgnoreOlder   time.Duration    `config:"ignore_older"`
	Paths         []string         `config:"paths"`
	ScanFrequency time.Duration    `config:"scan_frequency"`
	InputType     string           `config:"input_type"`
	CleanOlder    time.Duration    `config:"clean_older" validate:"min=0"`
	CleanRemoved  bool             `config:"clean_removed"`
}

func (config *prospectorConfig) Validate() error {

	if config.InputType == cfg.LogInputType && len(config.Paths) == 0 {
		return fmt.Errorf("No paths were defined for prospector")
	}
	return nil
}
