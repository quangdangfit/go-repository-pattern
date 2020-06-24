package models

type Model interface {
	BeforeCreate()
	BeforeUpdate()
	BeforeDelete()
}
