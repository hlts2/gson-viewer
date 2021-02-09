package gsonviewer

import (
	"io"
	"os"

	"github.com/hlts2/gson"
)

// Load load json from a given path and return gson instace if the loading is completed successfully.If it fails, it returns error
func Load(path string) (*gson.Gson, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return gson.CreateWithReader(file)
}

// LoadWithReader load json from a given the reader and return gson instace if the loading is completed successfully.If it fails, it returns error
func LoadWithReader(reader io.Reader) (*gson.Gson, error) {
	return gson.CreateWithReader(reader)
}
