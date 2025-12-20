package utils

import (
	"encoding/json"
	"os"
)

func ReadJsonFile[T any](filePath string) ([]T, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var items []T
	err = json.Unmarshal(data, &items)
	if err != nil {
		return nil, err
	}
	return items, nil
}
