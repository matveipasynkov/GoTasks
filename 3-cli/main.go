// Пример работы: go run main.go get --id=<BIN_ID>

package main

import (
	"3-cli/app/api"
	"3-cli/app/storage"
	"flag"
	"fmt"
	"strings"

	"github.com/fatih/color"
)

func main() {
	storage := storage.NewStorage("bins.json")
	flag.Parse()
	cmd := flag.Args()[0]
	switch cmd {
	case "create":
		api.CreateBin(storage, strings.Split(flag.Args()[1], "=")[1], strings.Split(flag.Args()[2], "=")[1])
	case "update":
		api.UpdateBin(storage, strings.Split(flag.Args()[1], "=")[1], strings.Split(flag.Args()[2], "=")[1])
	case "get":
		api.GetBin(storage, strings.Split(flag.Args()[1], "=")[1])
	case "delete":
		api.DeleteBin(storage, strings.Split(flag.Args()[1], "=")[1])
	case "list":
		api.GetList(storage)
	default:
		color.Red("Нет такой команды.")
	}
}

func promptData(text string) string {
	var answer string
	fmt.Print(text + ": ")
	fmt.Scan(&answer)
	return answer
}
