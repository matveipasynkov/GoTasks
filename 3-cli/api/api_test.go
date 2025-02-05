package api_test

import (
	"3-cli/app/api"
	"3-cli/app/storage"
	"testing"
)

func TestCreateBin(t *testing.T) {
	//Arrange
	storage := storage.NewStorage("bins.json")
	filename := "data.json"
	binName := "new_bin"
	var expected error = nil
	//Act
	id, err := api.CreateBin(storage, filename, binName)
	//Assert
	if err != expected {
		t.Errorf("Ожидалось %v, а получилось %v", expected, err)
	} else {
		api.DeleteBin(storage, *id)
	}
}

func TestUpdateBin(t *testing.T) {
	//Arrange
	storage := storage.NewStorage("bins.json")
	filename := "data.json"
	binName := "new_bin"
	var expected error = nil
	//Act
	id, err1 := api.CreateBin(storage, filename, binName)
	err2 := api.UpdateBin(storage, filename, *id)
	//Assert
	if err1 != expected {
		t.Errorf("Ожидалось %v, а получилось %v", expected, err1)
	}
	if err2 != expected {
		t.Errorf("Ожидалось %v, а получилось %v", expected, err2)
	}
	api.DeleteBin(storage, *id)
}

func TestDeleteBin(t *testing.T) {
	//Arrange
	storage := storage.NewStorage("bins.json")
	filename := "data.json"
	binName := "new_bin"
	var expected error = nil
	//Act
	id, err1 := api.CreateBin(storage, filename, binName)
	err2 := api.DeleteBin(storage, *id)
	//Assert
	if err1 != expected {
		t.Errorf("Ожидалось %v, а получилось %v", expected, err1)
	}
	if err2 != expected {
		t.Errorf("Ожидалось %v, а получилось %v", expected, err2)
	}
}

func TestGetBin(t *testing.T) {
	//Arrange
	storage := storage.NewStorage("bins.json")
	filename := "data.json"
	binName := "new_bin"
	var expected error = nil
	//Act
	id, err1 := api.CreateBin(storage, filename, binName)
	err2 := api.GetBin(storage, *id)
	//Assert
	if err1 != expected {
		t.Errorf("Ожидалось %v, а получилось %v", expected, err1)
	}
	if err2 != expected {
		t.Errorf("Ожидалось %v, а получилось %v", expected, err2)
	}
	api.DeleteBin(storage, *id)
}