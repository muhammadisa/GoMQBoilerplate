package foobar

import (
	"github.com/muhammadisa/go-mq-boilerplate/mq/models"
	uuid "github.com/satori/go.uuid"
)

// Usecase interface
type Usecase interface {
	Update(foobar *models.Foobar) error
	Store(foobar *models.Foobar) error
	Delete(id uuid.UUID) error
}
