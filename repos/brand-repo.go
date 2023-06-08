package repos

import (
	"fmt"

	"github.com/iamfio/crud-rest-api/entities"
)

type BrandRepository struct {
	brands []entities.Brand
}

func NewBrandRepository() *BrandRepository {
	var br = BrandRepository{make([]entities.Brand, 0)}
	return &br
}

func (b *BrandRepository) Create(partial entities.Brand) entities.Brand {
	newItem := entities.Brand{ID: uint(len(b.brands)) + 1, Name: partial.Name, Year: partial.Year}
	b.brands = append(b.brands, newItem)
	return newItem
}

func (b *BrandRepository) GetList() []entities.Brand {
	return b.brands
}

func (b *BrandRepository) GetOne(id uint) (entities.Brand, error) {
	for _, it := range b.brands {
		if it.ID == id {
			return it, nil
		}
	}
	return entities.Brand{}, fmt.Errorf("key '%d' not found", id)
}

func (b *BrandRepository) Update(id uint, amended entities.Brand) (entities.Brand, error) {
	for i, it := range b.brands {
		if it.ID == id {
			amended.ID = id
			b.brands = append(b.brands[:i], b.brands[i+1:]...)
			b.brands = append(b.brands, amended)
			return amended, nil
		}
	}
	return entities.Brand{}, fmt.Errorf("key '%d' not found", amended.ID)
}

func (b *BrandRepository) DeleteOne(id uint) (bool, error) {
	for i, it := range b.brands {
		if it.ID == id {
			b.brands = append(b.brands[:i], b.brands[i+1:]...)
			return true, nil
		}
	}
	return false, fmt.Errorf("key '%d' not found", id)
}
