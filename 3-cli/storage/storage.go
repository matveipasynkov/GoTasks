package storage

import (
	"3-cli/app/bins"
	"3-cli/app/file"
	"encoding/json"
	"os"

	"github.com/fatih/color"
)

type Storage struct {
	Path string
}

func NewStorage(newPath string) *Storage {
	return &Storage{
		Path: newPath,
	}
}

func (storage *Storage) SaveBins(bins *bins.BinList) {
	file, err := os.Create(storage.Path)
	if err != nil {
		color.Red(err.Error())
		return
	}
	bytes, err := bins.ToBytes()
	if err != nil {
		color.Red(err.Error())
		return
	}
	_, err = file.Write(bytes)
	if err != nil {
		color.Red(err.Error())
		return
	}
}

func (storage *Storage) ReadBins() *bins.BinList {
	err := file.CheckJsonType(storage.Path)
	if err != nil {
		color.Red("Файл неверного формата или не существует")
		return nil
	}
	bytes, err := os.ReadFile(storage.Path)
	if err != nil {
		color.Red(err.Error())
		return nil
	}
	var readedBins bins.BinList
	err = json.Unmarshal(bytes, &readedBins)
	if err != nil {
		color.Red(err.Error())
	}
	return &readedBins
}

func (storage *Storage) GetPath() string {
	return storage.Path
}
