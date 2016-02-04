package biz

import (
	"encoding/json"

	"github.com/BurntSushi/toml"
)

// Config
type Config struct {
	DebugLevel int `toml:"debug_level" json:"debugLevel"`

	ExpectedArray []ConfigAPI `toml:"expected" json:"expected"`
	ActualArray   []ConfigAPI `toml:"actual" json:"actual"`
}

type ConfigAPI struct {
	Addr string `toml:"addr" json:"addr"`
	DB   string `toml:"db" json:"db"`
	SQL  string `toml:"sql" json:"sql"`
}

func (e Config) String() string {
	s, _ := json.Marshal(e)

	return string(s)
}

func (e ConfigAPI) String() string {
	s, _ := json.Marshal(e)

	return string(s)
}

// Decode decodes a TOML-format string
func NewConfig(tomlStr string) (*Config, error) {
	var config Config
	if _, err := toml.Decode(tomlStr, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
