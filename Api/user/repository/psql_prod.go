package repository

import (
	"database/sql"
	"errors"

	"github.com/Eyuel-dev/PharmaSys/api/entity"
)

// ProdRepositoryImpl implements the menu.ProductRepository interface
type ProdRepositoryImpl struct {
	conn *sql.DB
}

// NewProdRepositoryImpl will create an object of PsqlProdRepository
func NewProdRepositoryImpl(Conn *sql.DB) *ProdRepositoryImpl {
	return &ProdRepositoryImpl{conn: Conn}
}

// Products returns all products from the database
func (cri *ProdRepositoryImpl) Products() ([]entity.Products, error) {

	rows, err := cri.conn.Query("SELECT * FROM products;")
	if err != nil {
		return nil, errors.New("Could not query the database")
	}
	defer rows.Close()

	prds := []entity.Products{}

	for rows.Next() {
		products := entity.Products{}
		err = rows.Scan(&products.ID, &products.Name, &products.Description, &products.Image, &products.Price)
		if err != nil {
			return nil, err
		}
		prds = append(prds, products)
	}

	return prds, nil
}

// Product returns a product with a given id
func (cri *ProdRepositoryImpl) Product(id int) (entity.Products, error) {

	row := cri.conn.QueryRow("SELECT * FROM products WHERE id = $1", id)

	p := entity.Products{}

	err := row.Scan(&p.ID, &p.Name, &p.Description, &p.Image, &p.Price)
	if err != nil {
		return p, err
	}

	return p, nil
}

// UpdateProduct updates a given object with a new data
func (cri *ProdRepositoryImpl) UpdateProduct(c entity.Products) error {

	_, err := cri.conn.Exec("UPDATE products SET name=$1,description=$2, image=$3, price=$4 WHERE id=$5", c.Name, c.Description, c.Image, c.Price, c.ID)
	if err != nil {
		return errors.New("Update has failed")
	}

	return nil
}

// DeleteProduct removes a product from a database by its id
func (cri *ProdRepositoryImpl) DeleteProduct(id int) error {

	_, err := cri.conn.Exec("DELETE FROM products WHERE id=$1", id)
	if err != nil {
		return errors.New("Delete has failed")
	}

	return nil
}

// StoreProduct stores new product information to database
func (cri *ProdRepositoryImpl) StoreProduct(c entity.Products) error {

	_, err := cri.conn.Exec("INSERT INTO products (name,description,image,price) values($1, $2, $3, $4)", c.Name, c.Description, c.Image, c.Price)
	if err != nil {
		return errors.New("Insertion has failed")
	}

	return nil
}
