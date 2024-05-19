package entity

import (
	"errors"
	"time"

	"github.com/eron97/project-fullCycle.git/pkg/entity"
)

var (
	ErrIDIsRequired    = errors.New("id is requerired")
	ErrInvalidID       = errors.New("invalid id")
	ErrNameisRequired  = errors.New("name is requiered")
	ErrPriceisRequired = errors.New("price is required")
	ErrInvalidPrice    = errors.New("invalid price")
)

type Product struct {
	ID       entity.ID `json:"id"`
	Name     string    `json:"name"`
	Price    int       `json:"price"`
	CreateAt string    `json:"created_at"`
}

func NewProduct(name string, price int) (*Product, error) {
	product := &Product{
		ID:       entity.NewID(),
		Name:     name,
		Price:    price,
		CreateAt: time.Now().String(),
	}

	err := product.Validate()
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *Product) Validate() error {

	if p.ID.String() == "" {
		return ErrIDIsRequired
	}
	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrInvalidID
	}
	if p.Name == "" {
		return ErrNameisRequired
	}
	if p.Price == 0 {
		return ErrPriceisRequired
	}
	if p.Price < 0 {
		return ErrInvalidPrice
	}
	return nil
}
