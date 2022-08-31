package main

import (
	"demo-go-basic-backend/model"
	"demo-go-basic-backend/database"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func getProducts() {
	db := database.Get()
	rows, err := db.Query("SELECT id, name, description, quantity FROM products_tab;")
	if err != nil {
		log.Fatalln(err)
	}
	// defer rows.Close()

	for rows.Next() {
		//Tampungan Data
		product := &model.Product{}

		//Isi data ke tampungan
		rows.Scan(&product.ID, &product.Name, &product.Description, &product.Quantity)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(product)
	}
}

func getProduct() {
	db := database.Get()
	product := &model.Product{}
	err := db.QueryRow("SELECT id, name FROM products_tab WHERE id = 1;").Scan(&product.ID, &product.Name)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(product)
}

func createProduct(product *model.Product) {
	db := database.Get()
	statement := `INSERT INTO products_tab (name, description, quantity) VALUES ($1, $2, $3);`
	_, err := db.Exec(statement, product.Name, product.Description, product.Quantity)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Create Product Done")
}

func deleteProduct(id int64) {
	db := database.Get()
	statement := `DELETE FROM products_tab WHERE id = $1;`
	_, err := db.Exec(statement, id)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Delete Product Done")
}

func updateProduct(id int64, product *model.Product) {
	db := database.Get()
	statement := `UPDATE products_tab SET name = $2, description = $3, quantity = $4 WHERE id = $1;`
	_, err := db.Exec(statement, id, product.Name, product.Description, product.Quantity)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Update Product Done")
}

func main() {
	database.Connect()
	
	// createProduct(&model.Product{
	// 	Name: "Product1",
	// 	Description: "Ini Product1",
	// 	Quantity: 50,
	// })

	// updateProduct(105, &model.Product{
	// 	Name: "ProductX",
	// 	Description: "Ini ProductX",
	// 	Quantity: 25,
	// })

	// deleteProduct(105)

	// getProduct()
	// getProducts()
}