package main

import "3-cli/app/bins"

type IStorage interface {
	ReadBins(path string) *bins.BinList
	SaveBins(bins *bins.BinList, path string)
}

func main() {

}
