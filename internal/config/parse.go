package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

func (c *Config) parseFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return errors.Wrap(err, "[configFile.Open] failed")
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	switch ext := strings.ToLower(filepath.Ext(path)); ext {
	case ".yaml", ".yml":
		err = c.parseYAML(file)
	default:
		return fmt.Errorf("file format '%s' doesn't supported by the parser", ext)
	}
	if err != nil {
		return fmt.Errorf("config file parsing errors: %s", err.Error())
	}
	return nil
}

func (c *Config) parseYAML(r io.Reader) error {
	return yaml.NewDecoder(r).Decode(c)
}

func (c *Config) parseEnvVars() error {
	return envconfig.Process("", c)
}

func Parse(configPath string) (*Config, error) {
	c := new(Config)
	if err := c.parseFile(configPath); err != nil {
		return nil, errors.Wrap(err, "parse config file failed")
	}
	if err := c.parseEnvVars(); err != nil {
		return nil, errors.Wrap(err, "parse env vars failed")
	}
	return c, nil
}
