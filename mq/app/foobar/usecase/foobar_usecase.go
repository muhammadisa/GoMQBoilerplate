package usecase

import (
	"time"

	"github.com/muhammadisa/go-mq-boilerplate/mq/app/foobar"
	"github.com/muhammadisa/go-mq-boilerplate/mq/models"
	uuid "github.com/satori/go.uuid"
)

// foobarUsecase struct
type foobarUsecase struct {
	foobarRepository foobar.Repository
}

// NewFoobarUsecase function
func NewFoobarUsecase(foobarUsecases foobar.Repository) foobar.Usecase {
	return &foobarUsecase{
		foobarRepository: foobarUsecases,
	}
}

func (foobarUsecases foobarUsecase) Store(foobar *models.Foobar) error {
	return foobarUsecases.foobarRepository.Store(foobar)
}

func (foobarUsecases foobarUsecase) Update(foobar *models.Foobar) error {
	foobar.UpdatedAt = time.Now()
	return foobarUsecases.foobarRepository.Update(foobar)
}

func (foobarUsecases foobarUsecase) Delete(id uuid.UUID) error {
	return foobarUsecases.foobarRepository.Delete(id)
}
