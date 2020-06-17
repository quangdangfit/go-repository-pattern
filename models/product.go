package models

type Product struct {
	Code string `json:"code,omitempty" bson:"code,omitempty"`
	Name string `json:"name,omitempty" bson:"name,omitempty"`
}
