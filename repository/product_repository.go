package repository

import (
	"api/model"
	"database/sql"
	"fmt"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT * FROM products"
	rows, err := pr.connection.Query(query)

	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}

	var productList []model.Product
	var productObj model.Product

	for rows.Next() {
		err = rows.Scan(
			&productObj.ID,
			&productObj.Name,
			&productObj.Price,
		)

		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}

		productList = append(productList, productObj)
	}

	rows.Close()

	fmt.Println(productList)

	return productList, nil
}

func (pr *ProductRepository) CreateProduct(product model.Product) (int, error) {
	var id int
	query, err := pr.connection.Prepare(`INSERT INTO products(name, price) VALUES($1,$2) RETURNING id`)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(product.Name, product.Price).Scan(&id)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()

	return id, nil

}

func (pr *ProductRepository) GetProduct(id int) (*model.Product, error) {
	var productObj = model.Product{}
	query := "SELECT * FROM products p WHERE p.id = $1"
	row := pr.connection.QueryRow(query, id)

	err := row.Scan(
		&productObj.ID,
		&productObj.Name,
		&productObj.Price,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Product Not Found")
			return nil, nil
		}

		fmt.Println(err)
		return nil, err
	}
	return &productObj, err
}
