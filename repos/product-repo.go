package repos

import (
	"fmt"

	"github.com/iamfio/crud-rest-api/entities"
)

type ProductRepository struct {
	products []entities.Product
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{make([]entities.Product, 0)}
}

func (p *ProductRepository) Create(partial entities.Product) entities.Product {
	newItem := partial
	newItem.ID = uint(len(p.products)) + 1
	p.products = append(p.products, newItem)
	return newItem
}

func (p *ProductRepository) GetList() []entities.Product {
	return p.products
}

func (p *ProductRepository) GetOne(id uint) (entities.Product, error) {
	for _, it := range p.products {
		if it.ID == id {
			return it, nil
		}
	}
	return entities.Product{}, fmt.Errorf("key '%d' not found", id)
}

func (p *ProductRepository) Update(id uint, amened entities.Product) (entities.Product, error) {
	for i, it := range p.products {
		if it.ID == id {
			amened.ID = id
			p.products = append(p.products[:i], p.products[i+1:]...)
			return amened, nil
		}
	}
	return entities.Product{}, fmt.Errorf("key '%d' not found", amened.ID)
}

func (p *ProductRepository) DeleteOne(id uint) (bool, error) {
	for i, it := range p.products {
		if it.ID == id {
			p.products = append(p.products[:i], p.products[i+1:]...)
			return true, nil
		}
	}
	return false, fmt.Errorf("key '%d' not found", id)
}
