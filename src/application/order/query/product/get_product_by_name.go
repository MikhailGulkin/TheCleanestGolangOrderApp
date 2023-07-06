package product

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/dto"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/persistence/dao"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/query/product"
)

type GetProductByNameImpl struct {
	DAO dao.ProductReader
	product.GetProductByName
}

func (interactor *GetProductByNameImpl) Get(query product.GetProductByNameQuery) (dto.Product, error) {
	productByName, err := interactor.DAO.GetProductByName(query.Name)
	if err != nil {
		return dto.Product{}, err
	}
	return productByName, nil
}
