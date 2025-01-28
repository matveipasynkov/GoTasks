package file

import (
	"encoding/json"
	"os"
	"errors"
	"github.com/fatih/color"
)

func ReadFile(path string) ([]byte, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		color.Red(err.Error())
		return nil, err
	}
	return bytes, nil
}

func CheckJsonType(path string) error {
	bytes, _:= ReadFile(path)
	if !json.Valid(bytes) {
		color.Red("Файл не JSON-формата")
		return errors.New("NOT_JSON_FORMAT")
	}
	return nil
}
