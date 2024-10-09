package usecase

import "api_go_rest/model"

type productUsecase struct {
	//Repository
}

func NewProductUseCase() productUsecase {
	return productUsecase{}
}

func (pu *productUsecase) GetProducts() ([]model.Product, error) {
	return []model.Product{}, nil
}
