package models

type Product struct {
	model
	Code string `json:"code,omitempty" bson:"code,omitempty"`
	Name string `json:"name,omitempty" bson:"name,omitempty"`
}

type ProductRequest struct {
	Code string `json:"code,omitempty" bson:"code,omitempty" validate:"required"`
	Name string `json:"name,omitempty" bson:"name,omitempty" validate:"required"`
}
