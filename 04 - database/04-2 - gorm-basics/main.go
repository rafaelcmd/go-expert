package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    string `gorm:"primaryKey"`
	Name  string
	Price float64
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert-gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	db.Create(&Product{Name: "Product 1", Price: 100})

	products := []Product{
		{ID: "1", Name: "Notebook", Price: 100},
		{ID: "2", Name: "Smartphone", Price: 200},
		{ID: "3", Name: "Tablet", Price: 300},
		{ID: "4", Name: "Smartwatch", Price: 400},
		{ID: "5", Name: "Headphones", Price: 500},
		{ID: "6", Name: "Keyboard", Price: 600},
		{ID: "7", Name: "Mouse", Price: 700},
		{ID: "8", Name: "Monitor", Price: 800},
		{ID: "9", Name: "Printer", Price: 900},
		{ID: "10", Name: "Scanner", Price: 1000},
	}
	db.Create(&products)

	var product Product
	db.First(&product, "id = ?", "1")
	fmt.Println(product)

	var product1 Product
	db.First(&product1, "name = ?", "Product 2")
	fmt.Println(product1)

	var findProducts []Product
	db.Find(&findProducts)
	fmt.Println(findProducts)

	var twoProducts []Product
	db.Limit(2).Find(&twoProducts)
	fmt.Println(twoProducts)

	var offsetProducts []Product
	db.Limit(4).Offset(2).Find(&offsetProducts)
	fmt.Println(offsetProducts)

	var whereProducts []Product
	db.Where("price > ?", 200).Find(&whereProducts)
	fmt.Println(whereProducts)

	var likeProducts []Product
	db.Where("name LIKE ?", "%Product%").Find(&likeProducts)
	fmt.Println(likeProducts)

	var deleteProduct Product
	db.First(&deleteProduct, "id = ?", "10")
	db.Delete(&deleteProduct)

	var updateProduct Product
	db.First(&updateProduct, "id = ?", "2")
	updateProduct.Price = 500
	db.Save(&updateProduct)
}
