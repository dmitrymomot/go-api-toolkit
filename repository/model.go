package repository

import (
	"time"
)

// Model base model
type Model struct {
	ID        string    `gorm:"primary_key;type:char(36);column:id" json:"id" validate:"required,uuid"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// PrimaryKey trigger
func (m *Model) PrimaryKey() string {
	return m.ID
}

// ModelSoftDelete base model
type ModelSoftDelete struct {
	Model
	DeletedAt *time.Time `sql:"index" json:"-"`
}

// ModelDates provides CreatedAt, UpdatedAt and DeletedAt properties
type ModelDates struct {
	CreatedAt *time.Time `json:"-"`
	UpdatedAt *time.Time `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

// ModelWithOwner provides UserID property
type ModelWithOwner struct {
	UserID string `gorm:"type:char(36);column:user_id;index" json:"-" validate:"omitempty,uuid"`
}

// OwnerID returns owner ID
func (m *ModelWithOwner) OwnerID() string {
	return m.UserID
}
