package models

import (
	"time"

	"github.com/google/uuid"
)

type Model interface {
	BeforeCreate()
	BeforeUpdate()
	BeforeDelete()
}

type model struct {
	ID string

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m *model) BeforeCreate() {
	m.ID = uuid.New().String()
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
}

func (m *model) BeforeUpdate() {
	m.UpdatedAt = time.Now()
}

func (m *model) BeforeDelete() {
	return
}
