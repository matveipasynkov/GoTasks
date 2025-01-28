package bins

import (
	"encoding/json"
	"time"
)

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}

func NewBin(id string, private bool, createdAt time.Time, name string) (Bin, error) {
	return Bin{
		Id:        id,
		Private:   private,
		CreatedAt: createdAt,
		Name:      name,
	}, nil
}

type BinList struct {
	Bins []Bin `json:"account"`
}

func NewBinList(bins ...Bin) (BinList, error) {
	list_bins := []Bin{}
	for _, bin := range bins {
		list_bins = append(list_bins, bin)
	}
	return BinList{
		Bins: list_bins,
	}, nil
}

func (binList *BinList) ToBytes() ([]byte, error) {
	bytes, err := json.Marshal(binList)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
