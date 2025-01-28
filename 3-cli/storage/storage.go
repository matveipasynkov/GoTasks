package storage

import (
	"3-cli/app/bins"
	"encoding/json"
	"os"

	"github.com/fatih/color"
)

func SaveBins(bins *bins.BinList, path string) {
	file, err := os.Create(path)
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

func ReadBins(path string) *bins.BinList {
	bytes, err := os.ReadFile(path)
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
