package api

import (
	"3-cli/app/bins"
	"3-cli/app/config"
	"3-cli/app/file"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/fatih/color"
)

type inputStruct struct {
	Metadata struct {
		Id        string    `json:"id"`
		Private   bool      `json:"private"`
		CreatedAt time.Time `json:"createdAt"`
	} `json:"metadata"`
	Record struct {
		Information string `json:"record"`
	}
}

func GetConfig() *config.Config {
	return config.NewConfig()
}

type IStorage interface {
	ReadBins() *bins.BinList
	SaveBins(bins *bins.BinList)
	GetPath() string
}

func CreateBin(storage IStorage, filename string, binName string) (*string, error) {
	config := GetConfig()
	_, err := file.ReadFile(storage.GetPath())
	var binList bins.BinList
	if err != nil {
		binList, _ = bins.NewBinList([]bins.Bin{}...)
	} else {
		binList = *storage.ReadBins()
	}
	err = file.CheckJsonType(filename)
	if err != nil {
		color.Red(err.Error())
		return nil, err
	}
	body, _ := file.ReadFile(filename)
	req, err := http.NewRequest("POST", "https://api.jsonbin.io/v3/b", bytes.NewBuffer(body))
	if err != nil {
		color.Red(err.Error())
		return nil, err
	}
	client := &http.Client{}
	req.Header = http.Header{
		"X-Master-Key": {config.Key},
		"Content-Type": {config.ContentType},
	}
	resp, err := client.Do(req)
	if err != nil {
		color.Red(err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		color.Red(err.Error())
		return nil, err
	}
	var informationInput inputStruct
	err = json.Unmarshal(body, &informationInput)
	if err != nil {
		color.Red(err.Error())
		return nil, err
	}
	binList.Bins = append(binList.Bins, bins.Bin{
		Id:        informationInput.Metadata.Id,
		CreatedAt: informationInput.Metadata.CreatedAt,
		Private:   informationInput.Metadata.Private,
		Name:      binName,
	})
	storage.SaveBins(&binList)
	color.Green("Запись успешна.")
	return &informationInput.Metadata.Id, nil
}

func GetBin(storage IStorage, id string) error {
	config := GetConfig()
	binList := storage.ReadBins()
	if binList == nil {
		color.Red("Список бинов не открылся.")
		return errors.New("EMPTY_LIST")
	}
	checkFlag := false
	for _, bin := range binList.Bins {
		if bin.Id == id {
			checkFlag = true
			break
		}
	}
	if !checkFlag {
		color.Red("Такого бина нет.")
		return errors.New("BIN_NOT_EXIST")
	}
	req, err := http.NewRequest("GET", "https://api.jsonbin.io/v3/b/"+fmt.Sprint(id), nil)
	if err != nil {
		color.Red(err.Error())
		return err
	}
	client := &http.Client{}
	req.Header = http.Header{
		"X-Master-Key": {config.Key},
		"Content-Type": {config.ContentType},
	}
	resp, err := client.Do(req)
	if err != nil {
		color.Red(err.Error())
		return err
	}
	defer resp.Body.Close()
	bytesList, err := io.ReadAll(resp.Body)
	color.Green("Получена следующая запись:")
	color.Green(string(bytesList))
	return nil
}

func UpdateBin(storage IStorage, filename string, id string) error {
	config := GetConfig()
	binList := storage.ReadBins()
	if binList == nil {
		color.Red("Список бинов не открылся.")
		return errors.New("EMPTY_LIST")
	}
	checkFlag := false
	for _, bin := range binList.Bins {
		if bin.Id == id {
			checkFlag = true
			break
		}
	}
	if !checkFlag {
		color.Red("Такого бина нет.")
		return errors.New("BIN_NOT_EXIST")
	}
	err := file.CheckJsonType(filename)
	if err != nil {
		color.Red(err.Error())
		return err
	}
	body, _ := file.ReadFile(filename)
	req, err := http.NewRequest("PUT", "https://api.jsonbin.io/v3/b/"+id, bytes.NewBuffer(body))
	if err != nil {
		color.Red(err.Error())
		return err
	}
	client := &http.Client{}
	req.Header = http.Header{
		"X-Master-Key": {config.Key},
		"Content-Type": {config.ContentType},
	}
	resp, err := client.Do(req)
	if err != nil {
		color.Red(err.Error())
		return err
	}
	if resp.StatusCode == 200 {
		color.Green("Обновление успешно.")
		return nil
	}
	color.Red("Обновление провалено, ошибка: " + fmt.Sprintln(resp.StatusCode))
	return errors.New("WRONG_REQUEST")
}

func DeleteBin(storage IStorage, id string) error {
	config := GetConfig()
	binList := storage.ReadBins()
	if binList == nil {
		color.Red("Список бинов не открылся.")
		return errors.New("EMPTY_LIST")
	}
	checkFlag := false
	for _, bin := range binList.Bins {
		if bin.Id == id {
			checkFlag = true
			break
		}
	}
	if !checkFlag {
		color.Red("Такого бина нет.")
		return errors.New("BIN_NOT_EXIST")
	}
	req, err := http.NewRequest("DELETE", "https://api.jsonbin.io/v3/b/"+fmt.Sprint(id), nil)
	if err != nil {
		color.Red(err.Error())
		return err
	}
	client := &http.Client{}
	req.Header = http.Header{
		"X-Master-Key": {config.Key},
		"Content-Type": {config.ContentType},
	}
	resp, err := client.Do(req)
	if err != nil {
		color.Red(err.Error())
		return err
	}
	if resp.StatusCode == 200 {
		color.Green("Удаление прошло успешно.")
		for index, bin := range binList.Bins {
			if bin.Id == id {
				binList.Bins = append(binList.Bins[:index], binList.Bins[index + 1:]...)
				storage.SaveBins(binList)
				break
			}
		}
		return nil
	}
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
	color.Red("Удаление провалено, ошибка: " + fmt.Sprintln(resp.StatusCode))
	return errors.New("WRONG_REQUEST")
}

func GetList(storage IStorage) {
	binList := storage.ReadBins()
	if binList == nil {
		color.Red("Список бинов не открылся.")
		return
	}
	for _, bin := range binList.Bins {
		fmt.Println(bin.Id, ":", bin.Name)
	}
}
