package database

import (
	"fmt"
	"github.com/rafaelcmd/go-expert/07-apis/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"math/rand"
	"testing"
)

func TestCreateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	err = db.AutoMigrate(&entity.Product{})
	if err != nil {
		t.Error(err)
	}
	product, err := entity.NewProduct("Product 1", 10)
	assert.NoError(t, err)
	productDB := NewProduct(db)
	err = productDB.Create(product)
	assert.NoError(t, err)
	assert.NotEmpty(t, product.ID)
}

func TestFindAllProducts(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	err = db.AutoMigrate(&entity.Product{})
	if err != nil {
		t.Error(err)
	}
	for i := 0; i < 10; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Int())
		assert.NoError(t, err)
		productDB := NewProduct(db)
		err = productDB.Create(product)
		assert.NoError(t, err)
	}
	products, err := NewProduct(db).FindAll(1, 5, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 5)
	products, err = NewProduct(db).FindAll(2, 5, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 5)
	products, err = NewProduct(db).FindAll(3, 5, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 0)
}

func TestFindProductByID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	err = db.AutoMigrate(&entity.Product{})
	if err != nil {
		t.Error(err)
	}
	product, err := entity.NewProduct("Product 1", 10)
	assert.NoError(t, err)
	productDB := NewProduct(db)
	err = productDB.Create(product)
	assert.NoError(t, err)
	productFound, err := productDB.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.NotNil(t, productFound)
}

func TestUpdateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	err = db.AutoMigrate(&entity.Product{})
	if err != nil {
		t.Error(err)
	}
	product, err := entity.NewProduct("Product 1", 10)
	assert.NoError(t, err)
	productDB := NewProduct(db)
	err = productDB.Create(product)
	assert.NoError(t, err)
	product.Price = 20
	err = productDB.Update(product)
	assert.NoError(t, err)
	productFound, err := productDB.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, 20, productFound.Price)
}

func TestDeleteProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	err = db.AutoMigrate(&entity.Product{})
	if err != nil {
		t.Error(err)
	}
	product, err := entity.NewProduct("Product 1", 10)
	assert.NoError(t, err)
	productDB := NewProduct(db)
	err = productDB.Create(product)
	assert.NoError(t, err)
	err = productDB.Delete(product.ID.String())
	assert.NoError(t, err)
	productFound, err := productDB.FindByID(product.ID.String())
	assert.Error(t, err)
	assert.Nil(t, productFound)
}
