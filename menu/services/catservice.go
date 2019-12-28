package services

import (
	"bytes"
	"encoding/gob"
	"errors"
	"io/ioutil"

	"gitlab.com/username/CareFirst/entity"
)

// CategoryService represents gob implementation of menu.CategoryService
type CategoryService struct {
	FileName string
}

// NewCategoryService returns new Category Service
func NewCategoryService(fileName string) *CategoryService {
	return &CategoryService{FileName: fileName}
}

// Categories returns all categories read from gob file
func (cs CategoryService) Categories() ([]entity.Categories, error) {

	raw, err := ioutil.ReadFile(cs.FileName)

	if err != nil {
		return nil, errors.New("File could not be read")
	}

	buffer := bytes.NewBuffer(raw)

	dec := gob.NewDecoder(buffer)

	var ctgs []entity.Categories

	err = dec.Decode(&ctgs)

	if err != nil {
		return nil, errors.New("Decoding error")
	}

	return ctgs, nil
}

// StoreCategories stores a batch of categories data to the a gob file
func (cs CategoryService) StoreCategories(ctgs []entity.Categories) error {

	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)

	err := encoder.Encode(ctgs)

	if err != nil {
		return errors.New("Data encoding has failed")
	}

	err = ioutil.WriteFile(cs.FileName, buffer.Bytes(), 0644)

	if err != nil {
		return errors.New("Writing to a file has failed")
	}

	return nil
}
