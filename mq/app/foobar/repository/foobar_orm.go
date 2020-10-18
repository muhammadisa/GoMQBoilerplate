package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/muhammadisa/go-mq-boilerplate/mq/app/foobar"
	"github.com/muhammadisa/go-mq-boilerplate/mq/models"
	"github.com/muhammadisa/go-mq-boilerplate/mq/utils/errhandler"
	uuid "github.com/satori/go.uuid"
)

type ormFoobarRepo struct {
	DB *gorm.DB
}

// NewORMFoobarRepo function
func NewORMFoobarRepo(db *gorm.DB) foobar.Repository {
	return &ormFoobarRepo{
		DB: db,
	}
}

func (foobarRepository *ormFoobarRepo) Store(foobar *models.Foobar) error {
	err := foobarRepository.DB.Model(
		&models.Foobar{},
	).Create(
		foobar,
	).Error
	return errhandler.HandleErrorThenReturn(err)
}

func (foobarRepository *ormFoobarRepo) Update(foobar *models.Foobar) error {
	err := foobarRepository.DB.Model(
		&models.Foobar{},
	).Where(
		"id = ?",
		foobar.ID.String(),
	).Update(models.Foobar{
		FoobarContent: foobar.FoobarContent,
		UpdatedAt:     foobar.UpdatedAt,
	}).Error
	return errhandler.HandleErrorThenReturn(err)
}

func (foobarRepository *ormFoobarRepo) Delete(id uuid.UUID) error {
	err := foobarRepository.DB.Model(
		&models.Foobar{},
	).Where(
		"id = ?",
		id,
	).Delete(
		&models.Foobar{},
	).Error
	return errhandler.HandleErrorThenReturn(err)
}
