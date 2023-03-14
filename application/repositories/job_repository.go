package repositories

import (
	"encoder/domain"
	"errors"

	"github.com/jinzhu/gorm"
)

type JobRepository interface {
	Insert(job *domain.Job) (*domain.Job, error)
	Find(id string) (*domain.Job, error)
	Update(job *domain.Job) (*domain.Job, error)
}

type JobRepositoryDb struct {
	Db *gorm.DB
}

func NewjobRepositoryDb(db *gorm.DB) *JobRepositoryDb {
	return &JobRepositoryDb{Db: db}
}

func (jr *JobRepositoryDb) Insert(job *domain.Job) (*domain.Job, error) {

	err := jr.Db.Create(job).Error

	if err != nil {
		return nil, err
	}

	return job, nil

}

func (jr *JobRepositoryDb) Find(id string) (*domain.Job, error) {

	var job domain.Job

	jr.Db.Preload("Video").Find(&job, "id = ?", id)

	if job.ID == "" {
		return nil, errors.New("job not found")
	}

	return &job, nil
}

func (jr *JobRepositoryDb) Update(job *domain.Job) (*domain.Job, error) {
	err := jr.Db.Save(&job).Error

	if err != nil {
		return nil, err
	}

	return job, nil
}
