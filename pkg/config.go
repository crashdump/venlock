package pkg

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Config[T Library] struct {
	Version   string                   `json:"version"`
	Catalogue map[string]LibrarySet[T] `json:"catalogue"`
}

type Generic interface{}

const version = "1.0"

func (c *Config[T]) Load(path string) (err error) {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	content, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(content, &c)
	if err != nil {
		return err
	}

	if c.Version != version {
		return fmt.Errorf("unknown config version %s", c.Version)
	}

	if c.Catalogue == nil {
		return fmt.Errorf("inventory field is missing")
	}

	return file.Close()
}

func (c *Config[T]) Save(path string) (err error) {
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	jsonBytes, err := json.Marshal(c)
	if err != nil {
		return err
	}

	_, err = file.Write(jsonBytes)
	if err != nil {
		return err
	}

	return file.Close()
}
