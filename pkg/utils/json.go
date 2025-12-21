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

func SaveJsonFile[T any](filePath string, data []T) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}
