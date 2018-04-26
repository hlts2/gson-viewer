package jsonviewer

import (
	"os"

	"github.com/hlts2/gson"
)

// LoadJSON load json from a given path and return gson instace if the loading is completed successfully.If it fails, it returns error
func LoadJSON(path string) (*gson.Gson, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return gson.NewGsonFromReader(file)
}
