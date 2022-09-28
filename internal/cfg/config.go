package cfg

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path"
)

type Config struct {
	ProjectDir string
}

func Save(filePath string, cfg *Config) error {
	marshalled, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	dir := path.Dir(filePath)
	if _, err := os.Stat(dir); errors.Is(err, os.ErrNotExist) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	if err := ioutil.WriteFile(filePath, marshalled, 0644); err != nil {
		return err
	}

	return nil
}
