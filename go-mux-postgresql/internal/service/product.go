package service

import (
	"demo-go-basic-backend/internal/database"
	"demo-go-basic-backend/internal/model"
	"fmt"
)

var Products []*model.Product

func GetProducts(page, size int, search, sort string) ([]*model.Product, error) {
	db := database.Get()
	rows, err := db.Query(fmt.Sprintf("SELECT id, name, description, quantity FROM products_tab WHERE name ILIKE $1 ORDER BY id %s LIMIT %d OFFSET %d;", sort, size,(page-1)*size), search)
	
	if err != nil {
		return nil, err
	}
	products := []*model.Product{}
	for rows.Next() {
		//Tampungan Data
		product := &model.Product{}

		//Isi data ke tampungan
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Quantity)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}
	return products, nil
}

func GetProduct(id int) (*model.Product, error) {
	db := database.Get()
	product := &model.Product{}
	err := db.QueryRow("SELECT id, name, description, quantity FROM products_tab WHERE id = $1;", id).Scan(&product.ID, &product.Name, &product.Description, &product.Quantity)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func CreateProduct(name, description string, quantity int) (*model.Product, error) {
	db := database.Get()
	product := &model.Product{}
	statement := `INSERT INTO products_tab (name, description, quantity) VALUES ($1, $2, $3);`
	_, err := db.Exec(statement, name, description, quantity)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func DeleteProduct(id int) (*model.Product, error) {
	db := database.Get()

	var checkID int
	selectStatement := `SELECT id FROM products_tab WHERE id = $1;`
	err := db.QueryRow(selectStatement, id).Scan(&checkID)
	if err != nil {
		return nil, err
	}

	product := &model.Product{}
	statement := `DELETE FROM products_tab WHERE id = $1 RETURNING id, name, description, quantity;`
	err = db.QueryRow(statement, id).Scan(&product.ID, &product.Name, &product.Description, &product.Quantity)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func UpdateProduct(id int, updatedProduct *model.Product) (*model.Product, error) {
	db := database.Get()

	var checkID int
	selectStatement := `SELECT id FROM products_tab WHERE id = $1;`
	err := db.QueryRow(selectStatement, id).Scan(&checkID)
	if err != nil {
		return nil, err
	}

	statement := `UPDATE products_tab SET name = $2, description = $3, quantity = $4 WHERE id = $1;`
	_, err = db.Exec(statement, id, updatedProduct.Name, updatedProduct.Description, updatedProduct.Quantity)
	if err != nil {
		return nil, err
	}
	
	return updatedProduct, nil
}