package services

import (
	"gitlab.com/username/CareFirst/entity"
)

// ProductService specifies products category services
type ProductService interface {
	Products() ([]entity.Products, error)
	Product(id int) (entity.Products, error)
	UpdateProduct(prod entity.Products) error
	DeleteProduct(id int) error
	StoreProduct(prod entity.Products) error
}

// ProductService represents gob implementation of menu.ProductService
// type ProductService struct {
// 	FileName string
// }

// // NewProductService returns new Product Service
// func NewProductService(fileName string) *ProductService {
// 	return &ProductService{FileName: fileName}
// }

// // Products returns all products read from gob file
// func (cs ProductService) Products() ([]entity.Products, error) {

// 	raw, err := ioutil.ReadFile(cs.FileName)

// 	if err != nil {
// 		return nil, errors.New("File could not be read")
// 	}

// 	buffer := bytes.NewBuffer(raw)

// 	dec := gob.NewDecoder(buffer)

// 	var ctgs []entity.Products

// 	err = dec.Decode(&ctgs)

// 	if err != nil {
// 		return nil, errors.New("Decoding error")
// 	}

// 	return ctgs, nil
// }

// // StoreProducts stores a batch of products data to the a gob file
// func (cs ProductService) StoreProducts(ctgs []entity.Products) error {

// 	buffer := new(bytes.Buffer)
// 	encoder := gob.NewEncoder(buffer)

// 	err := encoder.Encode(ctgs)

// 	if err != nil {
// 		return errors.New("Data encoding has failed")
// 	}

// 	err = ioutil.WriteFile(cs.FileName, buffer.Bytes(), 0644)

// 	if err != nil {
// 		return errors.New("Writing to a file has failed")
// 	}

// 	return nil
// }
