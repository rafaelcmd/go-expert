package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:products_categories;"`
}

type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	Categories   []Category `gorm:"many2many:products_categories;"`
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert-gorm-relationship?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	category1 := Category{Name: "Games"}
	db.Create(&category1)

	category2 := Category{Name: "Electronics"}
	db.Create(&category2)

	product := Product{Name: "Playstation", Price: 2000, Categories: []Category{category1, category2}}
	db.Create(&product)

	serialNumber := SerialNumber{Number: "123456", ProductID: 1}
	db.Create(&serialNumber)

	var products []Product
	db.Preload("Category").Preload("SerialNumber").Find(&products)
	for _, product := range products {
		fmt.Printf("%v | %v | %v\n", product.Name, product.SerialNumber.Number)
	}

	var categories []Category
	err = db.Model(&Category{}).Preload("Products.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		fmt.Printf("%v\n", category.Name)
		for _, product := range category.Products {
			fmt.Printf("  %v | %v\n", product.Name, product.SerialNumber.Number)
		}
	}
}
