package command

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/common/interfaces/persistence"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/product/interfaces/command"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/product/interfaces/persistence/repo"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/product/aggregate"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/product/vo"
)

type CreateProductImpl struct {
	repo.ProductRepo
	persistence.UoW
	command.CreateProduct
}

func (interactor *CreateProductImpl) Create(command command.CreateProductCommand) error {
	price, priceErr := vo.ProductPrice{}.Create(command.Price)
	if priceErr != nil {
		return priceErr
	}
	discount, discountErr := vo.ProductDiscount{}.Create(command.Discount)
	if discountErr != nil {
		return discountErr
	}

	productEntity, err := aggregate.Product{}.Create(
		vo.ProductID{Value: command.ProductID},
		price,
		discount,
		command.Description,
		command.Name,
	)
	if err != nil {
		return err
	}
	interactor.UoW.StartTx()
	err = interactor.ProductRepo.AddProduct(productEntity, interactor.UoW.GetTx())
	if err != nil {
		return err
	}
	err = interactor.UoW.Commit()
	if err != nil {
		return err
	}

	return nil
}
