package models

import (
	"time"
)

type Product struct {
	Code string `json:"code,omitempty" bson:"code,omitempty"`
	Name string `json:"name,omitempty" bson:"name,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type ProductRequest struct {
	Code string `json:"code,omitempty" bson:"code,omitempty" validate:"required"`
	Name string `json:"name,omitempty" bson:"name,omitempty" validate:"required"`
}

func (p *Product) BeforeCreate() {
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}
