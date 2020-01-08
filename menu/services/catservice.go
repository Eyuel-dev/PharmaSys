package services

import (
	"gitlab.com/username/CareFirst/entity"
	"gitlab.com/username/CareFirst/menu"
)

// ProdServiceImpl implements menu.ProductService interface
type ProdServiceImpl struct {
	prodRepo menu.ProductRepository
}

// NewProdServiceImpl will create new ProductService object
func NewProdServiceImpl(ProdRepo menu.ProductRepository) *ProdServiceImpl {
	return &ProdServiceImpl{prodRepo: ProdRepo}
}

// Products returns list of products
func (cs *ProdServiceImpl) Products() ([]entity.Products, error) {

	prods, err := cs.prodRepo.Products()

	if err != nil {
		return nil, err
	}

	return prods, nil
}

// StoreProduct persists new product information
func (cs *ProdServiceImpl) StoreProduct(prod entity.Products) error {

	err := cs.prodRepo.StoreProduct(prod)

	if err != nil {
		return err
	}

	return nil
}

// Product returns a product object with a given id
func (cs *ProdServiceImpl) Product(id int) (entity.Products, error) {

	c, err := cs.prodRepo.Product(id)

	if err != nil {
		return c, err
	}

	return c, nil
}

// UpdateProduct updates a product with new data
func (cs *ProdServiceImpl) UpdateProduct(product entity.Products) error {

	err := cs.prodRepo.UpdateProduct(product)

	if err != nil {
		return err
	}

	return nil
}

// DeleteProduct delete a product by its id
func (cs *ProdServiceImpl) DeleteProduct(id int) error {

	err := cs.prodRepo.DeleteProduct(id)
	if err != nil {
		return err
	}
	return nil
}

// // CategoryService represents gob implementation of menu.CategoryService
// type CategoryService struct {
// 	FileName string
// }

// // NewCategoryService returns new Category Service
// func NewCategoryService(fileName string) *CategoryService {
// 	return &CategoryService{FileName: fileName}
// }

// // Categories returns all categories read from gob file
// func (cs CategoryService) Categories() ([]entity.Categories, error) {

// 	raw, err := ioutil.ReadFile(cs.FileName)

// 	if err != nil {
// 		return nil, errors.New("File could not be read")
// 	}

// 	buffer := bytes.NewBuffer(raw)

// 	dec := gob.NewDecoder(buffer)

// 	var ctgs []entity.Categories

// 	err = dec.Decode(&ctgs)

// 	if err != nil {
// 		return nil, errors.New("Decoding error")
// 	}

// 	return ctgs, nil
// }

// // StoreCategories stores a batch of categories data to the a gob file
// func (cs CategoryService) StoreCategories(ctgs []entity.Categories) error {

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
