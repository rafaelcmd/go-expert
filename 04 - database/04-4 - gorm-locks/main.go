package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Product struct {
	ID    string `gorm:"primaryKey"`
	Name  string
	Price float64
}

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert-gorm-locks?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{})

	newProduct := Product{ID: "1", Name: "Product 1", Price: 100}
	db.Create(&newProduct)

	newCategory := Category{ID: 1, Name: "Category 1"}
	db.Create(&newCategory)

	tx := db.Begin()
	var category Category
	err = tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&category, "id = ?", 1).Error
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	category.Name = "Updated Category"
	err = tx.Debug().Save(&category).Error
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	tx.Commit()
}
