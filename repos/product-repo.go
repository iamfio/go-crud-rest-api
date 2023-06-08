package repos

import (
	"fmt"

	"github.com/iamfio/crud-rest-api/entities"
)

type ProductRepository struct {
	products []entities.Product
}

func NewProductRepository() *ProductRepository {
	var pr = ProductRepository{make([]entities.Product, 0)}
	return &pr
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

func (p *ProductRepository) Update(id uint, ameneded entities.Product) (entities.Product, error) {
	for i, it := range p.products {
		if it.ID == id {
			ameneded.ID = id
			p.products = append(p.products[:i], p.products[i+1:]...)
			p.products = append(p.products, ameneded)
			return ameneded, nil
		}
	}
	return entities.Product{}, fmt.Errorf("key '%d' not found", ameneded.ID)
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
