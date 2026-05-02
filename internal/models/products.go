package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Product struct {
	ID          int       `json:"-"`
	Name        string    `json:"name" gorm:"colum:name;type:varchar(255)" validate:"required"`
	Description string    `json:"description" gorm:"colum:name;type:text" validate:"required"`
	Price       float64   `json:"" gorm:"column:price;type:decimal(10,2)" validate:"required"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}

func (*Product) TableName() string {
	return "products"
}

func (l Product) Validate() error {
	v := validator.New()
	return v.Struct(l)
}

type ProductCategory struct {
	ID        int       `json:"-"`
	Name      string    `json:"name" gorm:"colum:name;type:varchar(255)" validate:"required"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (*ProductCategory) TableName() string {
	return "product_categories"
}

func (l ProductCategory) Validate() error {
	v := validator.New()
	return v.Struct(l)
}

type ProductVariant struct {
	ID        int       `json:"-"`
	ProductID int       `json:"product_id" gorm:"colum:product_id;type:int"`
	Color     string    `json:"colors" gorm:"colum:name;type:varchar(50)" validate:"required"`
	Size      int       `json:"size" gorm:"colum:name;type:varchar(10)" validate:"required"`
	Quantity  int       `json:"quantities" gorm:"colum:quantity;type:int"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (*ProductVariant) TableName() string {
	return "product_variant"
}

func (l ProductVariant) Validate() error {
	v := validator.New()
	return v.Struct(l)
}
