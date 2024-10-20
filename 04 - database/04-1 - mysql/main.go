package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{uuid.New().String(), name, price}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	product := NewProduct("Product 1", 100)
	err = insertProduct(db, product)
	if err != nil {
		panic(err)
	}

	product.Price = 500
	err = updateProduct(db, product)
	if err != nil {
		panic(err)
	}

	product, err = getProduct(db, product.ID)
	if err != nil {
		panic(err)
	}
	fmt.Println(product)

	products, err := getAllProducts(db)
	if err != nil {
		panic(err)
	}
	fmt.Println(products)

	err = deleteProduct(db, product.ID)
	if err != nil {
		panic(err)
	}
}

func insertProduct(db *sql.DB, product *Product) error {
	_, err := db.Exec("INSERT INTO products (id, name, price) VALUES (?, ?, ?)", product.ID, product.Name, product.Price)
	return err
}

func updateProduct(db *sql.DB, product *Product) error {
	_, err := db.Exec("UPDATE products SET name = ?, price = ? WHERE id = ?", product.Name, product.Price, "2f5335c4-d7ee-4925-aa02-b339767331f4")
	return err
}

func getProduct(db *sql.DB, id string) (*Product, error) {
	var product Product
	err := db.QueryRow("SELECT id, name, price FROM products WHERE id = ?", id).Scan(&product.ID, &product.Name, &product.Price)
	return &product, err
}

func getAllProducts(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("SELECT id, name, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		err = rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func deleteProduct(db *sql.DB, id string) error {
	_, err := db.Exec("DELETE FROM products WHERE id = ?", id)
	return err
}
