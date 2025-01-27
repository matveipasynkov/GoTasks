package bins

import (
	"time"
)

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

func newBin(id string, private bool, createdAt time.Time, name string) (Bin, error) {
	return Bin{
		id:        id,
		private:   private,
		createdAt: createdAt,
		name:      name,
	}, nil
}

type BinList struct {
	bins []Bin
}

func newBinList(bins ...Bin) (BinList, error) {
	list_bins := []Bin{}
	for _, bin := range bins {
		list_bins = append(list_bins, bin)
	}
	return BinList{
		bins: list_bins,
	}, nil
}
