package repository

import (
	"time"

	"github.com/gocraft/dbr/v2"
	"github.com/muhammadisa/go-mq-boilerplate/mq/app/foobar"
	"github.com/muhammadisa/go-mq-boilerplate/mq/models"
	"github.com/muhammadisa/go-mq-boilerplate/mq/utils/errhandler"
	uuid "github.com/satori/go.uuid"
)

type ormFoobarRepo struct {
	Session *dbr.Session
}

// NewORMFoobarRepo function
func NewORMFoobarRepo(sess *dbr.Session) foobar.Repository {
	return &ormFoobarRepo{
		Session: sess,
	}
}

func (foobarRepository *ormFoobarRepo) Store(foobar *models.Foobar) error {
	_, err := foobarRepository.Session.InsertInto("foobars").
		Columns("id", "foobar_content", "created_at").
		Record(foobar).
		Exec()
	return errhandler.HandleErrorThenReturn(err)
}

func (foobarRepository *ormFoobarRepo) Update(foobar *models.Foobar) error {
	_, err := foobarRepository.Session.Update("foobars").
		Where("id = ?", foobar.ID.String()).
		SetMap(map[string]interface{}{
			"foobar_content": foobar.FoobarContent,
			"updated_at":     time.Now(),
		}).
		Exec()
	return errhandler.HandleErrorThenReturn(err)
}

func (foobarRepository *ormFoobarRepo) Delete(id uuid.UUID) error {
	_, err := foobarRepository.Session.DeleteFrom("foobars").
		Where("id = ?", id.String()).
		Exec()
	return errhandler.HandleErrorThenReturn(err)
}
