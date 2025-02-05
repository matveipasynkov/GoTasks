package file

import (
	"encoding/json"
	"errors"
	"os"
)

func ReadFile(path string) ([]byte, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func CheckJsonType(path string) error {
	bytes, err := ReadFile(path)
	if err != nil {
		return err
	}
	if !json.Valid(bytes) {
		return errors.New("NOT_JSON_FORMAT")
	}
	return nil
}
