package product

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/entities/product"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/models"
	"gorm.io/gorm"
)

func CreateValidProductEntity() product.Product {
	entity, _ := product.Product{}.Create(100, 10, "", "some name")
	return entity
}
func CreateValidProductModel(entity product.Product) models.Product {
	return models.Product{
		Base:         models.Base{ID: entity.ProductID.Value},
		Price:        entity.Price,
		Name:         entity.Name,
		Discount:     entity.Discount,
		Quantity:     entity.Quantity,
		Description:  entity.Description,
		Availability: entity.Availability,
	}
}
func CreateProductInDB(model *models.Product, db *gorm.DB) {
	db.Create(&model)
}
